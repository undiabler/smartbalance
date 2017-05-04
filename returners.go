package smartb

const (
	UBalancer = 1 + iota
)

// object to return, can hold specific data
type returner_obj struct {
	holder *element
	node   interface{}
}

func (r *returner_obj) Item() interface{} {
	return r.holder.item
}
