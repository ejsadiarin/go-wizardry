package main

import "fmt"

/* January 21, 2024 4:58 PM
TOPIC: Pointer receiver method avoids runtime panic but value receiver method does not
------------------
Q:
I always presumed that a nil would cause a runtime panic if you attempted to call a method on the underlying type (in this a struct type).

Why is this possible to do with a pointer receiver but not a value receiver?

I imagine it has something to do with the syntactical sugar detailed in the spec here:https://go.dev/ref/spec#Method_expressions

But I'm not sure I understand it. Hoping someone here might be able to elaborate maybe for me.

Thanks!
------------------
A (detailed):
I might be oversimplifying or wrong but I think it has to do with dereferencing the pointer.
You get the panic when trying to access the underlying data and it is missing,
with the value receiver you call the method on the value, the pointer receiver is fine as null if you're not
interested in any data that is pointed to. When you have a field in that struct
and try to read/write it from a null pointer, that would cause panic from within the pointer receiver.

^ That's exactly right. The nil receiver is just a parameter - and just like passing
| nil in as any other parameter, it only panics on dereference.
|
| Now that I think about it, I wonder why it doesn't work in other languages? Maybe because in a language like Java,
| a "null pointer" is a pointer into the object store where the data and a pointer to the class metadata is stored,
| so it needs to dereference the pointer in order to get the method to call?

------------------
A (summary):
The explanation provided is accurate. In the given Go code, the difference between a pointer receiver method (PointerMethod)
and a value receiver method (ValueMethod) becomes apparent when a nil pointer is used.
When calling PointerMethod on a nil pointer, it does not result in a runtime panic because the
method is called on the pointer itself, and no attempt is made to access or modify the underlying data.
In contrast, calling ValueMethod on a nil pointer causes a runtime panic because the method attempts to
access the underlying data of the struct, resulting in a dereference panic.

The provided answer correctly states that the panic occurs when attempting to access the underlying data,
and with the value receiver, the method is called on the value itself, leading to a panic if the value is nil.
On the other hand, the pointer receiver is fine with a nil receiver when no data access is involved.

The additional speculation about why it might not work in other languages, like Java, suggests that in languages
where null pointers point to the object store, dereferencing might be necessary to access method calls.
This speculation highlights the language-specific nuances in handling nil or null pointers.


*/

type Example struct {
	Field int
}

// from: https://www.reddit.com/r/golang/comments/194rxwx/pointer_receiver_method_avoids_runtime_panic_but/?share_id=hyozBY74VxzeskuUFMAGR
func (e *Example) PointerMethod() {
	fmt.Printf("PointerMethod: %#v (%T)\n", e, e) // (*main.Example)(nil) (*main.Example)
}

func (e Example) ValueMethod() {
	fmt.Printf("ValueMethod: %#v (%T)\n", e, e)
}

// this does not create a new struct instance, but returns a nil pointer to an instance of the Example struct
func NewExample() *Example {
	return nil
}

func main() {
	e := NewExample()                    // this is a nil pointer
	fmt.Printf("main: %#v (%T)\n", e, e) // (*main.Example)(nil) (*main.Example)
	e.PointerMethod()                    // this is ok! no runtime panic
	e.ValueMethod()                      // this is a runtime panic
	fmt.Println(e.Field)                 // this is a runtime panic
}
