package main

import "fmt"

type LRUCache struct {
	Right *Node
	Left *Node
	Map  map[int]*Node
	Cap  int
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Map: make(map[int]*Node), Cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Map[key]
	if ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Map[key]

	if ok {
		this.Remove(node)
		node.Val = value
		this.Add(node)
	} else {
		if len(this.Map) == this.Cap {
			delete(this.Map, this.Left.Key)
			this.Remove(this.Left)
		}
		node = &Node{Key: key, Val: value}
		this.Map[key] = node
		this.Add(node)
	}
}

func (this *LRUCache) Add(node *Node) {
    node.Prev = nil
    node.Next = this.Right
    
    if this.Right != nil {
        this.Right.Prev = node
    }
    this.Right = node
    if this.Left == nil {
        this.Left = node
    }
    
// 	headNext := node.Next
// 	this.Head.Next = node

// 	node.Prev = this.Head
// 	node.Next = headNext

// 	headNext.Prev = node

}

func (this *LRUCache) Remove(node *Node) {
// 	nextNode := node.Next
// 	prevNode := node.Prev

// 	nextNode.Prev = prevNode
// 	nextNode.Next = nextNode
    
    if node != this.Right {
        node.Prev.Next = node.Next
    } else {
        this.Right = node.Next
    }
    if node != this.Left {
        node.Next.Prev = node.Prev
    } else {
        this.Left = node.Prev
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */