package server

import (
	"github.com/labstack/echo/v4"
	"github.com/madsilver/task-manager/internal/adapter/presenter"
	"net/http"
)

const (
	UserContext = "user"
	RoleContext = "role"
	AdminRole   = "manager"
)

func ValidateHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := ctx.Request().Header.Get("x-user-id")
		if user == "" {
			return ctx.JSON(http.StatusBadRequest, presenter.NewErrorResponse("x-user-id header missing", ""))
		}
		role := ctx.Request().Header.Get("x-role")
		if role == "" {
			return ctx.JSON(http.StatusBadRequest, presenter.NewErrorResponse("x-role header missing", ""))
		}
		ctx.Set(UserContext, user)
		ctx.Set(RoleContext, role)
		return next(ctx)
	}
}

func AuthAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Get(RoleContext) != AdminRole {
			return ctx.JSON(http.StatusForbidden, presenter.NewErrorResponse("forbidden", ""))
		}
		return next(ctx)
	}
}

func AuthAdminOrOwner(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userParam := ctx.QueryParam("user")
		if ctx.Get(RoleContext) == AdminRole ||
			ctx.Get(UserContext) == userParam {
			return next(ctx)
		}

		return ctx.JSON(http.StatusForbidden, presenter.NewErrorResponse("forbidden", "access to task not allowed"))
	}
}
