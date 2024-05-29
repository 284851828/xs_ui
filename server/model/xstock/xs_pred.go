package xstock

import (
	"fmt"
	"xs/global"
	"xs/utils"

	"go.uber.org/zap"
)

type Pred struct {
	global.GVA_BASE

	Symbol string  `json:"symbol" gorm:"index;comment:品种"` // 品种
	Model  string  `json:"model"  gorm:"comment:模型"`       // 模型
	Date   string  `json:"date" gorm:"comment:交易日"`        // 交易日
	Y      float64 `json:"y" gorm:"comment:y"`             // y
	Dt     string  `json:"dt" gorm:"comment:日期时间"`         // 日期时间
	Time   string  `json:"time" gorm:"comment:时间"`         // 时间
}

// {
//     "_id" : ObjectId("6646c996d75e926c26addc42"),
//     "symbol" : "IHm",
//     "model" : "IHm009",
//     "date" : "2024-05-17 11:05:58",
//     "dt" : "2024-05-17 11:05:58",
//     "y" : -0.15153813362121582
// }

func (Pred) TableName() string {
	return "xs_pred"
}

func (b *Pred) Save() error {

	if b.ID == "" {
		b.ID = "pred_" + utils.GenId().String()
	}
	return global.GVA_DB.Save(b).Error
}

const sql_PredList = `
	select * from xs_pred where date='%s' and symbol ='%s' order by dt 
	`

func GetPredList(tradeDay, symbol string) ([]Pred, error) {
	l := []Pred{}

	sql := fmt.Sprintf(sql_PredList, tradeDay, symbol)
	fmt.Println("GetPredList sql\n", sql)
	e := global.GVA_DB.Raw(sql).Scan(&l).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", sql))
		fmt.Println("GetPredList sql\n", sql)
		return l, e
	}

	return l, nil
}

func (i *PredListResp) GetData(r PredReq) error {
	i.Records = []Pred{}
	l, e := GetPredList(r.Date, r.Cls+"m")
	i.Records = l
	return e
}
