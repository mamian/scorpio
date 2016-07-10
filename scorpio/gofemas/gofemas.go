// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gofemas

//#cgo LDFLAGS: -fPIC -L. -L/workspace/QuantProject/qerio.scorpio.v0.3/deps/FemasAPI_V2.00/FemasNewMarker2.00.00API/lnx64 -Wl,-rpath=/workspace/QuantProject/qerio.scorpio.v0.3/deps/FemasAPI_V2.00/FemasNewMarker2.00.00API/lnx64 -lUSTPmduserapi -lUSTPtraderapi -lstdc++ -lssl
//#cgo CPPFLAGS: -fPIC -I. -I/workspace/QuantProject/qerio.scorpio.v0.3/deps/FemasAPI_V2.00/FemasNewMarker2.00.00API/include
import "C"

import (
	"fmt"
)

// type GoCallback struct{}

// func (p *GoCallback) Run() string {
// 	return "GoCallback.Run"
// }

type GoCUstpFtdcMdSpi struct{}

func (p *GoCUstpFtdcMdSpi) OnFrontConnected() {
	fmt.Println("GoCUstpFtdcMdSpi.OnFrontConnected.")
}

type GoCUstpFtdcTraderSpi struct{}

func (p *GoCUstpFtdcTraderSpi) OnFrontConnected() {
	fmt.Println("GoCUstpFtdcTraderSpi.OnFrontConnected.")
}

func (p *GoCUstpFtdcTraderSpi) OnHeartBeatWarning(nTimeLapse int) {
	fmt.Println("GoCUstpFtdcTraderSpi.OnHeartBeatWarning.")
}
