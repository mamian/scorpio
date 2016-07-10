// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goctp

//#cgo LDFLAGS: -fPIC -L. -L/workspace/GoProject/src/qerio/deps/ThostTraderApi_v6.3.5_tradeapi_linux64 -Wl,-rpath=/workspace/GoProject/src/qerio/deps/ThostTraderApi_v6.3.5_tradeapi_linux64 -lthostmduserapi -lthosttraderapi -lstdc++
//#cgo CPPFLAGS: -fPIC -I. -I/workspace/GoProject/src/qerio/deps/ThostTraderApi_v6.3.5_tradeapi_linux64
import "C"

import (
	"fmt"
)

type GoCallback struct{}

func (p *GoCallback) Run() string {
	return "GoCallback.Run"
}

type GoCThostFtdcMdSpi struct{}

func (p *GoCThostFtdcMdSpi) OnFrontConnected() {
	fmt.Println("GoCThostFtdcMdSpi.OnFrontConnected.")
}

func (p *GoCThostFtdcMdSpi) OnFrontConnected() {
	fmt.Println("GoCThostFtdcMdSpi.OnFrontConnected.")
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
