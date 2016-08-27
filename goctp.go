// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goctp

//#cgo windows LDFLAGS: -fPIC -L. -L D:/cygwin64/lib/w32api -L D:/cygwin64/lib/gcc/x86_64-pc-cygwin/5.4.0 -L ../../api/ThostTraderApi_v6.3.6_tradeapi_win64 ../../api/ThostTraderApi_v6.3.6_tradeapi_win64/thostmduserapi.lib ../../api/ThostTraderApi_v6.3.6_tradeapi_win64/thosttraderapi.lib -lthostmduserapi -lthosttraderapi
//#cgo windows CPPFLAGS: -fPIC -I. -I../../api/ThostTraderApi_v6.3.6_tradeapi_win64 -DWIN32 -DISLIB -DLIB_MD_API_EXPORT -DLIB_TRADER_API_EXPORT
//#cgo linux LDFLAGS: -fPIC -L. -L/workspace/GoProject/src/github.com/qerio/goctp/api/ThostTraderApi_v6.3.6_20160606_linux64 -Wl,-rpath=/workspace/GoProject/src/github.com/qerio/goctp/api/ThostTraderApi_v6.3.6_20160606_linux64 -lthostmduserapi -lthosttraderapi -lstdc++
//#cgo linux CPPFLAGS: -fPIC -I. -I/workspace/GoProject/src/github.com/qerio/goctp/api/ThostTraderApi_v6.3.6_20160606_linux64
import "C"
