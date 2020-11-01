/*
 * @Description: a base btree
 * @Date: 2020-08-29 14:01:45
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2020-11-01 17:51:50
 */

package btree

type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{Data: data}
}

func (node *Node) PreOrder() (res []interface{}) {
	if node == nil {
		return
	}
	res = append(res, node.Data)
	res = append(res, node.Left.PreOrder()...)
	res = append(res, node.Right.PreOrder()...)
	return
}

func (node *Node) MidOrder() (res []interface{}) {
	if node == nil {
		return
	}
	res = append(res, node.Left.MidOrder()...)
	res = append(res, node.Data)
	res = append(res, node.Right.MidOrder()...)
	return
}

func (node *Node) LastOrder() (res []interface{}) {
	if node == nil {
		return
	}
	res = append(res, node.Left.LastOrder()...)
	res = append(res, node.Right.LastOrder()...)
	res = append(res, node.Data)
	return
}

func (node *Node) LayerOrder() (res []interface{}) {
	if node == nil {
		return
	}
	nodes := []*Node{}
	nodes = append(nodes, node)
	for i := 0; i < len(nodes); i++ {
		node = nodes[i]
		res = append(res, node.Data)
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}
	return
}
