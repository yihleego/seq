package seq

import (
	"strconv"
	"sync"
	"testing"
)

func TestNewSeq(t *testing.T) {
	seq := NewSeq(64)
	v := seq.Next()
	t.Logf("v: %d, b: %s", v, ToBinary(v))
}

func TestRandomSeq(t *testing.T) {
	seq := RandomSeq()
	v := seq.Next()
	t.Logf("v: %d, b: %s", v, ToBinary(v))
}

func TestConcurrency(t *testing.T) {
	seq := NewSeq(0)
	num := 100
	group := make([][]int64, num)
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		group[i] = make([]int64, num)
		go func(seq *Seq, ids []int64, i int) {
			defer wg.Done()
			for j := 0; j < num; j++ {
				v := seq.Next()
				ids[j] = v
			}
		}(seq, group[i], i)
	}
	wg.Wait()
	m := make(map[int64]byte)
	for i := range group {
		for j := range group[i] {
			v := group[i][j]
			m[v] = 1
			t.Logf("v: %d, b: %s", v, ToBinary(v))
		}
	}
	if len(m) != num*num {
		t.Error("Conflict")
	} else {
		t.Log("OK")
	}
}

func TestHex(t *testing.T) {
	seq := NewSeq(0)
	h := seq.NextHex()
	v, _ := strconv.ParseInt(h, 16, 64)
	t.Logf("v: %d, b: %s, h: %s", v, ToBinary(v), strconv.FormatInt(v, 16))
}

func ToBinary(n int64) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.FormatInt(lsb, 10) + result
	}
	return result
}
