package smartb

import (
	"sync"
)

type (
	limitNode struct {
		sync.RWMutex
		counter int
	}

	limitBalancer struct {
		sync.RWMutex

		limit int
		ranks map[*element]*limitNode
	}
)

func (un *limitNode) get() int {
	un.RLock()
	defer un.RUnlock()
	return un.counter
}

func (un *limitNode) inc() {
	un.Lock()
	un.counter++
	un.Unlock()
}

func (un *limitNode) dec() {
	un.Lock()
	un.counter--
	un.Unlock()
}

func (ub *limitBalancer) set(el *element) {
	ub.Lock()
	if _, ok := ub.ranks[el]; !ok {
		ub.ranks[el] = new(limitNode)
	}
	ub.Unlock()

}

func (ub *limitBalancer) take(i int) (*element, interface{}) {

	var max_el *element
	var max_node *limitNode
	var max_val int

	ub.RLock()
	for k, val := range ub.ranks {

		tmp := val.get()

		if tmp >= ub.limit {
			continue
		}

		if max_el == nil {
			max_el = k
			max_node = ub.ranks[k]
			max_val = tmp
		} else {

			if tmp < max_val {
				max_el = k
				max_node = ub.ranks[k]
				max_val = tmp
			}
		}
	}
	ub.RUnlock()

	if max_node != nil {
		max_node.inc()
	}

	return max_el, max_node
}

// limit constructor
func NewLimitBalancer(limit int) *limitBalancer {

	ub := new(limitBalancer)
	ub.ranks = make(map[*element]*limitNode)
	ub.limit = limit
	return ub

}
