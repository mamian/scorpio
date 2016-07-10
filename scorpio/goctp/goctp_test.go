// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goctp

/*
BrokerID统一为：9999

第一套：

    标准CTP：

        第一组：Trade Front：180.168.146.187:10000，Market Front：180.168.146.187:10010；【电信】

        第二组：Trade Front：180.168.146.187:10001，Market Front：180.168.146.187:10011；【电信】

        第三组：Trade Front：218.202.237.33 :10002，Market Front：218.202.237.33 :10012；【移动】

        交易阶段(服务时间)：与实际生产环境保持一致

    CTPMini1：

        第一组：Trade Front：180.168.146.187:10003，Market Front：180.168.146.187:10013；【电信】

第二套：

    交易前置：180.168.146.187:10030，行情前置：180.168.146.187:10031；【7x24】

    第二套环境仅服务于CTP API开发爱好者，仅为用户提供CTP API测试需求，不提供结算等其它服务。

    新注册用户，需要等到第二个交易日才能使用第二套环境。

    账户、钱、仓跟第一套环境上一个交易日保持一致。

    交易阶段(服务时间)：交易日，16：00～次日09：00；非交易日，16：00～次日15：00。

    用户通过SimNow的账户（上一个交易日之前注册的账户都有效）接入环境，建议通过商业终端进行模拟交易的用户使用第一套环境。
*/

import (
	//"fmt"
	"testing"
)

func TestCall(t *testing.T) {
	c := NewCaller()
	cb := NewCallback()

	c.SetCallback(cb)
	s := c.Call()
	if s != "Callback::run" {
		t.Errorf("unexpected string from Call: %q", s)
	}
	c.DelCallback()
}

func TestCallback(t *testing.T) {
	c := NewCaller()
	cb := NewDirectorCallback(&GoCallback{})
	c.SetCallback(cb)
	s := c.Call()
	if s != "GoCallback.Run" {
		t.Errorf("unexpected string from Call with callback: %q", s)
	}
	c.DelCallback()
	DeleteDirectorCallback(cb)
}

func TestCThostFtdcMdSpi(t *testing.T) {
	pUserSpi := NewDirectorCThostFtdcMdSpi(&GoCThostFtdcMdSpi{})

	pUserApi = CThostFtdcMdApiCreateFtdcMdApi()
	pUserApi.RegisterSpi(pUserSpi)
	//pUserApi.RegisterFront("tcp://180.168.146.187:10011") // 第一套 Market Front
	pUserApi.RegisterFront("tcp://180.168.146.187:10031") // 第二套 Market Front
	pUserApi.Init()
	// wait for exit
	pUserApi.Join()
	pUserApi.Release()
}

func TestCThostFtdcTraderSpi(t *testing.T) {
	pUserSpi := NewDirectorCThostFtdcTraderSpi(&GoCThostFtdcTraderSpi{})

	pTraderApi = CThostFtdcTraderApiCreateFtdcTraderApi()
	pTraderApi.RegisterSpi(pUserSpi)                        // 注册事件类
	pTraderApi.SubscribePublicTopic(THOST_TERT_RESTART)     // 注册公有流
	pTraderApi.SubscribePrivateTopic(THOST_TERT_RESTART)    // 注册私有流
	pTraderApi.RegisterFront("tcp://180.168.146.187:10030") // connect
	//pTraderApi.RegisterFront("tcp://180.168.146.187:10000") // connect
	pTraderApi.Init()
	// wait for exit
	pTraderApi.Join()
	pTraderApi.Release()
}
