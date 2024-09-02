# Why is 5 the final value in the last index of 'j'?

A. j and g point to the same underlying array so g's append overwrote j
B. The Go team be trollin'
C. Because append only works properly when the number of elements is < 10
D. Because 5 was appended to i create j

Answer:

# Why doesn't the bug regarding slices 'j' and 'g' in example 2 occur in example 1 as well?

A. The array's cap() is exceeded so a new underlying array is allocated
B. The array in example 1 had a larger starting capacity
C. Because there are fewer elements and Go's runtime can't handle more than ~8

Answer:

# How can you best avoid these kinds of bugs?

A. Don't use the append() function  
B. Always assign the result of the append() function to a new slice
C. Always assign the result of the append() function back to the same slice

Answer:
