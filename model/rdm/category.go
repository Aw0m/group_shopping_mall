package rdm

// Category 商品分类信息表
type Category struct {
	CategoryId   int64  `gorm:"column:category_id;type:bigint(20);AUTO_INCREMENT;comment:商品分类的id;primary_key" json:"category_id,string"`
	CategoryName string `gorm:"column:category_name;type:varchar(16);comment:商品分类的名字" json:"category_name"`
	IsDeleted    bool   `gorm:"column:is_deleted;type:tinyint(1);default:0;comment:是否删除" json:"is_deleted"`
}

func (m *Category) TableName() string {
	return "category"
}
