package smartb

import (
	"fmt"
)

type holder interface {
	ToString() string
}

type Balancer interface {
	// getting elements
	take(int) (*element, interface{})

	// creating elements
	set(*element)
}

type Container struct {
	balancer Balancer
	current  *element
	all      map[string]*element
}
type element struct {
	next, prev *element
	item       holder
}

func NewContainer(bal Balancer, warming int) (*Container, error) {
	cnt := new(Container)

	cnt.balancer = bal
	cnt.all = make(map[string]*element)

	return cnt, nil
}

func (c *Container) Add(elem holder) {
	var key = elem.ToString()
	// TODO: mutex 2
	if obj, ok := c.all[key]; ok {
		fmt.Printf("Already in container: %p", obj.item)
	} else {
		el := new(element)
		el.item = elem
		c.all[key] = el
		c.balancer.set(el)
	}
}

func (c *Container) GetElem() *returner_obj {

	if len(c.all) == 0 {
		return nil
	}

	elem, node := c.balancer.take(0)

	if elem == nil {
		return nil
	}

	return &returner_obj{elem, node}
}

func (c *Container) Len() int {
	return len(c.all)
}
