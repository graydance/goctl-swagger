package action

import (
	"github.com/graydance/go-kit/tools/goctl/plugin"
	"github.com/urfave/cli/v2"
	"github.com/graydance/goctl-swagger/generate"
)

func Generator(ctx *cli.Context) error {
	fileName := ctx.String("filename")

	if len(fileName) == 0 {
		fileName = "rest.swagger.json"
	}

	p, err := plugin.NewPlugin()
	if err != nil {
		return err
	}
	basepath := ctx.String("basepath")
	host := ctx.String("host")
	return generate.Do(fileName, host, basepath, p)
}
