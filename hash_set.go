package hashset

type HashSet interface {
	Add(i string) (b bool)
	Has(i string) (b bool)
	Remove(i string)
}

type hashSet map[string]uint8

func New() HashSet {
	return make(hashSet)
}

func (set hashSet) Add(i string) (b bool) {
	_, b = set[i]
	set[i] = 0
	return !b //False if it existed already
}

func (set hashSet) Has(i string) (b bool) {
	_, b = set[i]
	return
}

func (set hashSet) Remove(i string) {
	delete(set, i)
}
