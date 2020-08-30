/*
 * @Description: Do not edit
 * @Date: 2020-08-29 15:08:31
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2020-08-30 16:40:46
 */
package btree

// BSTInsert 这里假设插入的数据没有重复
func (root *Node) BSTInsert(curNode *Node) {
	rint := root.Data.(int)
	cint := curNode.Data.(int)
	if rint > cint {
		if root.Left != nil {
			root.Left.BSTInsert(curNode)
		} else {
			root.Left = curNode
		}
	} else {
		if root.Right != nil {
			root.Right.BSTInsert(curNode)
		} else {
			root.Right = curNode
		}
	}
}

func (root *Node) BSTFind(target interface{}) *Node {
	if root == nil {
		return nil
	}
	rint := root.Data.(int)
	tint := target.(int)
	if rint == tint {
		return root
	} else if tint < rint {
		return root.Left.BSTFind(target)
	} else {
		return root.Right.BSTFind(target)
	}
}

func (tree *Node) BSTDel(target interface{}) {
	//先找到要删除的节点及父节点
	var parent *Node
	p := tree
	for p != nil {
		nint := p.Data.(int)
		cint := target.(int)
		if nint == cint {
			break
		} else if nint > cint {
			parent = p
			p = p.Left
		} else {
			parent = p
			p = p.Right
		}
	}

	if p == nil {
		return
	}

	// p节点有两个子节点，将右子树的最小节点替换
	if p.Left != nil && p.Right != nil {
		min := p.Right
		minParent := p
		for min.Left != nil {
			min = min.Left
			minParent = nil
		}
		p.Data = min.Data
		parent = minParent
	}

	// p节点是子节点或者只有一个右子节点
	var child *Node
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	} else {
		child = nil
	}

	if parent == nil {
		tree = child
	} else {
		if parent.Left.Data.(int) == p.Data.(int) {
			parent.Left = child
		} else {
			parent.Right = child
		}
	}
}
