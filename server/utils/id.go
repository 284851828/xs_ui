package utils

import (
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
