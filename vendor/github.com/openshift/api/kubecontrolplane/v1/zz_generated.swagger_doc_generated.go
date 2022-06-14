package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_AggregatorConfig = map[string]string{
	"":                "AggregatorConfig holds information required to make the aggregator function.",
	"proxyClientInfo": "proxyClientInfo specifies the client cert/key to use when proxying to aggregated API servers",
}

func (AggregatorConfig) SwaggerDoc() map[string]string {
	return map_AggregatorConfig
}

var map_KubeAPIServerConfig = map[string]string{
	"":                             "Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"authConfig":                   "authConfig configures authentication options in addition to the standard oauth token and client certificate authenticators",
	"aggregatorConfig":             "aggregatorConfig has options for configuring the aggregator component of the API server.",
	"kubeletClientInfo":            "kubeletClientInfo contains information about how to connect to kubelets",
	"servicesSubnet":               "servicesSubnet is the subnet to use for assigning service IPs",
	"servicesNodePortRange":        "servicesNodePortRange is the range to use for assigning service public ports on a host.",
	"consolePublicURL":             "consolePublicURL is an optional URL to provide a redirect from the kube-apiserver to the webconsole",
	"userAgentMatchingConfig":      "UserAgentMatchingConfig controls how API calls from *voluntarily* identifying clients will be handled.  THIS DOES NOT DEFEND AGAINST MALICIOUS CLIENTS!",
	"imagePolicyConfig":            "imagePolicyConfig feeds the image policy admission plugin",
	"projectConfig":                "projectConfig feeds an admission plugin",
	"serviceAccountPublicKeyFiles": "serviceAccountPublicKeyFiles is a list of files, each containing a PEM-encoded public RSA key. (If any file contains a private key, the public portion of the key is used) The list of public keys is used to verify presented service account tokens. Each key is tried in order until the list is exhausted or verification succeeds. If no keys are specified, no service account authentication will be available.",
	"oauthConfig":                  "oauthConfig, if present start the /oauth endpoint in this process",
}

func (KubeAPIServerConfig) SwaggerDoc() map[string]string {
	return map_KubeAPIServerConfig
}

var map_KubeAPIServerImagePolicyConfig = map[string]string{
	"internalRegistryHostname":  "internalRegistryHostname sets the hostname for the default internal image registry. The value must be in \"hostname[:port]\" format. For backward compatibility, users can still use OPENSHIFT_DEFAULT_REGISTRY environment variable but this setting overrides the environment variable.",
	"externalRegistryHostnames": "externalRegistryHostnames provides the hostnames for the default external image registry. The external hostname should be set only when the image registry is exposed externally. The first value is used in 'publicDockerImageRepository' field in ImageStreams. The value must be in \"hostname[:port]\" format.",
}

func (KubeAPIServerImagePolicyConfig) SwaggerDoc() map[string]string {
	return map_KubeAPIServerImagePolicyConfig
}

var map_KubeAPIServerProjectConfig = map[string]string{
	"defaultNodeSelector": "defaultNodeSelector holds default project node label selector",
}

func (KubeAPIServerProjectConfig) SwaggerDoc() map[string]string {
	return map_KubeAPIServerProjectConfig
}

var map_KubeControllerManagerConfig = map[string]string{
	"":                   "Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"serviceServingCert": "serviceServingCert provides support for the old alpha service serving cert signer CA bundle",
	"projectConfig":      "projectConfig is an optimization for the daemonset controller",
	"extendedArguments":  "extendedArguments is used to configure the kube-controller-manager",
}

func (KubeControllerManagerConfig) SwaggerDoc() map[string]string {
	return map_KubeControllerManagerConfig
}

var map_KubeControllerManagerProjectConfig = map[string]string{
	"defaultNodeSelector": "defaultNodeSelector holds default project node label selector",
}

func (KubeControllerManagerProjectConfig) SwaggerDoc() map[string]string {
	return map_KubeControllerManagerProjectConfig
}

var map_KubeletConnectionInfo = map[string]string{
	"":     "KubeletConnectionInfo holds information necessary for connecting to a kubelet",
	"port": "port is the port to connect to kubelets on",
	"ca":   "ca is the CA for verifying TLS connections to kubelets",
}

func (KubeletConnectionInfo) SwaggerDoc() map[string]string {
	return map_KubeletConnectionInfo
}

var map_MasterAuthConfig = map[string]string{
	"":                           "MasterAuthConfig configures authentication options in addition to the standard oauth token and client certificate authenticators",
	"requestHeader":              "requestHeader holds options for setting up a front proxy against the API.  It is optional.",
	"webhookTokenAuthenticators": "webhookTokenAuthenticators, if present configures remote token reviewers",
	"oauthMetadataFile":          "oauthMetadataFile is a path to a file containing the discovery endpoint for OAuth 2.0 Authorization Server Metadata for an external OAuth server. See IETF Draft: // https://tools.ietf.org/html/draft-ietf-oauth-discovery-04#section-2 This option is mutually exclusive with OAuthConfig",
}

func (MasterAuthConfig) SwaggerDoc() map[string]string {
	return map_MasterAuthConfig
}

var map_RequestHeaderAuthenticationOptions = map[string]string{
	"":                    "RequestHeaderAuthenticationOptions provides options for setting up a front proxy against the entire API instead of against the /oauth endpoint.",
	"clientCA":            "clientCA is a file with the trusted signer certs.  It is required.",
	"clientCommonNames":   "clientCommonNames is a required list of common names to require a match from.",
	"usernameHeaders":     "usernameHeaders is the list of headers to check for user information.  First hit wins.",
	"groupHeaders":        "groupHeaders is the set of headers to check for group information.  All are unioned.",
	"extraHeaderPrefixes": "extraHeaderPrefixes is the set of request header prefixes to inspect for user extra. X-Remote-Extra- is suggested.",
}

func (RequestHeaderAuthenticationOptions) SwaggerDoc() map[string]string {
	return map_RequestHeaderAuthenticationOptions
}

var map_ServiceServingCert = map[string]string{
	"":         "ServiceServingCert holds configuration for service serving cert signer which creates cert/key pairs for pods fulfilling a service to serve with.",
	"certFile": "CertFile is a file containing a PEM-encoded certificate",
}

func (ServiceServingCert) SwaggerDoc() map[string]string {
	return map_ServiceServingCert
}

var map_UserAgentDenyRule = map[string]string{
	"":                 "UserAgentDenyRule adds a rejection message that can be used to help a user figure out how to get an approved client",
	"rejectionMessage": "RejectionMessage is the message shown when rejecting a client.  If it is not a set, the default message is used.",
}

func (UserAgentDenyRule) SwaggerDoc() map[string]string {
	return map_UserAgentDenyRule
}

var map_UserAgentMatchRule = map[string]string{
	"":          "UserAgentMatchRule describes how to match a given request based on User-Agent and HTTPVerb",
	"regex":     "regex is a regex that is checked against the User-Agent. Known variants of oc clients 1. oc accessing kube resources: oc/v1.2.0 (linux/amd64) kubernetes/bc4550d 2. oc accessing openshift resources: oc/v1.1.3 (linux/amd64) openshift/b348c2f 3. openshift kubectl accessing kube resources:  openshift/v1.2.0 (linux/amd64) kubernetes/bc4550d 4. openshift kubectl accessing openshift resources: openshift/v1.1.3 (linux/amd64) openshift/b348c2f 5. oadm accessing kube resources: oadm/v1.2.0 (linux/amd64) kubernetes/bc4550d 6. oadm accessing openshift resources: oadm/v1.1.3 (linux/amd64) openshift/b348c2f 7. openshift cli accessing kube resources: openshift/v1.2.0 (linux/amd64) kubernetes/bc4550d 8. openshift cli accessing openshift resources: openshift/v1.1.3 (linux/amd64) openshift/b348c2f",
	"httpVerbs": "httpVerbs specifies which HTTP verbs should be matched.  An empty list means \"match all verbs\".",
}

func (UserAgentMatchRule) SwaggerDoc() map[string]string {
	return map_UserAgentMatchRule
}

var map_UserAgentMatchingConfig = map[string]string{
	"":                        "UserAgentMatchingConfig controls how API calls from *voluntarily* identifying clients will be handled.  THIS DOES NOT DEFEND AGAINST MALICIOUS CLIENTS!",
	"requiredClients":         "requiredClients if this list is non-empty, then a User-Agent must match one of the UserAgentRegexes to be allowed",
	"deniedClients":           "deniedClients if this list is non-empty, then a User-Agent must not match any of the UserAgentRegexes",
	"defaultRejectionMessage": "defaultRejectionMessage is the message shown when rejecting a client.  If it is not a set, a generic message is given.",
}

func (UserAgentMatchingConfig) SwaggerDoc() map[string]string {
	return map_UserAgentMatchingConfig
}

var map_WebhookTokenAuthenticator = map[string]string{
	"":           "WebhookTokenAuthenticators holds the necessary configuation options for external token authenticators",
	"configFile": "configFile is a path to a Kubeconfig file with the webhook configuration",
	"cacheTTL":   "cacheTTL indicates how long an authentication result should be cached. It takes a valid time duration string (e.g. \"5m\"). If empty, you get a default timeout of 2 minutes. If zero (e.g. \"0m\"), caching is disabled",
}

func (WebhookTokenAuthenticator) SwaggerDoc() map[string]string {
	return map_WebhookTokenAuthenticator
}

// AUTO-GENERATED FUNCTIONS END HERE
