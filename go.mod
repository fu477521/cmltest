module git.pinquest.cn/qlb/cmltest

go 1.15

replace (
	github.com/coreos/go-systemd => git.pinquest.cn/qlb/extrpkg/github.com/coreos/go-systemd v0.0.0-20210302022218-b498ef3236d3
	go.etcd.io/bbolt => github.com/coreos/bbolt v1.3.5
	google.golang.org/grpc/naming => git.pinquest.cn/qlb/extrpkg/google.golang.org/grpc/naming v0.0.0-20210302015939-5eeb3473985c
)