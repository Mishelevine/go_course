package main

import (
	"fmt"
)

type AATree struct {
	root *Node // root элемент
	size int // размер дерева
}

type Node struct {
	key int // значение вершины
	level int // высота вершины
	left, right *Node // указатели на потомков
}

// NewAATree: конструктор пустого дерева
func NewAATree() *AATree {
	return &AATree{}
}

// Вставка элемента
func Insert(t *AATree, key int) {
	inserted := false
	t.root = insertRec(t.root, key, &inserted)
	if inserted {
		t.size++
	}
}

// Удаление элемента
func Delete(t *AATree, key int) bool {
	deleted := false
	t.root = deleteRec(t.root, key, &deleted)
	if deleted {
		t.size--
	}
	return deleted
}

// Проверка налчия элемента
func Contains(t *AATree, key int) bool {
	cur := t.root
	for cur != nil {
		if key < cur.key {
			cur = cur.left
		} else if key > cur.key {
			cur = cur.right
		} else {
			return true
		}
	}
	return false
}

// Уровень узла
func levelOf(n *Node) int {
	if n == nil {
		return 0
	}
	return n.level
}

func rotateLeft(n *Node) *Node {
	r := n.right
	n.right = r.left
	r.left = n
	return r
}

func rotateRight(n *Node) *Node {
	l := n.left
	n.left = l.right
	l.right = n
	return l
}

func skew(n *Node) *Node {
	if n == nil {
		return n
	}
	if n.left != nil && n.left.level == n.level {
		n = rotateRight(n)
	}
	return n
}

func split(n *Node) *Node {
	if n == nil {
		return n
	}
	if n.right != nil && n.right.right != nil && n.level == n.right.right.level {
		n = rotateLeft(n)
		n.level++
	}
	return n
}

func decreaseLevel(n *Node) *Node {
	if n == nil {
		return n
	}
	ideal := min(levelOf(n.left), levelOf(n.right)) + 1
	if ideal < n.level {
		n.level = ideal
		if n.right != nil && n.right.level > ideal {
			n.right.level = ideal
		}
	}
	return n
}

func insertRec(n *Node, key int, inserted *bool) *Node {
	if n == nil {
		*inserted = true
		return &Node{key: key, level: 1}
	}
	if key < n.key {
		n.left = insertRec(n.left, key, inserted)
	} else if key > n.key {
		n.right = insertRec(n.right, key, inserted)
	} else {
		return n
	}

	n = skew(n)
	n = split(n)
	return n
}

func deleteRec(n *Node, key int, deleted *bool) *Node {
	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = deleteRec(n.left, key, deleted)
	} else if key > n.key {
		n.right = deleteRec(n.right, key, deleted)
	} else {
		*deleted = true
		if n.left == nil && n.right == nil {
			return nil
		} else if n.left == nil {
			succ := n.right
			for succ.left != nil {
				succ = succ.left
			}
			n.key = succ.key
			n.right = deleteRec(n.right, succ.key, new(bool))
		} else {
			pred := n.left
			for pred.right != nil {
				pred = pred.right
			}
			n.key = pred.key
			n.left = deleteRec(n.left, pred.key, new(bool))
		}
	}

	n = decreaseLevel(n)

	n = skew(n)
	if n != nil {
		n.right = skew(n.right)
		if n.right != nil {
			n.right.right = skew(n.right.right)
		}
		n = split(n)
		if n != nil {
			n.right = split(n.right)
		}
	}

	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Объод дерева
func InOrder(n *Node, visit func(int)) {
	if n == nil {
		return
	}
	InOrder(n.left, visit)
	visit(n.key)
	InOrder(n.right, visit)
}

func task10() {
	t := NewAATree()

	for _, x := range []int{10, 5, 15, 3, 7, 13, 17, 6, 8, 16, 18} {
		Insert(t, x)
	}
	fmt.Println("Contains 7: ", Contains(t, 7))
	fmt.Println("Contains 11: ", Contains(t, 11))

	fmt.Print("InOrder: ")
	InOrder(t.root, func(k int) { fmt.Printf("%d ", k) })
	fmt.Println()

	fmt.Println("Delete 7  ->", Delete(t, 7))
	fmt.Println("Delete 15 ->", Delete(t, 15))
	fmt.Println("Delete 11 ->", Delete(t, 11))

	fmt.Print("InOrder: ")
	InOrder(t.root, func(k int) { fmt.Printf("%d ", k) })
	fmt.Println()
	fmt.Println("Size:", t.size)
}
