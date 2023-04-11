package asyncx

import "sync"

type WaitGroup struct {
	tasks []*Task[any]
}

func (wg *WaitGroup) Add(task *Task[any]) {
	wg.tasks = append(wg.tasks, task)
}

func (wg *WaitGroup) Wait() {
	w := sync.WaitGroup{}
	for _, task := range wg.tasks {
		t := task
		w.Add(1)
		go func() {
			t.wait()
			w.Done()
		}()
	}
	w.Wait()
}

func (wg *WaitGroup) AddFunc(f func()) {
	wg.Add(NewTask(func() interface{} {
		f()
		return nil
	}))
}
