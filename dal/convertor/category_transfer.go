package convertor

import (
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
)

func CategoryBdmToRdm(c bdm.Category) rdm.Category {
	return rdm.Category{
		CategoryId:   c.CategoryId,
		CategoryName: c.CategoryName,
		IsDeleted:    c.IsDeleted,
	}
}

func CategoryRdmToBdm(c rdm.Category) bdm.Category {
	return bdm.Category{
		CategoryId:   c.CategoryId,
		CategoryName: c.CategoryName,
		IsDeleted:    c.IsDeleted,
	}
}
