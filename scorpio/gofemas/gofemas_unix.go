// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gofemas

//#cgo LDFLAGS: -fPIC -L. -L../../deps/FemasAPI_V2.00/FemasNewMarker2.00.00API/lnx64 -Wl,-rpath=../../deps/FemasAPI_V2.00/FemasNewMarker2.00.00API/lnx64 -lUSTPmduserapi -lUSTPtraderapi -lssl -lstdc++
//#cgo CPPFLAGS: -fPIC -I. -I../../deps/FemasAPI_V2.00/FemasNewMarker2.00.00API/include
import "C"

import (
	"fmt"
)

const (
	BROKER_ID   = "9999"   // 经纪公司代码
	INVESTOR_ID = "039813" // 投资者代码
	PASSWORD    = "123456" // 用户密码
)

var (
	iRequestID int                = 0 // 请求编号
	pUserApi   CUstpFtdcMduserApi     // USER_API参数
)

type GoCUstpFtdcMduserSpi struct{}

func (p *GoCUstpFtdcMduserSpi) OnRspError(pRspInfo CUstpFtdcRspInfoField, nRequestID int, bIsLast bool) {
	fmt.Println("GoCUstpFtdcMduserSpi.OnRspError.")
	p.IsErrorRspInfo(pRspInfo)
}

func (p *GoCUstpFtdcMduserSpi) OnFrontDisconnected(nReason int) {
	fmt.Println("GoCUstpFtdcMduserSpi.OnFrontDisconnected: %v", nReason)
}

func (p *GoCUstpFtdcMduserSpi) OnHeartBeatWarning(nTimeLapse int) {
	fmt.Println("GoCUstpFtdcMduserSpi.OnHeartBeatWarning: %v", nTimeLapse)
}

func (p *GoCUstpFtdcMduserSpi) OnFrontConnected() {
	fmt.Println("GoCUstpFtdcMduserSpi.OnFrontConnected.")
	p.ReqUserLogin()
}

func (p *GoCUstpFtdcMduserSpi) ReqUserLogin() {
	fmt.Println("GoCUstpFtdcMduserSpi.ReqUserLogin.")

	req := NewCUstpFtdcReqUserLoginField()

	req.SetBrokerID(BROKER_ID)
	req.SetUserID(INVESTOR_ID)
	req.SetPassword(PASSWORD)

	iRequestID++
	iResult := pUserApi.ReqUserLogin(req, iRequestID)
	if iResult != 0 {
		fmt.Println("--->>> 发送用户登录请求: 失败.")
	} else {
		fmt.Println("--->>> 发送用户登录请求: 成功.")
	}
}

func (p *GoCUstpFtdcMduserSpi) IsErrorRspInfo(pRspInfo CUstpFtdcRspInfoField) bool {
	// 如果ErrorID != 0, 说明收到了错误的响应
	bResult := (pRspInfo.GetErrorID() != 0)
	if bResult {
		fmt.Printf("--->>> ErrorID=%v ErrorMsg=%v\n", pRspInfo.GetErrorID(), pRspInfo.GetErrorMsg())
	}
	return bResult
}

func (p *GoCUstpFtdcMduserSpi) OnRspUserLogin(pRspUserLogin CUstpFtdcRspUserLoginField, pRspInfo CUstpFtdcRspInfoField, nRequestID int, bIsLast bool) {

	if bIsLast && !p.IsErrorRspInfo(pRspInfo) {

		fmt.Printf("--->>> 获取当前交易日   = %v\n", pUserApi.GetTradingDay())
		fmt.Printf("--->>> 获取用户登录信息 = %#v\n", pRspUserLogin)

		p.SubMarketData()
	}
}

func (p *GoCUstpFtdcMduserSpi) OnRspUserLogout(pRspUserLogout CUstpFtdcRspUserLogoutField, pRspInfo CUstpFtdcRspInfoField, nRequestID int, bIsLast bool) {
	fmt.Println("GoCUstpFtdcMduserSpi.OnRspUserLogout.")
}

func (p *GoCUstpFtdcMduserSpi) SubMarketData() {

	//ppInstrumentID := []string{"cu1609", "cu1610", "cu1611"}
	var name string = "cu1612"

	iResult := pUserApi.SubMarketData(&name, 1)
	if iResult != 0 {
		fmt.Println("--->>> 发送行情订阅请求: 失败.")
	} else {
		fmt.Println("--->>> 发送行情订阅请求: 成功.")
	}
}

func (p *GoCUstpFtdcMduserSpi) OnRspSubscribeTopic(pDissemination CUstpFtdcDisseminationField, pRspInfo CUstpFtdcRspInfoField, nRequestID int, bIsLast bool) {
	fmt.Println("GoCUstpFtdcMduserSpi.OnRspSubscribeTopic.")
}

func (p *GoCUstpFtdcMduserSpi) OnRspQryTopic(pDissemination CUstpFtdcDisseminationField, pRspInfo CUstpFtdcRspInfoField, nRequestID int, bIsLast bool) {
	fmt.Println("GoCUstpFtdcMduserSpi.OnRspQryTopic.")
}

func (p *GoCUstpFtdcMduserSpi) OnRtnDepthMarketData(pDepthMarketData CUstpFtdcDepthMarketDataField) {
	fmt.Println("GoCUstpFtdcMduserSpi.OnRtnDepthMarketData.")
}

func (p *GoCUstpFtdcMduserSpi) OnRspSubMarketData(pSpecificInstrument CUstpFtdcSpecificInstrumentField, pRspInfo CUstpFtdcRspInfoField, nRequestID int, bIsLast bool) {
	fmt.Println("GoCUstpFtdcMduserSpi.OnRspSubMarketData.")
}

func (p *GoCUstpFtdcMduserSpi) OnRspUnSubMarketData(pSpecificInstrument CUstpFtdcSpecificInstrumentField, pRspInfo CUstpFtdcRspInfoField, nRequestID int, bIsLast bool) {
	fmt.Println("GoCUstpFtdcMduserSpi.OnRspUnSubMarketData.")
}
