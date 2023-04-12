package resource

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/wechat-helper/model"
)

type Group struct {
	adminresource.Template
}

// 初始化
func (p *Group) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "群组"

	// 模型
	p.Model = &model.Group{}

	// 分页
	p.PerPage = 10

	return p
}

func (p *Group) Fields(ctx *builder.Context) []interface{} {
	field := &adminresource.Field{}

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("name", "昵称"),

		field.Switch("status", "状态").
			SetTrueValue("正常").
			SetFalseValue("禁用").
			SetEditable(true).
			SetDefault(true),
	}
}

// 搜索
func (p *Group) Searches(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&searches.Input{}).Init("name", "好友昵称"),
		(&searches.Status{}).Init(),
	}
}

// 行为
func (p *Group) Actions(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&actions.Delete{}).Init("批量删除"),
		(&actions.Disable{}).Init("批量禁用"),
		(&actions.Enable{}).Init("批量启用"),
		(&actions.Delete{}).Init("删除"),
	}
}
