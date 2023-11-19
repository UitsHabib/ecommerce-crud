package rest

import (
	"fmt"
	"net/http"

	"github.com/UitsHabib/ecommerce-crud/logger"
	"github.com/UitsHabib/ecommerce-crud/service"
	"github.com/UitsHabib/ecommerce-crud/util"
	"github.com/gin-gonic/gin"
)

func (s *Server) createCategory(ctx *gin.Context) {
	var req createCategoryReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	// TODO: need to check whether the parent category id exists

	ctgry := &service.Category{
		Name:      req.Name,
		ParentID:  req.ParentID,
		StatusID:  req.StatusID,
		CreatedAt: util.GetCurrentTimestamp(),
	}

	newCategory, err := s.svc.AddCategory(ctx, ctgry)
	if err != nil {
		logger.Error(ctx, "cannot add category", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", newCategory)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully created", newCategory))
}

func (s *Server) getCategory(ctx *gin.Context) {
	var req getCategoryReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	ctgry, err := s.svc.GetCategory(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot get category", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", ctgry)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully fetched", ctgry))
}

func (s *Server) getCategories(ctx *gin.Context) {
	var req getCategoriesReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	result, err := s.svc.GetCategories(ctx, req.Page, req.Limit)
	if err != nil {
		logger.Error(ctx, "cannot get categories", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", result)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Fetched categories", result))
}

func (s *Server) updateCategory(ctx *gin.Context) {
	var req updateCategoryReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	ctgryID := ctx.Param("id")
	if len(ctgryID) == 0 {
		logger.Error(ctx, "cannot pass validation", ctgryID)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Invalid brand ID", "Bad request"))
		return
	}

	logger.Info(ctx, fmt.Sprintf("req payload for categoryID: %s", ctgryID), req)

	ctgry, err := s.svc.GetCategory(ctx, ctgryID)
	if err != nil {
		logger.Error(ctx, "cannot get category", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if ctgry == nil {
		logger.Error(ctx, "category not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, "Category Not Found", "Not found"))
		return
	}

	// update category
	ctgry.Name = req.Name
	ctgry.StatusID = req.StatusID

	err = s.svc.UpdateCategory(ctx, ctgryID, ctgry)
	if err != nil {
		logger.Error(ctx, "cannot update category", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", ctgry)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully updated", ctgry))
}

func (s *Server) deleteCategory(ctx *gin.Context) {
	var req deleteCategoryReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Response(ctx, "Api parameter invalid", err))
		return
	}

	logger.Info(ctx, "req payload", req)

	ctgry, err := s.svc.GetCategory(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot get category", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	if ctgry == nil {
		logger.Error(ctx, "category not found", nil)
		ctx.JSON(http.StatusNotFound, s.svc.Response(ctx, "Category Not Found", "Not found"))
		return
	}

	err = s.svc.DeleteCategory(ctx, req.ID)
	if err != nil {
		logger.Error(ctx, "cannot delete category", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, "Internal Server Error", err))
		return
	}

	logger.Info(ctx, "res payload", ctgry)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully deleted", ctgry))
}
