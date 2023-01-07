package queue

import "testing"

func TestQueue(t *testing.T) {
	t.Run("Test Queue Array", func(t *testing.T) {
		t.Run("Test EnQueue", func(t *testing.T) {
			Enqueue(2)
			Enqueue(23)
			Enqueue(45)
			Enqueue(66)
			if FrontQueue() != 2 && BackQue() != 66 {
				t.Errorf("Test EnQueue is worng the result must be %v and %v but got %v and %v", 2, 66, FrontQueue(), BackQue())
			}
		})
	})
}
