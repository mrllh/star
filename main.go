package main

import (
	"errors"
	_ "star/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"star/internal/cmd"
)

func main() {
	// 设置默认语言为中文
	g.I18n().SetLanguage("zh-CN")

	// 连接数据库
	err := connDb()
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}

func connDb() error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("数据库连接失败")
	}
	return nil
}
