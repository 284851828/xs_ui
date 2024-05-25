package global

{{- if .HasGlobal }}

import "xs/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}