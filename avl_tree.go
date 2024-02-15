package main

import "fmt"

type Node struct {
	key    string
	left   *Node
	right  *Node
	height int
}

type Set struct {
	root *Node
}

// 1. Конструктор
func NewSet() *Set {
	return &Set{}
}

// 2.Конструктор копирования
func NewSetCopy(other *Set) *Set {
	if other == nil {
		return NewSet()
	}
	return &Set{root: copyNode(other.root)}
}

// Функция для глубокого копирования узла
func copyNode(node *Node) *Node {
	if node == nil {
		return nil
	}
	return &Node{
		key:    node.key,
		left:   copyNode(node.left),
		right:  copyNode(node.right),
		height: node.height,
	}
}

// 3.Оператор присваивания
func (s *Set) Assign(other *Set) {
	if s == other {
		return
	}
	s.root = copyNode(other.root)
}

// Перегрузка оператора равенства
func (s *Set) Equal(other *Set) bool {
	return isEqual(s.root, other.root)
}

// Рекурсивная функция для проверки равенства двух деревьев
func isEqual(node1, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 != nil && node2 != nil {
		return node1.key == node2.key &&
			isEqual(node1.left, node2.left) &&
			isEqual(node1.right, node2.right)
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(N *Node) int {
	if N == nil {
		return 0
	}
	return N.height
}

// Раставление указателей для нового элемента
func newNode(key string) *Node {
	node := &Node{key: key}
	node.left = nil
	node.right = nil
	node.height = 1
	return node
}

// Правый поворот
func rightRotate(y *Node) *Node {
	x := y.left
	T2 := x.right
	x.right = y
	y.left = T2
	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1
	return x
}

// Левый поворот
func leftRotate(x *Node) *Node {
	y := x.right
	T2 := y.left
	y.left = x
	x.right = T2
	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1
	return y
}

// Проверка на баланс
func getBalanceFactor(N *Node) int {
	if N == nil {
		return 0
	}
	return height(N.left) - height(N.right)
}

// 4. Добавление ключа
func insertNode(node *Node, key string) *Node {
	if node == nil {
		return newNode(key)
	}
	if key < node.key {
		node.left = insertNode(node.left, key)
	} else if key > node.key {
		node.right = insertNode(node.right, key)
	} else {
		return node
	}

	node.height = 1 + max(height(node.left), height(node.right))
	balanceFactor := getBalanceFactor(node)

	if balanceFactor > 1 {
		if key < node.left.key {
			return rightRotate(node)
		} else if key > node.left.key {
			node.left = leftRotate(node.left)
			return rightRotate(node)
		}
	}

	if balanceFactor < -1 {
		if key > node.right.key {
			return leftRotate(node)
		} else if key < node.right.key {
			node.right = rightRotate(node.right)
			return leftRotate(node)
		}
	}

	return node
}

func nodeWithMinimumValue(node *Node) *Node {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

// 5. Удаление узла из дерева
func deleteNode(root *Node, key string) *Node {
	if root == nil {
		return root
	}
	if key < root.key {
		root.left = deleteNode(root.left, key)
	} else if key > root.key {
		root.right = deleteNode(root.right, key)
	} else {
		if root.left == nil || root.right == nil {
			temp := root.left
			if temp == nil {
				temp = root.right
			}
			if temp == nil {
				temp = root
				root = nil
			} else {
				*root = *temp
			}
		} else {
			temp := nodeWithMinimumValue(root.right)
			root.key = temp.key
			root.right = deleteNode(root.right, temp.key)
		}
	}
	if root == nil {
		return root
	}
	root.height = 1 + max(height(root.left), height(root.right))
	balanceFactor := getBalanceFactor(root)

	if balanceFactor > 1 {
		if getBalanceFactor(root.left) >= 0 {
			return rightRotate(root)
		} else {
			root.left = leftRotate(root.left)
			return rightRotate(root)
		}
	}
	if balanceFactor < -1 {
		if getBalanceFactor(root.right) <= 0 {
			return leftRotate(root)
		} else {
			root.right = rightRotate(root.right)
			return leftRotate(root)
		}
	}
	return root
}

// 6. Поиск ключа
func searchKey(root *Node, key string) bool {
	found := false
	inOrderTraversal(root, func(currentKey string) {
		if currentKey == key {
			found = true
		}
	})
	return found
}

// 7. Печать дерева
func printTree(root *Node, indent string, last bool) {
	if root != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("R----")
			indent += " "
		} else {
			fmt.Print("L----")
			indent += "| "
		}
		fmt.Println(root.key)
		printTree(root.left, indent, false)
		printTree(root.right, indent, true)
	}
}

// 8. Обход дерева (Инфиксный)
func inOrderTraversal(root *Node, action func(string)) {
	if root != nil {
		inOrderTraversal(root.left, action)
		action(root.key)
		inOrderTraversal(root.right, action)
	}
}

func main() {
	set := NewSet()
	// Добавление 10 элементов, каждый на 2 больше предыдущего
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	for _, letter := range letters {
		set.root = insertNode(set.root, letter)
		fmt.Printf("After adding '%s':\n", letter)
		printTree(set.root, "", true)
		fmt.Println()
	}
	set1 := NewSet()
	for i := 1; i <= 10; i++ {
		set1.root = insertNode(set1.root, fmt.Sprintf("%d", i))
		fmt.Printf("After adding '%d':\n", i)
		printTree(set1.root, "", true)
		fmt.Println()
	}

	//проверим на равенство
	if set.Equal(set1) {
		fmt.Println("Sets are equal.")
	} else {
		fmt.Println("Sets are not equal.")
	}

	set.Assign(set1)

	printTree(set.root, "", true)
	printTree(set1.root, "", true)

	//проверим на равенство после set1:= set
	if set.Equal(set1) {
		fmt.Println("Sets are equal.")
	} else {
		fmt.Println("Sets are not equal.")
	}

	deleteNode(set1.root, "4")
	printTree(set1.root, "", true)
	// Создание множества и добавление элементов
	/*set := NewSet()
	set.root = insertNode(set.root, "33")
	set.root = insertNode(set.root, "13")
	set.root = insertNode(set.root, "53")
	set.root = insertNode(set.root, "9")
	set.root = insertNode(set.root, "21")
	set.root = insertNode(set.root, "61")
	set.root = insertNode(set.root, "8")
	set.root = insertNode(set.root, "11")

	// Печать AVL-дерева
	fmt.Println("AVL Tree:")
	printTree(set.root, "", true)

	// Удаление узла из AVL-дерева
	set.root = deleteNode(set.root, "13")
	fmt.Println("\nAfter deleting '13':")
	printTree(set.root, "", true)

	// Создание копии множества
	setCopy := NewSetCopy(set)
	fmt.Println("\nCopied AVL Tree:")
	printTree(setCopy.root, "", true)

	// Поиск ключа
	searchKey1 := "21"
	if searchKey(set.root, searchKey1) {
		fmt.Printf("\nKey '%s' found in the set.\n", searchKey1)
	} else {
		fmt.Printf("\nKey '%s' not found in the set.\n", searchKey1)
	}
	// Проверка на равенство двух множеств
	if set.Equal(setCopy) {
		fmt.Println("\nOriginal set is equal to the copied set.")
	} else {
		fmt.Println("\nOriginal set is not equal to the copied set.")
	}*/
}
