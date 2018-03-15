package jsont

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonFile struct {
	FileName string
	perms    os.FileMode
}

// Load will open a file and load it into our interface.
func (j *JsonFile) Load(d interface{}) error {
	fstat, err := os.Stat(j.FileName)
	if err != nil {
		return err
	}
	j.perms = fstat.Mode()
	jsonFile, err := ioutil.ReadFile(j.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonFile, d)
}

func (j *JsonFile) Save(d interface{}) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}
	fmt.Printf("Saving File: %s\n", j.FileName)
	return ioutil.WriteFile(j.FileName, data, j.perms)
}

func (j *JsonFile) NewSave(d interface{}) error {
	data, err := json.Marshal(d)
	if err != nil {
		return err
	}
	fmt.Printf("Writing initial file: %s\n", j.FileName)
	return ioutil.WriteFile(j.FileName, data, 0664)
}
