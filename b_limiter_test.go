package smartb

import "testing"
import "fmt"

func TestLimiter(t *testing.T) {

	const elems_len = 10

	bl := NewLimitBalancer(100)

	cont, err := NewContainer(bl /*warmup cache*/, 100)

	if err != nil {
		t.Fatalf("Error creating container: %s", err)
	}

	if cont.GetElem() != nil {
		t.Error("Non nil return on empty elems")
	}

	for i := 0; i < elems_len; i++ {

		var testing_str Link = Link(fmt.Sprintf("hi test me %d", i))

		cont.Add(testing_str)

	}

	counter := make(map[string]int)

	// вызвали дохера раз
	for i := 0; i < 1000; i++ {
		el := cont.GetElem()
		if el == nil {
			t.Fatal("Nil pointer to elem!")
		}
		counter[string(el.Item().(Link))]++
	}

	if len(counter) != elems_len {
		t.Fatalf("Invalid get elems, expect %d got %d, container len : %d", elems_len, len(counter), cont.Len())
	}

	f := true
	prev := 0
	// проверяем что распределение равномерное
	for k, val := range counter {

		if f {
			prev = val
			f = false
		} else {
			if prev != val {
				t.Fatalf("Invalid balance: %q %d <> %d", k, prev, val)
			}
		}

	}

	t.Logf("All items: %v", counter)
	t.Logf("All items has weight: %d", prev)

	new_item := cont.GetElem()

	// all elements is in limit, shoud get nil
	if new_item != nil {
		t.Errorf("Limit is not working, got : %+v", new_item)
	}

}
