// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This .cc file will be automatically compiled by the go tool and
// included in the package.

#include <string>
#include "goctp.h"

std::string Caller::call() {
	if (callback_ != 0)
		return callback_->run();
	return "";
}


// // 标准化以后的标准C库
// #include <cstdio>
// #include <cstring>
// #include <cmath>
// // 标准化以后的标准c++库
// #include <iostream>
// using namespace std;

// //#pragma warning(disable : 4996)

// // UserApi对象
// CThostFtdcMdApi* pUserApi;

// // 配置参数
// char FRONT_ADDR[128] = {0};									// 前置地址
// TThostFtdcBrokerIDType	BROKER_ID  = {0};					// 经纪公司代码
// TThostFtdcInvestorIDType INVESTOR_ID  = {0};				// 投资者代码
// TThostFtdcPasswordType  PASSWORD = {0};						// 用户密码
// char *ppInstrumentID[] = { "cu1601", "cu1602"};		    	// 行情订阅列表
// int iInstrumentID = 6;										// 行情订阅数量

// // 请求编号
// int iRequestID;

// void CMdSpi::OnRspError(CThostFtdcRspInfoField *pRspInfo,
// 		int nRequestID, bool bIsLast)
// {
// 	cout << "--->>> "<< "OnRspError" << endl;
// 	IsErrorRspInfo(pRspInfo);
// }

// void CMdSpi::OnFrontDisconnected(int nReason)
// {
// 	cout << "--->>> " << "OnFrontDisconnected" << endl;
// 	cout << "--->>> Reason = " << nReason << endl;
// }
		
// void CMdSpi::OnHeartBeatWarning(int nTimeLapse)
// {
// 	cout << "--->>> " << "OnHeartBeatWarning" << endl;
// 	cout << "--->>> nTimerLapse = " << nTimeLapse << endl;
// }

// void CMdSpi::OnFrontConnected()
// {
// 	cout << "--->>> " << "OnFrontConnected" << endl;
// 	///用户登录请求
// 	ReqUserLogin();
// }

// void CMdSpi::ReqUserLogin()
// {
// 	CThostFtdcReqUserLoginField req;
// 	memset(&req, 0, sizeof(req));
// 	strcpy(req.BrokerID, BROKER_ID);
// 	strcpy(req.UserID, INVESTOR_ID);
// 	strcpy(req.Password, PASSWORD);
// 	int iResult = pUserApi->ReqUserLogin(&req, ++iRequestID);
// 	cout << "--->>> 发送用户登录请求: " << ((iResult == 0) ? "成功" : "失败") << endl;
// }

// void CMdSpi::OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin,
// 		CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
// {
// 	cout << "--->>> " << "OnRspUserLogin" << endl;
// 	if (bIsLast && !IsErrorRspInfo(pRspInfo))
// 	{
// 		///获取当前交易日
// 		cout << "--->>> 获取当前交易日 = " << pUserApi->GetTradingDay() << endl;
// 		// 请求订阅行情
// 		SubscribeMarketData();	
// 		// 请求订阅询价,只能订阅郑商所的询价，其他交易所通过traderapi相应接口返回
// 		SubscribeForQuoteRsp();	
// 	}
// }

// void CMdSpi::SubscribeMarketData()
// {
// 	int iResult = pUserApi->SubscribeMarketData(ppInstrumentID, iInstrumentID);
// 	cout << "--->>> 发送行情订阅请求: " << ((iResult == 0) ? "成功" : "失败") << endl;
// }

// void CMdSpi::SubscribeForQuoteRsp()
// {
// 	int iResult = pUserApi->SubscribeForQuoteRsp(ppInstrumentID, iInstrumentID);
// 	cout << "--->>> 发送询价订阅请求: " << ((iResult == 0) ? "成功" : "失败") << endl;
// }

// void CMdSpi::OnRspSubMarketData(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
// {
// 	cout << "OnRspSubMarketData" << endl;
// }

// void CMdSpi::OnRspSubForQuoteRsp(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
// {
// 	cout << "OnRspSubForQuoteRsp" << endl;
// }

// void CMdSpi::OnRspUnSubMarketData(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
// {
// 	cout << "OnRspUnSubMarketData" << endl;
// }

// void CMdSpi::OnRspUnSubForQuoteRsp(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
// {
// 	cout << "OnRspUnSubForQuoteRsp" << endl;
// }

// void CMdSpi::OnRtnDepthMarketData(CThostFtdcDepthMarketDataField *pDepthMarketData)
// {
// 	cout << "OnRtnDepthMarketData " << 
//         pDepthMarketData->TradingDay <<" "<< 
//         pDepthMarketData->InstrumentID <<" "<< 
//         pDepthMarketData->ExchangeID<<" "<< 
//         pDepthMarketData->ExchangeInstID <<" "<< 
//         pDepthMarketData->LastPrice <<" "<< 
//         pDepthMarketData->PreSettlementPrice<<" "<< 
//         pDepthMarketData->PreClosePrice <<" "<< 
//         pDepthMarketData->PreOpenInterest <<" "<< 
//         pDepthMarketData->OpenPrice<<" "<< 
//         endl;
// }

// void CMdSpi::OnRtnForQuoteRsp(CThostFtdcForQuoteRspField *pForQuoteRsp)
// {
// 	cout << "OnRtnForQuoteRsp" << endl;
// }

// bool CMdSpi::IsErrorRspInfo(CThostFtdcRspInfoField *pRspInfo)
// {
// 	// 如果ErrorID != 0, 说明收到了错误的响应
// 	bool bResult = ((pRspInfo) && (pRspInfo->ErrorID != 0));
// 	if (bResult)
// 		cout << "--->>> ErrorID=" << pRspInfo->ErrorID << ", ErrorMsg=" << pRspInfo->ErrorMsg << endl;
// 	return bResult;
// }