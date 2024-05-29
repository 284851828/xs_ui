package v1

import (
	"xs/api/v1/example"
	"xs/api/v1/system"
	"xs/api/v1/xstock"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	XstockApiGroup  xstock.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
