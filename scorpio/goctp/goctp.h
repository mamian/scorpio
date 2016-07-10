#pragma once

// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifdef _WIN32
#include "../../deps/ThostTraderApi_v6.3.5_tradeapi_win64/ThostFtdcUserApiDataType.h" 
#include "../../deps/ThostTraderApi_v6.3.5_tradeapi_win64/ThostFtdcUserApiStruct.h"
#include "../../deps/ThostTraderApi_v6.3.5_tradeapi_win64/ThostFtdcMdApi.h"
#include "../../deps/ThostTraderApi_v6.3.5_tradeapi_win64/ThostFtdcTraderApi.h"
#else
#include "../../deps/ThostTraderApi_v6.3.5_tradeapi_linux64/ThostFtdcUserApiDataType.h" 
#include "../../deps/ThostTraderApi_v6.3.5_tradeapi_linux64/ThostFtdcUserApiStruct.h"
#include "../../deps/ThostTraderApi_v6.3.5_tradeapi_linux64/ThostFtdcMdApi.h"
#include "../../deps/ThostTraderApi_v6.3.5_tradeapi_linux64/ThostFtdcTraderApi.h"
#endif

class Callback {
public:
	virtual ~Callback() { }
	virtual std::string run() { return "Callback::run"; }
};

class Caller {
private:
	Callback *callback_;
public:
	Caller(): callback_(0) { }
	~Caller() { delCallback(); }
	void delCallback() { delete callback_; callback_ = 0; }
	void setCallback(Callback *cb) { delCallback(); callback_ = cb; }
	std::string call();
};
