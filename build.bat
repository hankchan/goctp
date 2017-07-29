go install -v -x -a -work runtime sync/atomic  #构建核心基本库

go build -v -x -work #构建GO动态库
go install

go build -v -x _example/goctp_md_example.go
go build -v -x _example/goctp_trader_example.go
