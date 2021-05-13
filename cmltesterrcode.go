package cmltest

import "git.pinquest.cn/qlb/brick/rpc"

const (
	ErrPermissionNotFound	=	347001
	ErrRoleNotFound	=	347002
	ErrUserRoleNotFound	=	347003
	ErrRoleResourceNotFound	=	347004

)
var (
	errCodeMap = map[int32]string{
	ErrPermissionNotFound	: "ErrPermissionNotFound",
	ErrRoleNotFound	: "ErrRoleNotFound",
	ErrUserRoleNotFound	: "ErrUserRoleNotFound",
	ErrRoleResourceNotFound	: "ErrRoleResourceNotFound",

	}
)

func RegisterError() {
	rpc.RegisterError(errCodeMap)
}
