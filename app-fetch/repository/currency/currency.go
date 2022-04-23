package currency

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"backend-engineer-test/app-fetch/config"
	"backend-engineer-test/app-fetch/model"
	"backend-engineer-test/app-fetch/repository"
)

type CurrencyRepository struct {
}

func NewCurrencyRepository() repository.CurrencyRepositoryInterface {
	return &CurrencyRepository{}
}
func (repo *CurrencyRepository) GetConversionRateCurrency(ratio string) (float64, error) {
	var rate model.CurrencyRate

	fmt.Println("Getting conversion rate for", ratio)

	url := fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=%s&compact=ultra&apiKey=%s", ratio, config.Key)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to get conversion rate :", err)
		err = errors.New("Failed to get conversion rate")
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &rate)
	if err != nil {
		fmt.Println("Failed to parse conversion rate data :", err)
		err = errors.New("Failed to parse conversion rate data")
		return 0, err
	} else if rate.Ratio == 0 {
		fmt.Println("Conversion Rate is 0")
		err = errors.New("Failed to get conversion rate data")
		return 0, err
	}

	return rate.Ratio, nil
}
