package main

import (
	"fmt"
	"sync"
	"time"
)

type SomeStruct struct {
	values map[string]int
	mutex  sync.Mutex
}

func (s *SomeStruct) Count(key string) {
	// 只要任务go例程调用lock方法，互斥锁就会被解锁并被锁定。另一个go例程试图锁定互斥锁，就会被阻塞
	s.mutex.Lock()
	defer s.mutex.Unlock()
	value, ok := s.values[key]
	// 竞争条件读取
	if !ok {
		// 竞争条件写入
		s.values[key] = 1
	} else {
		s.values[key] = value + 1
	}
}

func main() {
	s := &SomeStruct{
		values: make(map[string]int),
	}

	// 当多个go例程访问共享变量时，可能存在竞争条件
	for i := 0; i < 8; i++ {
		go s.Count("foo")
	}

	time.Sleep(100 * time.Millisecond)
	// 最终的计数器有时是8，但不是每次都是
	fmt.Printf("%d\n", s.values["foo"])
}
