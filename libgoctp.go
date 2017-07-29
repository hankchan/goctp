// +build cgo

// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goctp

// #cgo linux LDFLAGS: -fPIC -L${SRCDIR}/api/ThostTraderApi_v6.3.6/linux64 -Wl,-rpath=${SRCDIR}/api/ThostTraderApi_v6.3.6/linux64 -lthostmduserapi -lthosttraderapi -lstdc++
// #cgo linux CPPFLAGS: -fPIC -I${SRCDIR}/api/ThostTraderApi_v6.3.6/include
// #cgo windows LDFLAGS: -fPIC -L${SRCDIR}/api/ThostTraderApi_v6.3.6/win64 -Wl,-rpath=${SRCDIR}/api/ThostTraderApi_v6.3.6/win64 ${SRCDIR}/api/ThostTraderApi_v6.3.6/win64/thostmduserapi.lib ${SRCDIR}/api/ThostTraderApi_v6.3.6/win64/thosttraderapi.lib -lthostmduserapi -lthosttraderapi
// #cgo windows CPPFLAGS: -fPIC -I${SRCDIR}/api/ThostTraderApi_v6.3.6/include -DISLIB -DWIN32 -DLIB_MD_API_EXPORT -DLIB_TRADER_API_EXPORT
import "C"
