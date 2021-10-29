package rtda

import "gvm/rtda/heap"

type Slot struct {
	num int32        // 整数
	ref *heap.Object // 引用
}
