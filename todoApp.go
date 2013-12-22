package main

import (
	"clu"
	"fmt"
)

// fillout os flag data
func clu_init(a clu.ArgConf) {
	// -x: cross out
	// -a: show all
	// -d: show done
	// -lu: sort by last update instead of time created
	a.AddFlag("mark", "x", "mark todo Item as done")
	a.AddFlag("show_all", "a", "show all Items")
	a.AddFlag("show_done", "d", "show all done Items")
	a.AddFlag("last_updated", "lu", "changes ordering")
	a.AddFlag("update", "u", "updates todo Item")

	// future features
	// -sl: set current list
	// -al: show all todos on all Lists
	// -ald: show all done on all Lists
	// -ala: show all on all Lists
}

func main() {
	make_app()
	args := clu.Parse(clu_init)

	//fmt.Println("CurList:", app.CurList, "Lists:", app.Lists)

	if err := load_app(); err != nil {
		fmt.Println("Error loading app:", err.Error())

	}

	//fmt.Println("CurList:", app.CurList, "Lists:", app.Lists)
	list := app.List()
	//fmt.Println("List:", list)

	// TODO: put checking of args and doing stuff here
	note := ""
	if args.LenLoose() > 0 {
		//fmt.Println(args.Loosie(0))
		note = args.Loosie(0)
	}

	//fmt.Println("Mark Flag:", args.Flag("mark"))

	if args.Flag("mark") {
		itm, err := list.Find(note)
		if err != nil {
			fmt.Println("Error Marking:", err.Error())
		}
		itm.Mark(true)
	} else if note != "" {
		list.Add(NewItem(note))
	}

	var view Items
	if args.Flag("show_all") {
		view = *list
	} else if args.Flag("show_done") {
		view = list.Done()
	} else {
		view = list.Todo()
	}

	for _, e := range view {
		fmt.Println(e.Print())
	}

	//fmt.Println("CurList:", app.CurList, "Lists:", app.Lists)

	if err := save_app(); err != nil {
		fmt.Println("Error saving app! Possible loss of data;", err.Error())
	}
}
