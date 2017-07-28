// +build cgo

// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goctp

// #cgo linux LDFLAGS: -fPIC -L${SRCDIR}/api/ThostTraderApi_v6.3.6_20160606_linux64 -Wl,-rpath=${SRCDIR}/api/ThostTraderApi_v6.3.6_20160606_linux64 -lthostmduserapi -lthosttraderapi -lstdc++
// #cgo linux CPPFLAGS: -fPIC -I${SRCDIR}/api/ThostTraderApi_v6.3.6_20160606_linux64
// #cgo windows LDFLAGS: -fPIC -L${SRCDIR}/api/ThostTraderApi_v6.3.6_20160606_win64 -Wl,-rpath=${SRCDIR}/api/ThostTraderApi_v6.3.6_20160606_win64 ${SRCDIR}/api/ThostTraderApi_v6.3.6_20160606_win64/thostmduserapi.lib ${SRCDIR}/api/ThostTraderApi_v6.3.6_20160606_win64/thosttraderapi.lib -lthostmduserapi -lthosttraderapi -lstdc++
// #cgo windows CPPFLAGS: -fPIC -I${SRCDIR}/api/ThostTraderApi_v6.3.6_20160606_win64
import "C"