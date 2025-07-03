package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sakura/record-api/api"
	"github.com/sakura/record-api/models"
	"github.com/sakura/record-api/services"
)

// 创建分类服务实例
var categoryService = services.NewCategoryService()

// GetCategories 获取分类列表
func GetCategories(c *gin.Context) {
	// 获取查询参数
	categoryType := c.Query("type")

	var categories []models.Category
	var err error

	if categoryType != "" {
		// 根据类型获取分类
		categories, err = categoryService.GetCategoriesByType(categoryType)
	} else {
		// 获取所有分类
		categories, err = categoryService.GetAllCategories()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, api.SuccessResponse(gin.H{
		"categories": categories,
	}))
}

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	// 绑定请求参数
	var req api.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的请求参数: "+err.Error()))
		return
	}

	// 调用服务层创建分类
	category, err := categoryService.CreateCategory(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusCreated, api.SuccessResponse(gin.H{
		"category": category,
	}))
}

// UpdateCategory 更新分类
func UpdateCategory(c *gin.Context) {
	// 获取分类ID
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的分类ID"))
		return
	}

	// 绑定请求参数
	var req api.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的请求参数: "+err.Error()))
		return
	}

	// 调用服务层更新分类
	category, err := categoryService.UpdateCategory(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, api.SuccessResponse(gin.H{
		"category": category,
	}))
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	// 获取分类ID
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的分类ID"))
		return
	}

	// 调用服务层删除分类
	if err := categoryService.DeleteCategory(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, api.SuccessResponse(gin.H{
		"message": "分类删除成功",
	}))
}