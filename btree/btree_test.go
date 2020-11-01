package btree

import (
	"fmt"
	"strings"
	"testing"
)

func TestOrder(t *testing.T) {
	root := NewNode(9)
	root.BSTInsert(&Node{Data: 7})
	root.BSTInsert(&Node{Data: 8})
	root.BSTInsert(&Node{Data: 4})
	root.BSTInsert(&Node{Data: 10})
	root.BSTInsert(&Node{Data: 12})
	root.BSTInsert(&Node{Data: 11})
	root.BSTInsert(&Node{Data: 1})
	root.BSTInsert(&Node{Data: 6})
	//PreOrder
	except := "974168101211"
	testRes := strings.Replace(strings.Trim(fmt.Sprint(root.PreOrder()), "[]"), " ", "", -1)
	if except != testRes {
		t.Errorf("btree.PreOrder test failed, want %v, got %v", except, testRes)
	}
	//MidOrder
	except = "146789101112"
	testRes = strings.Replace(strings.Trim(fmt.Sprint(root.MidOrder()), "[]"), " ", "", -1)
	if except != testRes {
		t.Errorf("btree.PreOrder test failed, want %v, got %v", except, testRes)
	}
	// LastOrder
	except = "164871112109"
	testRes = strings.Replace(strings.Trim(fmt.Sprint(root.LastOrder()), "[]"), " ", "", -1)
	if except != testRes {
		t.Errorf("btree.PreOrder test failed, want %v, got %v", except, testRes)
	}
	//LayerOrder
	except = "971048121611"
	testRes = strings.Replace(strings.Trim(fmt.Sprint(root.LayerOrder()), "[]"), " ", "", -1)
	if except != testRes {
		t.Errorf("btree.PreOrder test failed, want %v, got %v", except, testRes)
	}
	root.BSTDel(8)
	// after del 8
	//PreOrder
	except = "97416101211"
	testRes = strings.Replace(strings.Trim(fmt.Sprint(root.PreOrder()), "[]"), " ", "", -1)
	if except != testRes {
		t.Errorf("btree.PreOrder test failed, want %v, got %v", except, testRes)
	}
	//MidOrder
	except = "14679101112"
	testRes = strings.Replace(strings.Trim(fmt.Sprint(root.MidOrder()), "[]"), " ", "", -1)
	if except != testRes {
		t.Errorf("btree.PreOrder test failed, want %v, got %v", except, testRes)
	}
	// LastOrder
	except = "16471112109"
	testRes = strings.Replace(strings.Trim(fmt.Sprint(root.LastOrder()), "[]"), " ", "", -1)
	if except != testRes {
		t.Errorf("btree.PreOrder test failed, want %v, got %v", except, testRes)
	}
	//LayerOrder
	except = "97104121611"
	testRes = strings.Replace(strings.Trim(fmt.Sprint(root.LayerOrder()), "[]"), " ", "", -1)
	if except != testRes {
		t.Errorf("btree.PreOrder test failed, want %v, got %v", except, testRes)
	}
}
