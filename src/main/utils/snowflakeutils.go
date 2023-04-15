package utils

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

var maxWorkerId int64

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

func getMaxWorkerId(workerId int64) (int64, error) {
	if workerId < 0 {
		maxWorkerId += 1
		if maxWorkerId > workerMax {
			return workerId, errors.New("WorkerId 超出最大限制")
		}
		return maxWorkerId, nil
	}
	if workerId < maxWorkerId {
		maxWorkerId += 1
		return maxWorkerId, nil
	} else {
		maxWorkerId = workerId
		return workerId, nil
	}
}

func NextWorker() (*Worker, error) {
	workerId, err := getMaxWorkerId(-1)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return nil, errors.New("获取节点id错误")
	}
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	workerId, err := getMaxWorkerId(workerId)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return nil, errors.New("获取节点id错误")
	}
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	ID := int64((now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number))
	return ID
}

func NextId() int64 {
	worker, err := NextWorker()
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return 0
	}
	return worker.GetId()
}

func main() {
	// 生成节点实例
	node, err := NewWorker(1)
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println(node.GetId())
	}
}
