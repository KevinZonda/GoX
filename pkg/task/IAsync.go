package task

type IAsync interface {
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

type IRunnable interface {
	Run()
}

func Run(task IRunnable) {
	task.Run()
}

func RunAsync(task IRunnable) {
	go task.Run()
}
