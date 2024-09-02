# Pointer Performance

Occasionally, new Go developers hear "pointers don't pass copies" and take that to a logical extreme, concluding:

> Pointers are always faster because copying is slow. I'll always use pointers!

_No. Bad. Stop._

**Here are my rules of thumb:**

1. First, worry about writing clear, correct, maintainable code.
2. If you have a performance problem, fix it.

Before even thinking about using pointers to optimize your code,

- use pointers when you need a shared reference to a value; otherwise, just use values.

**If you do have a performance problem, consider:**

1. [Stack vs. Heap](https://go.dev/doc/faq#stack_or_heap)
2. Copying

Interestingly, local non-pointer variables are generally faster to pass around than pointers because they're stored on the [stack](https://computersciencewiki.org/index.php?title=Stack_memory),

- which is faster to access than the [heap](https://computersciencewiki.org/index.php/Heap_memory). Even though copying is involved, the stack is so fast that it's no big deal.

Once the value becomes large enough that copying is the greater problem, it can be worth using a pointer to avoid copying.

That value will probably go to the heap, so the gain from avoiding copying needs to be greater than the loss from moving to the heap.

_One of the reasons Go programs tend to use less memory than Java and C# programs is that Go tends to allocate more on the stack._

I ran a [benchmark and wrote about it here](https://blog.boot.dev/golang/pointers-faster-than-values) if you're curious to dig deeper.
