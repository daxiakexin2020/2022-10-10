某个任务的耗时、内存统计组件

type Server struct {
src       *service
taskNames []string
sl        sync.Mutex
err       chan error
}
