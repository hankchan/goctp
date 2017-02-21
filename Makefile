# Copyright (c) 2013-2016 by Hank Chan. All Rights Reserved.
# Use of this source code is governed by a MIT-style license that can
# be found in the LICENSE file.

build:
	go install -v -x -a -buildmode=shared runtime sync/atomic #构建核心基本库
	go install -v -x -a -buildmode=shared -linkshared #构建GO动态库

example:
	go build -v -x -linkshared _example/goctp_md_example.go
	go build -v -x -linkshared _example/goctp_trader_example.go
