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

func prepHarFile(t *testing.T, filepath string) *File {
	har, err := ParseHar(filepath)
	if err != nil {
		t.Errorf("Error parsing %s har file: %v", filepath, err)
	}
	return har
}

func TestFailure(t *testing.T) {
	_, err := ParseHar("testdata/NotHere.har")
	if !os.IsNotExist(err) {
		t.Errorf("Wanted error for file not found got '%v'", err)
	}
}

func TestGoodRequest(t *testing.T) {
	fakeRequest := &Request{Method: "GET", Cookies: []NameValuePair{{Name: "CookieName", Value: "CookieValue"}}, Headers: []NameValuePair{{Name: "HeaderName", Value: "HeaderValue"}}}
	req, err := fakeRequest.CreateRequest()
	if err != nil {
		t.Errorf("Error creating creating good request")
	}
	assert(t, "create request header", "HeaderValue", req.Header.Get("HeaderName"))
	cookie, err := req.Cookie("CookieName")
	if err != nil {
		t.Errorf("Error getting good request cookie")
	}
	if cookie == nil { // nolint:staticcheck
		t.Errorf("Didn't find the cookie")
	}
	assert(t, "create request cookie", "CookieValue", cookie.Value) // nolint:staticcheck
}

func TestBadRequest(t *testing.T) {
	fakeRequest := &Request{Method: "("}
	_, err := fakeRequest.CreateRequest()
	if err == nil {
		t.Error("Wanted error when creating bad request and didn't get one")
	}
}

func TestBadFile(t *testing.T) {
	_, err := ParseHar("testdata/BadHar.har")
	if err.Error() != "unexpected end of JSON input" {
		t.Errorf("Expected bad JSON got %s", err)
	}
}

func TestBadJSON(t *testing.T) {
	file, err := os.Open("testdata/bad.json")
	if err != nil {
		t.Errorf("Error opening bad JSON file: %v", err)
	}
	defer file.Close()
	bytevalue, err := io.ReadAll(file)
	if err != nil {
		t.Errorf("Error reading bad json file: %v", err)
	}
	_, err = skipRoot(bytevalue)
	if err == nil {
		t.Errorf("Expected error but got none")
	}
}
