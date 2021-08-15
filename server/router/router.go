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
	case "/session/code/download":
		if string(ctx.Request.Header.Method()) != "POST" {
			handleError(&ctx.Response, apperror.UnsupportedMethodError)
			return
		}
		var sessionRetrieveBody UpdateSessionBody
		err := json.Unmarshal(ctx.Request.Body(), &sessionRetrieveBody)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		binaryBytes, err := r.sessionManager.DownloadSolanaCompiledProject(sessionRetrieveBody.SessionID)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		ctx.Response.SetBody(binaryBytes)
		ctx.Response.Header.Set("Content-Type", "application/octet-stream")
		ctx.Response.Header.Set("Content-Disposition", "attachment; filename=\"helloworld.so\"")
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	case "/session/code/compile":
		if string(ctx.Request.Header.Method()) != "POST" {
			handleError(&ctx.Response, apperror.UnsupportedMethodError)
			return
		}
		var sessionRetrieveBody UpdateSessionBody
		err := json.Unmarshal(ctx.Request.Body(), &sessionRetrieveBody)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		session, err := r.sessionManager.CompileSolanaProject(sessionRetrieveBody.SessionID)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		ctx.Response.SetBody(SuccessfulResponseFrom(session).Bytes())
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	case "/session/code/update":
		if string(ctx.Request.Header.Method()) != "POST" {
			handleError(&ctx.Response, apperror.UnsupportedMethodError)
			return
		}
		var sessionRetrieveBody UpdateSessionBody
		err := json.Unmarshal(ctx.Request.Body(), &sessionRetrieveBody)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		session, err := r.sessionManager.UpdateSessionData(sessionRetrieveBody.SessionID, sessionRetrieveBody.Tree)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		ctx.Response.SetBody(SuccessfulResponseFrom(session).Bytes())
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	case "/session/code/tree":
		if string(ctx.Request.Header.Method()) != "POST" {
			handleError(&ctx.Response, apperror.UnsupportedMethodError)
			return
		}
		var sessionRetrieveBody RequestSessionBody
		err := json.Unmarshal(ctx.Request.Body(), &sessionRetrieveBody)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		session, err := r.sessionManager.BuildSessionTreeFor(sessionRetrieveBody.SessionID)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		ctx.Response.SetBody(SuccessfulResponseFrom(session).Bytes())
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	case "/session/info":
		if string(ctx.Request.Header.Method()) != "POST" {
			handleError(&ctx.Response, apperror.UnsupportedMethodError)
			return
		}
		var sessionRetrieveBody RequestSessionBody
		err := json.Unmarshal(ctx.Request.Body(), &sessionRetrieveBody)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		session, err := r.sessionManager.RetrieveSession(sessionRetrieveBody.SessionID)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		ctx.Response.SetBody(SuccessfulResponseFrom(session).Bytes())
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	case "/session/new":
		if string(ctx.Request.Header.Method()) != "POST" {
			handleError(&ctx.Response, apperror.UnsupportedMethodError)
			return
		}
		newSession, err := r.sessionManager.NewSession()
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		ctx.Response.SetBody(SuccessfulResponseFrom(newSession).Bytes())
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
	
	// fmt.Fprint(ctx, string(resp))
}

func NewRouter(workDir, templatePath string) (*Router, error) {
	sessionManager, err := session.NewSessionManagerAt(workDir, templatePath)
	if err != nil {
		return nil, err
	}

	return &Router{
		sessionManager,
		// Session: cache.NewSession(),
	}, nil
}
