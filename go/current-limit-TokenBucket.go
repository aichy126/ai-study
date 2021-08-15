package main

//模拟令牌桶限流

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

//https://github.com/juju/ratelimit  应用比较广泛的令牌桶包
//https://en.wikipedia.org/wiki/Token_bucket 令牌桶算法

type BucketLimiter struct {
	lastRequestTime int64
	tokenSurplus    int64 //剩余令牌数
	qps             int64 //每秒请求数
}

//NewBucketLimiter 初始化令牌桶
func NewBucketLimiter(tokenSurplus, qps int64) *BucketLimiter {
	return &BucketLimiter{
		lastRequestTime: time.Now().Unix(),
		tokenSurplus:    tokenSurplus,
		qps:             qps,
	}
}

//getToken 获取令牌
func (B *BucketLimiter) getToken() bool {
	now := time.Now().Unix()
	temp := (now-B.lastRequestTime)*B.qps + B.tokenSurplus
	tokenNow := B.getMin(temp, B.qps)
	if tokenNow > 0 {
		B.lastRequestTime = now
		B.tokenSurplus--
		return true
	}
	return false
}

//getMin 比较两个值 输出最小
func (B *BucketLimiter) getMin(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func main() {
	tokenBucket := NewBucketLimiter(5, 5)

	for i := 0; i < 100; i++ {
		spew.Dump(tokenBucket.getToken())
	}

}
