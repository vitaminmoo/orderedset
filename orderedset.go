package orderedset

// OrderedSet represents a set as defined at https://infra.spec.whatwg.org/#ordered-set
type OrderedSet struct {
	set []string
}

// NewOrderedSet creates an OrderedSet with the specified contents
func NewOrderedSet(contents []string) *OrderedSet {
	return &OrderedSet{contents}
}

/*
Append adds an item, if it doesn't exist, to the end of the OrderedSet
> To append to an ordered set: if the set contains the given item, then do nothing; otherwise, perform the normal list append operation.
*/
func (os *OrderedSet) Append(item string) {
	for _, i := range os.set {
		if i == item {
			return
		}
	}
	os.set = append(os.set, item)
}

/*
Prepend adds an item, if it doesn't exist, to the start of the OrderedSet
> To prepend to an ordered set: if the set contains the given item, then do nothing; otherwise, perform the normal list prepend operation.
*/
func (os *OrderedSet) Prepend(item string) {
	for _, i := range os.set {
		if i == item {
			return
		}
	}
	os.set = append([]string{item}, os.set...)
}

/*
Replace replaces an item in the OrderedSet
> To replace within an ordered set set, given item and replacement: if set contains item or replacement, then replace the first instance of either with replacement and remove all other instances.
*/
func (os *OrderedSet) Replace(item string, replacement string) {
	replaced := false
	var toDelete []int
	for n, i := range os.set {
		if i == item || i == replacement {
			if replaced {
				toDelete = append(toDelete, n)
			} else {
				os.set[n] = replacement
				replaced = true
			}
		}
	}
	for _, i := range toDelete {
		os.set = removeIndex(os.set, i)
	}
}

/*
Intersect intersects the two OrderedSets, returning a new OrderedSet
> The intersection of ordered sets A and B, is the result of creating a new ordered set set and, for each item of A, if B contains item, appending item to set.
*/
func (os *OrderedSet) Intersect(other *OrderedSet) *OrderedSet {
	newos := &OrderedSet{}
	for _, i := range os.set {
		for _, a := range other.set {
			if a == i {
				newos.Append(i)
			}
		}
	}
	return newos
}

/*
Union returns a union of both OrderedSets, as a new OrderedSet
> The union of ordered sets A and B, is the result of cloning A as set and, for each item of B, appending item to set."
*/
func (os *OrderedSet) Union(other *OrderedSet) *OrderedSet {
	newos := *os
	newos.set = make([]string, len(os.set))
	copy(newos.set, os.set)
	if other == nil {
		return &newos
	}
	for _, i := range other.set {
		newos.Append(i)
	}
	return &newos
}

/*
Range returns a chunk of an OrderedSet
> The range n to m, inclusive, creates a new ordered set containing all of the integers from n up to and including m in consecutively increasing order, as long as m is greater than or equal to n.
*/
func (os *OrderedSet) Range(start int, end int) *OrderedSet {
	ns := &OrderedSet{}
	ns.set = os.set[start : end+1]
	return ns
}

// AsSlice returns a copy of the Ordered Set as a slice for iteration
func (os *OrderedSet) AsSlice() []string {
	sl := make([]string, len(os.set))
	copy(sl, os.set)
	return sl
}
