package internal

import (
	"sync"
	"testing"
)

func TestMySetCon(t *testing.T) {
	set := NewMySet()
	wg := sync.WaitGroup{}
	var count int = 1000

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < count; i++ {
				set.Add(i)
			}
		}()
	}

	wg.Wait()

	if len(set.inner) != count {
		t.Fail()
	}
}

func TestMySetAdd(t *testing.T) {
	set := NewMySet()

	if !set.Add(1) || len(set.inner) != 1 {
		t.Fail()
	}

	if set.Add(1) || len(set.inner) != 1 {
		t.Fail()
	}
}

func TestMySetDel(t *testing.T) {
	set := NewMySet()
	set.Add(1)
	set.Del(1)

	if len(set.inner) != 0 {
		t.Fail()
	}
}

func TestMySetIsExist(t *testing.T) {
	set := NewMySet()
	set.Add(1)

	if !set.IsExist(1) {
		t.Fail()
	}
}
