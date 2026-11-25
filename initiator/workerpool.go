package initiator

import "folderstructure/platform/workerpool"

type WorkerPool struct {
	Pool workerpool.Pool
}

func initWorkerPool(maxWorkers, buffer int) *WorkerPool {
	pool := workerpool.New(maxWorkers, buffer)
	pool.Start()
	return &WorkerPool{Pool: pool}
}
