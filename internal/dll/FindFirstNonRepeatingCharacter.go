package dll

/*
Given a stream of characters, find the first non-repeating character from the stream.
You need to tell the first non-repeating character in O(1) time at any moment.

Input: S= "cleartaxisclear"

clear -> c
cleartaxis -> c
cleartaxisc -> l
cleartaxiscl -> e
cleartaxiscle -> r
cleartaxisclea -> r
cleartaxisclear -> t
*/

type DllNode struct {
	data byte
	prev *DllNode
	next *DllNode
}

func newNode(data byte) *DllNode {
	return &DllNode{data: data, prev: nil, next: nil}
}

type Dll struct {
	head *DllNode
}

func newDll(data byte) *Dll {
	return &Dll{head: newNode(data)}
}

func (dll *Dll) insert(node *DllNode) {
	if dll.head == nil {
		dll.head = node
	} else {
		t := dll.head
		for t.next != nil {
			t = t.next
		}
		t.next = node
		node.prev = t
	}
}

func (dll *Dll) insertData(data byte) {
	dll.insert(newNode(data))
}

func (dll *Dll) delete(node *DllNode) {
	if node == nil {
		return
	}

	prev := node.prev
	next := node.next

	if prev == nil && next == nil { // deleting the head node
		node = nil
	} else if prev == nil { // deleting the head node
		dll.head = node.next
		dll.head.prev = nil
	} else if next == nil { // last node
		prev.next = nil
		node.prev = nil
	} else {
		prev.next = next
		next.prev = prev
		node.prev = nil
		node.next = nil
	}
	node = nil
}

func (dll *Dll) toString() string {
	t := dll.head
	var byteArray []byte
	for t != nil {
		byteArray = append(byteArray, t.data)
		t = t.next
	}
	return string(byteArray)
}

func firstNonRepeatingCharecter(s string) byte {
	var dll *Dll
	lookup := map[byte]*DllNode{}

	for i := 0; i < len(s); i++ {
		if v, ok := lookup[s[i]]; ok {
			if dll != nil && v != nil {
				dll.delete(v)
			}
		} else {
			if dll == nil {
				dll = newDll(s[i])
				lookup[s[i]] = dll.head
			} else {
				node := newNode(s[i])
				dll.insert(node)
				lookup[s[i]] = node
			}
		}
	}
	return dll.head.data
}
