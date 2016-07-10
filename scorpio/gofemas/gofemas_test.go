// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gofemas

import (
	"testing"
)

// func TestCall(t *testing.T) {
// 	c := NewCaller()
// 	cb := NewCallback()

// 	c.SetCallback(cb)
// 	s := c.Call()
// 	if s != "Callback::run" {
// 		t.Errorf("unexpected string from Call: %q", s)
// 	}
// 	c.DelCallback()
// }

// func TestCallback(t *testing.T) {
// 	c := NewCaller()
// 	cb := NewDirectorCallback(&GoCallback{})
// 	c.SetCallback(cb)
// 	s := c.Call()
// 	if s != "GoCallback.Run" {
// 		t.Errorf("unexpected string from Call with callback: %q", s)
// 	}
// 	c.DelCallback()
// 	DeleteDirectorCallback(cb)
// }

// func TestCThostFtdcMdSpi(t *testing.T) {

// // 初始化UserApi
// pUserApi = CThostFtdcMdApi::CreateFtdcMdApi();			// 创建UserApi
// CThostFtdcMdSpi* pUserSpi = new CMdSpi();
// pUserApi->RegisterSpi(pUserSpi);						// 注册事件类
// pUserApi->RegisterFront(FRONT_ADDR);					// connect
// pUserApi->Init();

// pUserApi->Join();
// //	pUserApi->Release();

// 	pUserApi := CThostFtdcMdApiCreateFtdcMdApi()
// 	pUserSpi := NewDirectorCThostFtdcMdSpi(&GoCThostFtdcMdSpi{})
// 	pUserApi.RegisterSpi(pUserSpi)
// 	pUserApi.RegisterFront("tcp://180.168.146.187:10011")
// 	pUserApi.Init()

// 	pUserApi.Join()
// }

func TestCUstpFtdcMdSpi(t *testing.T) {
	pUserApi := CUstpFtdcTraderApiCreateFtdcTraderApi()

	pUserSpi := NewDirectorCUstpFtdcTraderSpi(&GoCUstpFtdcMdSpi{})

	pUserApi.RegisterSpi(pUserSpi)                        // 注册事件类
	pUserApi.SubscribePublicTopic(USTP_TERT_RESTART)      // 注册公有流
	pUserApi.SubscribePrivateTopic(USTP_TERT_RESTART)     // 注册私有流
	pUserApi.RegisterFront("tcp://180.168.146.187:10000") // connect
	pUserApi.Init()
	// wait for exit
	pUserApi.Join()
}
