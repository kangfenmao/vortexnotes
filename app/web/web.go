package web

import "embed"

//go:embed index.html
var IndexHtml embed.FS

//go:embed assets/*
var Assets embed.FS

//go:embed index.html
var IndexByte []byte

//go:embed public/favicon.svg
var Favicon embed.FS
