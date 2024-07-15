package utils

import (
	"math"
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
)

var __n *snowflake.Node

func init() {
	__n, _ = snowflake.NewNode(time.Now().UnixMilli() % 1024)
}

func GenId() snowflake.ID {
	id := __n.Generate()
	return id
	// id.String()
	// return fmt.Sprintf("%d", id.Int64()), id
}

// Round
/**
 * @Description 用于对浮点数进行四舍五入，并保留指定的小数位数
 * @Date 2023-05-17 16:42:35
 */
func Round(value float64, decimalPlaces int) float64 {
	decimalPow := math.Pow10(decimalPlaces)
	rounded := math.Round(value*decimalPow) / decimalPow
	rounded, _ = strconv.ParseFloat(strconv.FormatFloat(rounded, 'f', decimalPlaces, 64), 64)
	return rounded
}
