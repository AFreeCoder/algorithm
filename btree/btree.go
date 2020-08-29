/*
 * @Description: a base btree
 * @Date: 2020-08-29 14:01:45
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2020-08-30 01:19:19
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

func PreOrder(node *Node) (res []interface{}) {
	if node == nil {
		return
	}
	res = append(res, node.Data)
	res = append(res, PreOrder(node.Left)...)
	res = append(res, PreOrder(node.Right)...)
	return
}

func MidOrder(node *Node) (res []interface{}) {
	if node == nil {
		return
	}
	res = append(res, MidOrder(node.Left)...)
	res = append(res, node.Data)
	res = append(res, MidOrder(node.Right)...)
	return
}

func LastOrder(node *Node) (res []interface{}) {
	if node == nil {
		return
	}
	res = append(res, LastOrder(node.Left)...)
	res = append(res, LastOrder(node.Right)...)
	res = append(res, node.Data)
	return
}

func LayerOrder(node *Node) (res []interface{}) {
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
