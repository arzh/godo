package main

import (
	"os"
	"encoding/gob"
	//"time"
	//"todo"
	"io/ioutil"
	"clu"
	"bytes"
	"fmt"
)

const file = ".todo"

type appData struct {
	curList string
	lists map[string]items
}

func (a appData) List() items {
	return a.lists[a.curList]
}

// fillout os flag data
func init_flags(a clu.ArgSet) {
	// -x: cross out
	// -a: show all
	// -d: show done
	// -lu: sort by last update instead of time created
	a.SetFlag("mark", "x", "mark todo item as done")
	a.SetFlag("show_all", "a", "show all items")
	a.SetFlag("show_done", "d", "show all done items")
	a.SetFlag("last_updated", "lu", "changes ordering")
	a.SetFlag("update", "u", "updates todo item")

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

func main() {
	make_app()
	args := clu.Parse(init_flags)

	if err := load_app(); err != nil {
		fmt.Println("Error loading app:", err.Error())

	}

	// TODO: put checking of args and doing stuff here
	if args.LenLoose() > 0 {
		fmt.Println(args.Loosie(0))
	}

	view := app.List().Todo()
	fmt.Println("Length", len(view))
	for _, e := range view {
		fmt.Println(e.note)
	}



	if err := save_app(); err != nil {
		fmt.Println("Error saving app! Possible loss of data;", err.Error())
	}
}
