package main

import (
	"os"
	"encoding/gob"
	"io/ioutil"
	"bytes"
	"fmt"
)

const file = ".todo"
type ItemsMap map[string]*Items

type persis struct {
	CurList string
	Lists ItemsMap
}

func (p persis) List() *Items {
	list := p.Lists[p.CurList]
	return list
}


var app persis

func make_app() {
	app.CurList = ""
	app.Lists = make(ItemsMap)
}

func init_default() {
	app.CurList = "default"
	app.Lists["default"] = new(Items)
}

// Try to load the gob file
func load_app() (err error) {
	n, err := ioutil.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("init app")
			init_default() // No file, we need to init the default list
			err = nil  // we want to eat that error
		}
		return
	}

	// Lets decode the file
	buf := bytes.NewBuffer(n)
	d := gob.NewDecoder(buf)
	err = d.Decode(&app)

	return
}

// Writes out as a gob file
func save_app() error {
	// encode the appData
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	enc.Encode(app)

	// write out the persistence file
	err := ioutil.WriteFile(file, buf.Bytes(), 0600)
	return err
}