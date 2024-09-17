package sgp

import (
	"fmt"
	"runtime"
	"testing"
)

func init() {
	println("using MAXPROC")
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
}

func BenchmarkPool(b *testing.B) {
	goRoutineLimit := 10
	taskCount := 400

	pool := NewPool(goRoutineLimit)
	pool.Run()
	defer pool.Close()

	myPrint := func() {
		fmt.Println("doing task")
	}
	task := NewTask(myPrint)
	for i := 0; i < taskCount; i++ {
		pool.AddTask(task)
	}
}
