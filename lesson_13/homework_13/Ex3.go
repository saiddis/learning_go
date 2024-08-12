package main

type PointerOperations interface {
	Increment()
	Decrement()
}

type IntPointer struct {
	n *int
}

func (ip IntPointer) Increment() {
	*ip.n++
}

func (ip IntPointer) Decrement() {
	*ip.n--
}

type Pointer struct {
	PointerOperations
}

func NewPointer(po PointerOperations) Pointer {
	return Pointer{
		po,
	}
}
