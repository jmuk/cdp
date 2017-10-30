// Code generated by cdpgen. DO NOT EDIT.

package network

// EnableArgs represents the arguments for Enable in the Network domain.
type EnableArgs struct {
	// MaxTotalBufferSize Buffer size in bytes to use when preserving network payloads (XHRs, etc).
	//
	// Note: This property is experimental.
	MaxTotalBufferSize *int `json:"maxTotalBufferSize,omitempty"`
	// MaxResourceBufferSize Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
	//
	// Note: This property is experimental.
	MaxResourceBufferSize *int `json:"maxResourceBufferSize,omitempty"`
}

// NewEnableArgs initializes EnableArgs with the required arguments.
func NewEnableArgs() *EnableArgs {
	args := new(EnableArgs)

	return args
}

// SetMaxTotalBufferSize sets the MaxTotalBufferSize optional argument. Buffer size in bytes to use when preserving network payloads (XHRs, etc).
//
// Note: This property is experimental.
func (a *EnableArgs) SetMaxTotalBufferSize(maxTotalBufferSize int) *EnableArgs {
	a.MaxTotalBufferSize = &maxTotalBufferSize
	return a
}

// SetMaxResourceBufferSize sets the MaxResourceBufferSize optional argument. Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
//
// Note: This property is experimental.
func (a *EnableArgs) SetMaxResourceBufferSize(maxResourceBufferSize int) *EnableArgs {
	a.MaxResourceBufferSize = &maxResourceBufferSize
	return a
}

// SetUserAgentOverrideArgs represents the arguments for SetUserAgentOverride in the Network domain.
type SetUserAgentOverrideArgs struct {
	UserAgent string `json:"userAgent"` // User agent to use.
}

// NewSetUserAgentOverrideArgs initializes SetUserAgentOverrideArgs with the required arguments.
func NewSetUserAgentOverrideArgs(userAgent string) *SetUserAgentOverrideArgs {
	args := new(SetUserAgentOverrideArgs)
	args.UserAgent = userAgent
	return args
}

// SetExtraHTTPHeadersArgs represents the arguments for SetExtraHTTPHeaders in the Network domain.
type SetExtraHTTPHeadersArgs struct {
	Headers Headers `json:"headers"` // Map with extra HTTP headers.
}

// NewSetExtraHTTPHeadersArgs initializes SetExtraHTTPHeadersArgs with the required arguments.
func NewSetExtraHTTPHeadersArgs(headers Headers) *SetExtraHTTPHeadersArgs {
	args := new(SetExtraHTTPHeadersArgs)
	args.Headers = headers
	return args
}

// GetResponseBodyArgs represents the arguments for GetResponseBody in the Network domain.
type GetResponseBodyArgs struct {
	RequestID RequestID `json:"requestId"` // Identifier of the network request to get content for.
}

// NewGetResponseBodyArgs initializes GetResponseBodyArgs with the required arguments.
func NewGetResponseBodyArgs(requestID RequestID) *GetResponseBodyArgs {
	args := new(GetResponseBodyArgs)
	args.RequestID = requestID
	return args
}

// GetResponseBodyReply represents the return values for GetResponseBody in the Network domain.
type GetResponseBodyReply struct {
	Body          string `json:"body"`          // Response body.
	Base64Encoded bool   `json:"base64Encoded"` // True, if content was sent as base64.
}

// SetBlockedURLsArgs represents the arguments for SetBlockedURLs in the Network domain.
type SetBlockedURLsArgs struct {
	URLs []string `json:"urls"` // URL patterns to block. Wildcards ('*') are allowed.
}

// NewSetBlockedURLsArgs initializes SetBlockedURLsArgs with the required arguments.
func NewSetBlockedURLsArgs(urls []string) *SetBlockedURLsArgs {
	args := new(SetBlockedURLsArgs)
	args.URLs = urls
	return args
}

// ReplayXHRArgs represents the arguments for ReplayXHR in the Network domain.
type ReplayXHRArgs struct {
	RequestID RequestID `json:"requestId"` // Identifier of XHR to replay.
}

// NewReplayXHRArgs initializes ReplayXHRArgs with the required arguments.
func NewReplayXHRArgs(requestID RequestID) *ReplayXHRArgs {
	args := new(ReplayXHRArgs)
	args.RequestID = requestID
	return args
}

// CanClearBrowserCacheReply represents the return values for CanClearBrowserCache in the Network domain.
type CanClearBrowserCacheReply struct {
	Result bool `json:"result"` // True if browser cache can be cleared.
}

// CanClearBrowserCookiesReply represents the return values for CanClearBrowserCookies in the Network domain.
type CanClearBrowserCookiesReply struct {
	Result bool `json:"result"` // True if browser cookies can be cleared.
}

// GetCookiesArgs represents the arguments for GetCookies in the Network domain.
type GetCookiesArgs struct {
	URLs []string `json:"urls,omitempty"` // The list of URLs for which applicable cookies will be fetched
}

// NewGetCookiesArgs initializes GetCookiesArgs with the required arguments.
func NewGetCookiesArgs() *GetCookiesArgs {
	args := new(GetCookiesArgs)

	return args
}

// SetURLs sets the URLs optional argument. The list of URLs for which applicable cookies will be fetched
func (a *GetCookiesArgs) SetURLs(urls []string) *GetCookiesArgs {
	a.URLs = urls
	return a
}

// GetCookiesReply represents the return values for GetCookies in the Network domain.
type GetCookiesReply struct {
	Cookies []Cookie `json:"cookies"` // Array of cookie objects.
}

// GetAllCookiesReply represents the return values for GetAllCookies in the Network domain.
type GetAllCookiesReply struct {
	Cookies []Cookie `json:"cookies"` // Array of cookie objects.
}

// DeleteCookiesArgs represents the arguments for DeleteCookies in the Network domain.
type DeleteCookiesArgs struct {
	Name   string  `json:"name"`             // Name of the cookies to remove.
	URL    *string `json:"url,omitempty"`    // If specified, deletes all the cookies with the given name where domain and path match provided URL.
	Domain *string `json:"domain,omitempty"` // If specified, deletes only cookies with the exact domain.
	Path   *string `json:"path,omitempty"`   // If specified, deletes only cookies with the exact path.
}

// NewDeleteCookiesArgs initializes DeleteCookiesArgs with the required arguments.
func NewDeleteCookiesArgs(name string) *DeleteCookiesArgs {
	args := new(DeleteCookiesArgs)
	args.Name = name
	return args
}

// SetURL sets the URL optional argument. If specified, deletes all the cookies with the given name where domain and path match provided URL.
func (a *DeleteCookiesArgs) SetURL(url string) *DeleteCookiesArgs {
	a.URL = &url
	return a
}

// SetDomain sets the Domain optional argument. If specified, deletes only cookies with the exact domain.
func (a *DeleteCookiesArgs) SetDomain(domain string) *DeleteCookiesArgs {
	a.Domain = &domain
	return a
}

// SetPath sets the Path optional argument. If specified, deletes only cookies with the exact path.
func (a *DeleteCookiesArgs) SetPath(path string) *DeleteCookiesArgs {
	a.Path = &path
	return a
}

// SetCookieArgs represents the arguments for SetCookie in the Network domain.
type SetCookieArgs struct {
	Name     string         `json:"name"`               // Cookie name.
	Value    string         `json:"value"`              // Cookie value.
	URL      *string        `json:"url,omitempty"`      // The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
	Domain   *string        `json:"domain,omitempty"`   // Cookie domain.
	Path     *string        `json:"path,omitempty"`     // Cookie path.
	Secure   *bool          `json:"secure,omitempty"`   // True if cookie is secure.
	HTTPOnly *bool          `json:"httpOnly,omitempty"` // True if cookie is http-only.
	SameSite CookieSameSite `json:"sameSite,omitempty"` // Cookie SameSite type.
	Expires  TimeSinceEpoch `json:"expires,omitempty"`  // Cookie expiration date, session cookie if not set
}

// NewSetCookieArgs initializes SetCookieArgs with the required arguments.
func NewSetCookieArgs(name string, value string) *SetCookieArgs {
	args := new(SetCookieArgs)
	args.Name = name
	args.Value = value
	return args
}

// SetURL sets the URL optional argument. The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
func (a *SetCookieArgs) SetURL(url string) *SetCookieArgs {
	a.URL = &url
	return a
}

// SetDomain sets the Domain optional argument. Cookie domain.
func (a *SetCookieArgs) SetDomain(domain string) *SetCookieArgs {
	a.Domain = &domain
	return a
}

// SetPath sets the Path optional argument. Cookie path.
func (a *SetCookieArgs) SetPath(path string) *SetCookieArgs {
	a.Path = &path
	return a
}

// SetSecure sets the Secure optional argument. True if cookie is secure.
func (a *SetCookieArgs) SetSecure(secure bool) *SetCookieArgs {
	a.Secure = &secure
	return a
}

// SetHTTPOnly sets the HTTPOnly optional argument. True if cookie is http-only.
func (a *SetCookieArgs) SetHTTPOnly(httpOnly bool) *SetCookieArgs {
	a.HTTPOnly = &httpOnly
	return a
}

// SetSameSite sets the SameSite optional argument. Cookie SameSite type.
func (a *SetCookieArgs) SetSameSite(sameSite CookieSameSite) *SetCookieArgs {
	a.SameSite = sameSite
	return a
}

// SetExpires sets the Expires optional argument. Cookie expiration date, session cookie if not set
func (a *SetCookieArgs) SetExpires(expires TimeSinceEpoch) *SetCookieArgs {
	a.Expires = expires
	return a
}

// SetCookieReply represents the return values for SetCookie in the Network domain.
type SetCookieReply struct {
	Success bool `json:"success"` // True if successfully set cookie.
}

// SetCookiesArgs represents the arguments for SetCookies in the Network domain.
type SetCookiesArgs struct {
	Cookies []CookieParam `json:"cookies"` // Cookies to be set.
}

// NewSetCookiesArgs initializes SetCookiesArgs with the required arguments.
func NewSetCookiesArgs(cookies []CookieParam) *SetCookiesArgs {
	args := new(SetCookiesArgs)
	args.Cookies = cookies
	return args
}

// CanEmulateNetworkConditionsReply represents the return values for CanEmulateNetworkConditions in the Network domain.
type CanEmulateNetworkConditionsReply struct {
	Result bool `json:"result"` // True if emulation of network conditions is supported.
}

// EmulateNetworkConditionsArgs represents the arguments for EmulateNetworkConditions in the Network domain.
type EmulateNetworkConditionsArgs struct {
	Offline            bool           `json:"offline"`                  // True to emulate internet disconnection.
	Latency            float64        `json:"latency"`                  // Minimum latency from request sent to response headers received (ms).
	DownloadThroughput float64        `json:"downloadThroughput"`       // Maximal aggregated download throughput (bytes/sec). -1 disables download throttling.
	UploadThroughput   float64        `json:"uploadThroughput"`         // Maximal aggregated upload throughput (bytes/sec).  -1 disables upload throttling.
	ConnectionType     ConnectionType `json:"connectionType,omitempty"` // Connection type if known.
}

// NewEmulateNetworkConditionsArgs initializes EmulateNetworkConditionsArgs with the required arguments.
func NewEmulateNetworkConditionsArgs(offline bool, latency float64, downloadThroughput float64, uploadThroughput float64) *EmulateNetworkConditionsArgs {
	args := new(EmulateNetworkConditionsArgs)
	args.Offline = offline
	args.Latency = latency
	args.DownloadThroughput = downloadThroughput
	args.UploadThroughput = uploadThroughput
	return args
}

// SetConnectionType sets the ConnectionType optional argument. Connection type if known.
func (a *EmulateNetworkConditionsArgs) SetConnectionType(connectionType ConnectionType) *EmulateNetworkConditionsArgs {
	a.ConnectionType = connectionType
	return a
}

// SetCacheDisabledArgs represents the arguments for SetCacheDisabled in the Network domain.
type SetCacheDisabledArgs struct {
	CacheDisabled bool `json:"cacheDisabled"` // Cache disabled state.
}

// NewSetCacheDisabledArgs initializes SetCacheDisabledArgs with the required arguments.
func NewSetCacheDisabledArgs(cacheDisabled bool) *SetCacheDisabledArgs {
	args := new(SetCacheDisabledArgs)
	args.CacheDisabled = cacheDisabled
	return args
}

// SetBypassServiceWorkerArgs represents the arguments for SetBypassServiceWorker in the Network domain.
type SetBypassServiceWorkerArgs struct {
	Bypass bool `json:"bypass"` // Bypass service worker and load from network.
}

// NewSetBypassServiceWorkerArgs initializes SetBypassServiceWorkerArgs with the required arguments.
func NewSetBypassServiceWorkerArgs(bypass bool) *SetBypassServiceWorkerArgs {
	args := new(SetBypassServiceWorkerArgs)
	args.Bypass = bypass
	return args
}

// SetDataSizeLimitsForTestArgs represents the arguments for SetDataSizeLimitsForTest in the Network domain.
type SetDataSizeLimitsForTestArgs struct {
	MaxTotalSize    int `json:"maxTotalSize"`    // Maximum total buffer size.
	MaxResourceSize int `json:"maxResourceSize"` // Maximum per-resource size.
}

// NewSetDataSizeLimitsForTestArgs initializes SetDataSizeLimitsForTestArgs with the required arguments.
func NewSetDataSizeLimitsForTestArgs(maxTotalSize int, maxResourceSize int) *SetDataSizeLimitsForTestArgs {
	args := new(SetDataSizeLimitsForTestArgs)
	args.MaxTotalSize = maxTotalSize
	args.MaxResourceSize = maxResourceSize
	return args
}

// GetCertificateArgs represents the arguments for GetCertificate in the Network domain.
type GetCertificateArgs struct {
	Origin string `json:"origin"` // Origin to get certificate for.
}

// NewGetCertificateArgs initializes GetCertificateArgs with the required arguments.
func NewGetCertificateArgs(origin string) *GetCertificateArgs {
	args := new(GetCertificateArgs)
	args.Origin = origin
	return args
}

// GetCertificateReply represents the return values for GetCertificate in the Network domain.
type GetCertificateReply struct {
	TableNames []string `json:"tableNames"` // No description.
}

// SetRequestInterceptionArgs represents the arguments for SetRequestInterception in the Network domain.
type SetRequestInterceptionArgs struct {
	Patterns []RequestPattern `json:"patterns"` // Requests matching any of these patterns will be forwarded and wait for the corresponding continueInterceptedRequest call.
}

// NewSetRequestInterceptionArgs initializes SetRequestInterceptionArgs with the required arguments.
func NewSetRequestInterceptionArgs(patterns []RequestPattern) *SetRequestInterceptionArgs {
	args := new(SetRequestInterceptionArgs)
	args.Patterns = patterns
	return args
}

// ContinueInterceptedRequestArgs represents the arguments for ContinueInterceptedRequest in the Network domain.
type ContinueInterceptedRequestArgs struct {
	InterceptionID        InterceptionID         `json:"interceptionId"`                  // No description.
	ErrorReason           ErrorReason            `json:"errorReason,omitempty"`           // If set this causes the request to fail with the given reason. Passing Aborted for requests marked with isNavigationRequest also cancels the navigation. Must not be set in response to an authChallenge.
	RawResponse           *string                `json:"rawResponse,omitempty"`           // If set the requests completes using with the provided base64 encoded raw response, including HTTP status line and headers etc... Must not be set in response to an authChallenge.
	URL                   *string                `json:"url,omitempty"`                   // If set the request url will be modified in a way that's not observable by page. Must not be set in response to an authChallenge.
	Method                *string                `json:"method,omitempty"`                // If set this allows the request method to be overridden. Must not be set in response to an authChallenge.
	PostData              *string                `json:"postData,omitempty"`              // If set this allows postData to be set. Must not be set in response to an authChallenge.
	Headers               Headers                `json:"headers,omitempty"`               // If set this allows the request headers to be changed. Must not be set in response to an authChallenge.
	AuthChallengeResponse *AuthChallengeResponse `json:"authChallengeResponse,omitempty"` // Response to a requestIntercepted with an authChallenge. Must not be set otherwise.
}

// NewContinueInterceptedRequestArgs initializes ContinueInterceptedRequestArgs with the required arguments.
func NewContinueInterceptedRequestArgs(interceptionID InterceptionID) *ContinueInterceptedRequestArgs {
	args := new(ContinueInterceptedRequestArgs)
	args.InterceptionID = interceptionID
	return args
}

// SetErrorReason sets the ErrorReason optional argument. If set this causes the request to fail with the given reason. Passing Aborted for requests marked with isNavigationRequest also cancels the navigation. Must not be set in response to an authChallenge.
func (a *ContinueInterceptedRequestArgs) SetErrorReason(errorReason ErrorReason) *ContinueInterceptedRequestArgs {
	a.ErrorReason = errorReason
	return a
}

// SetRawResponse sets the RawResponse optional argument. If set the requests completes using with the provided base64 encoded raw response, including HTTP status line and headers etc... Must not be set in response to an authChallenge.
func (a *ContinueInterceptedRequestArgs) SetRawResponse(rawResponse string) *ContinueInterceptedRequestArgs {
	a.RawResponse = &rawResponse
	return a
}

// SetURL sets the URL optional argument. If set the request url will be modified in a way that's not observable by page. Must not be set in response to an authChallenge.
func (a *ContinueInterceptedRequestArgs) SetURL(url string) *ContinueInterceptedRequestArgs {
	a.URL = &url
	return a
}

// SetMethod sets the Method optional argument. If set this allows the request method to be overridden. Must not be set in response to an authChallenge.
func (a *ContinueInterceptedRequestArgs) SetMethod(method string) *ContinueInterceptedRequestArgs {
	a.Method = &method
	return a
}

// SetPostData sets the PostData optional argument. If set this allows postData to be set. Must not be set in response to an authChallenge.
func (a *ContinueInterceptedRequestArgs) SetPostData(postData string) *ContinueInterceptedRequestArgs {
	a.PostData = &postData
	return a
}

// SetHeaders sets the Headers optional argument. If set this allows the request headers to be changed. Must not be set in response to an authChallenge.
func (a *ContinueInterceptedRequestArgs) SetHeaders(headers Headers) *ContinueInterceptedRequestArgs {
	a.Headers = headers
	return a
}

// SetAuthChallengeResponse sets the AuthChallengeResponse optional argument. Response to a requestIntercepted with an authChallenge. Must not be set otherwise.
func (a *ContinueInterceptedRequestArgs) SetAuthChallengeResponse(authChallengeResponse AuthChallengeResponse) *ContinueInterceptedRequestArgs {
	a.AuthChallengeResponse = &authChallengeResponse
	return a
}
