package toolPack

import (
	"os"
	"io"
	"sync"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

var lock sync.Mutex

// Marshal is a function that marshals the object into an
// io.Reader.
// By default, it uses the JSON marshaller.
var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}


// Save saves a representation of v to the file at path.
func SaveStruct(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}


// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func LoadStruct(path string, v interface{}) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
		return json.Unmarshal(file, v)
}