package advanced

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 定义任务类型
type Task func()

type TaskResult struct {
	taskID int
	time   int
}

type TaskScheduler struct {
	tasks   []Task
	results []TaskResult
	ws      sync.WaitGroup
	me      sync.Mutex
	value   int
}

func (ts *TaskScheduler) AddTask(task Task) {
	ts.me.Lock()
	ts.tasks = append(ts.tasks, task)
	ts.me.Unlock()
}

func Job13() {
	taskScheduler := TaskScheduler{}
	fmt.Println("任务调度器开始执行任务", time.Now())
	//添加十个任务
	for i := 0; i < 10; i++ {
		taskScheduler.AddTask(func() {
			defer taskScheduler.ws.Done()
			satrtTime := time.Now()
			//每个任务执行1000000次自增
			for i := 0; i < 1000000; i++ {
				taskScheduler.me.Lock()
				taskScheduler.value++
				taskScheduler.me.Unlock()
			}
			endTime := time.Now()
			taskScheduler.me.Lock()
			taskScheduler.results = append(taskScheduler.results, TaskResult{taskID: i, time: int(endTime.Sub(satrtTime).Milliseconds())})
			taskScheduler.me.Unlock()
		})
	}
	for i := 0; i < len(taskScheduler.tasks); i++ {
		taskScheduler.ws.Add(1)
		go taskScheduler.tasks[i]()
	}
	taskScheduler.ws.Wait()
	fmt.Println("任务调度器执行完毕", time.Now())
	for i := 0; i < len(taskScheduler.results); i++ {
		fmt.Printf("任务%d执行时间：%d毫秒\n", taskScheduler.results[i].taskID, taskScheduler.results[i].time)
	}
}
