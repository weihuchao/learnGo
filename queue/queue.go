package queue

import "fmt"

/*
此时Queue队列的值只能是int, 如果需要可以传任意类型的, 可以使用 interface {}
*/

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() (int, error) {
	queue := *q
	if queue.IsEmpty() {
		return -1, fmt.Errorf("%s", "queue is empty.")
	}
	head := queue[0]
	*q = queue[1:]
	return head, nil

}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
