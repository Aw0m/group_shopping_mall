package rdm

// User 用户信息表
type User struct {
	UserId    int64  `json:"user_id"`    // 用户自增id
	OpenId    string `json:"open_id"`    // 用户唯一标识
	Gender    string `json:"gender"`     // 用户性别，枚举值m/f/u 男/女/未知
	Username  string `json:"username"`   // 用户名，长度需不大于16
	AvatarUrl string `json:"avatar_url"` // 用户头像地址
	PhoneNum  string `json:"phone_num"`  // 电话号码
	AddressId int    `json:"address_id"` // 用户自提点id
}

func (m *User) TableName() string {
	return "user"
}
