package queue

type Queue[T any] struct {
	vals []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		vals: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(val T) {
	q.vals = append(q.vals, val)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	front, ok := q.Front()
	if ok {
		q.vals = q.vals[1:]
	}
	return front, ok
}

func (q Queue[T]) Front() (T, bool) {
	var t T
	if q.IsEmpty() {
		return t, false
	}
	return q.vals[0], true
}

func (q Queue[T]) Len() int {
	return len(q.vals)
}

func (q Queue[T]) IsEmpty() bool {
	return len(q.vals) == 0
}
