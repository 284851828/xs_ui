package response

import "xs/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
