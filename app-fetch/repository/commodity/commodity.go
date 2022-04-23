package commodity

import (
	"backend-engineer-test/app-fetch/helper"
	"backend-engineer-test/app-fetch/model"
	"backend-engineer-test/app-fetch/repository"
	currencyRepo "backend-engineer-test/app-fetch/repository/currency"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type tempData struct {
	province string
	amount   int
	year     string
	week     string
}

type CommodityRepository struct {
}

func NewCommodityRepository() repository.CommodityRepositoryInterface {
	return &CommodityRepository{}
}
func (repo *CommodityRepository) FetchingDataCommodity() (res []model.Commodity, err error) {
	resp, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		fmt.Println("Failed get list data commodity :", err)
		err = errors.New("Failed to get list data commodity from site")
		return res, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Println("Failed to parse commodities list data :", err)
		err = errors.New("Failed to parse commodities list data")
		return res, err
	}

	return res, nil
}

func (repo *CommodityRepository) GetListCommodity() (listCommodity []model.Commodity, err error) {
	commodities, err := repo.FetchingDataCommodity()
	if err != nil || commodities == nil {
		fmt.Println("Failed to get list data commodity :", err)
		err = errors.New("Failed to get list data commodity")
		return listCommodity, err
	}

	listCommodity, err = repo.AddUSDPrice(commodities)
	if err != nil || listCommodity == nil {
		fmt.Println("Failed to get USD price :", err)
		err = errors.New("Failed to get USD price for list data commodities")
		return listCommodity, err
	}

	return listCommodity, nil
}

func (repo *CommodityRepository) AddUSDPrice(listData []model.Commodity) (res []model.Commodity, err error) {
	var rate float64

	rate, err = currencyRepo.NewCurrencyRepository().GetConversionRateCurrency("IDR_USD")
	if err != nil {
		fmt.Println("Failed to get conversion rate :", err)
		err = errors.New("Failed to get conversion rate currency")
		return res, err
	}

	for _, val := range listData {
		if val.Price == "" {
			continue
		}
		priceFloat, err := strconv.ParseFloat(val.Price, 64)
		if err != nil {
			fmt.Println(err, "Failed to convert price to float for uuid:", val.ID)
		}

		dollar := priceFloat * rate

		val.PriceUSD = fmt.Sprintf("$%f", dollar)
		res = append(res, val)
	}

	return res, nil
}

func (repo *CommodityRepository) GetCommodityAggregate() (res []model.DataCommodity, err error) {
	commodities, err := repo.FetchingDataCommodity()
	if err != nil || commodities == nil {
		fmt.Println("Failed to get list data commodity :", err)
		err = errors.New("Failed to get list data commodity")
		return res, err
	}

	// Parse the only needed data
	var data []tempData

	for _, commodity := range commodities {
		if commodity.Date == "" || commodity.Price == "" || commodity.Size == "" || commodity.Province == "" {
			continue
		}
		dateTime := helper.ConvertTimeStringToTime(commodity.Date)
		year, week := dateTime.ISOWeek()

		price, _ := strconv.Atoi(commodity.Price)
		size, _ := strconv.Atoi(commodity.Size)
		amount := price * size

		temp := tempData{
			province: commodity.Province,
			amount:   amount,
			year:     fmt.Sprintf("Tahun %s", strconv.Itoa(year)),
			week:     fmt.Sprintf("Minggu ke %s", strconv.Itoa(week)),
		}

		data = append(data, temp)
	}

	// Compiling data
	tempMap := make(map[string]map[string]map[string]int)

	for _, val := range data {
		// Create data set if data province is not exist yet
		if prov, ok := tempMap[val.province]; !ok {
			week := map[string]int{val.week: val.amount}
			year := map[string]map[string]int{val.year: week}
			tempMap[val.province] = year
		} else {
			// Create data set if data province is not exist yet
			if year, ok := prov[val.year]; !ok {
				week := map[string]int{val.week: val.amount}
				prov[val.year] = week
			} else {
				// Create data set if data profit is not exist yet
				// Add the amount if data profit is exist already
				if profit, ok := year[val.week]; !ok {
					year[val.week] = val.amount
				} else {
					year[val.week] += profit
				}
			}
		}
	}

	var result []model.DataCommodity

	for key, val := range tempMap {

		data := model.DataCommodity{
			ProvinceArea:  key,
			Profit:        val,
			MaximumProfit: repo.FindMaxProfit(val),
			MinimumProfit: repo.FindMinProfit(val),
			AvgProfit:     repo.FindAvgProfit(val),
			MedianProfit:  repo.FindMedianProfit(val),
		}

		result = append(result, data)
	}

	return result, nil
}

func (repo *CommodityRepository) FindMaxProfit(data map[string]map[string]int) float64 {
	var max int
	for _, val := range data {
		for _, amount := range val {
			if amount >= max {
				max = amount
			}
		}
	}

	return float64(max)
}

func (repo *CommodityRepository) FindMinProfit(data map[string]map[string]int) float64 {
	min := int(^uint(0) >> 1)
	for _, val := range data {
		for _, amount := range val {
			if amount <= min {
				min = amount
			}
		}
	}

	return float64(min)
}

func (repo *CommodityRepository) FindAvgProfit(data map[string]map[string]int) float64 {
	var sum, counter int
	for _, val := range data {
		for _, amount := range val {
			sum += amount
			counter++
		}
	}

	return float64(sum / counter)
}

func (repo *CommodityRepository) FindMedianProfit(data map[string]map[string]int) float64 {
	var arr []int
	for _, val := range data {
		for _, amount := range val {
			arr = append(arr, amount)
		}
	}

	counter := len(arr)

	if counter+1%2 == 0 {
		a := arr[(counter / 2)]
		b := arr[(counter/2)+1]
		return float64((a + b) / 2)
	} else {
		return float64(arr[counter/2])
	}
}
