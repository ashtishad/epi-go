Go maps are implemented by hash tables and have efficient add, get and delete operations.
The map that’s built-in to Go is a hash map.

A map (or dictionary) is an unordered collection of key-value pairs, where each key is unique.
You create a new map with a make statement or a map literal.
The default zero value of a map is nil. A nil map is equivalent to an empty map except that elements can’t be added.
The len function returns the size of a map, which is the number of key-value pairs.

#Warning: 
If you try to add an element to an uninitialized map you get the mysterious run-time error Assignment to entry in nil map.
