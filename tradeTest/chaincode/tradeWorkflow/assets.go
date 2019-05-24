package main

type TradeAgreement struct {
	Amount             int    `json:"amount"`
	DescriptionOfGoods string `json:"descriptionOfGoods"`
	Status             string `json:"status"`
	Payment            int    `json:"payment"`
}

type LetterOfCredit struct {
	Id             string   `json:"id"`
	ExpirationDate string   `json:"expirationDate"`
	Beneficiary    string   `json:"beneficiary"`
	Amount         int      `json:"amount"`
	Documents      []string `json:"documents"`
	Status         string   `json:"status"`
}

type ExportLicense struct {
	Id                 string `json:"id"`
	ExpirationDate     string `json:"expirationDate"`
	Exporter           string `json:"exporter"`
	Carrier            string `json:"carrier"`
	DescriptionOfGoods string `json:"descriptionOfGoods"`
	Approver           string `json:"approver"`
	Status             string `json:"status"`
}

type BillOfLading struct {
	Id                 string `json:"id"`
	ExpirationDate     string `json:"expirationDate"`
	Exporter           string `json:"exporter"`
	Carrier            string `json:"carrier"`
	DescriptionOfGoods string `json:"descriptionOfGoods"`
	Amount             int    `json:"amount"`
	Beneficiary        string `json:"beneficiary"`
	SourcePort         string `json:"sourcePort"`
	DestinationPort    string `json:"destinationPort"`
}
