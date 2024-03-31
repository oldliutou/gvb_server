package models
import (
	"gvb_server/models/ctype"
)


// AuthModel用户表
type AuthModel struct {
	MODEL
	Nickname   string           `gorm:"size:36" json:"nickname"`             // 昵称
	Username   string           `gorm:"size:36" json:"username"`             // 用户名
	Password   string           `gorm:"size:128" json:"password"`            //  密码
	Avatar     string           `gorm:"size:256" json:"avatar"`               // 头像
	Email      string           `gorm:"size:128" json:"email"`               // 邮箱
	Tel        string           `gorm:"size:18" json:"tel"`                  // 手机号
	Addr       string           `gorm:"size:64" json:"addr"`                 // 地址
	Token      string           `gorm:"size:64" json:"token"`                // 其他平台的唯一id
	IP         string           `gorm:"size:20" json:"ip"`                   // ip地址
	Role       ctype.Role       `gorm:"size:4;default:1" json:"role"`        // 权限  1 管理员  2 普通用户  3 游客
	SignStatus ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"` // 注册来源
	// Integral   int              `gorm:"default:0" json:"integral"`           // 积分
	// Sign       string           `gorm:"size:128" json:"sign"`                // 签名
	// Link       string           `gorm:"size:128" json:"link"`                // 链接地址
	ArticleModels	[]ArticleModel	`gorm:"foreignKey:AuthID" json:"-"` // 发布的文章列表
	CollectModels	[]ArticleModel	`gorm:"many2many:auth2_collects;joinForeignKey:AuthID;JoinReferences:ArticleID" json:"-"` // 收藏的文章列表

}
