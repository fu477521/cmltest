module github.com/fu477521/cmltest

go 1.15

replace (
	github.com/coreos/go-systemd => git.pinquest.cn/qlb/extrpkg/github.com/coreos/go-systemd v0.0.0-20210302022218-b498ef3236d3
	go.etcd.io/bbolt => github.com/coreos/bbolt v1.3.5
	google.golang.org/grpc/naming => git.pinquest.cn/qlb/extrpkg/google.golang.org/grpc/naming v0.0.0-20210302015939-5eeb3473985c
)

require (
	git.pinquest.cn/qlb/brick v0.0.0-20210512090112-f216ab7526e1
	git.pinquest.cn/qlb/core v0.0.0-20210509035748-1c97f8c652ec
	git.pinquest.cn/qlb/quan v0.0.0-20210511163729-fd37729fbc50
	git.pinquest.cn/qlb/schedulerv2 v0.0.0-20210423121255-b07bbb048155 // indirect
	git.pinquest.cn/qlb/wework v0.0.0-20210511070459-af61eec0483d // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/klauspost/compress v1.12.2 // indirect
	github.com/prometheus/common v0.24.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/sys v0.0.0-20210511113859-b0526f3d8744 // indirect
	google.golang.org/genproto v0.0.0-20210510173355-fb37daa5cd7a // indirect
	google.golang.org/protobuf v1.26.0
)
