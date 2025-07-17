package main

import (
	"log"
)

// 定义并发连接数限制器结构体，用于限制并发连接数。
type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

// 创建并发连接数限制器实例。
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

// 获取连接。如果并发数达到限制，则返回false；否则返回true。
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}
	cl.bucket <- 1
	return true
}

// 释放连接。
func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New connction coming: %d", c)
}
