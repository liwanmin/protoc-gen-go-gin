// Code generated by github.com/mohuishou/protoc-gen-go-gin. DO NOT EDIT.

package v1

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	metadata "google.golang.org/grpc/metadata"
)

import (
	kgin "github.com/go-kratos/gin"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the mohuishou/protoc-gen-go-gin package it is being compiled against.
// context.metadata.
//gin.

type BlogServiceHTTPServer interface {
	CreateArticle(context.Context, *Article) (*Article, error)

	GetArticles(context.Context, *GetArticlesReq) (*GetArticlesResp, error)
}

func RegisterBlogServiceHTTPServer(r gin.IRouter, srv BlogServiceHTTPServer) {
	s := BlogService{
		server: srv,
		router: r,
		resp:   defaultBlogServiceResp{},
	}
	s.RegisterService()
}

type BlogService struct {
	server BlogServiceHTTPServer
	router gin.IRouter
	resp   interface {
		Success(ctx *gin.Context, data interface{})
	}
}

// Resp 返回值
type defaultBlogServiceResp struct{}

func (resp defaultBlogServiceResp) response(ctx *gin.Context, status, code int, msg string, data interface{}) {
	ctx.JSON(status, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// Success 返回成功信息
func (resp defaultBlogServiceResp) Success(ctx *gin.Context, data interface{}) {
	resp.response(ctx, 200, 0, "成功", data)
}

func (s *BlogService) GetArticles_0(ctx *gin.Context) {
	var in GetArticlesReq

	if err := ctx.ShouldBindUri(&in); err != nil {
		kgin.Error(ctx, ErrorParamError(err.Error()))
		return
	}

	if err := ctx.ShouldBindQuery(&in); err != nil {
		kgin.Error(ctx, ErrorParamError(err.Error()))
		return
	}

	if err := in.ValidateAll(); err != nil {
		kgin.Error(ctx, ErrorParamError(err.Error()))
		return
	}
	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(BlogServiceHTTPServer).GetArticles(newCtx, &in)
	if err != nil {
		kgin.Error(ctx, err)
		return
	}

	s.resp.Success(ctx, out)
}

func (s *BlogService) GetArticles_1(ctx *gin.Context) {
	var in GetArticlesReq

	if err := ctx.ShouldBindQuery(&in); err != nil {
		kgin.Error(ctx, ErrorParamError(err.Error()))
		return
	}

	if err := in.ValidateAll(); err != nil {
		kgin.Error(ctx, ErrorParamError(err.Error()))
		return
	}
	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(BlogServiceHTTPServer).GetArticles(newCtx, &in)
	if err != nil {
		kgin.Error(ctx, err)
		return
	}

	s.resp.Success(ctx, out)
}

func (s *BlogService) CreateArticle_0(ctx *gin.Context) {
	var in Article

	if err := ctx.ShouldBindUri(&in); err != nil {
		kgin.Error(ctx, ErrorParamError(err.Error()))
		return
	}

	if err := ctx.ShouldBindJSON(&in); err != nil {
		kgin.Error(ctx, ErrorParamError(err.Error()))
		return
	}

	if err := in.ValidateAll(); err != nil {
		kgin.Error(ctx, ErrorParamError(err.Error()))
		return
	}
	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(BlogServiceHTTPServer).CreateArticle(newCtx, &in)
	if err != nil {
		kgin.Error(ctx, err)
		return
	}

	s.resp.Success(ctx, out)
}

func (s *BlogService) RegisterService() {

	s.router.Handle("GET", "/v1/author/:author_id/articles", s.GetArticles_0)

	s.router.Handle("GET", "/v1/articles", s.GetArticles_1)

	s.router.Handle("POST", "/v1/author/:author_id/articles", s.CreateArticle_0)

}
