package har

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func assert(t *testing.T, action string, expected, tested interface{}) {
	if fmt.Sprintf("%T", expected) != fmt.Sprintf("%T", tested) {
		t.Errorf("Need matching types for %s. Expected type '%T' Tested type '%T'", action, expected, tested)
	}
	if tested != expected {
		t.Errorf("Expected '%v' for '%s'. Got '%v'", expected, action, tested)
	}
}

func TestRun(t *testing.T) {
	har, _ := parseHar("examples/FireFox.har")
	assert(t, "har version", "1.2", har.Version)
	if har.Browser.Name != "Firefox" && har.Browser.Name == har.Creator.Name {
		t.Errorf("Invalid browser name. Wanted FireFox got %s", har.Browser.Name)
	}
	if har.Browser.Version != "99.0" && har.Browser.Version == har.Creator.Version {
		t.Errorf("Invalid browser version. Wanted FireFox got %s", har.Browser.Version)
	}
}

func TestPage(t *testing.T) {
	har, err := parseHar("examples/FireFox.har")
	if err != nil {
		t.Errorf("Error parsing har file :%v", err)
	}
	assert(t, "length of har pages", 1, len(har.Pages))
	page := har.Pages[0]
	assert(t, "page start time", "2022-04-09T14:47:41.176-04:00", page.StartTime)
	assert(t, "page ID", "page_1", page.ID)
	assert(t, "page title", "Cyber Jake", page.Title)
	assert(t, "page timings onContentLoad", float64(2035), page.PageTimings.ContentLoad)
	assert(t, "page timings onLoad", float64(2452), page.PageTimings.Load)

}

func TestEntry(t *testing.T) {
	har, err := parseHar("examples/FireFox.har")
	if err != nil {
		t.Errorf("Error parsing har file :%v", err)
	}
	entry := har.Entries[0]
	assert(t, "entry IP", "104.22.40.104", entry.IP)
	assert(t, "entry page ID", "page_1", entry.PageID)
	assert(t, "entry port", "443", entry.Port)
	assert(t, "entry secure status", "secure", entry.Secure)
	assert(t, "entry startedTime", "2022-04-09T14:47:41.176-04:00", entry.StartedTime)
	assert(t, "entry Time", float64(444), entry.Time)
}

func TestTimingFirefox(t *testing.T) {
	har, err := parseHar("examples/FireFox.har")
	if err != nil {
		t.Errorf("Got error trying to parse firefox file: %v", err)
	}
	timings := har.Entries[0].Timing
	assert(t, "firefox timings blocked", float64(0), timings.Blocked)
	assert(t, "firefox timings DNS", float64(86), timings.DNS)
	assert(t, "firefox timings connect", float64(12), timings.Connect)
	assert(t, "firefox timings SSL", float64(69), timings.SSL)
	assert(t, "firefox timings send", float64(0), timings.Send)
	assert(t, "firefox timings wait", float64(277), timings.Wait)
	assert(t, "firefox timings receive", float64(0), timings.Receive)
}

func TestTimingsChrome(t *testing.T) {
	har, err := parseHar("examples/Chrome.har")
	if err != nil {
		t.Errorf("Got error when trying to parse chrome file: '%v'", err)
	}
	timings := har.Entries[0].Timing
	assert(t, "chrome timings blocked", 2.7729999983466698, timings.Blocked)
	assert(t, "chrome timings DNS", 78.15, timings.DNS)
	assert(t, "chrome timings SSL", 66.09899999999999, timings.SSL)
	assert(t, "chrome timings connect", 164.974, timings.Connect)
	assert(t, "chrome timings send", 0.23400000000000887, timings.Send)
	assert(t, "chrome timings wait", 396.6860000003886, timings.Wait)
	assert(t, "chrome timings receive", 1.5050000001792796, timings.Receive)
	assert(t, "chrome timings blocked queueing", 2.4559999983466696, timings.BlockedQueueing)
}

func TestRequest(t *testing.T) {
	har, _ := parseHar("examples/FireFox.har")
	request := har.Entries[0].Request

	assert(t, "request URL", "https://cyberjake.xyz/", request.URL)
	assert(t, "request cookies length", 0, len(request.Cookies))
	assert(t, "request headers length", 12, len(request.Headers))
	assert(t, "request method", "GET", request.Method)
	assert(t, "request body size", 0, request.BodySize)
	assert(t, "request http version", "HTTP/2", request.HTTPVersion)
	assert(t, "request header size", 451, request.HeaderSize)
	assert(t, "request query string length", 0, len(request.QueryString))

	_, err := request.CreateRequest()
	if err != nil {
		t.Errorf("Got error '%v' when creating request", err)
	}
}

func TestResponse(t *testing.T) {
	har, _ := parseHar("examples/FireFox.har")
	response := har.Entries[0].Response

	assert(t, "response status code", 200, response.Status)
	assert(t, "response status text", "OK", response.StatusText)
	assert(t, "response HTTP version", "HTTP/2", response.HTTPVersion)
	assert(t, "response cookies length", 0, len(response.Cookies))
	assert(t, "response headers size", 1767, response.HeaderSize)
	assert(t, "response body size", 7978, response.BodySize)
	assert(t, "response redirect URL", "", response.RedirectURL)
	assert(t, "response MIME type",  "text/html; charset=utf-8", response.Content.MIMEType)
	assert(t, "response content text length", 19841, len(response.Content.Text))
	assert(t, "response content size", 19844, response.Content.Size)
}

func TestFailure(t *testing.T) {
	_, err := parseHar("examples/NotHere.har")
	if !os.IsNotExist(err) {
		t.Errorf("Wanted error for file not found got '%v'", err)
	}
}

func TestBadRequest(t *testing.T) {
	fakeRquest := &Request{Method: ""}
	_, err := fakeRquest.CreateRequest()
	if err != nil {
		t.Error(err)
	}
}

func TestBadFile(t *testing.T) {
	_, err := parseHar("examples/BadHar.har")
	if err.Error() != "unexpected end of JSON input" {
		t.Errorf("Expected bad JSON got %s", err)
	}
}

func TestBadJSON(t *testing.T) {
	file, err := os.Open("examples/bad.json")
	if err != nil {
		t.Errorf("Error opening bad JSON file: %v", err)
	}
	defer file.Close()
	bytevalue, err := io.ReadAll(file)
	if err != nil {
		t.Errorf("Error reading bad json file: %v", err)
	}
	_, err = skiproot(bytevalue)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
