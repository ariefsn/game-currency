package helper

import "github.com/unrolled/render"

var rdr *render.Render

func InitRender() {
	rdr = render.New()
}

func Render() *render.Render {
	return rdr
}
