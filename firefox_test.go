package har

import "testing"

const firefoxHARPath = "examples/Firefox.har"

func TestFirefoxRun(t *testing.T) {
	har := prepHarFile(t, firefoxHARPath)

	assert(t, "firefox har version", "1.2", har.Version)
	assert(t, "firefox har creator name", "Firefox", har.Browser.Name)
	assert(t, "firefox har creator version", "99.0", har.Creator.Version)
	assert(t, "firefox match har browser and creator name", har.Browser.Name, har.Creator.Name)
	assert(t, "firefox match har browser and creator version", har.Browser.Version, har.Creator.Version)
}

func TestTimingFirefox(t *testing.T) {
	har := prepHarFile(t, firefoxHARPath)

	timings := har.Entries[0].Timing
	assert(t, "firefox timings blocked", float64(0), timings.Blocked)
	assert(t, "firefox timings DNS", float64(86), timings.DNS)
	assert(t, "firefox timings connect", float64(12), timings.Connect)
	assert(t, "firefox timings SSL", float64(69), timings.SSL)
	assert(t, "firefox timings send", float64(0), timings.Send)
	assert(t, "firefox timings wait", float64(277), timings.Wait)
	assert(t, "firefox timings receive", float64(0), timings.Receive)
}

func TestFirefoxPage(t *testing.T) {
	har := prepHarFile(t, firefoxHARPath)

	assert(t, "length of har pages", 1, len(har.Pages))
	page := har.Pages[0]
	assert(t, "firefox page start time", "2022-04-09T14:47:41.176-04:00", page.StartTime)
	assert(t, "firefox page ID", "page_1", page.ID)
	assert(t, "firefox page title", "Cyber Jake", page.Title)
	assert(t, "firefox page timings onContentLoad", float64(2035), page.PageTimings.ContentLoad)
	assert(t, "firefox page timings onLoad", float64(2452), page.PageTimings.Load)

}

func TestFirefoxEntry(t *testing.T) {
	har := prepHarFile(t, firefoxHARPath)

	entry := har.Entries[0]
	assert(t, "firefox entry IP", "104.22.40.104", entry.IP)
	assert(t, "firefox entry page ID", "page_1", entry.PageID)
	assert(t, "firefox entry port", "443", entry.Port)
	assert(t, "firefox entry secure status", "secure", entry.Secure)
	assert(t, "firefox entry startedTime", "2022-04-09T14:47:41.176-04:00", entry.StartedTime)
	assert(t, "firefox entry Time", float64(444), entry.Time)
}

func TestFirefoxRequest(t *testing.T) {
	har := prepHarFile(t, firefoxHARPath)

	request := har.Entries[0].Request

	assert(t, "firefox request URL", "https://cyberjake.xyz/", request.URL)
	assert(t, "firefox request cookies length", 0, len(request.Cookies))
	assert(t, "firefox request headers length", 12, len(request.Headers))
	assert(t, "firefox request method", "GET", request.Method)
	assert(t, "firefox request body size", 0, request.BodySize)
	assert(t, "firefox request http version", "HTTP/2", request.HTTPVersion)
	assert(t, "firefox request header size", 451, request.HeaderSize)
	assert(t, "firefox request query string length", 0, len(request.QueryString))

	_, err := request.CreateRequest()
	if err != nil {
		t.Errorf("Got error '%v' when creating request", err)
	}
}

func TestFirefoxResponse(t *testing.T) {
	har := prepHarFile(t, firefoxHARPath)

	response := har.Entries[0].Response

	assert(t, "firefox response status code", 200, response.Status)
	assert(t, "firefox response status text", "OK", response.StatusText)
	assert(t, "firefox response HTTP version", "HTTP/2", response.HTTPVersion)
	assert(t, "firefox response cookies length", 0, len(response.Cookies))
	assert(t, "firefox response headers size", 1767, response.HeaderSize)
	assert(t, "firefox response body size", 7978, response.BodySize)
	assert(t, "firefox response redirect URL", "", response.RedirectURL)
	assert(t, "firefox response MIME type", "text/html; charset=utf-8", response.Content.MIMEType)
	assert(t, "firefox response content text length", 19841, len(response.Content.Text))
	assert(t, "firefox response content size", 19844, response.Content.Size)
}
