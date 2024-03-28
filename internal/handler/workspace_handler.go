package handler

import (
	"net/http"

	"github.com/gimmefear/dswv3/internal/domain"
	"github.com/gimmefear/dswv3/internal/services"
	"github.com/labstack/echo/v4"
)

type WokspaceHandler struct {
	service *services.WorkspaceService
}

func NewWorkspaceHandler(e *echo.Echo, service *services.WorkspaceService) {
	handler := &WokspaceHandler{service: service}

	e.POST("/workspace", handler.CreateWorkspace)
}

func (h *WokspaceHandler) CreateWorkspace(c echo.Context) error {
	workspace := &domain.Workspace{}
	if err := c.Bind(workspace); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid workspace data")
	}

	err := h.service.CreateWorkspace(c.Request().Context(), *workspace)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Workspace creation failed")
	}

	return c.JSON(http.StatusOK, workspace)
}
