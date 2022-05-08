package seq

import (
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"
)

const (
	// Offset Unix timestamp offset. (2020-01-01 00:00:00 UTC)
	Offset = 1577836800
	// MaxWorkerId Maximum worker-id is 1023.
	MaxWorkerId = (1 << 10) - 1
	// MaxSequence Maximum sequence is 4194303.
	MaxSequence = (1 << 22) - 1
)

// Adder with initial random number.
var adder = rand.Int31n(MaxSequence)

// A Seq is a globally, 64 bits, thread-safe identifier. It can generate 4,194,303 numbers per second.
//	┌--------┬--------┬--------┬--------┬--------┬--------┬--------┬--------┐
//	|11111111|11111111|11111111|11111111|11111111|11111111|11111111|11111111| FORMAT: 64 bits
//	├--------┼--------┼--------┼--------┼--------┼--------┼--------┼--------┤
//	|XXXXXXXX|XXXXXXXX|XXXXXXXX|XXXXXXXX|        |        |        |        | TIMESTAMP: 32 bits
//	├--------┼--------┼--------┼--------┼--------┼--------┼--------┼--------┤
//	|        |        |        |        |XXXXXXXX|XX      |        |        | WORKER ID: 10 bits
//	├--------┼--------┼--------┼--------┼--------┼--------┼--------┼--------┤
//	|        |        |        |        |        |  XXXXXX|XXXXXXXX|XXXXXXXX| SEQUENCE: 22 bits
//	└--------┴--------┴--------┴--------┴--------┴--------┴--------┴--------┘
type Seq struct {
	workerId int32
}

// Next returns a new sequence.
func (s *Seq) Next() int64 {
	t := time.Now().Unix() - Offset
	w := int64(s.workerId)
	v := int64(atomic.AddInt32(&adder, 1) & MaxSequence)
	return t<<32 | w&0x3FF<<22 | v&0x3FFFFF
}

// NextHex returns an 8-byte hexadecimal string representation of the Seq.
func (s *Seq) NextHex() string {
	return strconv.FormatInt(s.Next(), 16)
}

// NewSeq returns a new Seq with the worker identifier, and the worker identifier should be between 0 and 1023.
func NewSeq(workerId int32) *Seq {
	if workerId < 0 {
		return &Seq{0}
	} else if workerId > MaxWorkerId {
		return &Seq{MaxWorkerId}
	} else {
		return &Seq{workerId}
	}
}

// RandomSeq returns a new Seq with a random worker identifier.
func RandomSeq() *Seq {
	return &Seq{rand.Int31n(MaxWorkerId)}
}
