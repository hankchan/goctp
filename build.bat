go install -v -x -a -work runtime sync/atomic  #构建核心基本库
go install -v -x -a -work -tags no_pkgconfig #构建GO动态库
go build -v -x -linkshared _example/goctp_md_example.go
go build -v -x -linkshared _example/goctp_trader_example.go
