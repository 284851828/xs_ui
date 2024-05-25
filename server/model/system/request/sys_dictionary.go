package request

import (
	"xs/model/common/request"
	"xs/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
