package router

import (
	"swi/server/cache"

	fasthttp "github.com/valyala/fasthttp"
)



type Router struct {
	Session *cache.SessionManager
}

func (r *Router) bindHeaders(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Add("Content-Type", "application/json")
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
}

func (r *Router) Route(ctx *fasthttp.RequestCtx) {
	r.bindHeaders(ctx)

	switch string(ctx.Path()) {
	case "/fetch":
		// fooHandler(ctx)
		r.Session.Fetch(ctx.Request.Body())
	// case "/read":
	// 	fooHandler(ctx)
	// case "/write":
	// 	fooHandler(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
	// fmt.Fprint(ctx, string(resp))
}

func NewRouter() *Router {
	return &Router{
		Session: cache.NewSession(),
	}
}
