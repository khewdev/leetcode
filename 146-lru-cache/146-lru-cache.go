type LRUCache struct {
    capacity int
    head, tail *Node
    values map[int]*Node
}

type Node struct {
    key, value int
    prev, next *Node
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        values: map[int]*Node{},
        capacity: capacity,
    }
}

func (this *LRUCache) Get(key int) int {
    node, ok := this.values[key]
    if !ok {return -1}
    this.moveToLast(node)
    return node.value
}

func (this *LRUCache) moveToLast(node *Node) {
    if node == this.tail {return}
    if node == this.head {
        this.head = this.head.next
        this.head.prev = nil        
    } else {
        node.prev.next = node.next
        node.next.prev = node.prev
    }
    this.tail.next = node
    node.prev = this.tail
    this.tail = this.tail.next
    this.tail.next = nil    
}

func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.values[key]; ok {
        this.values[key].value = value
        this.moveToLast(this.values[key])
        return
    }
    if len(this.values) < this.capacity {
        this.append(key, value)
        return
    }
    node := this.head
    this.moveToLast(node)
    delete(this.values, node.key)

    this.values[key] = node
    node.key = key
    node.value = value
}

func (this *LRUCache) append(key, value int) {
    node := &Node{
        key: key,
        value: value,
    }
    if this.tail == nil {
        this.tail = node
        this.head = node
    } else {
        this.tail.next = node
        node.prev = this.tail
        this.tail = node
    }
    this.values[key] = node 
}