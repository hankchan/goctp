// +build linux,cgo windows,cgo

// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goctp

/*
#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/api/SFIT_CTP_6.3.11_20180109_tradeapi/v6.3.11_20180109_api_tradeapi_linux64 -Wl,-rpath,${SRCDIR}/api/SFIT_CTP_6.3.11_20180109_tradeapi/v6.3.11_20180109_api_tradeapi_linux64 -lthostmduserapi -lthosttraderapi -lstdc++
#cgo linux CPPFLAGS: -fPIC -I${SRCDIR}/api/SFIT_CTP_6.3.11_20180109_tradeapi/v6.3.11_20180109_api_tradeapi_linux64
#cgo windows LDFLAGS: -fPIC -L${SRCDIR}/api/SFIT_CTP_6.3.11_20180109_tradeapi/20180109_tradeapi64_windows -Wl,-rpath,${SRCDIR}/api/SFIT_CTP_6.3.11_20180109_tradeapi/20180109_tradeapi64_windows ${SRCDIR}/api/SFIT_CTP_6.3.11_20180109_tradeapi/20180109_tradeapi64_windows/thostmduserapi.lib ${SRCDIR}/api/SFIT_CTP_6.3.11_20180109_tradeapi/20180109_tradeapi64_windows/thosttraderapi.lib -lthostmduserapi -lthosttraderapi
#cgo windows CPPFLAGS: -fPIC -I${SRCDIR}/api/SFIT_CTP_6.3.11_20180109_tradeapi/20180109_tradeapi64_windows -DISLIB -DWIN32 -DLIB_MD_API_EXPORT -DLIB_TRADER_API_EXPORT
*/
import "C"
