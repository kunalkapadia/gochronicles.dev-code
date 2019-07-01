package singlylinkedlist

// node represents a single item in the list, where each item has a value and points to another item
type node struct {
	value interface{}
	next  *node
}

// List has all the items of the list beginning at first and ending at last
type List struct {
	first *node
	last  *node
	size  int
}

// New instantiates a new list and adds values to the list if any
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Append(values...)
	}

	return list
}

// Append appends values to the list
func (list *List) Append(values ...interface{}) {
	for _, value := range values {
		newNode := &node{value: value}
		if list.size == 0 {
			list.first = newNode
			list.last = newNode
		} else {
			list.last.next = newNode
			list.last = newNode
		}
		list.size++
	}
}

// Prepend prepends values to the list in the reverse order
// So if list is [1, 2] and Prepend(3, 4) is called then updated list is [3, 4, 1, 2]
func (list *List) Prepend(values ...interface{}) {
	for i := len(values) - 1; i >= 0; i-- {
		newNode := &node{value: values[i], next: list.first}
		list.first = newNode
		if list.size == 0 {
			list.last = newNode
		}
		list.size++
	}
}

// Remove removes the node at given index
func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	var prevNode *node
	node := list.first
	for i := 0; i != index; i, node = i+1, node.next {
		prevNode = node
	}

	// If node to be removed is first node, update list.first to next one
	if node == list.first {
		list.first = node.next
	}

	// If node to be removed is last node, update list.last to prev node
	if node == list.last {
		list.last = prevNode
	}

	// Update prevNode to point to node next to node to be removed
	if prevNode != nil {
		prevNode.next = node.next
	}

	node = nil

	list.size--
}

// Get gets the value at given index
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}

	node := list.first
	for i := 0; i != index; i, node = i+1, node.next {
	}

	return node.value, true
}

// Values returns all values in the list
func (list *List) Values() []interface{} {
	values := make([]interface{}, list.size, list.size)
	for i, node := 0, list.first; node != nil; i, node = i+1, node.next {
		values[i] = node.value
	}
	return values
}

// Size returns the size of the list
func (list *List) Size() int {
	return list.size
}

// IsEmpty returns if the list is empty
func (list *List) IsEmpty() bool {
	return list.size == 0
}

// Clear clears the list
func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

// withinRange checks if given index is within list's range
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}
