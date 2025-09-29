package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta `path:"/register" method:"post" tags:"用户" summary:"用户注册"`
}
