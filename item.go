package main

import (
	"time"
	"sort"
	"errors"
	"strings"
)


// A basic 'todo' Item
// Note: the body of the Item
// Created: time created, used for basic sorting
// Lu: last time updated, used for special sorting
//		{User defined '-Lu' and moving to the Done list after a while (probably a day)}
type Item struct {
	Note string
	Done bool
	Created time.Time
	Lu time.Time
}

func NewItem(s string) *Item {
	return &Item{
		Note: s,
		Done: false,
		Created: time.Now(),
		Lu: time.Now(),
	}
}

// Update the Note
func (i *Item) Update(s string) {
	i.Note = s
	i.Lu = time.Now()
}

// Set Done
func (i *Item) Mark(d bool) {
	i.Done = d
	i.Lu = time.Now()
}

// Item sort func-type
type byFunc func(i1, i2 *Item) bool

// Sorting function on Created
func byCreated(i1, i2 *Item) bool {
	return i1.Created.Before(i2.Created)
}

// Sorting function on Lu
func byUpdated(i1, i2 *Item) bool {
	return i1.Lu.Before(i2.Lu)
}

/////// Item list ///////////////////////////

type Items []*Item

// Good helper to order the list
func (li *Items) Sort(by byFunc) {
	sorter := &ItemSorter {
		i: *li,
		by: by,
	}
	sort.Sort(sorter)
}

// A little cleaner for adding an Item
func (li *Items) Add(i *Item) {
	*li = append(*li, i)
}

// Find helper 
func (li *Items) Find(note string) (*Item, error) {
	found := make([]int, 0)
	for i, e := range *li {
		if strings.HasPrefix(e.Note, note) {
			found = append(found, i);
		}
	}

	if len(found) > 1 {
		return nil, errors.New("Found more than one")
	} else if len(found) < 1 {
		return nil, errors.New("Didn't find any")
	} else {
		it := (*li)[found[0]]
		return it, nil
	}
}

// Helper functions to build a list based on the contents of the item
type buildSwitch func(*Item) bool

func buildList(li Items, swt buildSwitch) Items {
	list := make(Items, 0)

	for _, e := range li {
		if swt(e) {
			list.Add(e)
		}
	}

	return list
}

// Creates a list of all unmarked and recently marked Items
func (li *Items) Todo() Items {
	return buildList(*li, isTodo)
}

func isTodo(i *Item) bool {
	return !i.Done
}

// Creates a lite of all marked Items
func (li *Items) Done() Items {
	return buildList(*li, isDone)
}

func isDone(i *Item) bool {
	return i.Done
}


func (li *Items) Arch() Items {
	return buildList(*li, isArch)
}

func isArch(i *Item) bool {
	return (i.Done && (time.Since(i.Lu) > (24*time.Hour)))
}


/////// Sorter /////////////////////////////

// Sorter for a list of Items
// implements 'Sort' interface
// uses 'byCreated' or 'byUpdated' to sort
type ItemSorter struct {
	i Items
	by byFunc
}

func (s *ItemSorter) Len() int {
	return len(s.i)
}

func (s *ItemSorter) Swap(j, k int) {
	s.i[j], s.i[k] = s.i[k], s.i[j]
}

func (s *ItemSorter) Less(j, k int) bool {
	return s.by((s.i[j]), (s.i[k]))
}
