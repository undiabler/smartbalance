package smartb

import (
	"sync"
)

type leaf0 struct {
	sync.RWMutex
	leaf map[*element]*groupNode
}

type leaf1 struct {
	sync.RWMutex
	leaf map[string]*leaf0
}

type leaf2 struct {
	sync.RWMutex
	leaf map[string]*leaf1
}

type (
	groupNode struct {
		sync.RWMutex
		// сколько одновременно юзают этот элемент сейчас
		counter int
		// сколько элемент был выдан раз / вес
		weight int
	}

	GroupedLimitBalancer struct {
		limit int

		groups leaf2

		sync.RWMutex
		ranks map[*element]bool
	}
)

func (un *groupNode) get() (int, int) {
	un.RLock()
	defer un.RUnlock()
	return un.counter, un.weight
}

func (un *groupNode) inc() {
	un.Lock()
	un.weight++
	un.counter++
	un.Unlock()
}

func (un *groupNode) dec() {
	un.Lock()
	un.counter--
	un.Unlock()
}

func (ub *GroupedLimitBalancer) set(el *element) {
	ub.Lock()
	if _, ok := ub.ranks[el]; !ok {
		ub.ranks[el] = new(groupNode)
	}
	ub.Unlock()

}

func (ub *GroupedLimitBalancer) take(i int) (*element, interface{}) {

	var max_el *element
	var max_node *groupNode
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
func NewGroupedLimitBalancer(limit int) *GroupedLimitBalancer {

	ub := new(GroupedLimitBalancer)
	ub.ranks = make(map[*element]*groupNode)
	ub.limit = limit
	return ub

}
