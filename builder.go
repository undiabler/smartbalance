package smartb

import (
	"fmt"
)

type Elem interface {
	Less(interface{}) bool

	Prev() Elem
	Next() Elem

	Swap(interface{})
	Insert(p interface{}, n interface{})

	Item() interface{}
}

type Ring struct {
	head Elem
	last Elem
	all  map[interface{}]Elem
}

func (e *element) Less(e1 interface{}) bool {
	if e1 == nil {
		return false
	}
	if obj, ok := e1.(*element); ok {
		return e.weight <= obj.weight
	}
	return false
}

func (r *Ring) Head() Elem {
	return r.head
}

func (r *Ring) Len() int {
	if r == nil || r.all == nil {
		return 0
	}
	return len(r.all)
}

func (r *Ring) NewElem(el Elem) {

	if r.head == nil {
		r.last = el
		r.head = el
		fmt.Println("first elem")
		return
	}

	if !el.Less(r.head) {
		fmt.Println("second case")
		el.Insert(nil, r.head)
		r.head = el
		return
	}

	pos := r.head.Next()
	var prev Elem = nil
	for el.Less(pos) {
		prev = pos
		pos = pos.Next()
	}

	el.Insert(prev, pos)

	if pos == nil {
		r.last = el
	}

}

func (r *Ring) Get() interface{} {
	if r.head == nil {
		return nil
	}
	return r.head.Item()

}

// // GetCurrent return value of element that curent ring point to
// func (r *Ring) GetCurrent() interface{} {
// 	if r.Len() > 0 {
// 		return r.current.value
// 	}
// 	return nil
// }

// // GetCurrent return value of element that curent ring point to and rotate forward
// func (r *Ring) GetNext() interface{} {
// 	if r.Len() > 0 {
// 		r.current = r.current.next
// 		return r.current.prev.value
// 	}
// 	return nil
// }

// // Append add value to one step back of the ring
// func (r *Ring) Append(value interface{}) {
// 	if value != nil {
// 		if r.all == nil {
// 			r.all = make(map[interface{}]*element)
// 			e := &element{value: value}
// 			e.next, e.prev, r.current, r.all[value] = e, e, e, e
// 		} else if r.Len() == 0 {
// 			e := &element{value: value}
// 			e.next, e.prev, r.current, r.all[value] = e, e, e, e
// 		} else if _, ok := r.all[value]; !ok {
// 			e := &element{value: value}
// 			e.next = r.current
// 			e.prev = r.current.prev
// 			r.current.prev.next = e
// 			r.current.prev = e
// 			r.all[value] = e
// 		}
// 	}
// }

// // Pop return and remove value
// func (r *Ring) Pop() interface{} {
// 	if r.Len() == 1 {
// 		v := r.current.value
// 		r.current = nil
// 		delete(r.all, v)
// 		return v
// 	} else if r.Len() > 0 {
// 		v := r.current.value
// 		r.current.prev.next = r.current.next
// 		r.current.next.prev = r.current.prev
// 		r.current = r.current.next
// 		delete(r.all, v)
// 		return v
// 	}
// 	return nil
// }

// // Remove value from the ring. Return true if find
// func (r *Ring) Remove(value interface{}) bool {
// 	if r.Len() == 0 {
// 		return false
// 	}
// 	if r.current.value == value {
// 		r.Pop()
// 		return true
// 	}
// 	if e, ok := r.all[value]; ok {
// 		e.next.prev = e.prev
// 		e.prev.next = e.next
// 		delete(r.all, value)
// 		return true
// 	}
// 	return false
// }
