package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/madsilver/task-manager/internal/adapter/presenter"
	"github.com/madsilver/task-manager/internal/domain"
	"github.com/madsilver/task-manager/internal/infrastructure/env"
	"github.com/madsilver/task-manager/internal/usecase"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

type UserHandler interface {
	Create(c ContextHTTP) error
	Login(c ContextHTTP) error
}

type userHandler struct {
	uc  usecase.UserUsecase
	log Logger
}

func NewUserHandler(uc usecase.UserUsecase, log Logger) UserHandler {
	return &userHandler{uc, log}
}

// Create godoc
// @Summary Create user.
// @Description Create new users.
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param User body domain.User true "User payload"
// @Success 201 {object} presenter.UserResponse
// @Failure 400 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /users [post]
func (h *userHandler) Create(c ContextHTTP) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		h.log.Error(c, "failed to bind user", err)
		return c.JSON(http.StatusBadRequest, presenter.ErrorResponse{Error: "invalid request"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		h.log.Error(c, "failed to hash password", err)
		return c.JSON(http.StatusInternalServerError, presenter.ErrorResponse{Error: "failed to process request"})
	}
	user.Password = string(hashedPassword)

	if err := h.uc.CreateUser(c, &user); err != nil {
		if h.EmailExistsError(c, err) {
			return c.JSON(http.StatusBadRequest, presenter.ErrorResponse{Error: "email already exists"})
		}
		h.log.Error(c, "failed to create user", err)
		return c.JSON(http.StatusInternalServerError, presenter.ErrorResponse{Error: "failed to create user"})
	}

	return c.JSON(http.StatusCreated, presenter.ToUserResponse(&user))
}

// Login godoc
// @Summary Authenticate user and return JWT.
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body presenter.LoginRequest true "User credentials"
// @Success 200 {object} presenter.LoginResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /login [post]
func (h *userHandler) Login(c ContextHTTP) error {
	var req presenter.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, presenter.ErrorResponse{Error: "invalid credentials"})
	}

	user, err := h.uc.GetUserByEmail(c, req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, presenter.ErrorResponse{Error: "invalid email or password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, presenter.ErrorResponse{Error: "invalid email or password"})
	}

	signed, err := h.GetSignedToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, presenter.ErrorResponse{Error: "failed to sign token"})
	}

	return c.JSON(http.StatusOK, presenter.LoginResponse{Token: signed})
}

// EmailExistsError checks if the error indicates that the email already exists.
// c: the context for logging
// err: the error to check
// returns: true if the email already exists, false otherwise
func (h *userHandler) EmailExistsError(c ContextHTTP, err error) bool {
	m := err.Error()
	if strings.Contains(m, "duplicate key") && strings.Contains(m, "email") {
		h.log.Error(c, "email already exists", err)
		return true
	}
	return false
}

// GetSignedToken generates a signed JWT token for the user.
// user: the user for whom the token is generated
// returns: the signed JWT token or an error if signing fails
func (h *userHandler) GetSignedToken(user *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"roles": user.Roles,
		"exp":   time.Now().Add(env.App.JwtExpiresMinutes * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(env.App.JwtSecret))
}
