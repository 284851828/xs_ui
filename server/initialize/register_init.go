package initialize

import (
	_ "xs/source/example"
	_ "xs/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
