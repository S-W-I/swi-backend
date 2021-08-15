package app

import (
	"fmt"
	"log"
	"os"
	"swi/server/router"

	"github.com/urfave/cli/v2"
	"github.com/valyala/fasthttp"
)

type AppConfig struct {
	SessionWorkDir, TemplatePath string
	Port int
}

func Run() {
	inputCfg := &AppConfig{}

	app := &cli.App{
		Name:  "Blockchain Target Chains proxy",
		Usage: "",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "port",
				Value:       8081,
				Usage:       "port for allocator",
				Destination: &inputCfg.Port,
			},
			&cli.StringFlag{
				Name:        "workdir",
				Value:       "",
				Usage:       "coding sessions workdir",
				Destination: &inputCfg.SessionWorkDir,
			},
			&cli.StringFlag{
				Name:        "tmp",
				Value:       "",
				Usage:       "init code template",
				Destination: &inputCfg.TemplatePath,
			},
		},
		Action: func(c *cli.Context) error {
			swiRouter, err := router.NewRouter(inputCfg.SessionWorkDir, inputCfg.TemplatePath)
			if err != nil {
				return err
			}
			fastHTTPHandler := func (ctx *fasthttp.RequestCtx) {
				swiRouter.Route(ctx)
			}
			
			fmt.Printf("started swi backend on port: %v\n", inputCfg.Port)
			fasthttp.ListenAndServe(":" + fmt.Sprint(inputCfg.Port), fastHTTPHandler)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
