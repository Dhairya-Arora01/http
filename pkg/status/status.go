// Package status contains a collection of HTTP status.
package status

// Status represents whether the http request succeeded or failed.
type Status struct {
	Code int

	// Text is the textual description of the status code.
	Text string
}

var (
	// --- 1xx Informational ---

	// Request received, please continue.
	Continue = Status{100, "Continue"}

	// Server switching to a different protocol.
	SwitchingProtocols = Status{101, "Switching Protocols"}

	// WebDAV: Server has received and is processing the request.
	Processing = Status{102, "Processing"}

	// Used for preloading resources before final response.
	EarlyHints = Status{103, "Early Hints"}

	// --- 2xx Success ---

	// Request succeeded.
	OK = Status{200, "OK"}

	// Resource created successfully.
	Created = Status{201, "Created"}

	// Request accepted for processing, but not completed.
	Accepted = Status{202, "Accepted"}

	// Response from a third-party or proxy.
	NonAuthoritativeInfo = Status{203, "Non-Authoritative Information"}

	// Request succeeded but no content to return.
	NoContent = Status{204, "No Content"}

	// Reset the document view.
	ResetContent = Status{205, "Reset Content"}

	// Partial response due to range header.
	PartialContent = Status{206, "Partial Content"}

	// WebDAV: Multiple status codes for sub-requests.
	MultiStatus = Status{207, "Multi-Status"}

	// WebDAV: Results previously reported.
	AlreadyReported = Status{208, "Already Reported"}

	// --- 3xx Redirection ---

	// Multiple options for the resource.
	MultipleChoices = Status{300, "Multiple Choices"}

	// Resource moved permanently to a new URI.
	MovedPermanently = Status{301, "Moved Permanently"}

	// Temporary redirect to another URI.
	Found = Status{302, "Found"}

	// Redirect to another URI using GET.
	SeeOther = Status{303, "See Other"}

	// Resource not modified since last request.
	NotModified = Status{304, "Not Modified"}

	// Deprecated: Use specified proxy.
	UseProxy = Status{305, "Use Proxy"}

	// Temporarily redirected to another URI.
	TemporaryRedirect = Status{307, "Temporary Redirect"}

	// Permanent redirection, method not changed.
	PermanentRedirect = Status{308, "Permanent Redirect"}

	// --- 4xx Client Errors ---

	// Malformed request syntax or invalid parameters.
	BadRequest = Status{400, "Bad Request"}

	// Authentication is required and has failed.
	Unauthorized = Status{401, "Unauthorized"}

	// Reserved for future use (e.g., digital payments).
	PaymentRequired = Status{402, "Payment Required"}

	// Server understands but refuses to authorize.
	Forbidden = Status{403, "Forbidden"}

	// Requested resource not found.
	NotFound = Status{404, "Not Found"}

	// Method not allowed for this resource.
	MethodNotAllowed = Status{405, "Method Not Allowed"}

	// Requested resource not acceptable according to Accept headers.
	NotAcceptable = Status{406, "Not Acceptable"}

	// Authentication with a proxy is required.
	ProxyAuthRequired = Status{407, "Proxy Authentication Required"}

	// Server timed out waiting for the request.
	RequestTimeout = Status{408, "Request Timeout"}

	// Request conflict with current state of the server.
	Conflict = Status{409, "Conflict"}

	// Resource is no longer available and will not be available again.
	Gone = Status{410, "Gone"}

	// Content-Length header is required.
	LengthRequired = Status{411, "Length Required"}

	// Preconditions given in request headers failed.
	PreconditionFailed = Status{412, "Precondition Failed"}

	// Request entity too large.
	PayloadTooLarge = Status{413, "Payload Too Large"}

	// Request URI too long.
	URITooLong = Status{414, "URI Too Long"}

	// Unsupported media type in the request.
	UnsupportedMediaType = Status{415, "Unsupported Media Type"}

	// Requested range not satisfiable.
	RangeNotSatisfiable = Status{416, "Range Not Satisfiable"}

	// Expectation in request headers could not be met.
	ExpectationFailed = Status{417, "Expectation Failed"}

	// April Fools: I'm a teapot.
	Teapot = Status{418, "I'm a teapot"}

	// Request was directed at a server that is not able to produce a response.
	MisdirectedRequest = Status{421, "Misdirected Request"}

	// Request was well-formed but unable to be followed due to semantic errors.
	UnprocessableEntity = Status{422, "Unprocessable Entity"}

	// Resource is locked.
	Locked = Status{423, "Locked"}

	// Failed dependency.
	FailedDependency = Status{424, "Failed Dependency"}

	// Server is unwilling to risk processing a request that might be replayed.
	TooEarly = Status{425, "Too Early"}

	// Client must upgrade to a different protocol.
	UpgradeRequired = Status{426, "Upgrade Required"}

	// Origin server requires the request to be conditional.
	PreconditionRequired = Status{428, "Precondition Required"}

	// User has sent too many requests in a given amount of time.
	TooManyRequests = Status{429, "Too Many Requests"}

	// Request header fields too large.
	RequestHeaderFieldsTooLarge = Status{431, "Request Header Fields Too Large"}

	// Resource is unavailable for legal reasons.
	UnavailableForLegalReasons = Status{451, "Unavailable For Legal Reasons"}

	// --- 5xx Server Errors ---

	// Generic internal server error.
	InternalServerError = Status{500, "Internal Server Error"}

	// Server does not support the functionality required.
	NotImplemented = Status{501, "Not Implemented"}

	// Invalid response from upstream server.
	BadGateway = Status{502, "Bad Gateway"}

	// Server is currently unavailable.
	ServiceUnavailable = Status{503, "Service Unavailable"}

	// Server did not receive a timely response from upstream.
	GatewayTimeout = Status{504, "Gateway Timeout"}

	// Server does not support HTTP version used.
	HTTPVersionNotSupported = Status{505, "HTTP Version Not Supported"}

	// Server has internal configuration error: content negotiation failed.
	VariantAlsoNegotiates = Status{506, "Variant Also Negotiates"}

	// WebDAV: Server is out of storage.
	InsufficientStorage = Status{507, "Insufficient Storage"}

	// WebDAV: Infinite loop detected.
	LoopDetected = Status{508, "Loop Detected"}

	// Further extensions to the request are required.
	NotExtended = Status{510, "Not Extended"}

	// Client needs to authenticate to gain network access.
	NetworkAuthenticationRequired = Status{511, "Network Authentication Required"}
)
