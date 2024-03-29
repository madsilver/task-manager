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
	TechRole    = "technician"
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

func AuthRole(roles ...string) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			for _, role := range roles {
				if ctx.Get(RoleContext) == role {
					return next(ctx)
				}
			}
			return ctx.JSON(http.StatusForbidden, presenter.NewErrorResponse("forbidden", ""))
		}
	}
}

func AuthOwnerTask(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userParam := ctx.QueryParam("user")
		if ctx.Get(RoleContext) == AdminRole ||
			ctx.Get(UserContext) == userParam {
			return next(ctx)
		}

		return ctx.JSON(http.StatusForbidden, presenter.NewErrorResponse("forbidden", "access to task not allowed"))
	}
}
