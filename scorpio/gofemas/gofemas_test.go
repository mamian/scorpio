// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gofemas

import (
	"testing"
)

func TestCUstpFtdcMdSpi(t *testing.T) {
	pUserSpi := NewDirectorCUstpFtdcMduserSpi(&GoCUstpFtdcMduserSpi{})

	pUserApi = CUstpFtdcMduserApiCreateFtdcMduserApi()
	pUserApi.RegisterSpi(pUserSpi)
	//pUserApi.RegisterFront("tcp://180.168.146.187:10011") // 第一套 Market Front
	pUserApi.RegisterFront("tcp://180.168.146.187:10031") // 第二套 Market Front
	pUserApi.Init()
	// wait for exit
	pUserApi.Join()
	pUserApi.Release()
}

// func TestCUstpFtdcMdSpi(t *testing.T) {
// 	pUserApi := CUstpFtdcTraderApiCreateFtdcTraderApi()

// 	pUserSpi := NewDirectorCUstpFtdcTraderSpi(&GoCUstpFtdcTraderSpi{})

// 	pUserApi.RegisterSpi(pUserSpi)                        // 注册事件类
// 	pUserApi.SubscribePublicTopic(USTP_TERT_RESTART)      // 注册公有流
// 	pUserApi.SubscribePrivateTopic(USTP_TERT_RESTART)     // 注册私有流
// 	pUserApi.RegisterFront("tcp://180.168.146.187:10030") // connect
// 	pUserApi.Init()
// 	// wait for exit
// 	pUserApi.Join()
// 	pUserApi.Release()
// }
