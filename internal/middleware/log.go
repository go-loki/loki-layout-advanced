package middleware

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/helper/md5"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/helper/uuid"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/log"
	"go.uber.org/zap"
)

func RequestLogMiddleware(logger *log.Logger) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {

		// 每次请求都初始化一次配置
		trace := md5.Md5(uuid.GenUUID())
		logger.NewContext(ctx, zap.String("trace", trace))
		logger.NewContext(ctx, zap.String("request_method", string(ctx.Request.Method())))
		headers := ctx.Request.Header.Header()
		logger.NewContext(ctx, zap.String("request_headers", string(headers)))
		logger.NewContext(ctx, zap.String("request_url", string(ctx.Request.RequestURI())))
		if len(ctx.Request.Body()) > 0 {
			bodyBytes := ctx.GetRawData()
			ctx.Request.SwapBody(bodyBytes) // 关键点
			logger.NewContext(ctx, zap.String("request_params", string(bodyBytes)))
		}
		logger.WithContext(ctx).Info("Request")
		ctx.Next(c)
	}
}
func ResponseLogMiddleware(logger *log.Logger) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		//blw := &bodyLogWriter{body: bytes.NewBufferString(""), Writer: ctx.GetWriter()}
		//ctx.Writer = blw
		//startTime := time.Now()
		ctx.Next(c)
		//duration := int(time.Since(startTime).Milliseconds())
		//ctx.Header("X-Response-Time", strconv.Itoa(duration))
		//logger.WithContext(ctx).Info("Response", zap.Any("response_body", blw.body.String()), zap.Any("time", fmt.Sprintf("%sms", strconv.Itoa(duration))))
	}
}

type bodyLogWriter struct {
	network.Writer
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.Writer.WriteBinary(b)
}
