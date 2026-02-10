package others

import (
	"errors"
	"sync"
	"time"
)

/**

实现一个map：
1、面向高并发
2、只存在插入和查询操作 O(1)
3、查询时，若key存在，直接返回val；若key不存在，阻塞直到key val对被放入后，获取val返回；等待指定时长仍未放入，返回超时错误
4、写出真实代码，不能有死锁或者panic风险

*/

type MyConcurrentMap struct {
	lock sync.Mutex
	m    map[int]int
	chm  map[int]chan int
}

func (m *MyConcurrentMap) Put(key, val int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.m == nil {
		m.m = make(map[int]int)
	}
	m.m[key] = val
	ch, ok := m.chm[key]
	if !ok {
		return
	}
	select {
	case <-ch:
	default:
		close(ch)
	}
}

func (m *MyConcurrentMap) Get(key int, maxWaitingDuration time.Duration) (int, error) {
	var val int
	var err error
	m.lock.Lock()
	if val, ok := m.m[key]; ok {
		m.lock.Unlock()
		return val, nil
	}
	if m.chm == nil {
		m.chm = make(map[int]chan int)
	}
	if _, ok := m.chm[key]; !ok {
		m.chm[key] = make(chan int)
	}
	m.lock.Unlock()
	select {
	case <-time.After(maxWaitingDuration):
		err = errors.New("timeout")
	case <-m.chm[key]:
		m.lock.Lock()
		val = m.m[key]
		m.lock.Unlock()
	}
	return val, err
}
