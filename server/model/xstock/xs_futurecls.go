package xstock

import (
	"fmt"
	"xs/global"
	"xs/utils"

	"go.uber.org/zap"
)

type FutureCls struct {
	global.GVA_BASE
	// 用户UUID
	Class        string  `json:"class" gorm:"index;comment:类型"`                // 类型
	Exchange     string  `json:"exchange"  gorm:"comment:交易所"`                 // 交易所
	Pricetick    float64 `json:"pricetick" gorm:"comment:最小跳价"`                // 最小跳价
	Size         int     `json:"size" gorm:"comment:倍数"`                       // 倍数
	Mainsymbol   string  `json:"mainsymbol" gorm:"comment:主力合约"`               // 主力合约
	Riye         int     `json:"riye" gorm:"comment:夜盘类型"`                     // 夜盘类型
	Maxvolsymbol string  `json:"maxvolsymbol" gorm:"comment:最大量合约"`            // 最大量合约
	Showindex    uint    `json:"showindex" gorm:"default:0;comment:展示排序"`      // 展示排序
	Enable       int     `json:"enable" gorm:"default:1;comment:是否启用 1正常 0禁用"` // 1正常 2禁用
}

/*
{
    "_id" : ObjectId("60dea87748d5d0dc1fc2c445"),
    "class" : "ZC",
    "exchange" : "CZCE",
    "pricetick" : 0.2,
    "size" : NumberInt(100),
    "mainsymbol" : "ZC301",
    "riye" : NumberInt(2),
    "maxvolsymbol" : "ZC301"
}
*/

func (FutureCls) TableName() string {
	return "xs_futurecls"
}

func (b *FutureCls) Save() error {

	if b.ID == "" {
		b.ID = "cls_" + utils.GenId().String()
	}
	return global.GVA_DB.Save(b).Error
}

func (b *FutureCls) DelById(id string) (err error) {
	return global.GVA_DB.Where("id = ?", id).Delete(b).Error
}

func (b *FutureCls) GetById(id string) error {
	return global.GVA_DB.First(b, "id = ?", id).Error
}

const sql_FutureClsList = `
	select * from xs_futurecls where enable =1  order by showindex
	`

func GetFutureClsList() ([]FutureCls, error) {
	l := []FutureCls{}

	sql := fmt.Sprintf(sql_FutureClsList)
	e := global.GVA_DB.Raw(sql).Scan(&l).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", sql))
		fmt.Println("GetFutureClsList sql\n", sql)
		return l, e
	}

	return l, nil
}

func (i *FutureClsResp) GetData() error {
	i.Records = []FutureCls{}
	l, e := GetFutureClsList()
	i.Records = l
	return e
}
