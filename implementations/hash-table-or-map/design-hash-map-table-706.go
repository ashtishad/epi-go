package hash_table_or_map

const Len int = 100000

type HashNode struct {
	key   int
	value int
	next  *HashNode
}
type MyHashMap struct {
	data [Len]*HashNode
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func (this *HashNode) Put(key, value int) {
	if this.key == key {
		this.value = value
		return
	}
	if this.next == nil {
		this.next = &HashNode{key, value, nil}
		return
	}

	this.next.Put(key, value)
}

func (this *HashNode) Get(key int) int {
	if this.key == key {
		return this.value
	}
	if this.next == nil {
		return -1
	}

	return this.next.Get(key)
}

func (this *HashNode) Remove(key int) *HashNode {
	if this.key == key {
		p := this.next
		this.next = nil
		return p
	}
	if this.next != nil {
		return this.next.Remove(key)
	}
	return nil
}

func (this *MyHashMap) Put(key int, value int) {
	node := this.data[this.Hash(key)]
	if node == nil {
		this.data[this.Hash(key)] = &HashNode{key: key, value: value, next: nil}
		return
	}
	node.Put(key, value)
}

func (this *MyHashMap) Get(key int) int {
	node := this.data[this.Hash(key)]
	if node == nil {
		return -1
	}
	return node.Get(key)
}

func (this *MyHashMap) Remove(key int) {
	node := this.data[this.Hash(key)]
	if node == nil {
		return
	}
	this.data[this.Hash(key)] = node.Remove(key)

}

func (this *MyHashMap) Hash(value int) int {
	return value % Len
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
