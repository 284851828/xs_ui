package xstock

import (
	"fmt"
	"strings"
	"xs/global"
	"xs/utils"

	"github.com/jinzhu/copier"
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
	select id, symbol, model, y, dt, date, DATE_FORMAT(dt, '%%H:%%i:00') as time  from xs_pred where date='%s' and symbol ='%s' order by dt 
	`
const sql_PredListYepan = `
	select id, symbol, model, y, dt, date, DATE_FORMAT(dt, '%%H:%%i:00') as time  from xs_pred where date='%s' and symbol ='%s' and time >"20:40:00"  order by dt 
	`

func GetPredList(tradeDay, symbol string) ([]Pred, error) {
	l := []Pred{}

	sql := fmt.Sprintf(sql_PredList, tradeDay, symbol)
	if GetGroupByCls(GetClsBySymbolm(symbol)) <= 3 {
		sql = fmt.Sprintf(sql_PredListYepan, tradeDay, symbol)
	}
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

	(&i.Moc900).GetData(r.Cls+"m", r.Date)
	return e
}

const sql_Pred0000List_v1 = `
	select id, symbol, model, y, dt, date, DATE_FORMAT(dt, '%%H:%%i:00') as time  
	from xs_pred0000 where  symbol ='%s' and model ='%s008'   order by dt 
	`

const sql_Pred0000List = `
with tt as (
(select id, symbol, model, y, dt, date, dt as time  from xs_pred0000 
where  symbol =':symbol' and model =':symbol008'   order by dt desc limit 60 )
UNION
(select id, symbol, model, y, dt, date, dt as time from xs_day_close
where  symbol =':symbol'   order by dt desc limit 60 )

UNION
(select id, symbol, model, y, dt, date, dt as time from xs_day_zf
where  symbol =':symbol'   order by dt desc limit 60 )

)
select * from tt order by dt

`

func GetPred0000List(symbol string) ([]Pred, error) {
	l := []Pred{}

	// sql := fmt.Sprintf(sql_Pred0000List, symbol, symbol)
	sql := strings.ReplaceAll(sql_Pred0000List, ":symbol", symbol)
	fmt.Println("sql_Pred0000List sql\n", sql)
	e := global.GVA_DB.Raw(sql).Scan(&l).Error
	if e != nil {
		global.GVA_LOG.Error("Raw(mainSql) select", zap.String("err", e.Error()), zap.String("sql", sql))
		fmt.Println("sql_Pred0000List sql\n", sql)
		return l, e
	}

	return l, nil
}

func (i *Pred0000ListResp) GetData(r PredReq) error {
	i.Records = []Pred{}
	l, e := GetPred0000List(r.Cls + "m")
	// i.Records = l
	i.Records = i.AppendAvg20(l)
	return e
}

func (i *Pred0000ListResp) AppendAvg20(l []Pred) []Pred {
	var numbers []float64
	var list008 []Pred
	for _, item := range l {
		if strings.Contains(item.Model, "008") {
			numbers = append(numbers, item.Y)
			list008 = append(list008, item)
		}
	}
	avg20 := slidingAverage(numbers)
	fmt.Println("", len(numbers), len(avg20))
	var result []Pred

	for idx, item := range list008 {
		// result = append(result, item)
		newItem := Pred{}
		copier.Copy(&newItem, &item)
		newItem.Model = "avg20"
		newItem.Y = avg20[idx]
		result = append(result, newItem)
	}
	l = append(l, result...)
	return l
}

// slidingAverage 计算输入切片中每个元素的滑动窗口平均值，
// 并返回一个与输入切片相同长度的切片，其中每个元素代表了对应位置的平均值。
func slidingAverage(numbers []float64) []float64 {
	n := len(numbers)
	averages := make([]float64, n)

	for i := 0; i < n; i++ {
		sum := 0.0
		count := 0

		// 确定窗口的下限，防止越界
		lowerLimit := i - 19
		if lowerLimit < 0 {
			lowerLimit = 0
		}

		// 计算当前窗口内所有元素的总和
		for j := lowerLimit; j <= i; j++ {
			sum += numbers[j]
			count++
		}

		// 防止除以零错误
		if count == 0 {
			averages[i] = 0 // 或者可以考虑使用一个默认值，比如numbers[i]
			continue
		}

		// 计算平均值并添加到结果切片中
		average := sum / float64(count)
		averages[i] = average
	}

	return averages
}

// func main() {
// 	numbers := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0, 21.0, 22.0, 23.0, 24.0, 25.0}
// 	averages := averageWindow(numbers)
// 	fmt.Println(averages)
// }
