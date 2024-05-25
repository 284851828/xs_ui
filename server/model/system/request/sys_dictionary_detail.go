package request

import (
	"xs/model/common/request"
	"xs/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
