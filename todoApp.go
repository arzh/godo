package main

import (
	"os"
	"encoding/gob"
	"time"
	"todo"
	"io/ioutil"
)

const file ".todo"

type appData struct {
	curList string
	lists map[string]items
}

// fillout os flag data
func init_flags() {
	// -x: cross out
	// -a: show all
	// -d: show done
	// -lu: sort by last update instead of time created
	
	// future features
	// -sl: set current list
	// -al: show all todos on all lists
	// -ald: show all done on all lists
	// -ala: show all on all lists
}

var app appData

func make_app() {
	app.curList = ""
	app.lists = make(map[string]items)
}

func init_default() {
	app.curList = "default"
	app.lists["default"] = make(items, 0)
}

// Try to load the gob file
func load_app() (err error) {
	n, err := ioutil.ReadFile(file)
	if err != nil {
		if err == os.ErrNotExist {
			init_default() // No file, we need to init the default list
			err = nil  // we want to eat that error
		}
		return 
	}

	// Lets decode the file
	buf := bytes.NewBuffer(n)
	d := gob.Decoder(buf)
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

func main() {
	make_app()
	init_flags()
	if err := load_app(); err != nil {
		fmt.Println("Error loading app:", err.Error())
		return
	}




	if err := save_app(); err != nil {
		fmt.Println("Error saving app! Possible loss of data;", err.Error())
	}
}