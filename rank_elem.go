package smartb

type element struct {
	next, prev *element

	weight int
	value  interface{}
}

func (e *element) Prev() Elem {
	if e == nil {
		return nil
	}
	if e.prev == nil {
		return nil
	}
	return e.prev
}

func (e *element) Next() Elem {
	if e == nil {
		return nil
	}
	if e.next == nil {
		return nil
	}
	return e.next
}

func (e *element) Swap(e1 interface{}) {
	if obj, ok := e1.(*element); ok {
		o_prev := obj.prev
		obj.prev = e.prev
		e.prev = o_prev

		o_next := obj.next
		obj.next = e.next
		e.next = o_next
	}
}

func (e *element) Insert(e_p interface{}, e_n interface{}) {

	o_p, ok1 := e_p.(*element)
	o_n, ok2 := e_n.(*element)

	if ok2 {
		e.next = o_n
		o_n.prev = e
	}

	if ok1 {
		e.prev = o_p
		o_p.next = e
	}

}

func (e *element) Item() interface{} {
	if e == nil {
		return nil
	}
	return e.value

}
