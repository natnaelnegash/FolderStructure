package workerpool

import "context"

type Pool interface {
	Start()
	Submit(task Task) error
	SubmitWithKey(key string, maxConcurrent int, task Task) error
	Stop(ctx context.Context) error
	Drain()
}
