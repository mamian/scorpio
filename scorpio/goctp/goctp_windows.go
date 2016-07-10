// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goctp

//#cgo LDFLAGS: -fPIC -L. -L../../deps/ThostTraderApi_v6.3.5_tradeapi_win64 -Wl,-rpath=../../deps/ThostTraderApi_v6.3.5_tradeapi_win64 -lthostmduserapi -lthosttraderapi -lstdc++
//#cgo CPPFLAGS: -fPIC -I. -I../../deps/ThostTraderApi_v6.3.5_tradeapi_win64
import "C"

import (
	"fmt"
)

// char FRONT_ADDR[] = "tcp://180.168.146.187:10031";		// 前置地址
// TThostFtdcBrokerIDType	BROKER_ID = "9999";				// 经纪公司代码
// TThostFtdcInvestorIDType INVESTOR_ID = "039813";			// 投资者代码
// TThostFtdcPasswordType  PASSWORD = "123456";			// 用户密码
// char *ppInstrumentID[] = {"cu1512", "cu1601"};			// 行情订阅列表
// int iInstrumentID = 2;									// 行情订阅数量

const (
	BROKER_ID   = "9999"   // 经纪公司代码
	INVESTOR_ID = "039813" // 投资者代码
	PASSWORD    = "123456" // 用户密码
)

var (
	iRequestID int             = 0 // 请求编号
	pUserApi   CThostFtdcMdApi     // USER_API参数
)

type GoCallback struct{}

func (p *GoCallback) Run() string {
	return "GoCallback.Run"
}

type GoCThostFtdcMdSpi struct{}

func (p *GoCThostFtdcMdSpi) OnFrontConnected() {
	fmt.Println("GoCThostFtdcMdSpi.OnFrontConnected.")
	p.ReqUserLogin()
}

func (p *GoCThostFtdcMdSpi) ReqUserLogin() {
	fmt.Println("GoCThostFtdcMdSpi.ReqUserLogin.")

	req := CThostFtdcReqUserLoginField{
		BrokerID: BROKER_ID,
		UserID:   INVESTOR_ID,
		Password: PASSWORD,
	}
	iRequestID++
	iResult := pUserApi.ReqUserLogin(req, iRequestID)
	if iResult != 0 {
		fmt.Println("pUserApi.ReqUserLogin faild.")
	} else {
		fmt.Println("--->>> 发送用户登录请求: 成功.")
	}
}

// type GoCThostFtdcTraderSpi struct {
// 	UserApi GoCThostFtdcMdApi
// }

// func (p *GoCThostFtdcTraderSpi) OnFrontConnected() {
// 	fmt.Println("GoCThostFtdcTraderSpi.OnFrontConnected.")
// }

// func (p *GoCThostFtdcTraderSpi) OnHeartBeatWarning(nTimeLapse int) {
// 	fmt.Println("GoCThostFtdcTraderSpi.OnHeartBeatWarning.")
// }
