// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v3.14.0
// source: article/v1/article.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationArticleCreateArticle = "/api.article.v1.Article/CreateArticle"
const OperationArticleDeleteArticle = "/api.article.v1.Article/DeleteArticle"
const OperationArticleFavoriteArticle = "/api.article.v1.Article/FavoriteArticle"
const OperationArticleFeedArticle = "/api.article.v1.Article/FeedArticle"
const OperationArticleGetArticle = "/api.article.v1.Article/GetArticle"
const OperationArticleListArticle = "/api.article.v1.Article/ListArticle"
const OperationArticleUnFavoriteArticle = "/api.article.v1.Article/UnFavoriteArticle"
const OperationArticleUpdateArticle = "/api.article.v1.Article/UpdateArticle"

type ArticleHTTPServer interface {
	// CreateArticle 创建文章，需要身份验证
	CreateArticle(context.Context, *CreateArticleRequest) (*ArticleReply, error)
	// DeleteArticle 删除文章，需要身份验证
	DeleteArticle(context.Context, *SlugRequest) (*EmptyReply, error)
	FavoriteArticle(context.Context, *SlugRequest) (*ArticleReply, error)
	// FeedArticle 返回关注用户的多篇文章，需要身份验证
	FeedArticle(context.Context, *ListArticleRequest) (*ListArticleReply, error)
	// GetArticle 获取单篇文章，无需身份验证
	GetArticle(context.Context, *SlugRequest) (*ArticleReply, error)
	// ListArticle 获取多篇文章，需要身份验证
	ListArticle(context.Context, *ListArticleRequest) (*ListArticleReply, error)
	UnFavoriteArticle(context.Context, *SlugRequest) (*ArticleReply, error)
	// UpdateArticle 更新文章，需要身份验证
	UpdateArticle(context.Context, *SaveArticleRequest) (*ArticleReply, error)
}

func RegisterArticleHTTPServer(s *http.Server, srv ArticleHTTPServer) {
	r := s.Route("/")
	r.GET("/api/articles/feed", _Article_FeedArticle0_HTTP_Handler(srv))
	r.GET("/api/articles", _Article_ListArticle0_HTTP_Handler(srv))
	r.GET("/api/articles/{slug}", _Article_GetArticle0_HTTP_Handler(srv))
	r.POST("/api/articles", _Article_CreateArticle0_HTTP_Handler(srv))
	r.PUT("/api/articles/{slug}", _Article_UpdateArticle0_HTTP_Handler(srv))
	r.DELETE("/api/articles/{slug}", _Article_DeleteArticle0_HTTP_Handler(srv))
	r.POST("/api/articles/{slug}/favorite", _Article_FavoriteArticle0_HTTP_Handler(srv))
	r.DELETE("/api/articles/{slug}/favorite", _Article_UnFavoriteArticle0_HTTP_Handler(srv))
}

func _Article_FeedArticle0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleFeedArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FeedArticle(ctx, req.(*ListArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Article_ListArticle0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleListArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListArticle(ctx, req.(*ListArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Article_GetArticle0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SlugRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleGetArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArticle(ctx, req.(*SlugRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Article_CreateArticle0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleCreateArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateArticle(ctx, req.(*CreateArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Article_UpdateArticle0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleUpdateArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateArticle(ctx, req.(*SaveArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Article_DeleteArticle0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SlugRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleDeleteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteArticle(ctx, req.(*SlugRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _Article_FavoriteArticle0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SlugRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleFavoriteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FavoriteArticle(ctx, req.(*SlugRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Article_UnFavoriteArticle0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SlugRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleUnFavoriteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UnFavoriteArticle(ctx, req.(*SlugRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ArticleReply)
		return ctx.Result(200, reply)
	}
}

type ArticleHTTPClient interface {
	CreateArticle(ctx context.Context, req *CreateArticleRequest, opts ...http.CallOption) (rsp *ArticleReply, err error)
	DeleteArticle(ctx context.Context, req *SlugRequest, opts ...http.CallOption) (rsp *EmptyReply, err error)
	FavoriteArticle(ctx context.Context, req *SlugRequest, opts ...http.CallOption) (rsp *ArticleReply, err error)
	FeedArticle(ctx context.Context, req *ListArticleRequest, opts ...http.CallOption) (rsp *ListArticleReply, err error)
	GetArticle(ctx context.Context, req *SlugRequest, opts ...http.CallOption) (rsp *ArticleReply, err error)
	ListArticle(ctx context.Context, req *ListArticleRequest, opts ...http.CallOption) (rsp *ListArticleReply, err error)
	UnFavoriteArticle(ctx context.Context, req *SlugRequest, opts ...http.CallOption) (rsp *ArticleReply, err error)
	UpdateArticle(ctx context.Context, req *SaveArticleRequest, opts ...http.CallOption) (rsp *ArticleReply, err error)
}

type ArticleHTTPClientImpl struct {
	cc *http.Client
}

func NewArticleHTTPClient(client *http.Client) ArticleHTTPClient {
	return &ArticleHTTPClientImpl{client}
}

func (c *ArticleHTTPClientImpl) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...http.CallOption) (*ArticleReply, error) {
	var out ArticleReply
	pattern := "/api/articles"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationArticleCreateArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ArticleHTTPClientImpl) DeleteArticle(ctx context.Context, in *SlugRequest, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleDeleteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ArticleHTTPClientImpl) FavoriteArticle(ctx context.Context, in *SlugRequest, opts ...http.CallOption) (*ArticleReply, error) {
	var out ArticleReply
	pattern := "/api/articles/{slug}/favorite"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationArticleFavoriteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ArticleHTTPClientImpl) FeedArticle(ctx context.Context, in *ListArticleRequest, opts ...http.CallOption) (*ListArticleReply, error) {
	var out ListArticleReply
	pattern := "/api/articles/feed"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleFeedArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ArticleHTTPClientImpl) GetArticle(ctx context.Context, in *SlugRequest, opts ...http.CallOption) (*ArticleReply, error) {
	var out ArticleReply
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleGetArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ArticleHTTPClientImpl) ListArticle(ctx context.Context, in *ListArticleRequest, opts ...http.CallOption) (*ListArticleReply, error) {
	var out ListArticleReply
	pattern := "/api/articles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleListArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ArticleHTTPClientImpl) UnFavoriteArticle(ctx context.Context, in *SlugRequest, opts ...http.CallOption) (*ArticleReply, error) {
	var out ArticleReply
	pattern := "/api/articles/{slug}/favorite"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleUnFavoriteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ArticleHTTPClientImpl) UpdateArticle(ctx context.Context, in *SaveArticleRequest, opts ...http.CallOption) (*ArticleReply, error) {
	var out ArticleReply
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationArticleUpdateArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
