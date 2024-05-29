package router

import (
	"xs/router/example"
	"xs/router/system"
	"xs/router/xstock"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Xstock  xstock.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
