/*
 * @Description: Do not edit
 * @Date: 2020-10-20 00:52:14
 * @LastEditors: wanghaijie01
 * @LastEditTime: 2020-11-01 17:39:47
 */

package heap

import (
	"fmt"
	"strings"
)

// 默认是小顶堆
type Heap struct {
	meta []int
}

func (h *Heap) insert(n int) {
	h.meta = append(h.meta, n)
	h.HeapifyFromBottom()
}

func (h *Heap) Insert(array ...int) {
	for _, val := range array {
		h.insert(val)
	}
}
func (h *Heap) String() string {
	return strings.Replace(strings.Trim(fmt.Sprint(h.meta), "[]"), " ", "", -1)
}

//HeapFyFromTop 自上而下的堆化
func (h *Heap) HeapfyFromTop() error {
	if len(h.meta) == 0 {
		return fmt.Errorf("the heap is nil")
	}
	if len(h.meta) == 1 {
		return nil
	}
	minPos := 1
	for i := 1; i <= len(h.meta)/2; {
		if 2*i <= len(h.meta) && h.meta[i-1] > h.meta[2*i-1] {
			minPos = 2 * i
		}
		if 2*i+1 <= len(h.meta) && h.meta[minPos] > h.meta[2*i] {
			minPos = 2*i + 1
		}
		if minPos == i {
			break
		}
		h.meta[i-1], h.meta[minPos-1] = h.meta[minPos-1], h.meta[i-1]
		i = minPos
	}
	return nil
}

// HeapifyFromBottom 自下而上的堆化
func (h *Heap) HeapifyFromBottom() {
	for i := len(h.meta); i >= 2; {
		// 和父节点比较
		if h.meta[i-1] < h.meta[i/2-1] {
			h.meta[i/2-1], h.meta[i-1] = h.meta[i-1], h.meta[i/2-1]
			i = i / 2
		} else {
			break
		}
	}
}

//DeleteTop 删除堆顶元素
func (h *Heap) DeleteTop() (int, error) {
	if len(h.meta) == 0 {
		return 0, fmt.Errorf("the heap is nil")
	}
	top := h.meta[0]
	if len(h.meta) == 1 {
		h.meta = []int{}
		return top, nil
	}
	h.meta[0] = h.meta[len(h.meta)-1]
	h.meta = h.meta[:(len(h.meta) - 1)]
	return top, h.HeapfyFromTop()
}
