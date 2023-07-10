package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/helper/md5"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/helper/resp"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/log"
	"github.com/spf13/viper"
	"net/http"
	"sort"
	"strings"
)

func SignMiddleware(logger *log.Logger, conf *viper.Viper) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		requiredHeaders := []string{"Timestamp", "Nonce", "Sign", "App-Version"}

		for _, header := range requiredHeaders {
			value := ctx.Request.Header.Get(header)
			if len(value) == 0 {
				resp.HandleError(ctx, http.StatusBadRequest, 1, "sign error.", nil)
				ctx.Abort()
				return
			}
		}

		data := map[string]string{
			"AppKey":     conf.GetString("security.api_sign.app_key"),
			"Timestamp":  ctx.Request.Header.Get("Timestamp"),
			"Nonce":      ctx.Request.Header.Get("Nonce"),
			"AppVersion": ctx.Request.Header.Get("App-Version"),
		}

		var keys []string
		for k := range data {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })

		var str string
		for _, k := range keys {
			str += k + data[k]
		}
		str += conf.GetString("security.api_sign.app_security")

		if ctx.Request.Header.Get("Sign") != strings.ToUpper(md5.Md5(str)) {
			resp.HandleError(ctx, http.StatusBadRequest, 1, "sign error.", nil)
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}
}
