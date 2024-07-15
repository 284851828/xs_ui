package xstock

type FutureClsResp struct {
	Records []FutureCls `json:"records"`
}

type PredListResp struct {
	Records []Pred `json:"records"`
	Moc900  MOC900 `json:"moc900"`
}

type Pred0000ListResp struct {
	Records []Pred `json:"records"`
}
