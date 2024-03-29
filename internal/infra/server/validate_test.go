package server

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/madsilver/task-manager/internal/adapter/presenter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidateHeader(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("x-user-id", "1")
	req.Header.Set("x-role", "manager")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err := ValidateHeader(func(ctx echo.Context) error {
		assert.Equal(t, "1", ctx.Get("user"))
		assert.Equal(t, "manager", ctx.Get("role"))
		return nil
	})(ctx)

	response := &presenter.ErrorResponse{}
	_ = json.Unmarshal(rec.Body.Bytes(), response)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestValidateHeader_UserMissing(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("x-role", "admin")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err := ValidateHeader(nil)(ctx)

	response := &presenter.ErrorResponse{}
	_ = json.Unmarshal(rec.Body.Bytes(), response)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "x-user-id header missing", *response.Error)
}

func TestValidateHeader_RoleMissing(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("x-user-id", "1")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err := ValidateHeader(nil)(ctx)

	response := &presenter.ErrorResponse{}
	_ = json.Unmarshal(rec.Body.Bytes(), response)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "x-role header missing", *response.Error)
}

func TestAuthAdmin(t *testing.T) {
	e := echo.New()
	ctx := e.NewContext(nil, nil)
	ctx.Set("role", "manager")

	err := AuthAdmin(func(ctx echo.Context) error {
		return nil
	})(ctx)

	assert.NoError(t, err)
}

func TestAuthAdmin_Error(t *testing.T) {
	e := echo.New()
	rec := httptest.NewRecorder()
	ctx := e.NewContext(httptest.NewRequest(http.MethodGet, "/", nil), rec)
	ctx.Set("role", "technician")

	err := AuthAdmin(func(ctx echo.Context) error {
		return nil
	})(ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, rec.Code)
}

func TestAuthAdminOrOwner_Manager(t *testing.T) {
	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, rec)
	ctx.Set(RoleContext, "manager")
	ctx.Set(UserContext, "1")
	ctx.QueryParams().Set("user", "99")

	err := AuthAdminOrOwner(func(ctx echo.Context) error {
		return nil
	})(ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestAuthAdminOrOwner_Technician(t *testing.T) {
	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, rec)
	ctx.Set(RoleContext, "technician")
	ctx.Set(UserContext, "1")
	ctx.QueryParams().Set("user", "1")

	err := AuthAdminOrOwner(func(ctx echo.Context) error {
		return nil
	})(ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestAuthAdminOrOwner_Error(t *testing.T) {
	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := e.NewContext(req, rec)
	ctx.Set(RoleContext, "technician")
	ctx.Set(UserContext, "1")
	ctx.QueryParams().Set("user", "99")

	err := AuthAdminOrOwner(func(ctx echo.Context) error {
		return nil
	})(ctx)

	response := &presenter.ErrorResponse{}
	_ = json.Unmarshal(rec.Body.Bytes(), response)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, rec.Code)
	assert.Equal(t, "forbidden", *response.Error)
	assert.Equal(t, "access to task not allowed", *response.Message)
}
