package handler

import (
	"net/http"
	"strconv"

	"plots/internal/model"

	"github.com/gin-gonic/gin"
)

// ListArticles 获取文章列表
func (h *Handler) ListArticles(c *gin.Context) {
	var query model.PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		query.Page = 1
		query.PageSize = 10
	}

	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 || query.PageSize > 100 {
		query.PageSize = 10
	}

	articles, total, err := h.repos.Article.List(query, true)
	if err != nil {
		Error(c, http.StatusInternalServerError, "获取文章列表失败")
		return
	}

	totalPages := int(total) / query.PageSize
	if int(total)%query.PageSize > 0 {
		totalPages++
	}

	Success(c, model.PaginatedResponse{
		Data:       articles,
		Total:      total,
		Page:       query.Page,
		PageSize:   query.PageSize,
		TotalPages: totalPages,
	})
}

// GetArticle 获取文章详情
func (h *Handler) GetArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		Error(c, http.StatusBadRequest, "无效的文章ID")
		return
	}

	article, err := h.repos.Article.GetByID(uint(id))
	if err != nil {
		Error(c, http.StatusNotFound, "文章不存在")
		return
	}

	// 增加浏览量
	_ = h.repos.Article.IncrementViews(uint(id))

	Success(c, article)
}

// CreateArticle 创建文章
func (h *Handler) CreateArticle(c *gin.Context) {
	var req model.CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	userID := c.GetUint("user_id")

	article := &model.Article{
		Title:     req.Title,
		Content:   req.Content,
		Summary:   req.Summary,
		Cover:     req.Cover,
		Category:  req.Category,
		Tags:      req.Tags,
		Published: req.Published,
		AuthorID:  userID,
	}

	if err := h.repos.Article.Create(article); err != nil {
		Error(c, http.StatusInternalServerError, "创建文章失败")
		return
	}

	Success(c, article)
}

// UpdateArticle 更新文章
func (h *Handler) UpdateArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		Error(c, http.StatusBadRequest, "无效的文章ID")
		return
	}

	var req model.UpdateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	article, err := h.repos.Article.GetByID(uint(id))
	if err != nil {
		Error(c, http.StatusNotFound, "文章不存在")
		return
	}

	// 更新字段
	if req.Title != nil {
		article.Title = *req.Title
	}
	if req.Content != nil {
		article.Content = *req.Content
	}
	if req.Summary != nil {
		article.Summary = *req.Summary
	}
	if req.Cover != nil {
		article.Cover = *req.Cover
	}
	if req.Category != nil {
		article.Category = *req.Category
	}
	if req.Tags != nil {
		article.Tags = *req.Tags
	}
	if req.Published != nil {
		article.Published = *req.Published
	}

	if err := h.repos.Article.Update(article); err != nil {
		Error(c, http.StatusInternalServerError, "更新文章失败")
		return
	}

	Success(c, article)
}

// DeleteArticle 删除文章
func (h *Handler) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		Error(c, http.StatusBadRequest, "无效的文章ID")
		return
	}

	if err := h.repos.Article.Delete(uint(id)); err != nil {
		Error(c, http.StatusInternalServerError, "删除文章失败")
		return
	}

	Success(c, nil)
}
