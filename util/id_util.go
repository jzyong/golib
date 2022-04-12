package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"sync/atomic"
	"time"
)

//
type Uint32IdAllocator struct {
	id uint32
}

//@
func NewUint32IdAllocator() *Uint32IdAllocator {
	return &Uint32IdAllocator{}
}

//@
func (a *Uint32IdAllocator) Get() uint32 {
	id := atomic.AddUint32(&a.id, 1)
	if id == 0 {
		id = atomic.AddUint32(&a.id, 1)
	}
	return id
}

//
func GenerateSessionId() (string, error) {
	k := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return "", nil
	}
	return hex.EncodeToString(k), nil
}

/*
组成：0(1 bit) | timestamp in milli second (41 bit) | machine id (10 bit) | index (12 bit)
每毫秒最多生成4096个id，集群机器最多1024台
*/

type Snowflake struct {
	lastTimestamp int64
	index         int16
	machId        int16
}

func NewSnowflake(id int16) *Snowflake {
	sf := &Snowflake{}
	sf.Init(id)
	return sf
}

func (s *Snowflake) Init(id int16) error {
	if id > 0xff {
		return errors.New("illegal machine id")
	}

	s.machId = id
	s.lastTimestamp = time.Now().UnixNano() / 1e6
	s.index = 0
	return nil
}

func (s *Snowflake) GetId() (int64, error) {
	curTimestamp := time.Now().UnixNano() / 1e6
	if curTimestamp == s.lastTimestamp {
		s.index++
		if s.index > 0xfff {
			s.index = 0xfff
			return -1, errors.New("out of range")
		}
	} else {
		//fmt.Printf("id/ms:%d -- %d\n", s.lastTimestamp, s.index)
		s.index = 0
		s.lastTimestamp = curTimestamp
	}
	return int64((0x1ffffffffff&s.lastTimestamp)<<22) + int64(0xff<<10) + int64(0xfff&s.index), nil
}

var UUID *Snowflake

//计算字符串java hash值 int java byte为有符号（-128~127），go为无符号（0~256）
func GetJavaIntHash(s string) int32 {
	if len(s) < 1 {
		return 0
	}

	sum := md5.Sum([]byte(s))
	var hash int32 = 5381
	for i := 0; i < len(sum); i++ {
		cc := int32(sum[i])
		if cc > 127 { // java中需要进行转换
			cc = cc - 256
		}
		hash = hash + (hash << 5) + cc
		//fmt.Printf("%v \t %v \n", cc, hash)
	}
	if hash < 0 {
		hash = -hash
	}
	return hash
}
