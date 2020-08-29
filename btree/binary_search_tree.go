/*
 * @Description: Do not edit
 * @Date: 2020-08-29 15:08:31
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2020-08-30 01:35:15
 */
package btree

func BSTInsert(root, curNode *Node) {
	rint := root.Data.(int)
	cint := curNode.Data.(int)
	if rint > cint {
		if root.Left != nil {
			BSTInsert(root.Left, curNode)
		} else {
			root.Left = curNode
		}
	} else {
		if root.Right != nil {
			BSTInsert(root.Right, curNode)
		} else {
			root.Right = curNode
		}
	}
}

func BSTFind(root *Node, target interface{}) *Node {
	if root == nil {
		return nil
	}
	rint := root.Data.(int)
	tint := target.(int)
	if rint == tint {
		return root
	} else if tint < rint {
		return BSTFind(root.Left, target)
	} else {
		return BSTFind(root.Right, target)
	}
}
