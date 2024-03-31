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
    e.GET("/workspace", handler.GetAllWorkspaces)
    e.DELETE("/workspace", handler.DeleteWorkspace)
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

func (h *WokspaceHandler) GetAllWorkspaces(c echo.Context) (error) {
	workspaces, err := h.service.GetAllWorkspaces(c.Request().Context())

	if err != nil {
        echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, workspaces)
}

func (h *WokspaceHandler) DeleteWorkspace(c echo.Context) (error) {
    workspace := &domain.Workspace{}

	if err := c.Bind(workspace); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid workspace data")
	}

	err := h.service.DeleteWorkspace(c.Request().Context(), *workspace)

	if err != nil {
        echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, workspace)
}
