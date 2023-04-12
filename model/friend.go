package model

import (
	"time"

	appmodel "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

// 好友模型
type Friend struct {
	Id        int            `json:"id" gorm:"autoIncrement"`
	WechatId  string         `json:"wechat_id" gorm:"autoIncrement"`
	NickName  string         `json:"nick_name" gorm:"size:200;not null"`
	Avatar    string         `json:"avatar" gorm:"size:200;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Seeder
func (m *Friend) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if (&appmodel.Menu{}).IsExist(18) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 18, Name: "微信管理", GuardName: "admin", Icon: "icon-comment", Type: "default", Pid: 0, Sort: 0, Path: "/wechat", Show: 1, Status: 1},
		{Id: 19, Name: "好友列表", GuardName: "admin", Icon: "", Type: "engine", Pid: 18, Sort: 0, Path: "/api/admin/friend/index", Show: 1, Status: 1},
	}
	db.Client.Create(&menuSeeders)
}

// 插入数据
func (m *Friend) Insert(friend *Friend) {
	db.Client.Create(&friend)
}

// 获取信息
func (m *Friend) GetInfoByWechatId(wechatId string) *Friend {
	var friend = &Friend{}
	db.Client.Where("wechat_id = ?", wechatId).First(friend)

	return friend
}
