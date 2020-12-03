package models

// ModlishkaConfig struct will handle parameters for modlishka json config file
type ModlishkaConfig struct {
	/** @type {String} Proxy domain name that will be used (ie. campaign.com) */
	ProxyDomain *string `json:"proxyDomain"`

	/** @type {String} Listening address (ie. 0.0.0.0) */
	ListeningAddress *string `json:"listeningAddress"`

	/** @type {String} Proxy that should be used (socks/https/http) (ie. http://0.0.0.1:8080) */
	ProxyAddress *string `json:"proxyAddress"`

	/** @type {String} Target domain name (ie. nsa.gov) */
	Target *string `json:"target"`

	/** @type {String} Comma separated list of domains that were not translated
	automatically. Use this to force domain translation - (ie. static.nsa.gov) */
	TargetResources *string `json:"targetResources"`

	/** @type {String} Comma separated list of 'string' patterns and their
	replacements - (ie. base64(new):base64(old)) */
	Rules *string `json:"rules"`

	/** @type {string} Session termination: Comma separated list of URLs from target's
	origin which will trigger session termination */
	TerminateTriggers *string `json:"terminateTriggers"`

	/** @type {String} URL to which a client will be redirected after Session
	Termination rules apply */
	TerminateRedirectURL *string `json:"terminateRedirectUrl"`

	/** @type {String} Name of the HTTP cookie used to track the client */
	TrackingCookie *string `json:"trackingCookie"`

	/** @type {String} Name of the HTTP parameter used to track the client */
	TrackingParam *string `json:"trackingParam"`

	/** @type {String} Comma separated list of URL patterns and JS base64 encoded
	payloads that will be injected - (ie. nsa.gov:base64(alert(1))) */
	JsRules *string `json:"jsRules"`

	/** @type {String} todo ++ */
	JsReflectParam *string `json:"jsReflectParam"`

	/** @type {Bool} Print debug logs */
	Debug *bool `json:"debug"`

	/** @type {Bool} Strip all clear-text from the traffic and proxy through HTTPS only */
	ForceHTTPS *bool `json:"forceHTTPS"`

	/** @type {Bool} Strip all TLS from the traffic and proxy through HTTP only */
	ForceHTTP *bool `json:"forceHTTP"`

	/** @type {Bool} Enable dynamic mode for 'Client Domain Hooking' */
	DynamicMode *bool `json:"dynamicMode"`

	/** @type {Bool} todo ++ */
	LogPostOnly *bool `json:"logPostOnly"`

	/** @type {Bool} todo ++ */
	DisableSecurity *bool `json:"disableSecurity"`

	/** @type {String} Local file to which fetched requests will be written (appended) */
	Log *string `json:"log"`

	/** @type {String} Comma seperated list of enabled plugin names */
	Plugins *string `json:"plugins"`

	/** @type {String} Credential regexp with matching groups */
	CredParams *string `json:"credParams"`

	/**
	 * Certificate extension
	 * @type {+}
	 */
	Cert     *string `json:"cert"`
	CertKey  *string `json:"certKey"`
	CertPool *string `json:"certPool"`
}
