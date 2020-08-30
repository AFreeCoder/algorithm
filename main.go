/*
 * @Description: Do not edit
 * @Date: 2020-08-29 15:43:29
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2020-08-30 16:46:46
 */
package main

import (
	"algorithm/btree"
	"fmt"
)

func main() {
	root := btree.NewNode(9)
	root.BSTInsert(&btree.Node{Data: 7})
	root.BSTInsert(&btree.Node{Data: 8})
	root.BSTInsert(&btree.Node{Data: 4})
	root.BSTInsert(&btree.Node{Data: 10})
	root.BSTInsert(&btree.Node{Data: 12})
	root.BSTInsert(&btree.Node{Data: 11})
	root.BSTInsert(&btree.Node{Data: 1})
	root.BSTInsert(&btree.Node{Data: 6})
	fmt.Println("PreOrder: ", root.PreOrder())
	fmt.Println("MidOrder: ", root.MidOrder())
	fmt.Println("LastOrder: ", root.LastOrder())
	fmt.Println("LayerOrder: ", root.LayerOrder())
	root.BSTDel(8)
	fmt.Println("PreOrder: ", root.PreOrder())
	fmt.Println("MidOrder: ", root.MidOrder())
	fmt.Println("LastOrder: ", root.LastOrder())
	fmt.Println("LayerOrder: ", root.LayerOrder())
}
