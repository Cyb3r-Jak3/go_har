package har

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

// HAR files have a root level log and this is used to get rid of it.
func skipRoot(jsonBlob []byte) (json.RawMessage, error) {
	var root map[string]json.RawMessage

	if err := json.Unmarshal(jsonBlob, &root); err != nil {
		return nil, err
	}
	for _, v := range root {
		return v, nil
	}
	return nil, nil
}

// CreateRequest will return a *http.Request for an Entry.Request.
func (harReq *Request) CreateRequest() (*http.Request, error) {
	req, err := http.NewRequest(harReq.Method, harReq.URL, strings.NewReader(harReq.Text))
	if err != nil {
		return nil, err
	}
	for _, cookie := range harReq.Cookies {
		req.AddCookie(&http.Cookie{Name: cookie.Name, Value: cookie.Value})
	}
	for _, header := range harReq.Headers {
		req.Header.Add(header.Name, header.Value)
	}
	return req, nil
}

// ParseHar reads a har file and returns a *File.
func ParseHar(filename string) (*File, error) {
	harFile := &File{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	message, err := skipRoot(byteValue)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(message, harFile)
	if err != nil {
		return nil, err
	}
	return harFile, nil
}
