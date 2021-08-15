package router

import (
	"encoding/json"
	"swi/server/apperror"
	"swi/server/cache"
	"swi/server/session"

	fasthttp "github.com/valyala/fasthttp"
)



var _ cache.CacheManager


type Router struct {
	sessionManager *session.SessionManager
}

func (r *Router) bindHeaders(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Add("Content-Type", "application/json")
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
}

func handleError(ctx *fasthttp.Response, err error) {
	if err != nil {
		resp := FailedResponseFrom(err)
		respBytes, _ := json.Marshal(resp)

		ctx.SetBody(respBytes)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
	}
}


func (r *Router) Route(ctx *fasthttp.RequestCtx) {
	r.bindHeaders(ctx)

	var err error
	_ = err
	switch string(ctx.Path()) {
	case "/fetch":



	case "/session/info":
		if string(ctx.Request.Header.Method()) != "POST" {
			handleError(&ctx.Response, apperror.UnsupportedMethodError)
			return
		}
		newSession := r.sessionManager.NewSession()

		ctx.Response.SetBody(SuccessfulResponseFrom(newSession).Bytes())
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	case "/session/new":
		if string(ctx.Request.Header.Method()) != "POST" {
			handleError(&ctx.Response, apperror.UnsupportedMethodError)
			return
		}
		newSession := r.sessionManager.NewSession()

		ctx.Response.SetBody(SuccessfulResponseFrom(newSession).Bytes())
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
	
	// fmt.Fprint(ctx, string(resp))
}

func NewRouter() *Router {
	return &Router{
		sessionManager: session.NewSessionManager(),
		// Session: cache.NewSession(),
	}
}
