package main

import (
	"fmt"
	"reflect"

	"git.pinquest.cn/qlb/brick/log"
	"git.pinquest.cn/qlb/brick/rpc"
	"github.com/fu477521/cmltest"
)

func main() {
	type args struct {
		ctx *rpc.Context
		req *GetPermissionReq
	}
	tests := []struct {
		name    string
		args    args
		want    *GetPermissionReq
		wantErr bool
	}{
		{
			name:    "",
			args:    args{},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := cmltest.GetPermission(tt.args.ctx, tt.args.req)
		fmt.Println(got)
		if (err != nil) != tt.wantErr {
			log.Errorf("AddFriendsTpl() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			log.Errorf("AddFriendsTpl() got = %v, want %v", got, tt.want)
		}

	}
}
