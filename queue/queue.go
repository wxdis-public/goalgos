package queue

// Queue : A circular FIFO queue.
type Queue[T any] struct {
	Vals  []T
	front int
	back  int
}

// New : Returns a circular FIFO queue, with a maximum capacity
// of the supplied length.
func New[T any](length int) Queue[T] {
	return Queue[T]{
		Vals:  make([]T, length),
		front: -1,
		back:  -1,
	}
}

// Enqueue : Add values to the back of the queue. Returns
// the number of values added and false if the queue filled
// up before all the values could be added.
func (q *Queue[T]) Enqueue(vals ...T) (int, bool) {
	for v := range vals {
		if q.IsEmpty() {
			q.front++
		} else if q.IsFull() {
			return v + 1, false
		}
		q.back = (q.back + 1) % len(q.Vals)
		q.Vals[q.back] = vals[v]
	}
	return len(vals), true
}

// Dequeue : Remove the value at the head of the queue.
func (q *Queue[T]) Dequeue() bool {
	if q.IsEmpty() {
		return false
	}
	q.Vals[q.front] = *new(T)
	if q.front == q.back {
		q.front = -1
		q.back = -1
	} else {
		q.front = (q.front + 1) % len(q.Vals)
	}
	return true
}

// IsEmpty : Returns true if the queue is empty.
func (q Queue[T]) IsEmpty() bool {
	return q.front == -1
}

// IsFull : Returns true if the queue is full.
func (q Queue[T]) IsFull() bool {
	return ((q.back + 1) % len(q.Vals)) == q.front
}

// Front : Fetch the value at the head of the queue. Will
// return the type arguments default value and 'false' if
// the queue is empty.
func (q Queue[T]) Front() (T, bool) {
	if q.IsEmpty() {
		return *new(T), false
	}
	return q.Vals[q.front], true
}

// Back : Fetch the value at the back of the queue. Will
// return the type arguments default value and 'false' if
// the queue is empty.
func (q Queue[T]) Back() (T, bool) {
	if q.IsEmpty() {
		return *new(T), false
	}
	return q.Vals[q.back], true
}
