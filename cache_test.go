package cache

import (
	"pmo-test4.yz-intelligence.com/kit/cache/adapter"
	"pmo-test4.yz-intelligence.com/kit/cache/config"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	cfg := config.New(adapter.Redis,
		"",
		"",
		"3306",
		"")
	Instancer, err := New(cfg)
	if err != nil {
		println("init fail , err :", err.Error())
		return
	}

	get, err := Instancer.Get("HelloWorld")
	if err != nil {
		println("get fail, err :", err.Error())
	}
	println(get)

	set, err := Instancer.Set("HelloWorld", "Fuck", time.Duration(time.Second)*10)
	if err != nil {
		println("set fail, err :", err.Error())
		return
	}
	if set {
		println("ok")
		return
	}
	println("fail")
}
