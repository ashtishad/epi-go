#Basics
A slice doesn't store any data, it just describes a section of an underlying array.

When you change an element of a slice, you modify the corresponding element of its underlying array, and other slices that share the same underlying array will see the change.
A slice can grow and shrink within the bounds of the underlying array.
Slices are indexed in the usual way: s[i] accesses the ith element, starting from zero.

#Built-in Functions
1. Append: The append function appends elements to a slice. It will automatically allocate a larger backing array if the capacity is exceeded.
2. Copy: The copy function copies elements into a destination slice dst from a source slice src. The number of elements copied is the minimum of len(dst) and len(src).
3. Make: which takes the length and an optional capacity as arguments.
4. Len and Cap: The built-in len and cap functions retrieve the length and capacity.