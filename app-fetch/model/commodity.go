package model

type Commodity struct {
	ID        string `json:"uuid"`
	Commodity string `json:"komoditas"`
	Province  string `json:"area_provinsi"`
	City      string `json:"area_kota"`
	Size      string `json:"size"`
	Price     string `json:"price"`
	Date      string `json:"tgl_parsed"`
	Timestamp string `json:"timestamp"`
	PriceUSD  string `json:"price_usd"`
}

type DataCommodity struct {
	ProvinceArea  string `json:"area_provinsi"`
	Profit        map[string]map[string]int
	MaximumProfit float64 `json:"max_profit"`
	MinimumProfit float64 `json:"min_profit"`
	AvgProfit     float64 `json:"average_profit"`
	MedianProfit  float64 `json:"median_profit"`
}
