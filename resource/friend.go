package resource

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/wechat-helper/action"
	"github.com/quarkcms/wechat-helper/model"
)

type Friend struct {
	adminresource.Template
}

// 初始化
func (p *Friend) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "好友"

	// 模型
	p.Model = &model.Friend{}

	// 分页
	p.PerPage = 10

	return p
}

func (p *Friend) Fields(ctx *builder.Context) []interface{} {
	field := &adminresource.Field{}

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("wechat_id", "微信ID"),

		field.Image("avatar", "头像"),

		field.Text("nick_name", "昵称"),
	}
}

// 搜索
func (p *Friend) Searches(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&searches.Input{}).Init("nick_name", "好友昵称"),
		(&searches.Status{}).Init(),
	}
}

// 行为
func (p *Friend) Actions(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&action.Sync{}).Init("同步数据"),
		(&actions.Delete{}).Init("批量删除"),
		(&actions.Disable{}).Init("批量禁用"),
		(&actions.Enable{}).Init("批量启用"),
		(&actions.Delete{}).Init("删除"),
	}
}
