package cmltest

import "git.pinquest.cn/qlb/brick/rpc"

const (

)
var (
	errCodeMap = map[int32]string{

	}
)

func RegisterError() {
	rpc.RegisterError(errCodeMap)
}
