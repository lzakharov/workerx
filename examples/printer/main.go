package main

import (
	"fmt"

	"github.com/lzakharov/workerx"
)

const limit uint = 3

func process(s string) error {
	fmt.Println(s)
	return nil
}

func main() {
	worker := workerx.NewWorkerPool[string](limit, process)

	worker.Add("foo")
	worker.Add("bar")
	worker.Add("baz")

	worker.Close()
}
