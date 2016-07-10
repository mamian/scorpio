// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goctp

//#cgo LDFLAGS: -fPIC -L. -L../../deps/ThostTraderApi_v6.3.5_tradeapi_linux64 -Wl,-rpath=../../deps/ThostTraderApi_v6.3.5_tradeapi_linux64 -lthostmduserapi -lthosttraderapi -lstdc++
//#cgo CPPFLAGS: -fPIC -I. -I../../deps/ThostTraderApi_v6.3.5_tradeapi_linux64
import "C"

import (
	"fmt"
	//"unsafe"
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
	iRequestID int                 = 0 // 请求编号
	pUserApi   CThostFtdcMdApi         // User API
	pTraderApi CThostFtdcTraderApi     // Trader AAPI
)

type GoCallback struct{}

func (p *GoCallback) Run() string {
	return "GoCallback.Run"
}

type GoCThostFtdcMdSpi struct{}

func (p *GoCThostFtdcMdSpi) OnRspError(pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	fmt.Println("GoCThostFtdcMdSpi.OnRspError.")
	p.IsErrorRspInfo(pRspInfo)
}

func (p *GoCThostFtdcMdSpi) OnFrontDisconnected(nReason int) {
	fmt.Println("GoCThostFtdcMdSpi.OnFrontDisconnected: %v", nReason)
}

func (p *GoCThostFtdcMdSpi) OnHeartBeatWarning(nTimeLapse int) {
	fmt.Println("GoCThostFtdcMdSpi.OnHeartBeatWarning: %v", nTimeLapse)
}

func (p *GoCThostFtdcMdSpi) OnFrontConnected() {
	fmt.Println("GoCThostFtdcMdSpi.OnFrontConnected.")
	p.ReqUserLogin()
}

func (p *GoCThostFtdcMdSpi) ReqUserLogin() {
	fmt.Println("GoCThostFtdcMdSpi.ReqUserLogin.")

	req := NewCThostFtdcReqUserLoginField()
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

func (p *GoCThostFtdcMdSpi) IsErrorRspInfo(pRspInfo CThostFtdcRspInfoField) bool {
	// 如果ErrorID != 0, 说明收到了错误的响应
	bResult := (pRspInfo.GetErrorID() != 0)
	if bResult {
		fmt.Printf("--->>> ErrorID=%v ErrorMsg=%v\n", pRspInfo.GetErrorID(), pRspInfo.GetErrorMsg())
	}
	return bResult
}

func (p *GoCThostFtdcMdSpi) OnRspUserLogin(pRspUserLogin CThostFtdcRspUserLoginField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	if bIsLast && !p.IsErrorRspInfo(pRspInfo) {

		fmt.Printf("--->>> 获取当前交易日  : %v\n", pUserApi.GetTradingDay())
		fmt.Printf("--->>> 获取用户登录信息: %v\n", pRspUserLogin)

		p.SubscribeMarketData()
		p.SubscribeForQuoteRsp()
	}
}

func (p *GoCThostFtdcMdSpi) SubscribeMarketData() {

	//ppInstrumentID := []string{"cu1609", "cu1610", "cu1611"}
	var name string = "cu1612"

	iResult := pUserApi.SubscribeMarketData(&name, 1)
	if iResult != 0 {
		fmt.Println("--->>> 发送行情订阅请求: 失败.")
	} else {
		fmt.Println("--->>> 发送行情订阅请求: 成功.")
	}
}

func (p *GoCThostFtdcMdSpi) SubscribeForQuoteRsp() {

	//ppInstrumentID := new("cu1611")

	var name string = "cu1612"

	iResult := pUserApi.SubscribeForQuoteRsp(&name, 1)
	if iResult != 0 {
		fmt.Println("--->>> 发送询价订阅请求: 失败.")
	} else {
		fmt.Println("--->>> 发送询价订阅请求: 成功.")
	}
}

func (p *GoCThostFtdcMdSpi) OnRspSubMarketData(pSpecificInstrument CThostFtdcSpecificInstrumentField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	fmt.Printf("OnRspSubMarketData:%#v\n", pRspInfo)
}

func (p *GoCThostFtdcMdSpi) OnRspSubForQuoteRsp(pSpecificInstrument CThostFtdcSpecificInstrumentField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	fmt.Printf("OnRspSubForQuoteRsp:%#v\n", pRspInfo)
}

func (p *GoCThostFtdcMdSpi) OnRspUnSubMarketData(pSpecificInstrument CThostFtdcSpecificInstrumentField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	fmt.Printf("OnRspUnSubMarketData:%#v\n", pRspInfo)
}

func (p *GoCThostFtdcMdSpi) OnRspUnSubForQuoteRsp(pSpecificInstrument CThostFtdcSpecificInstrumentField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	fmt.Printf("OnRspUnSubForQuoteRsp:%#v\n", pRspInfo)
}

func (p *GoCThostFtdcMdSpi) OnRtnDepthMarketData(pDepthMarketData CThostFtdcDepthMarketDataField) {
	fmt.Printf("OnRtnDepthMarketData:%#v\n", pDepthMarketData)
}

func (p *GoCThostFtdcMdSpi) OnRtnForQuoteRsp(pForQuoteRsp CThostFtdcForQuoteRspField) {
	fmt.Printf("OnRtnForQuoteRsp:%#v\n", pForQuoteRsp)
}

type GoCThostFtdcTraderSpi struct {
}

func (p *GoCThostFtdcTraderSpi) OnRspError(pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
	fmt.Println("GoCThostFtdcTraderSpi.OnRspError.")
	p.IsErrorRspInfo(pRspInfo)
}

func (p *GoCThostFtdcTraderSpi) OnFrontDisconnected(nReason int) {
	fmt.Println("GoCThostFtdcTraderSpi.OnFrontDisconnected: %v", nReason)
}

func (p *GoCThostFtdcTraderSpi) OnHeartBeatWarning(nTimeLapse int) {
	fmt.Println("GoCThostFtdcTraderSpi.OnHeartBeatWarning: %v", nTimeLapse)
}

func (p *GoCThostFtdcTraderSpi) OnFrontConnected() {
	fmt.Println("GoCThostFtdcTraderSpi.OnFrontConnected.")
	p.ReqUserLogin()
}

func (p *GoCThostFtdcTraderSpi) ReqUserLogin() {
	fmt.Println("GoCThostFtdcTraderSpi.ReqUserLogin.")

	req := NewCThostFtdcReqUserLoginField()
	req.SetBrokerID(BROKER_ID)
	req.SetUserID(INVESTOR_ID)
	req.SetPassword(PASSWORD)

	iRequestID++
	iResult := pTraderApi.ReqUserLogin(req, iRequestID)
	if iResult != 0 {
		fmt.Println("--->>> 发送用户登录请求: 失败.")
	} else {
		fmt.Println("--->>> 发送用户登录请求: 成功.")
	}
}

func (p *GoCThostFtdcTraderSpi) IsErrorRspInfo(pRspInfo CThostFtdcRspInfoField) bool {
	// 如果ErrorID != 0, 说明收到了错误的响应
	bResult := (pRspInfo.GetErrorID() != 0)
	if bResult {
		fmt.Printf("--->>> ErrorID=%v ErrorMsg=%v\n", pRspInfo.GetErrorID(), pRspInfo.GetErrorMsg())
	}
	return bResult
}

func (p *GoCThostFtdcTraderSpi) OnRspUserLogin(pRspUserLogin CThostFtdcRspUserLoginField, pRspInfo CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {

	if bIsLast && !p.IsErrorRspInfo(pRspInfo) {

		fmt.Printf("--->>> 获取当前交易日  : %v\n", pTraderApi.GetTradingDay())
		fmt.Printf("--->>> 获取用户登录信息: %#v\n", pRspUserLogin)

		//p.SubscribeMarketData()
		//p.SubscribeForQuoteRsp()
	}
}
