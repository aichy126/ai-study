package main

//模拟令牌桶限流

import (
	"math"
	"sync"
	"time"
)

func mian() {

}

// 定义令牌桶结构
type tokenBucket struct {
	timestamp time.Time // 当前时间戳
	capacity  float64   // 桶的容量（存放令牌的最大量）
	rate      float64   // 令牌放入速度
	tokens    float64   // 当前令牌总量
	lock      sync.Mutex
}

// 判断是否获取令牌（若能获取，则处理请求）
func (s *tokenBucket) getToken() bool {
	now := time.Now()
	s.lock.Lock()
	defer s.lock.Unlock()
	// 先添加令牌
	leftTokens := math.Max(s.capacity, s.tokens+now.Sub(s.timestamp).Seconds()*s.rate)
	if leftTokens < 1 {
		// 若桶中一个令牌都没有了，则拒绝
		return false
	} else {
		// 桶中还有令牌，领取令牌
		s.tokens -= 1
		s.timestamp = now
		return true
	}
}
