package bdm

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
)

// Commodity 商品信息表
type Commodity struct {
	CommodityId     int64          `json:"commodity_id,string"`
	CommodityName   string         `json:"commodity_name"`
	Price           float64        `json:"price"`
	SellerId        int64          `json:"seller_id,string"`
	CategoryId      int64          `json:"category_id,string"`
	Content         datatypes.JSON `json:"content"`
	AddressList     string         `json:"address_list"`
	CommodityStatus int            `json:"commodity_status,string"`
	ImageURL        string         `gorm:"column:image_url;type:varchar(255);comment:商品展示图片" json:"image_url"`
	IsDeleted       bool           `json:"is_deleted"`
}

type CommodityContent struct {
	Intro string `json:"intro"`
}

func (c *Commodity) GetContent() (CommodityContent, error) {
	// 将datatypes.JSON转换为自定义的结构体CommodityContent
	if len(c.Content) == 0 {
		return CommodityContent{}, nil
	}
	var content CommodityContent
	err := jsoniter.Unmarshal(c.Content, &content)
	if err != nil {
		err = errors.Errorf("content unmarshal from datatypes.JSON fail!, errMsg: %s", err.Error())
	}
	return content, err
}

func (c *Commodity) SetContent(content CommodityContent) error {
	// 将自定义的结构体CommodityContent转换为datatypes.JSON
	contentByte, err := jsoniter.Marshal(content)
	if err != nil {
		err = errors.Errorf("content marshal to datatypes.JSON fail!, errMsg: %s", err.Error())
		return err
	}
	c.Content = contentByte
	return nil
}
