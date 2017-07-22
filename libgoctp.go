// +build cgo

// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goctp

/*
#cgo linux LDFLAGS: -fPIC -L. -L$GOPATH/src/github.com/qerio/goctp/api/ThostTraderApi_v6.3.6_20160606_linux64 -Wl,-rpath=$GOPATH/src/github.com/qerio/goctp/api/ThostTraderApi_v6.3.6_20160606_linux64 -lthostmduserapi -lthosttraderapi -lstdc++
#cgo linux CPPFLAGS: -fPIC -I. -I$GOPATH/src/github.com/qerio/goctp/api/ThostTraderApi_v6.3.6_20160606_linux64

#cgo windows LDFLAGS: -lmsvcrt -lm -mthread
#cgo windows LDFLAGS: -fPIC -L. -L$GOPATH/src/github.com/qerio/goctp/api/ThostTraderApi_v6.3.6_20160606_win64 -lthostmduserapi -lthosttraderapi
#cgo windows CPPFLAGS: -fPIC -I. -I$GOPATH/src/github.com/qerio/goctp/api/ThostTraderApi_v6.3.6_20160606_win64 -D_WIN64 -DISLIB -DLIB_MD_API_EXPORT -DLIB_TRADER_API_EXPORT

#cgo CFLAGS: -Wall -Werror
*/

import "C"
