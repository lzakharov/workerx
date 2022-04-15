package workerx

import (
	"sync"
)

// WorkerPool is a limited worker pool.
type WorkerPool[T any] struct {
	wg            *sync.WaitGroup
	limit         uint
	size          uint
	tasks         chan T
	before, after func(T)
	process       func(T) error
	handleErr     func(T, error)
}

// NewWorkerPool creates a new worker pool.
func NewWorkerPool[T any](
	limit uint,
	process func(T) error,
	opts ...Option[T],
) *WorkerPool[T] {
	worker := &WorkerPool[T]{
		wg:        &sync.WaitGroup{},
		limit:     limit,
		size:      0,
		tasks:     make(chan T, limit),
		before:    func(_ T) {},
		process:   process,
		after:     func(_ T) {},
		handleErr: func(_ T, _ error) {},
	}

	for _, opt := range opts {
		opt(worker)
	}

	return worker
}

// Add adds task and starts one more worker process if necessary.
func (w *WorkerPool[T]) Add(task T) {
	if w.size < w.limit {
		w.wg.Add(1)
		w.size++

		go w.run()
	}

	w.tasks <- task
}

// Close closes worker and waits for all processed to be done.
func (w *WorkerPool[T]) Close() {
	close(w.tasks)
	w.wg.Wait()
}

func (w *WorkerPool[T]) run() {
	defer w.wg.Done()

	for task := range w.tasks {
		w.before(task)

		err := w.process(task)
		if err != nil {
			w.handleErr(task, err)
		}

		w.after(task)
	}
}
