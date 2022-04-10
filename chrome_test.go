package har

import "testing"

const chromeHARPath = "testdata/Chrome.har"

func TestChromeRun(t *testing.T) {
	har := prepHarFile(t, chromeHARPath)
	assert(t, "chrome har version", "1.2", har.Version)
	assert(t, "chrome har creator name", "WebInspector", har.Creator.Name)
	assert(t, "chrome har creator version", "537.36", har.Creator.Version)

}

func TestChromePage(t *testing.T) {
	har := prepHarFile(t, chromeHARPath)

	assert(t, "chrome length of har pages", 1, len(har.Pages))

	page := har.Pages[0]
	assert(t, "chrome page start time", "2022-04-09T18:48:12.680Z", page.StartTime)
	assert(t, "chrome page ID", "page_1", page.ID)
	assert(t, "chrome page title", "https://cyberjake.xyz/", page.Title)
	assert(t, "chrome page timings onContentLoad", 1146.9810000016878, page.PageTimings.ContentLoad)
	assert(t, "chrome page timings onLoad", 1337.6230000012583, page.PageTimings.Load)
}

func TestTimingsChrome(t *testing.T) {
	har := prepHarFile(t, chromeHARPath)

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

func TestChromeEntry(t *testing.T) {
	har := prepHarFile(t, chromeHARPath)

	entry := har.Entries[0]
	assert(t, "chrome entry IP", "104.22.40.104", entry.IP)
	assert(t, "chrome entry page ID", "page_1", entry.PageID)
	assert(t, "chrome entry port", "489", entry.Port)
	assert(t, "chrome entry startedTime", "2022-04-09T18:48:12.678Z", entry.StartedTime)
	assert(t, "chrome entry Time", 644.3219999989146, entry.Time)
}

func TestChromeRequest(t *testing.T) {
	har := prepHarFile(t, chromeHARPath)

	request := har.Entries[0].Request

	assert(t, "chrome request URL", "https://cyberjake.xyz/", request.URL)
	assert(t, "chrome request cookies length", 0, len(request.Cookies))
	assert(t, "chrome request headers length", 18, len(request.Headers))
	assert(t, "chrome request method", "GET", request.Method)
	assert(t, "chrome request body size", 0, request.BodySize)
	assert(t, "chrome request http version", "http/2.0", request.HTTPVersion)
	assert(t, "chrome request header size", -1, request.HeaderSize)
	assert(t, "chrome request query string length", 0, len(request.QueryString))

	_, err := request.CreateRequest()
	if err != nil {
		t.Errorf("Got error '%v' when creating request", err)
	}
}

func TestChromeResponse(t *testing.T) {
	har := prepHarFile(t, chromeHARPath)

	response := har.Entries[0].Response

	assert(t, "chrome response status code", 200, response.Status)
	assert(t, "chrome response status text", "", response.StatusText)
	assert(t, "chrome response HTTP version", "http/2.0", response.HTTPVersion)
	assert(t, "chrome response cookies length", 0, len(response.Cookies))
	assert(t, "chrome response headers size", -1, response.HeaderSize)
	assert(t, "chrome response body size", -1, response.BodySize)
	assert(t, "chrome response redirect URL", "", response.RedirectURL)
	assert(t, "chrome response MIME type", "text/html", response.Content.MIMEType)
	assert(t, "chrome response content text length", 19841, len(response.Content.Text))
	assert(t, "chrome response content size", 19844, response.Content.Size)
}
