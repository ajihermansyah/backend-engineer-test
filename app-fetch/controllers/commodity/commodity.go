package commodity

import (
	httpHelper "backend-engineer-test/app-fetch/helper/http"
	"backend-engineer-test/app-fetch/repository"

	"github.com/gin-gonic/gin"
)

type CommodityController struct {
	Helper        httpHelper.HTTPHelper
	CommodityRepo repository.CommodityRepositoryInterface
	TokenRepo     repository.TokenRepositoryInterface
}

func (c *CommodityController) GetListCommodityController(ctx *gin.Context) {
	commodities, err := c.CommodityRepo.GetListCommodity()
	if err != nil {
		c.Helper.SendError(ctx, 500, err.Error(), "error", c.Helper.EmptyJsonMap())
		return
	}

	c.Helper.SendSuccess(ctx, "Get list data commodity successfully", commodities)
}

func (c *CommodityController) GetCommodityAggregateController(ctx *gin.Context) {
	commodities, err := c.CommodityRepo.GetCommodityAggregate()
	if err != nil {
		c.Helper.SendError(ctx, 500, "This API can access for role admin only", "error", c.Helper.EmptyJsonMap())
		return
	}

	c.Helper.SendSuccess(ctx, "Get list data commodity aggregate successfully", commodities)
}
