package router

import (
	"encoding/json"
	"swi/protobuf/transport"
	"swi/server/cache"
	"swi/server/errors"
	"swi/server/session"

	fasthttp "github.com/valyala/fasthttp"
)



var _ cache.CacheManager


type Router struct {
	sessionManager *session.SessionManager
}

func (r *Router) bindHeaders(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Add("Content-Type", "application/json")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Keep-Alive,User-Agent,Cache-Control,Content-Type,Authorization")
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

	if string(ctx.Request.Header.Method()) == "OPTIONS" {
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
		// ctx.Response.Header.Del("Content-Type")
		return
	}
	// var err error
	// _ = err
	switch string(ctx.Path()) {
	case "/session/code/download":
		if string(ctx.Request.Header.Method()) != "POST" {
			handleError(&ctx.Response, errors.UnsupportedMethodError)
			return
		}
		var sessionRetrieveBody transport.RequestSessionBody
		err := json.Unmarshal(ctx.Request.Body(), &sessionRetrieveBody)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		binaryBytes, err := r.sessionManager.DownloadSolanaCompiledProject(sessionRetrieveBody.SessionId)
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
			handleError(&ctx.Response, errors.UnsupportedMethodError)
			return
		}
		var sessionRetrieveBody transport.RequestSessionBody
		err := json.Unmarshal(ctx.Request.Body(), &sessionRetrieveBody)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		session, err := r.sessionManager.CompileSolanaProject(sessionRetrieveBody.SessionId)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		ctx.Response.SetBody(SuccessfulResponseFrom(session).Bytes())
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	case "/session/code/update":
		if string(ctx.Request.Header.Method()) != "POST" {
			handleError(&ctx.Response, errors.UnsupportedMethodError)
			return
		}
		var sessionRetrieveBody transport.UpdateSessionBody
		err := json.Unmarshal(ctx.Request.Body(), &sessionRetrieveBody)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		session, err := r.sessionManager.UpdateSessionData(sessionRetrieveBody.SessionId, sessionRetrieveBody.Tree)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		ctx.Response.SetBody(SuccessfulResponseFrom(session).Bytes())
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	case "/session/code/tree":
		if string(ctx.Request.Header.Method()) != "POST" {
			handleError(&ctx.Response, errors.UnsupportedMethodError)
			return
		}
		var sessionRetrieveBody transport.RequestSessionBody
		err := json.Unmarshal(ctx.Request.Body(), &sessionRetrieveBody)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		session, err := r.sessionManager.BuildSessionTreeFor(sessionRetrieveBody.SessionId)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		ctx.Response.SetBody(SuccessfulResponseFrom(session).Bytes())
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	case "/session/code/legacy/tree":
		if string(ctx.Request.Header.Method()) != "POST" {
			handleError(&ctx.Response, errors.UnsupportedMethodError)
			return
		}
		var sessionRetrieveBody transport.RequestSessionBody
		err := json.Unmarshal(ctx.Request.Body(), &sessionRetrieveBody)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		session, err := r.sessionManager.BuildSessionLegacyTreeFor(sessionRetrieveBody.SessionId)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		ctx.Response.SetBody(SuccessfulResponseFrom(session).Bytes())
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	case "/session/info":
		if string(ctx.Request.Header.Method()) != "POST" {
			handleError(&ctx.Response, errors.UnsupportedMethodError)
			return
		}
		var sessionRetrieveBody transport.RequestSessionBody
		err := json.Unmarshal(ctx.Request.Body(), &sessionRetrieveBody)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		session, err := r.sessionManager.RetrieveSession(sessionRetrieveBody.SessionId)
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		ctx.Response.SetBody(SuccessfulResponseFrom(session).Bytes())
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	case "/session/new":
		if string(ctx.Request.Header.Method()) != "POST" {
			handleError(&ctx.Response, errors.UnsupportedMethodError)
			return
		}
		newSession, err := r.sessionManager.NewSession()
		if err != nil {
			handleError(&ctx.Response, err)
			return
		}

		ctx.Response.SetBody(SuccessfulResponseFrom(newSession).Bytes())
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	// default:
	// 	ctx.Error("Unsupported path", fasthttp.StatusNotFound)
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
