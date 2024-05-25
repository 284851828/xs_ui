package request

import (
	"xs/model/common/request"
	"xs/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
