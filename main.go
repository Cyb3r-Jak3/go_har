package har

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

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

func parseHar(filename string) (File, error) {
	harFile := File{}
	file, err := os.Open(filename)
	if err != nil {
		return harFile, err
	}
	defer file.Close()
	bytevalue, berr := ioutil.ReadAll(file)
	if berr != nil {
		return harFile, berr
	}

	json.Unmarshal(skiproot(bytevalue), &harFile)
	return harFile, nil
}
