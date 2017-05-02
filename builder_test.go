package smartb

import "testing"
import "fmt"

func TestSomething(t *testing.T) {
	r := new(Ring)
	t.Logf("Item: %v", r.Get())
}

func TestOrder(t *testing.T) {
	r := new(Ring)

	for i := 0; i < 10; i++ {

		testing_str := fmt.Sprintf("hi test me %d", i)
		new_elem := element{nil, nil, i, &testing_str}
		r.NewElem(&new_elem)
	}

	head := r.Head()
	for head != nil {
		itm := head.Item()
		if itm != nil {
			t.Logf("Item: %p <- (%p %s) -> %p", head.Prev(), head, *itm.(*string), head.Next())
		} else {
			t.Logf("Item: %+v", head.Item())
		}
		head = head.Next()
	}
}
