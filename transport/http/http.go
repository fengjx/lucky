package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/fengjx/go-halo/json"
	"github.com/fengjx/luchen/env"
	"github.com/fengjx/luchen/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"

	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/connom/errno"
	"github.com/fengjx/lucky/current"
)

const (
	StatusOK = 0

	// AdminAPI 管理后台接口前缀
	AdminAPI = "/admin"
	// API 接口前缀
	API = "/api"
	// OpenAPI 不需要鉴权的接口前缀
	OpenAPI = API + "/open"

	RequestHeaderAdminUID   = "X-Admin-UID"
	RequestHeaderAdminToken = "X-Admin-Token"

	ResponseHeaderServer       = "Server"
	ResponseHeaderRefreshToken = "X-Refresh-Token"
)

type result struct {
	Status int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func ResponseWrapper(data interface{}) interface{} {
	res := &result{
		Status: StatusOK,
		Msg:    "ok",
		Data:   data,
	}
	return res
}

// ErrorEncoder 统一异常处理
func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	var errn *luchen.Errno
	ok := errors.As(err, &errn)
	if !ok {
		log.ErrorCtx(ctx, "handler error", zap.Error(err))
		msg := errno.SystemErr.Msg
		if !env.IsProd() {
			msg = err.Error()
		}
		res := &result{
			Status: errno.SystemErr.Code,
			Msg:    msg,
		}
		WriteData(ctx, w, errno.SystemErr.HTTPCode, res)
		return
	}
	httpCode := 500
	if errn.HTTPCode > 0 {
		httpCode = errn.HTTPCode
	}
	res := &result{
		Status: errn.Code,
		Msg:    errn.Msg,
	}
	WriteData(ctx, w, httpCode, res)
}

func WriteData(ctx context.Context, w http.ResponseWriter, httpCode int, data any) {
	w.WriteHeader(httpCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.ErrorCtx(ctx, "write http response err", zap.Error(err))
	}
}

// NewHandler 创建 http handler
func NewHandler(e luchen.Endpoint,
	dec httptransport.DecodeRequestFunc,
	enc httptransport.EncodeResponseFunc,
	options ...httptransport.ServerOption) *luchen.HTTPTransportServer {

	options = append(options, httptransport.ServerErrorEncoder(ErrorEncoder))
	targetEndpoint := luchen.AccessMiddleware(&luchen.AccessLogOpt{
		MaxDay: 15,
		ContextFields: map[string]luchen.GetValueFromContext{
			"uid": func(ctx context.Context) any {
				return current.UID(ctx)
			},
		},
	})(e)
	return luchen.NewHTTPTransportServer(
		targetEndpoint,
		dec,
		enc,
		options...,
	)
}
