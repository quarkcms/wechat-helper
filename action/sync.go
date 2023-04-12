package action

import (
	"os"
	"time"

	"github.com/eatmoreapple/openwechat"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource/actions"
	"github.com/quarkcms/quark-go/pkg/rand"
	"github.com/quarkcms/wechat-helper/model"
	"gorm.io/gorm"
)

type Sync struct {
	actions.Action
}

// 初始化
func (p *Sync) Init(name string) *Sync {

	// 初始化父结构
	p.ParentInit()

	// 行为名称，当行为在表格行展示时，支持js表达式
	p.Name = name

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "primary"

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 是否具有loading，当action 的作用类型为ajax,submit时有效
	p.WithLoading = true

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	// 行为类型
	p.ActionType = "ajax"

	// 当行为在表格行展示时，支持js表达式
	p.WithConfirm("确定要同步吗？", "同步数据需要您扫码登录！", "modal")

	return p
}

// 执行行为句柄
func (p *Sync) Handle(ctx *builder.Context, query *gorm.DB) interface{} {

	// 桌面模式
	bot := openwechat.DefaultBot(openwechat.Desktop)

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		return ctx.SimpleError(err.Error())
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		return ctx.SimpleError(err.Error())
	}

	// 获取所有的好友
	friends, err := self.Friends()
	if err != nil {
		return ctx.SimpleError(err.Error())
	}
	for _, friend := range friends {

		// 已存在跳出本次循环
		if (&model.Friend{}).IsExist(friend.ID()) {
			continue
		}

		fileName := rand.MakeAlphanumeric(40) + ".png"

		// url
		fileUrl := "/storage/images/" + time.Now().Format("20060102") + "/"

		// 设置保存路径
		savePath := "./website" + fileUrl

		if !p.isExist(savePath) {
			err := os.MkdirAll(savePath, os.ModeDir)
			if err != nil {
				return err
			}
		}
		friendData := &model.Friend{
			WechatId: friend.ID(),
			NickName: friend.NickName,
		}

		// 保存头像到本地
		err := friend.SaveAvatar(savePath + fileName)
		if err == nil {
			friendData.Avatar = fileUrl + fileName
		}

		// 入库
		(&model.Friend{}).Insert(friendData)
	}

	// 获取所有的群组
	groups, err := self.Groups()
	if err != nil {
		return ctx.SimpleError(err.Error())
	}
	for _, group := range groups {

		// 已存在跳出本次循环
		if (&model.Group{}).IsExist(group.ID()) {
			continue
		}

		fileName := rand.MakeAlphanumeric(40) + ".png"

		// url
		fileUrl := "/storage/images/" + time.Now().Format("20060102") + "/"

		// 设置保存路径
		savePath := "./website" + fileUrl

		if !p.isExist(savePath) {
			err := os.MkdirAll(savePath, os.ModeDir)
			if err != nil {
				return err
			}
		}

		groupData := &model.Group{
			WechatId: group.ID(),
			Name:     group.NickName,
		}

		// 保存头像到本地
		err := group.SaveAvatar(savePath + fileName)
		if err == nil {
			groupData.Cover = fileUrl + fileName
		}

		// 入库
		(&model.Group{}).Insert(groupData)
	}

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	// bot.Block()

	return ctx.SimpleSuccess("操作成功")
}

// 检查路径是否存在
func (p *Sync) isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
