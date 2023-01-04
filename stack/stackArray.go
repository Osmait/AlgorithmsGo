package stack

var StackArray []any

// stackPush push to  first index of array

func StaclPush(n any) {
	StackArray = append([]any{n}, StackArray...)
}

// StackLength return length of array
func StackLength() int {
	return len(StackArray)
}

// StackPeak return last input of array

func StackPeak() any {
	return StackArray[0]
}

// StackEmpty cheack array is empty or not
func StackEmpty() bool {
	return len(StackArray) == 0
}

// StackPop return last input and remove it in array
func StackPop() any {
	pop := StackArray[0]
	StackArray = StackArray[1:]
	return pop
}
