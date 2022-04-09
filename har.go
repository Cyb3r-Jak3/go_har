package har

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// HAR files have a root level log and this is used to get rid of it
func skiproot(jsonBlob []byte) json.RawMessage {
	var root map[string]json.RawMessage

	if err := json.Unmarshal(jsonBlob, &root); err != nil {
		panic(err)
	}
	for _, v := range root {
		return v
	}
	return nil
}

// CreateRequest will return a *Request for a RequestStruct
func (hareq *RequestStruct) CreateRequest() (*http.Request, error) {
	req, err := http.NewRequest(hareq.Method, hareq.URL, strings.NewReader(hareq.PostData.Text))
	if err != nil {
		return req, err
	}
	for _, cookie := range hareq.Cookies {
		req.AddCookie(&http.Cookie{Name: cookie.Name, Value: cookie.Value})
	}
	for _, header := range hareq.Headers {
		req.Header.Add(header.Name, header.Value)
	}
	return req, nil
}

// parseHar will return a File struct
func parseHar(filename string) (*File, error) {
	harFile := &File{}
	file, err := os.Open(filename)
	if err != nil {
		return harFile, err
	}
	defer file.Close()
	bytevalue, berr := ioutil.ReadAll(file)
	if berr != nil {
		return harFile, berr
	}

	err = json.Unmarshal(skiproot(bytevalue), harFile)
	if err != nil {
		return harFile, err
	}
	return harFile, nil
}
