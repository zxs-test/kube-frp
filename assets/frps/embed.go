package frpc

import (
	"embed"

	"github.com/imneov/kube-frp/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
