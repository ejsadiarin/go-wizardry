# Why does the empty anonymous struct have two pairs of braces? 'struct{}{}'

A. It doesn't, it's a syntax error.
B. Because the Go developers like to flex their 200 WPM typing speed
C. 'struct{}' is the type (empty struct) and `{}` is the value (empty struct literal)

Answer: C.

# Which is ordered from least -> most memory usage?

A. uint16, bool, int64, struct{}
B. struct{}, uint16, bool, int64
C. struct{}, bool, unint16, int64
D. bool, struct{}, uint16, int64

Answer: C.
