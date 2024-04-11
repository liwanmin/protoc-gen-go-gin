type {{ $.InterfaceName }} interface {
{{range .MethodSet}}
	{{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{end}}
}
func Register{{ $.InterfaceName }}(r gin.IRouter, srv {{ $.InterfaceName }}) {
	s := {{.Name}}{
		server: srv,
		router: r,
		resp: default{{$.Name}}Resp{},
	}
	s.RegisterService()
}

type {{$.Name}} struct{
	server {{ $.InterfaceName }}
	router gin.IRouter
	resp  interface {
		Success(ctx *gin.Context, data interface{})
	}
}

// Resp 返回值
type default{{$.Name}}Resp struct {}

func (resp default{{$.Name}}Resp) response(ctx *gin.Context, status, code int, message string, data interface{}) {
	ctx.JSON(status, map[string]interface{}{
		"code": code, 
		"message": message,
		"data": data,
	})
}

// Success 返回成功信息
func (resp default{{$.Name}}Resp) Success(ctx *gin.Context, data interface{}) {
	resp.response(ctx, 200, 0, "success", data)
}


{{range .Methods}}
func (s *{{$.Name}}) {{ .HandlerName }} (ctx *gin.Context) {
	var in {{.Request}}
{{if .HasPathParams }}
	if err := ctx.ShouldBindUri(&in); err != nil {
		kgin.Error(ctx, ErrorParamError(err.Error()))
		return
	}
{{end}}
{{if eq .Method "GET" "DELETE" }}
	if err := ctx.ShouldBindQuery(&in); err != nil {
		kgin.Error(ctx, ErrorParamError(err.Error()))
		return
	}
{{else if eq .Method "POST" "PUT" }}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		kgin.Error(ctx, ErrorParamError(err.Error()))
		return
	}
{{else}}
	if err := ctx.ShouldBind(&in); err != nil {
		kgin.Error(ctx, ErrorParamError(err.Error()))
		return
	}
{{end}}

   	if err := in.ValidateAll(); err != nil {
   		kgin.Error(ctx, ErrorParamError(err.Error()))
   		return
   	}
	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.({{ $.InterfaceName }}).{{.Name}}(newCtx, &in)
	if err != nil {
		kgin.Error(ctx, err)
		return
	}

	s.resp.Success(ctx, out)
}
{{end}}

func (s *{{$.Name}}) RegisterService() {
{{range .Methods}}
		s.router.Handle("{{.Method}}", "{{.Path}}", s.{{ .HandlerName }})
{{end}}
}