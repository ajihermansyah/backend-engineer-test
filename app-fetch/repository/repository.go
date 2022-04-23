package repository

import "backend-engineer-test/app-fetch/model"

type CurrencyRepositoryInterface interface {
	GetConversionRateCurrency(ratio string) (float64, error)
}

type TokenRepositoryInterface interface {
	ParseToken(token string) (tokenClaim model.TokenClaims, err error)
}

type CommodityRepositoryInterface interface {
	FetchingDataCommodity() (res []model.Commodity, err error)
	GetListCommodity() (listCommodity []model.Commodity, err error)
	GetCommodityAggregate() (res []model.DataCommodity, err error)
}
