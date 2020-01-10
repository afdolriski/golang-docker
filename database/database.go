package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type (
	animal struct {
		Name  string `json:"name"`
		Sound string `json:"sound"`
	}
)

var animals []animal

func init() {
	wd, err := os.Getwd()
	if err != nil {
		er(err)
	}

	file, err := os.Open(fmt.Sprintf("%v/database/data.json", wd))
	if err != nil {
		er(err)
	}

	byteFile, _ := ioutil.ReadAll(file)
	if err != nil {
		er(err)
	}

	err = json.Unmarshal(byteFile, &animals)
	if err != nil {
		er(err)
	}
}

func GetAnimal(name string) (*animal, error) {
	for _, v := range animals {
		if name == v.Name {
			return &v, nil
		}
	}
	return nil, errors.New("No animal found")
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
