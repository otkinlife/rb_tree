package rb_tree

// rotateLeft 左旋操作
func (t *RBTree) rotateLeft(x *Node) {
	// 需要考虑x的右孩子和父节点的关系
	y := x.Right
	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

// rotateRight 右旋操作
func (t *RBTree) rotateRight(x *Node) {
	// 需要考虑x的左孩子和父节点的关系
	y := x.Left
	x.Left = y.Right
	if y.Right != nil {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Right = x
	x.Parent = y
}

// insertFixup 插入修复操作
func (t *RBTree) insertFixup(x *Node) {
	// 如果插入节点的父节点是红色，需要进行调整
	for x.Parent != nil && x.Parent.Color == RED {
		if x.Parent == x.Parent.Parent.Left {
			y := x.Parent.Parent.Right
			if y != nil && y.Color == RED {
				x.Parent.Color = BLACK
				y.Color = BLACK
				x.Parent.Parent.Color = RED
				x = x.Parent.Parent
			} else {
				if x == x.Parent.Right {
					x = x.Parent
					t.rotateLeft(x)
				}
				x.Parent.Color = BLACK
				x.Parent.Parent.Color = RED
				t.rotateRight(x.Parent.Parent)
			}
		} else {
			y := x.Parent.Parent.Left
			if y != nil && y.Color == RED {
				x.Parent.Color = BLACK
				y.Color = BLACK
				x.Parent.Parent.Color = RED
				x = x.Parent.Parent
			} else {
				if x == x.Parent.Left {
					x = x.Parent
					t.rotateRight(x)
				}
				x.Parent.Color = BLACK
				x.Parent.Parent.Color = RED
				t.rotateLeft(x.Parent.Parent)
			}
		}
	}
	// 根节点必须是黑色
	t.Root.Color = BLACK
}

// deleteFixup 删除修复操作
func (t *RBTree) deleteFixup(x *Node) {
	// 如果删除的是黑色节点，需要进行调整
	for x != t.Root && x.Color == BLACK {
		if x == x.Parent.Left {
			w := x.Parent.Right
			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color = RED
				t.rotateLeft(x.Parent)
				w = x.Parent.Right
			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				x = x.Parent
			} else {
				if w.Right.Color == BLACK {
					w.Left.Color = BLACK
					w.Color = RED
					t.rotateRight(w)
					w = x.Parent.Right
				}
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				t.rotateLeft(x.Parent)
				x = t.Root
			}
		} else {
			w := x.Parent.Left
			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color = RED
				t.rotateRight(x.Parent)
				w = x.Parent.Left
			}
			if w.Right.Color == BLACK && w.Left.Color == BLACK {
				w.Color = RED
				x = x.Parent
			} else {
				if w.Left.Color == BLACK {
					w.Right.Color = BLACK
					w.Color = RED
					t.rotateLeft(w)
					w = x.Parent.Left
				}
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Left.Color = BLACK
				t.rotateRight(x.Parent)
				x = t.Root
			}
		}
	}
	x.Color = BLACK
}

// transplant 用一棵子树替换另一棵子树
func (t *RBTree) transplant(u *Node, v *Node) {
	if u.Parent == nil {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
}

// min 找到一棵树的最小节点
func (t *RBTree) min(x *Node) *Node {
	for x.Left != nil {
		x = x.Left
	}
	return x
}
