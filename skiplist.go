package main

import (
	"math/rand"
	"time"
	"fmt"
)

const (
	maxLevel    = 16
	probability = 0.5
	SENTINAL    = -1
)

type Node struct {
	key  int
	next []*Node
}

// Creates new node
func NewNode(key, level int) *Node {
	return &Node{
		key:  key,
		next: make([]*Node, level),
	}
}

type SkipList struct {
	head *Node
	level int
}

// Creates new skiplist datastructure
func NewSkipList() *SkipList {
	return &SkipList{
		head:  NewNode(SENTINAL, maxLevel),
		level: 1,
	}
}

// Creates a random level
func randomLevel() int {
	level := 1
	for rand.Float64() < probability && level < maxLevel {
		level++
	}

	return level
}

// Searches key in the skiplist
func (sl *SkipList) Search(key int) bool {
	cur := sl.head

	// start from top level
	for i:=sl.level-1; i>=0; i--{

		// if current.next.key is < then we move to right
		// we do this because we assume there's a sentinal value
		// to the left and to the right
		for cur.next[i]!=nil && cur.next[i].key<key{
			cur = cur.next[i]
		}
	}
	// now we must be at the last level
	// just to be safe we manually enforce the level
	cur = cur.next[0]

	return cur!=nil && cur.key==key
}

func (sl *SkipList) Insert(key int) bool {
	update := make([]*Node, maxLevel) // Track pointers to update

	curr := sl.head

	// Find the position to insert the new node
	for i := sl.level - 1; i >= 0; i-- {
		for curr.next[i] != nil && curr.next[i].key < key {
			curr = curr.next[i]
		}
		update[i] = curr
	}

	// Create new node with a random level
	newLevel := randomLevel()
	if newLevel > sl.level {
		// Extend the update pointers
		for i := sl.level; i < newLevel; i++ {
			update[i] = sl.head
		}
		sl.level = newLevel
	}
	newNode := NewNode(key, newLevel)

	// Insert node at each level
	for i := 0; i < newLevel; i++ {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}

	return true
}

func (sl *SkipList) Delete(key int) bool{
	update := make([]*Node, maxLevel)
	curr := sl.head

	// Find node to delete
	for i := sl.level - 1; i >= 0; i-- {
		// move right if next key is smaller
		for curr.next[i] != nil && curr.next[i].key < key {
			curr = curr.next[i]
		}
		update[i] = curr
		// descend
	}

	// I somehow feel what if next[0] is null it could be ?
	target := curr.next[0]

	if target != nil && target.key == key {
		for i := 0; i < sl.level; i++ {
			if update[i].next[i] != target {
				break
			}
			update[i].next[i] = target.next[i]
		}

		// Reduce level if needed
		for sl.level > 1 && sl.head.next[sl.level-1] == nil {
			sl.level--
		}
		return true
	}

	return false
}

// Display prints the skip list for debugging.
func (sl *SkipList) Display() {
	// Collect all keys from level 0 for alignment
	keys := []int{}
	curr := sl.head.next[0]
	for curr != nil {
		keys = append(keys, curr.key)
		curr = curr.next[0]
	}

	// Print each level
	for i := sl.level - 1; i >= 0; i-- {
		curr := sl.head.next[i]
		fmt.Printf("Level %2d: ", i)

		for _, key := range keys {
			if curr != nil && curr.key == key {
				fmt.Printf("%3d ", key)
				curr = curr.next[i] // Move to the next node in this level
			} else {
				fmt.Printf(" -> ") // Print spaces for missing values
			}
		}

		fmt.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	sl := NewSkipList()

	sl.Insert(3)
	sl.Insert(6)
	sl.Insert(7)
	sl.Insert(9)
	sl.Insert(12)
	sl.Insert(19)
	for i:=0;i<10;i++{
		sl.Insert(10*i+2)
	}

	sl.Display()

	fmt.Println("Search 6:", sl.Search(6))
	fmt.Println("Search 15:", sl.Search(15))

	sl.Delete(6)
	fmt.Println("After deleting 6:")
	sl.Display()
}
