package har

// File represents a HarFile with various properties
type File struct {
	Version string          `json:"version"`
	Creator CreatorStruct   `json:"creator"`
	Browser BrowserStruct   `json:"browser"`
	Pages   []Page          `json:"pages"`
	Entries []EntriesStruct `json:"entries"`
}

// CreatorStruct represents a the section for HarFile creator
type CreatorStruct struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// BrowserStruct represents a the section for HarFile browser
type BrowserStruct struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Page represents a page in a HarFile
type Page struct {
	StartTime   string            `json:"startedDateTime"`
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	PagerTiming PageTimingsStruct `json:"pageTimings"`
}

// PageTimingsStruct represents the timings to load a page
type PageTimingsStruct struct {
	ContentLoad uint16 `json:"onContentLoad"`
	Load        uint16 `json:"onLoad"`
}

// NameValuePair represents a named value pair. Used for Cookies and Headers
type NameValuePair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// RequestStruct represents a request in an Entry
type RequestStruct struct {
	BodySize    uint16          `json:"bodySize"`
	Cookies     []NameValuePair `json:"cookies"`
	Headers     []NameValuePair `json:"headers"`
	HeaderSize  uint16          `json:"headersSize"`
	HTTPVersion string          `json:"httpVersion"`
	Method      string          `json:"method"`
	QueryString []NameValuePair `json:"queryString"`
	URL         string          `json:"url"`
}

// ContentStruct represents the returned content in a response
type ContentStruct struct {
	MIMEType string `json:"mimeType"`
	Size     uint64 `json:"size"`
	Text     string `json:"text"`
}

// ResponseStruct represents a response in a har entry
type ResponseStruct struct {
	BodySize    uint64          `json:"bodySize"`
	Content     ContentStruct   `json:"content"`
	Cookies     []NameValuePair `json:"cookies"`
	Headers     []NameValuePair `json:"headers"`
	Status      uint16          `json:"status"`
	StatusText  string          `json:"statusText"`
	HTTPVersion string          `json:"httpVersion"`
	RedirectURL string          `json:"redirectURL"`
	HeaderSize  uint16          `json:"headersSize"`
}

// Timings represent the timing for a har page
type Timings struct {
	Blocked uint16 `json:"blocked"`
	DNS     uint16 `json:"dns"`
	Connect uint16 `json:"connect"`
	SSL     uint16 `json:"ssl"`
	Send    uint16 `json:"send"`
	Wait    uint16 `json:"wait"`
	Receive uint16 `json:"receive"`
}

// EntriesStruct represents an entry in a Page
type EntriesStruct struct {
	PageRef     string         `json:"pageref"`
	StartedTime string         `json:"startedDateTime"`
	Time        uint16         `json:"time"`
	Secure      string         `json:"_securityState"`
	IP          string         `json:"serverIPAddress"`
	Port        string         `json:"connection"`
	Request     RequestStruct  `json:"request"`
	Response    ResponseStruct `json:"response"`
	Timing      Timings        `json:"timings"`
}
