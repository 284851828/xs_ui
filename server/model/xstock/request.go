package xstock

type BankReq struct {
}

type IdReq struct {
	ID string `json:"id"`
}

type PredReq struct {
	Cls  string `json:"cls"`
	Date string `json:"date"`
}
