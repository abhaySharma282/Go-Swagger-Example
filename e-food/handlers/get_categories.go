package handlers

import (
	"e-food/models"
	"e-food/restapi/operations/menu"
	"github.com/go-openapi/runtime/middleware"
)

type menuListImpl struct{}

func NewMenuCategoryListHandler() menu.CategoryListHandler {
	return &menuListImpl{}
}

func (impl *menuListImpl) Handle(params menu.CategoryListParams) middleware.Responder {
	responseVal := &models.Categories{
		&models.Category{
			BcID:       2001,
			BcName:     "Fruits",
			BcIsActive: true,
			BcImageURL: "",
			SubCategories: []*models.SubCategory{
				{
					ScID:       1002,
					ScName:     "Apple",
					ScImageURL: "",
					ScIsActive: true,
				},
			},
		},
	}
	return menu.NewCategoryListOK().WithPayload(*responseVal)
}
