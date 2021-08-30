# cahe

封装缓存层的一系列操作，不用对正常 [Action]Cmd 进行重复代码的操作

并且封装相同的缓存层实现，统一存放相应的客户端，屏蔽实现。

## Redis使用规范建议

[Redis规范建议](https://www.cnblogs.com/winson-317/p/12684874.html)

## Quick start

```go
package cache

import (
	"pmo-test4.yz-intelligence.com/kit/cache/adapter"
	"pmo-test4.yz-intelligence.com/kit/cache/config"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	TestSet(t)

	TestGet(t)
}

var Instancer ICache

func init() {
	cfg := config.DefaultNew(adapter.Redis,
		"",
		"audit.yz-intelligence.com",
		"6379",
		"Tyyz@redis2019")

	var err error
	Instancer, err = New(cfg)
	if err != nil {
		println("init fail , err :", err.Error())
		return
	}
}

func TestSet(t *testing.T) {
	ok, err := Instancer.Set("HelloWorld", "2021-08-30 14:41:54", time.Second*10)
	if err != nil {
		println("set fail, err :", err.Error())
		return
	}
	println("isOK:", ok)
}

func TestGet(t *testing.T) {
	result, err := Instancer.Get("HelloWorld")
	if err != nil {
		println("get fail, err :", err.Error())
		return
	}
	println("result is :", result)
}

```