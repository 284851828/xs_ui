package request

import (
	"xs/model/common/request"
	"xs/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
