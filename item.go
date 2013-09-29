package main

// A basic 'todo' item
// note: the body of the item
// created: time created, used for basic sorting
// lu: last time updated, used for special sorting
//		{User defined '-lu' and moving to the done list after a while (probably a day)}
type item struct {
	note string
	done bool
	created time.Time
	lu time.Time
}

func NewItem(s string) item {
	return item{
		note: s,
		done: false
		created: time.Now(),
		lu: time.Now(),
	}
}

// Update the note
func (i item) Update(s string) {
	i.note = s
	i.lu = time.Now()
}

// Set done
func (i item) Mark(d bool) {
	i.done = d
	i.lu = time.Now()
}

// item sort func-type
type byFunc func(i1, i2 *item) bool

// Sorting function on created
func byCreated(i1, i2 *item) bool {
	return (i1.created < i2.created)
}

// Sorting function on lu
func byUpdated(i1, i2 *item) bool {
	return (i1.lu < i2.lu)
}

/////// Item list ///////////////////////////

type items []item

// Good helper to order the list
func (li items) Sort(by byFunc) {
	sorter := itemSorter { 
		i: li, 
		by: by, 
	}
	sort.Sort(sorter)
}

// A little cleaner for adding an item
func (li items) Add(i item) {
	li = append(li, i)
}

// Creates a list of all unmarked and recently marked items
func (li items) Todo() (ret items) {
	ret = make(items, 0)

	for _, e := range li {
		if !isDone(e) {
			ret.Add(e)
		}
	}

	return
}

// Creates a lite of all marked items
func (li items) Done() (ret items) {
	ret = make(items, 0)

	for _, e := range li {
		if isDone(e) {
			ret.Add(e)
		}
	}
}

func isDone(i item) bool {
	return (e.done && (time.Since(d.lu) > (24*time.Hour)))
}

/////// Sorter /////////////////////////////

// Sorter for a list of items
// implements 'Sort' interface
// uses 'byCreated' or 'byUpdated' to sort
type itemSorter struct {
	i items
	by byFunc
}

func (s *itemSorter) Len() int {
	return len(s.i)
}

func (s *itemSorter) Swap(j, k int) {
	s.i[j], s.i[k] = s.i[k], s.i[j]
}

func (s *itemsSorter) Less(j, k int) bool {
	return s.by(s.i[j], s.i[k])
}