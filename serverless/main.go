package main

import (
	"fmt"
	"net/http"

	"github.com/linkxzhou/gongx/interp_go"
	"github.com/linkxzhou/gongx/interp_go/loader"
	"github.com/linkxzhou/gongx/packages/server"
)

var initApiMap = make(map[string]*interp_go.Interp, 0)
var interpc = loader.NewContext(loader.EnableDumpImports)
var rootPathList = []string{"api"}

func init() {
	for _, path := range rootPathList {
		initApiMap[path], _ = interp_go.LoadDirWithCache(interpc, path)
	}
}

func main() {
	app := server.New()
	for path, initApiV := range initApiMap {
		iv := initApiV
		app.Any(path, func(c server.Context) error {
			fn := c.QueryParams().Get("fn")
			result, err := iv.RunAny(fn, c)
			if err != nil {
				return c.String(http.StatusInternalServerError, "fn: "+fn+", err: "+err.Error())
			} else {
				return c.String(http.StatusOK, fmt.Sprintf("%v", result))
			}
		})
	}

	app.Start(":3001")
}