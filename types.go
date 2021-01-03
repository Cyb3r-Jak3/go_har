package har

// File represents a HarFile with various properties
type File struct {
	// Version represents the version of HAR
	Version string `json:"version"`
	// Creator represents the creator of the archive
	Creator NameVersionPair `json:"creator"`
	// Browser represents the Browser that the HAR file was created for
	Browser NameVersionPair `json:"browser"`
	// Pages represents the pages in the HAR file
	Pages []Page `json:"pages"`
	// Entries represents the entries in the HAR file
	Entries []Entry `json:"entries"`
}

// NameVersionPair represents a JSON object that has a name and a version. Used for Creator and Browser
type NameVersionPair struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// PageTimingsStruct represents the timings to load a page
type PageTimingsStruct struct {
	ContentLoad float64 `json:"onContentLoad"`
	Load        float64 `json:"onLoad"`
}

// NameValuePair represents a named value pair. Used for Cookies and Headers
type NameValuePair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// PostData represents postdata for a RequestStruct
type PostData struct {
	MineType string          `json:"mineType"`
	Params   []NameValuePair `json:"params"`
	Text     string          `json:"text"`
}

// Page represents a page in a HarFile
type Page struct {
	// StartTime represents when the first request for the page was made
	StartTime string `json:"startedDateTime"`
	// ID represents the page id
	ID string `json:"id"`
	// Title represents the page title
	Title string `json:"title"`
	// PageTimings represents the timings for the page
	PageTimings PageTimingsStruct `json:"pageTimings"`
}

// RequestStruct represents a request in an Entry
type RequestStruct struct {
	// BodySize represents the size in bits of the request body
	BodySize uint16 `json:"bodySize"`
	// Cookies represent the cookies send in the request.
	Cookies []NameValuePair `json:"cookies"`
	// Headers represents the headers send in the request
	Headers []NameValuePair `json:"headers"`
	// HeaderSize represent the size in bites of the headers
	HeaderSize uint16 `json:"headersSize"`
	// HTTPVersion represents the version of HTTP used in the request
	HTTPVersion string `json:"httpVersion"`
	// Method represents the HTTP method of the request
	Method string `json:"method"`
	//PostData represents the post data of the request
	PostData
	// QueryString represents any query string that was sent in the request
	QueryString []NameValuePair `json:"queryString"`
	// URL represents the URL of the request
	URL string `json:"url"`
}

// ContentStruct represents the returned content in a response
type ContentStruct struct {
	// MIMETYPE represents the MIMEType in the response
	MIMEType string `json:"mimeType"`
	// Size represents the size in bites of the response content
	Size uint64 `json:"size"`
	// Text represents the raw text in the response
	Text string `json:"text"`
}

// ResponseStruct represents a response in a har entry
type ResponseStruct struct {
	// BodySize represents the size in bits of the response body
	BodySize uint64 `json:"bodySize"`
	// Content represents the content of the response
	Content ContentStruct `json:"content"`
	// Cookies represents the cookies in the response
	Cookies []NameValuePair `json:"cookies"`
	// Headers represents the headers of the response
	Headers []NameValuePair `json:"headers"`
	// Status represents the status code of the response
	Status uint16 `json:"status"`
	// StatusText represents the status test of the response
	StatusText string `json:"statusText"`
	// HTTPVersion represents the HTTP version in the response
	HTTPVersion string `json:"httpVersion"`
	// RedirectURL represents the returned redirect URL if any
	RedirectURL string `json:"redirectURL"`
	// HeaderSize represents the size in bits of the response headers
	HeaderSize uint16 `json:"headersSize"`
}

// Timings represent the timing for a har page
type Timings struct {
	Blocked float64 `json:"blocked"`
	DNS     float64 `json:"dns"`
	Connect float64 `json:"connect"`
	SSL     float64 `json:"ssl"`
	Send    float64 `json:"send"`
	Wait    float64 `json:"wait"`
	Receive float64 `json:"receive"`
}

// Entry represents an entry in a Page
type Entry struct {
	// PageID represents the page ID of the entry
	PageID string `json:"pageref"`
	// StartedTime represents the start time of the entry
	StartedTime string `json:"startedDateTime"`
	// Time represents the time taken to complete the entry
	Time uint16 `json:"time"`
	// Secure represents if the request was completed securely
	Secure string `json:"_securityState"`
	// IP represents the server IP address
	IP string `json:"serverIPAddress"`
	// Port represents the port on the server that the connection was made to
	Port string `json:"connection"`
	// Request represents the request of the entry
	Request RequestStruct `json:"request"`
	// Response represents the response of the entry
	Response ResponseStruct `json:"response"`
	// Timing represents the times of the entry load
	Timing Timings `json:"timings"`
}
