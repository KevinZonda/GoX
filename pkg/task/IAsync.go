package task

type IAsync interface {
	// RunAsync runs the task asynchronously.
	Async()
}

func Async(task IAsync) {
	task.Async()
}

type IWaitable interface {
	Wait()
}

func Wait(task IWaitable) {
	task.Wait()
}
