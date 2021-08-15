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
	SessionWorkDir string
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
		},
		Action: func(c *cli.Context) error {
			swiRouter := router.NewRouter()
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
