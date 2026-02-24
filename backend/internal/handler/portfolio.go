package handler

import (
	"net/http"
	"strconv"

	"plots/internal/model"

	"github.com/gin-gonic/gin"
)

// ListPortfolios 获取作品集列表
func (h *Handler) ListPortfolios(c *gin.Context) {
	portfolios, err := h.repos.Portfolio.List(true)
	if err != nil {
		Error(c, http.StatusInternalServerError, "获取作品集列表失败")
		return
	}

	Success(c, portfolios)
}

// GetPortfolio 获取作品集详情
func (h *Handler) GetPortfolio(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		Error(c, http.StatusBadRequest, "无效的作品集ID")
		return
	}

	portfolio, err := h.repos.Portfolio.GetByID(uint(id))
	if err != nil {
		Error(c, http.StatusNotFound, "作品集不存在")
		return
	}

	Success(c, portfolio)
}

// CreatePortfolio 创建作品集
func (h *Handler) CreatePortfolio(c *gin.Context) {
	var req model.CreatePortfolioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	portfolio := &model.Portfolio{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		ImageURL:    req.ImageURL,
		ProjectURL:  req.ProjectURL,
		GithubURL:   req.GithubURL,
		TechStack:   req.TechStack,
		SortOrder:   req.SortOrder,
		Published:   req.Published,
	}

	if err := h.repos.Portfolio.Create(portfolio); err != nil {
		Error(c, http.StatusInternalServerError, "创建作品集失败")
		return
	}

	Success(c, portfolio)
}

// UpdatePortfolio 更新作品集
func (h *Handler) UpdatePortfolio(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		Error(c, http.StatusBadRequest, "无效的作品集ID")
		return
	}

	var req model.UpdatePortfolioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	portfolio, err := h.repos.Portfolio.GetByID(uint(id))
	if err != nil {
		Error(c, http.StatusNotFound, "作品集不存在")
		return
	}

	// 更新字段
	if req.Title != nil {
		portfolio.Title = *req.Title
	}
	if req.Description != nil {
		portfolio.Description = *req.Description
	}
	if req.Content != nil {
		portfolio.Content = *req.Content
	}
	if req.ImageURL != nil {
		portfolio.ImageURL = *req.ImageURL
	}
	if req.ProjectURL != nil {
		portfolio.ProjectURL = *req.ProjectURL
	}
	if req.GithubURL != nil {
		portfolio.GithubURL = *req.GithubURL
	}
	if req.TechStack != nil {
		portfolio.TechStack = *req.TechStack
	}
	if req.SortOrder != nil {
		portfolio.SortOrder = *req.SortOrder
	}
	if req.Published != nil {
		portfolio.Published = *req.Published
	}

	if err := h.repos.Portfolio.Update(portfolio); err != nil {
		Error(c, http.StatusInternalServerError, "更新作品集失败")
		return
	}

	Success(c, portfolio)
}

// DeletePortfolio 删除作品集
func (h *Handler) DeletePortfolio(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		Error(c, http.StatusBadRequest, "无效的作品集ID")
		return
	}

	if err := h.repos.Portfolio.Delete(uint(id)); err != nil {
		Error(c, http.StatusInternalServerError, "删除作品集失败")
		return
	}

	Success(c, nil)
}
