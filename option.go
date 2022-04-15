package workerx

// Option is a worker pool option.
type Option[T any] func(*WorkerPool[T])

// WithBefore is an Option to set the before function for the WorkerPool.
func WithBefore[T any](before func(T)) Option[T] {
	return func(w *WorkerPool[T]) {
		w.before = before
	}
}

// WithAfter is an Option to set the after function for the WorkerPool.
func WithAfter[T any](after func(T)) Option[T] {
	return func(w *WorkerPool[T]) {
		w.after = after
	}
}

// WithHandleErr is an Option to set the error handler for the WorkerPool.
func WithHandleErr[T any](handleErr func(T, error)) Option[T] {
	return func(w *WorkerPool[T]) {
		w.handleErr = handleErr
	}
}
