package http

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fengjx/luchen/env"
	"github.com/fengjx/luchen/log"
	"go.uber.org/zap"

	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/common/auth"
	"github.com/fengjx/lucky/current"
	"github.com/fengjx/lucky/logic/admin"
)

var (
	noAuthPaths = []string{
		OpenAPI,
		"/static",
		"/dashboard",
		"/admin",
		"/debug/pprof",
	}
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(ResponseHeaderServer, "lucky")
		next.ServeHTTP(w, r)
	})
}

func adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, admin.APIPrefix) {
			next.ServeHTTP(w, r)
			return
		}
		ctx := r.Context()
		l := log.GetLogger(ctx)
		var uid int64
		token := r.Header.Get(RequestHeaderAdminToken)
		if len(token) > 0 {
			payload, expiresAt, err := auth.Parse(token)
			if err != nil {
				l.Warn("parse token err", zap.String("token", token), zap.Error(err))
				luchen.WriteError(ctx, w, luchen.ErrUnauthorized)
				return
			}
			uid = payload.UID
			if expiresAt > 0 && time.Unix(expiresAt, 0).Sub(time.Now()) < (time.Hour*24*6) {
				refreshToken, _ := auth.GenToken(payload)
				// 刷新 token
				w.Header().Set(ResponseHeaderRefreshToken, refreshToken)
			}
		}
		if uid == 0 && !env.IsProd() {
			uidHeader := r.Header.Get(RequestHeaderAdminUID)
			if len(uidHeader) > 0 {
				debugUID, err := strconv.ParseInt(uidHeader, 10, 64)
				if err == nil {
					uid = debugUID
					l.Info("set debug uid", zap.Int64("uid", uid))
				}
			}
		}
		if uid > 0 {
			ctx := log.WithLogger(r.Context(), zap.Int64("admin_uid", uid))
			ctx = current.WithAdminUID(ctx, uid)
			r = r.WithContext(ctx)
		}

		if isNoAuthPath(r) {
			next.ServeHTTP(w, r)
			return
		}
		if uid == 0 {
			l.Warn("admin request unauthorized", zap.String("path", r.URL.Path))
			errn := luchen.ErrBadRequest
			luchen.WriteError(r.Context(), w, errn)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isNoAuthPath(r *http.Request) bool {
	p := r.URL.Path
	for _, prefix := range noAuthPaths {
		if strings.HasPrefix(p, prefix) {
			return true
		}
	}
	return false
}
