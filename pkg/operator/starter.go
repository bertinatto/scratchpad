package operator

import (
	"context"
	"fmt"
	"strings"
	"time"

	gcfg "gopkg.in/gcfg.v1"

	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/dynamic"
	kubeclient "k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"k8s.io/legacy-cloud-providers/gce"

	opv1 "github.com/openshift/api/operator/v1"
	configclient "github.com/openshift/client-go/config/clientset/versioned"
	configinformers "github.com/openshift/client-go/config/informers/externalversions"
	"github.com/openshift/gcp-filestore-csi-driver-operator/assets"
	"github.com/openshift/library-go/pkg/controller/controllercmd"
	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/csi/csicontrollerset"
	"github.com/openshift/library-go/pkg/operator/csi/csidrivercontrollerservicecontroller"
	"github.com/openshift/library-go/pkg/operator/csi/csidrivernodeservicecontroller"
	"github.com/openshift/library-go/pkg/operator/csi/csistorageclasscontroller"
	goc "github.com/openshift/library-go/pkg/operator/genericoperatorclient"
	"github.com/openshift/library-go/pkg/operator/v1helpers"
)

const (
	// Operand and operator run in the same namespace
	defaultNamespace     = "openshift-cluster-csi-drivers"
	operatorName         = "gcp-filestore-csi-driver-operator"
	operandName          = "gcp-filestore-csi-driver"
	secretName           = "gcp-filestore-cloud-credentials"
	trustedCAConfigMap   = "gcp-filestore-csi-driver-trusted-ca-bundle"
	cloudConfigNamespace = "openshift-config"
	cloudConfigName      = "cloud-provider-config"
	cloudConfigKey       = "config"
)

func RunOperator(ctx context.Context, controllerConfig *controllercmd.ControllerContext) error {
	// Create core clientset and informers
	kubeClient := kubeclient.NewForConfigOrDie(rest.AddUserAgent(controllerConfig.KubeConfig, operatorName))
	kubeInformersForNamespaces := v1helpers.NewKubeInformersForNamespaces(kubeClient, defaultNamespace, cloudConfigNamespace, "")
	secretInformer := kubeInformersForNamespaces.InformersFor(defaultNamespace).Core().V1().Secrets()
	configMapInformer := kubeInformersForNamespaces.InformersFor(defaultNamespace).Core().V1().ConfigMaps()
	nodeInformer := kubeInformersForNamespaces.InformersFor("").Core().V1().Nodes()

	// Create config clientset and informer. This is used to get the cluster ID
	configClient := configclient.NewForConfigOrDie(rest.AddUserAgent(controllerConfig.KubeConfig, operatorName))
	configInformers := configinformers.NewSharedInformerFactory(configClient, 20*time.Minute)
	infraInformer := configInformers.Config().V1().Infrastructures()

	// Create GenericOperatorclient. This is used by the library-go controllers created down below
	gvr := opv1.SchemeGroupVersion.WithResource("clustercsidrivers")
	operatorClient, dynamicInformers, err := goc.NewClusterScopedOperatorClientWithConfigName(controllerConfig.KubeConfig, gvr, string(opv1.GCPFilestoreCSIDriver))
	if err != nil {
		return err
	}

	dynamicClient, err := dynamic.NewForConfig(controllerConfig.KubeConfig)
	if err != nil {
		return err
	}

	csiControllerSet := csicontrollerset.NewCSIControllerSet(
		operatorClient,
		controllerConfig.EventRecorder,
	).WithLogLevelController().WithManagementStateController(
		operandName,
		false,
	).WithStaticResourcesController(
		"GCPFilestoreDriverStaticResourcesController",
		kubeClient,
		dynamicClient,
		kubeInformersForNamespaces,
		assets.ReadFile,
		[]string{
			"volumesnapshotclass.yaml",
			"csidriver.yaml",
			"controller_sa.yaml",
			"controller_pdb.yaml",
			"node_sa.yaml",
			"service.yaml",
			"cabundle_cm.yaml",
			"rbac/attacher_role.yaml",
			"rbac/attacher_binding.yaml",
			"rbac/privileged_role.yaml",
			"rbac/controller_privileged_binding.yaml",
			"rbac/node_privileged_binding.yaml",
			"rbac/provisioner_role.yaml",
			"rbac/provisioner_binding.yaml",
			"rbac/resizer_role.yaml",
			"rbac/resizer_binding.yaml",
			"rbac/snapshotter_role.yaml",
			"rbac/snapshotter_binding.yaml",
			"rbac/kube_rbac_proxy_role.yaml",
			"rbac/kube_rbac_proxy_binding.yaml",
			"rbac/prometheus_role.yaml",
			"rbac/prometheus_rolebinding.yaml",
		},
	).WithCSIConfigObserverController(
		"GCPFilestoreDriverCSIConfigObserverController",
		configInformers,
	).WithCSIDriverControllerService(
		"GCPFilestoreDriverControllerServiceController",
		assets.ReadFile,
		"controller.yaml",
		kubeClient,
		kubeInformersForNamespaces.InformersFor(defaultNamespace),
		configInformers,
		[]factory.Informer{
			nodeInformer.Informer(),
			infraInformer.Informer(),
			secretInformer.Informer(),
			configMapInformer.Informer(),
		},
		csidrivercontrollerservicecontroller.WithObservedProxyDeploymentHook(),
		csidrivercontrollerservicecontroller.WithCABundleDeploymentHook(
			defaultNamespace,
			trustedCAConfigMap,
			configMapInformer,
		),
		csidrivercontrollerservicecontroller.WithSecretHashAnnotationHook(
			defaultNamespace,
			secretName,
			secretInformer,
		),
		csidrivercontrollerservicecontroller.WithReplicasHook(nodeInformer.Lister()),
	).WithCSIDriverNodeService(
		"GCPFilestoreDriverNodeServiceController",
		assets.ReadFile,
		"node.yaml",
		kubeClient,
		kubeInformersForNamespaces.InformersFor(defaultNamespace),
		[]factory.Informer{configMapInformer.Informer()},
		csidrivernodeservicecontroller.WithObservedProxyDaemonSetHook(),
		csidrivernodeservicecontroller.WithCABundleDaemonSetHook(
			defaultNamespace,
			trustedCAConfigMap,
			configMapInformer,
		),
	).WithServiceMonitorController(
		"GCPFilestoreDriverServiceMonitorController",
		dynamicClient,
		assets.ReadFile,
		"servicemonitor.yaml",
	).WithStorageClassController(
		"GCPFilestoreStorageClassController",
		assets.ReadFile,
		"storageclass.yaml",
		kubeClient,
		kubeInformersForNamespaces.InformersFor(""),
		// withCustomStorageClass(kubeInformersForNamespaces.ConfigMapLister().ConfigMaps(cloudConfigNamespace)),
	)

	if err != nil {
		return err
	}

	klog.Info("Starting the informers")
	go kubeInformersForNamespaces.Start(ctx.Done())
	go dynamicInformers.Start(ctx.Done())
	go configInformers.Start(ctx.Done())

	klog.Info("Starting controllerset")
	go csiControllerSet.Run(ctx, 1)

	<-ctx.Done()

	return fmt.Errorf("stopped")
}

// withCustomStorageClass gets the network name from the openshift-config/cloud-provider-config
// ConfigMap and populate the StorageClass.Parameters.Network field with it.
// This hook will not return any error if it fails to get the network name.
func withCustomStorageClass(configMapLister corev1listers.ConfigMapNamespaceLister) csistorageclasscontroller.StorageClassHookFunc {
	return func(_ *opv1.OperatorSpec, sc *storagev1.StorageClass) error {
		cm, err := configMapLister.Get(cloudConfigName)
		if errors.IsNotFound(err) {
			klog.Warningf("Couldn't find the %s/%s ConfigMap: %v", cloudConfigNamespace, cloudConfigName, err)
			return nil
		}
		if err != nil {
			klog.Errorf("Failed to get the %s/%s ConfigMap: %v", cloudConfigNamespace, cloudConfigName, err)
			return nil
		}

		data, ok := cm.Data[cloudConfigKey]
		if !ok {
			klog.Errorf("Failed to get the key %q in ConfigMap %s/%s", cloudConfigKey, cloudConfigNamespace, cloudConfigName)
			return nil
		}

		cfg := &gce.ConfigFile{}
		if err := gcfg.FatalOnly(gcfg.ReadInto(cfg, strings.NewReader(data))); err != nil {
			klog.Errorf("Failed to parse GCE config: %v", err)
			return nil
		}

		if cfg.Global.NetworkName != "" {
			fmt.Printf("----------------------------------------%s\n", cfg.Global.NetworkName)

			// sc.Parameters["network"] = cfg.Global.NetworkName
		}
		return nil
	}
}
