package model

type Receipt struct {
	//represents the receipt object.
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []Item `json:"items"`
	Total        string
}

type Item struct {
	//represents the Items in the receipt object.
	ShortDescription string
	Price            float64 `json:",string"`
}
