package main

type Example2 struct {
	Field int
}

func PointerFunc(e *Example2) {
}

func ValueFunc(e Example2) {
}

func other() {
	var nilPtr *Example2

	PointerFunc(nilPtr) // equivalent to your e.PointerMethod()
	ValueFunc(*nilPtr)  // equivalent to your e.ValueMethod()
}
