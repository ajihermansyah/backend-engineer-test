package cache

import (
	"backend-engineer-test/app-fetch/config"
	"backend-engineer-test/app-fetch/repository"
	"backend-engineer-test/app-fetch/repository/currency"
	"time"
)

var (
	RateCurrencyValue float64
	CacheExpired      time.Time
)

type CacheController struct {
	CurrencyRepo repository.CurrencyRepositoryInterface
}

func NewCache() repository.CacheRepositoryInterface {
	return &CacheController{
		CurrencyRepo: currency.NewCurrencyRepository(),
	}
}

func (c *CacheController) SetCacheRateCurrency(rateCurrency float64) {
	if RateCurrencyValue == 0 {
		RateCurrencyValue = rateCurrency
		CacheExpired = time.Now().Add(time.Duration(config.CacheExpired) * time.Second)
	}

	c.ClearCacheRateCurrency(config.CacheExpired)
}

func (c *CacheController) ClearCacheRateCurrency(cacheDuration int) {
	now := time.Now()
	expiredTime := now.Add(time.Duration(cacheDuration) * time.Second)
	diff := expiredTime.Sub(CacheExpired)

	remainingExpired := int(diff / time.Second)
	if remainingExpired > cacheDuration {
		RateCurrencyValue = 0
		CacheExpired = time.Time{}
	}
}
