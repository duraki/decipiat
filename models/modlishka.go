package models

type ModlishkaConfig struct {
	ProxyDomain          *string `json:"proxyDomain"`
	ListeningAddress     *string `json:"listeningAddress"`
	ProxyAddress         *string `json:"proxyAddress"`
	Target               *string `json:"target"`
	TargetResources      *string `json:"targetResources"`
	Rules                *string `json:"rules"`
	TerminateTriggers    *string `json:"terminateTriggers"`
	TerminateRedirectUrl *string `json:"terminateRedirectUrl"`
	TrackingCookie       *string `json:"trackingCookie"`
	TrackingParam        *string `json:"trackingParam"`
	JsRules              *string `json:"jsRules"`
	JsReflectParam       *string `json:"jsReflectParam"`
	Debug                *bool   `json:"debug"`
	ForceHTTPS           *bool   `json:"forceHTTPS"`
	ForceHTTP            *bool   `json:"forceHTTP"`
	DynamicMode          *bool   `json:"dynamicMode"`
	LogPostOnly          *bool   `json:"logPostOnly"`
	DisableSecurity      *bool   `json:"disableSecurity"`
	Log                  *string `json:"log"`
	Plugins              *string `json:"plugins"`
	CredParams           *string `json:"credParams"`
	Cert                 *string `json:"cert"`
	CertKey              *string `json:"certKey"`
	CertPool             *string `json:"certPool"`
}
