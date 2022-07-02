package sparrow

// source: https://github.com/golang/go/blob/master/src/container/list/list.go

type element struct {
	next, prev *element
	list       *ConnList

	conn Conn
}

// Next returns the next list element or nil.
func (e *element) Next() *element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (e *element) Prev() *element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// ConnList represents a doubly linked list.
// The zero value for ConnList is an empty list ready to use.
type ConnList struct {
	root element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *ConnList) Init() *ConnList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// NewConnList returns an initialized list.
func NewConnList() *ConnList { return new(ConnList).Init() }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *ConnList) Len() int { return l.len }

// Front returns the first element of list l or nil if the list is empty.
func (l *ConnList) Front() *element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last element of list l or nil if the list is empty.
func (l *ConnList) Back() *element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit lazily initializes a zero List value.
func (l *ConnList) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (l *ConnList) insert(e, at *element) *element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *ConnList) insertValue(conn Conn, at *element) *element {
	return l.insert(&element{conn: conn}, at)
}

// remove removes e from its list, decrements l.len
func (l *ConnList) remove(e *element) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
}

// move moves e to next to at.
func (l *ConnList) move(e, at *element) {
	if e == at {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *ConnList) Remove(e *element) Conn {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.conn
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *ConnList) PushFront(conn Conn) *element {
	l.lazyInit()
	return l.insertValue(conn, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *ConnList) PushBack(conn Conn) *element {
	l.lazyInit()
	return l.insertValue(conn, l.root.prev)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *ConnList) InsertBefore(conn Conn, mark *element) *element {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(conn, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *ConnList) InsertAfter(conn Conn, mark *element) *element {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(conn, mark)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *ConnList) MoveToFront(e *element) {
	if e.list != l || l.root.next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, &l.root)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *ConnList) MoveToBack(e *element) {
	if e.list != l || l.root.prev == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, l.root.prev)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *ConnList) MoveBefore(e, mark *element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *ConnList) MoveAfter(e, mark *element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

// PushBackList inserts a copy of another list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *ConnList) PushBackList(other *ConnList) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.conn, l.root.prev)
	}
}

// PushFrontList inserts a copy of another list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *ConnList) PushFrontList(other *ConnList) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.conn, &l.root)
	}
}
