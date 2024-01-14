package main

func embeddedInterface() {
	t := A(T{})
	t.X()
}

type T struct{}

func (t T) X() {}

type A interface {
	B // embeded
}

type B = interface { // B is not a defined type
	X()
}