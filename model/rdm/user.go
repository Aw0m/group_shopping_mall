package rdm

// User 用户信息表
type User struct {
	UserId      int64   `gorm:"column:user_id;type:bigint(20);primary_key;AUTO_INCREMENT;comment:用户自增id" json:"user_id"`
	OpenId      string  `gorm:"column:open_id;type:varchar(255);NOT NULL" json:"open_id"`
	Gender      string  `gorm:"column:gender;type:enum('m','f','u');comment:用户性别，枚举值m/f/u 男/女/未知;NOT NULL" json:"gender"`
	Username    string  `gorm:"column:username;type:varchar(16);comment:用户名，长度需不大于16;NOT NULL" json:"username"`
	ClubId      int64   `gorm:"column:club_id;type:bigint(20);comment:所属俱乐部id" json:"club_id"`
	Intro       string  `gorm:"column:intro;type:varchar(255);comment:个人简介" json:"intro"`
	AdminType   int     `gorm:"column:admin_type;type:int(11);default:0;comment:管理员类型，非管理员1，普通管理员2，董事长3;NOT NULL" json:"admin_type"`
	UserType    int     `gorm:"column:user_type;type:int(11);default:0;comment:陪玩类型, 娱乐1、技术2等;NOT NULL" json:"user_type"`
	Price       float64 `gorm:"column:price;type:double;default:0;comment:陪玩的单价;NOT NULL" json:"price"`
	Commissions float64 `gorm:"column:commissions;type:double;default:0;comment:抽成比例;NOT NULL" json:"commissions"`
	AvatarURL   string  `gorm:"column:avatar_url;type:varchar(255);comment:用户头像地址" json:"avatar_url"`
}

func (m *User) TableName() string {
	return "user"
}
