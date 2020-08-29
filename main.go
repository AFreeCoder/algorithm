/*
 * @Description: Do not edit
 * @Date: 2020-08-29 15:43:29
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2020-08-30 01:37:54
 */
package main

import (
	"algorithm/btree"
	"fmt"
)

func main() {
	root := btree.NewNode(9)
	btree.BSTInsert(root, &btree.Node{Data: 8})
	btree.BSTInsert(root, &btree.Node{Data: 4})
	btree.BSTInsert(root, &btree.Node{Data: 10})
	btree.BSTInsert(root, &btree.Node{Data: 12})
	btree.BSTInsert(root, &btree.Node{Data: 1})
	btree.BSTInsert(root, &btree.Node{Data: 9})
	btree.BSTInsert(root, &btree.Node{Data: 6})
	fmt.Println("PreOrder: ", btree.PreOrder(root))
	fmt.Println("MidOrder: ", btree.MidOrder(root))
	fmt.Println("LastOrder: ", btree.LastOrder(root))
	fmt.Println("LayerOrder: ", btree.LayerOrder(root))
	fmt.Println("find 4: ", btree.BSTFind(root, 10))
}
