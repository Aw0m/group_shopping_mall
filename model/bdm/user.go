package bdm

import (
	"github.com/golang-jwt/jwt"
	"group_shopping_mall/utils/constant"
	"group_shopping_mall/utils/middleware"
	"time"
)

type User struct {
	UserId      int64   `json:"user_id,string"`
	OpenId      string  `json:"open_id"`
	Gender      string  `json:"gender"`
	Username    string  `json:"username"`
	ClubId      int64   `json:"club_id,string"`
	Intro       string  `json:"intro"`
	AdminType   int     `json:"admin_type"`
	UserType    int     `json:"user_type"`
	Price       float64 `json:"price"`
	Commissions float64 `json:"commissions"`
	AvatarURL   string  `json:"avatar_url"`
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
