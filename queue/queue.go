package queue

var ListQueue []any

func Enqueue(n any) {
	ListQueue = append(ListQueue, n)
}

func DeQueue() any {
	data := ListQueue[0]
	ListQueue = ListQueue[1:]
	return data
}

func FrontQueue() any {
	return ListQueue[0]
}
func BackQue() any {
	return ListQueue[len(ListQueue)-1]
}
func LenQueue() int {
	return len(ListQueue)
}
func IsEmptyQueue() bool {
	return len(ListQueue) == 0
}
