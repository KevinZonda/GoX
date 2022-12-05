package task

type ITask interface {
	Async()
	Wait()
}

func Async(task ITask) {
	task.Async()
}

func Wait(task ITask) {
	task.Wait()
}

func WaitAll(tasks []ITask) {
	for _, task := range tasks {
		task.Async()
	}
	for _, task := range tasks {
		task.Wait()
	}
}

type IRunnable interface {
	Run()
}

func Run(task IRunnable) {
	task.Run()
}

func RunAsync(task IRunnable) {
	go task.Run()
}
