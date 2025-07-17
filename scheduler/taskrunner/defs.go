package taskrunner

// 预定义类型定义
const (
	// 控制信号
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "e"
	CLOSE             = "c"

	// 任务状态
	// TaskStateRunning  = "running"
	// TaskStateFinished = "finished"
	// TaskStateFailed   = "failed"
)

type controlChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error
