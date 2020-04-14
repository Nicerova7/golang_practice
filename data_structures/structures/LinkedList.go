package structures

type List struct {
	head *node
}
type node struct {
	prev *node
	key  int
	next *node
}

func NewList() List {
	L := new(List)
	L.head = nil
	return *L
}

func NewNode(x int) node {
	n := new(node)
	n.key = x
	return *n
}

func Empty_list(list List) bool {
	return list.head == nil
}

func (list *List) List_insert(x node) {
	x.next = list.head // .head or .next are pointers !!!
	if list.head != nil {
		list.head.prev = &x
	}
	list.head = &x
	x.prev = nil
}

func (list *List) List_delete(x node) {
	if x.prev != nil {
		x.prev.next = x.next
	} else {
		list.head = x.next
	}
	if x.next != nil {
		x.next.prev = x.prev
	}
}

func (list *List) List_search(k int) *node {
	x := list.head
	for x != nil && x.key != k {
		x = x.next
	}
	return x
}

/*
func main() {
	a := newList()
	x := newNode(2)
	y := newNode(3)
	z := newNode(4)
	a.list_insert(x)
	fmt.Println(empty_list(a))
	fmt.Println(a.head.key)
	a.list_delete(x)
	fmt.Println(empty_list(a))
	a.list_insert(x)
	a.list_insert(y)
	a.list_insert(z)
	fmt.Println(a.list_search(3)) // return pointer
	fmt.Println(a.list_search(23))
}
*/
