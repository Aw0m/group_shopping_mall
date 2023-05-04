package bdm

import (
	"time"

	"github.com/golang-jwt/jwt"

	"group_shopping_mall/utils/constant"
	"group_shopping_mall/utils/middleware"
)

// User 用户信息表
type User struct {
	UserId    int64  `gorm:"column:user_id;type:bigint(20);AUTO_INCREMENT;comment:用户自增id;primary_key" json:"user_id,string"`
	OpenId    string `gorm:"column:open_id;type:varchar(255);NOT NULL" json:"open_id"`
	Gender    string `gorm:"column:gender;type:enum('m','f','u');comment:用户性别，枚举值m/f/u 男/女/未知;NOT NULL" json:"gender"`
	Username  string `gorm:"column:username;type:varchar(16);comment:用户名，长度需不大于16;NOT NULL" json:"username"`
	AvatarUrl string `gorm:"column:avatar_url;type:varchar(255);comment:用户头像地址" json:"avatar_url"`
	PhoneNum  string `gorm:"column:phone_num;type:varchar(16);comment:电话号码" json:"phone_num"`
	AddressId int64  `gorm:"column:address_id;type:bigint(20);comment:用户自提点id" json:"address_id,string"`
}

func (u *User) CreateToken(key string) string {
	myToken := jwt.NewWithClaims(jwt.SigningMethodHS256, middleware.MyCustomClaims{
		OpenID:     u.OpenId,
		UserID:     u.UserId,
		UserName:   u.Username,
		SessionKey: key,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * constant.Duration).Unix(),
			NotBefore: time.Now().Unix(),
		},
	})
	ss, _ := myToken.SignedString(constant.MySigningKey)
	return ss
}
