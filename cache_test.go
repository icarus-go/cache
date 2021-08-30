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
