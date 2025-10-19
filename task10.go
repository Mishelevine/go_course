package main

import (
	"fmt"
)

type AATree struct {
	root *Node // указатель на root элемент
	size int   // размер дерева
}

type Node struct {
	key         int   // значение вершины
	level       int   // высота вершины
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

// Повороты для балансировок
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

// Балансировки
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

// Корректировка уровня при удалении
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

// Рекурсивная функция вставки
func insertRec(n *Node, key int, inserted *bool) *Node {
	// Если дошли до конца
	if n == nil {
		*inserted = true
		return &Node{key: key, level: 1}
	}
	if key < n.key {
		n.left = insertRec(n.left, key, inserted)
	} else if key > n.key {
		n.right = insertRec(n.right, key, inserted)
	} else { // если есть элемент равный вставке
		return n
	}

	// Балансировки
	n = skew(n)
	n = split(n)

	return n
}

// Рекурсивная функция удаления
func deleteRec(n *Node, key int, deleted *bool) *Node {
	// Если элемента нет
	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = deleteRec(n.left, key, deleted)
	} else if key > n.key {
		n.right = deleteRec(n.right, key, deleted)
	} else { // Нашли элемент, удаляем
		*deleted = true
		// Если у элемента нет детей просто удаляем
		if n.left == nil && n.right == nil {
			return nil
		} else if n.left == nil { // если есть только правый ребенок
			// ищем наименьший элемент в правом поддереве
			succ := n.right
			for succ.left != nil {
				succ = succ.left
			}
			// Ставим найденный элемент в текущий узел
			n.key = succ.key
			// Удаляем найденный элемент
			n.right = deleteRec(n.right, succ.key, new(bool)) // new используется чтобы передать незначимый нам указатель
		} else { // если есть левый ребенок
			// ицем наибольший элемент в левом поддереве
			pred := n.left
			for pred.right != nil {
				pred = pred.right
			}
			// Ставим найденный элемент в текущий узел
			n.key = pred.key
			n.left = deleteRec(n.left, pred.key, new(bool))
		}
	}

	// Вычисляет новый уровень узла
	n = decreaseLevel(n)

	// Балансировки
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
	InOrder(n.left, visit) // левое поддерево
	visit(n.key)
	InOrder(n.right, visit) // Правое поддерево
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
