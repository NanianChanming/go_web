package utils

import (
	"fmt"
	"testing"
)

func TestNextWorker(t *testing.T) {
	worker, err := NextWorker()
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	fmt.Println(worker.GetId())
}

func TestNewWorker(t *testing.T) {
	worker, err := NewWorker(1)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	fmt.Println(worker.GetId())
}
