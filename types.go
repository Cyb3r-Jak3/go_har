package main

type HarFile struct {
	Version string           `json:"version"`
	Creator Creator_struct   `json:"creator"`
	Browser Browser_struct   `json:"browser"`
	Pages   []Page           `json:"pages"`
	Entries []Entries_struct `json:"entries"`
}

type Creator_struct struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Browser_struct struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Page struct {
	StartTime   string             `json:"startedDateTime"`
	ID          string             `json:"id"`
	Title       string             `json:"title"`
	PagerTiming PageTimings_struct `json:"pageTimings"`
}

type PageTimings_struct struct {
	ContentLoad uint16 `json:"onContentLoad"`
	Load        uint16 `json:"onLoad"`
}

type NameValuePair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Request_struct struct {
	BodySize    uint16          `json:"bodySize"`
	Cookies     []NameValuePair `json:"cookies"`
	Headers     []NameValuePair `json:"headers"`
	HeaderSize  uint16          `json:"headersSize"`
	HTTPVersion string          `json:"httpVersion"`
	Method      string          `json:"method"`
	QueryString []NameValuePair `json:"queryString"`
	URL         string          `json:"url"`
}

type Content_struct struct {
	MIMEType string `json:"mimeType"`
	Size     uint64 `json:"size"`
	Text     string `json:"text"`
}

type Response_struct struct {
	BodySize    uint64          `json:"bodySize"`
	Content     Content_struct  `json:"content"`
	Cookies     []NameValuePair `json:"cookies"`
	Headers     []NameValuePair `json:"headers"`
	Status      uint16          `json:"status"`
	StatusText  string          `json:"statusText"`
	HTTPVersion string          `json:"httpVersion"`
	RedirectURL string          `json:"redirectURL"`
	HeaderSize  uint16          `json:"headersSize"`
}

type Timings struct {
	Blocked uint16 `json:"blocked"`
	DNS     uint16 `json:"dns"`
	Connect uint16 `json:"connect"`
	SSL     uint16 `json:"ssl"`
	Send    uint16 `json:"send"`
	Wait    uint16 `json:"wait"`
	Receive uint16 `json:"receive"`
}

type Entries_struct struct {
	PageRef     string          `json:"pageref"`
	StartedTime string          `json:"startedDateTime"`
	Time        uint16          `json:"time"`
	Secure      string          `json:"_securityState"`
	IP          string          `json:"serverIPAddress"`
	Port        string          `json:"connection"`
	Request     Request_struct  `json:"request"`
	Response    Response_struct `json:"response"`
	Timing      Timings         `json:"timings"`
}
