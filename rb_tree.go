// Package rb_tree
// golang 红黑树实现
// 红黑树适合需要有序性和支持范围查询的场景
// 顺序性：红黑树是一种自平衡二叉查找树，保持了有序性，可以支持范围查询和有序遍历。
// 高效的插入和删除操作：红黑树通过自平衡的调整可以保持树的平衡，插入和删除操作的时间复杂度为 O(log n)。
// 内存占用较小：相比哈希表，红黑树不需要额外的空间来存储哈希函数和哈希冲突，内存占用较小。
package rb_tree

// Color 定义颜色类型
type Color bool

// 定义红色和黑色常量
const (
	RED   = false
	BLACK = true
)

// Node 红黑树节点
type Node struct {
	Key    int
	Value  int
	Color  Color
	Left   *Node
	Right  *Node
	Parent *Node
}

type RBTree struct {
	Root *Node
}

// NewNode 创建新的节点
func NewNode(key int, value int, color Color, left *Node, right *Node, parent *Node) *Node {
	return &Node{key, value, color, left, right, parent}
}

// NewRBTree 创建新的红黑树
func NewRBTree() *RBTree {
	return &RBTree{nil}
}

// Insert 插入操作
func (t *RBTree) Insert(key int, value int) {
	// 如果根节点为空，直接创建新的黑色节点作为根节点
	if t.Root == nil {
		t.Root = NewNode(key, value, BLACK, nil, nil, nil)
		return
	}

	// 否则从根节点开始，找到合适的插入位置
	curr := t.Root
	for {
		if key < curr.Key {
			if curr.Left == nil {
				curr.Left = NewNode(key, value, RED, nil, nil, curr)
				t.insertFixup(curr.Left)
				return
			}
			curr = curr.Left
		} else if key > curr.Key {
			if curr.Right == nil {
				curr.Right = NewNode(key, value, RED, nil, nil, curr)
				t.insertFixup(curr.Right)
				return
			}
			curr = curr.Right
		} else {
			// 如果找到了相同的key，直接更新value
			curr.Value = value
			return
		}
	}
}

// Delete 删除操作
func (t *RBTree) Delete(key int) {
	// 先找到要删除的节点
	node := t.Search(key)
	if node == nil {
		// 如果没找到，直接返回
		return
	}

	// 否则进行删除操作，需要考虑多种情况
	var child *Node
	color := node.Color

	// 如果左孩子为空，那么用右孩子替换当前节点
	if node.Left == nil {
		child = node.Right
		t.transplant(node, child)
		// 如果右孩子为空，那么用左孩子替换当前节点
	} else if node.Right == nil {
		child = node.Left
		t.transplant(node, child)
		// 如果两个孩子都不为空，那么找到右子树的最小节点来替换当前节点
	} else {
		successor := t.min(node.Right)
		color = successor.Color
		child = successor.Right

		if successor.Parent != node {
			t.transplant(successor, child)
			successor.Right = node.Right
			successor.Right.Parent = successor
		} else {
			if child != nil {
				child.Parent = successor
			}
		}

		t.transplant(node, successor)
		successor.Left = node.Left
		successor.Left.Parent = successor
		successor.Color = node.Color
	}

	// 如果删除的节点是黑色，需要进行调整
	if color == BLACK && child != nil {
		t.deleteFixup(child)
	}
}

// Search 查找操作
func (t *RBTree) Search(key int) *Node {
	// 从根节点开始查找
	curr := t.Root
	for curr != nil {
		if key < curr.Key {
			curr = curr.Left
		} else if key > curr.Key {
			curr = curr.Right
		} else {
			// 如果找到了，直接返回
			return curr
		}
	}
	// 如果没找到，返回nil
	return nil
}
