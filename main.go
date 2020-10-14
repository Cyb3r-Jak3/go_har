package har

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func parseHar(filename string) File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytevalue, berr := ioutil.ReadAll(file)
	if berr != nil {
		log.Fatal(berr)
	}
	harFile := File{}
	json.Unmarshal(skiproot(bytevalue), &harFile)
	return harFile
}
