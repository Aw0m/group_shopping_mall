package bdm

type Category struct {
	CategoryId   int64  `json:"category_id,string"`
	CategoryName string `json:"category_name"`
	IsDeleted    bool   `json:"is_deleted"`
}
