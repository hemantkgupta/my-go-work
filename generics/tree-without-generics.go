package main

type Tree1 struct {
	val         Orderable
	left, right *Tree1
}

func (t *Tree1) Insert(val Orderable) *Tree1 {
	if t == nil {
		return &Tree1{val: val}
	}

	switch comp := val.Order(t.val); {
	case comp < 0:
		t.left = t.left.Insert(val)
	case comp > 0:
		t.right = t.right.Insert(val)
	}
	return t
}

type OrderableInt int

func (oi OrderableInt) Order(val any) int {
	return int(oi - val.(OrderableInt))
}

func main() {
	var it *Tree1
	it = it.Insert(OrderableInt(5))
	it = it.Insert(OrderableInt(3))
}

type Orderable interface {
	// Order returns:
	// a value < 0 when the Orderable is less than the supplied value,
	// a value > 0 when the Orderable is greater than the supplied value,
	// and 0 when the two values are equal.
	Order(any) int
}
