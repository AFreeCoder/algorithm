/*
 * @Description: heap test
 * @Date: 2020-10-20 01:12:42
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2020-11-01 17:35:50
 */

package heap

import (
	"fmt"
	"testing"
)

func TestHeapfiyFromBottom(t *testing.T) {
	h := Heap{}
	h.Insert(6, 5, 9, 1, 8)
	want := "15968"
	fmt.Println(h.String())
	if h.String() != want {
		t.Errorf("test failed! want:%v, got:%v", want, h.String())
	}
}

func TestDeleteTop(t *testing.T) {
	h := Heap{}
	h.Insert(6, 5, 9, 1, 8)
	except := []int{1, 5, 6, 8, 9}
	for _, v := range except {
		top, err := h.DeleteTop()
		fmt.Println(h.meta)
		if err != nil {
			t.Errorf(err.Error())
		} else if top != v {
			t.Errorf("DeleteTop test failed, want %v, get %v", v, top)
		}
	}
}
