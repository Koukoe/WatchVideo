package service

import (
	"errors"
	"strconv"
	"strings"

	"WatchVideo/biz/dao/db"
	"WatchVideo/biz/model/api"
	"WatchVideo/biz/model/store"
	"WatchVideo/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidParams = errors.New("invalid params")
	ErrUserExists    = errors.New("user already exists")
	ErrUnauthorized  = errors.New("unauthorized")
)

// 用户注册
func Register(req *api.RegisterRequest) (*api.RegisterResponse, error) {
	username := strings.TrimSpace(req.Username)
	password := strings.TrimSpace(req.Password)

	if username == "" || password == "" {
		return &api.RegisterResponse{}, ErrInvalidParams
	}

	existing, err := db.GetUserByUsername(username)
	if err != nil {
		return &api.RegisterResponse{}, err
	}
	if existing != nil {
		return &api.RegisterResponse{}, ErrUserExists
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return &api.RegisterResponse{}, err
	}

	nextID, err := db.NextUserID()
	if err != nil {
		return &api.RegisterResponse{}, err
	}
	userID := strconv.FormatInt(nextID, 10)

	user := &store.User{
		ID:           userID,
		Username:     username,
		PasswordHash: string(hashed),
	}

	if err := db.CreateUser(user); err != nil {
		return &api.RegisterResponse{}, err
	}

	return &api.RegisterResponse{
		UserID: userID,
	}, nil
}

// 用户登录
func Login(req *api.LoginRequest) (*api.LoginResponse, error) {
	username := strings.TrimSpace(req.Username)
	password := strings.TrimSpace(req.Password)

	if username == "" || password == "" {
		return &api.LoginResponse{}, ErrInvalidParams
	}

	user, err := db.GetUserByUsername(username)
	if err != nil {
		return &api.LoginResponse{}, err
	}
	if user == nil {
		return &api.LoginResponse{}, ErrInvalidParams
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return &api.LoginResponse{}, ErrInvalidParams
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(user.ID)
	if err != nil {
		return &api.LoginResponse{}, err
	}

	return &api.LoginResponse{
		UserID:       user.ID,
		Username:     user.Username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// 根据用户ID获取用户信息
func GetUserInfoByID(userID string) (*api.UserInfoResponse, error) {
	if strings.TrimSpace(userID) == "" {
		return nil, ErrInvalidParams
	}

	user, err := db.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidParams
	}

	resp := &api.UserInfoResponse{
		UserID:    user.ID,
		Username:  user.Username,
		AvatarURL: user.AvatarURL,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	if user.DeletedAt.Valid {
		resp.DeletedAt = user.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return resp, nil
}

// 上传用户头像
func UploadAvatar(userID, avatarURL string) (*api.UserInfoResponse, error) {
	if strings.TrimSpace(userID) == "" || strings.TrimSpace(avatarURL) == "" {
		return nil, ErrInvalidParams
	}

	if err := db.UpdateUserAvatar(userID, avatarURL); err != nil {
		return nil, err
	}

	user, err := db.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidParams
	}

	resp := &api.UserInfoResponse{
		UserID:    user.ID,
		Username:  user.Username,
		AvatarURL: user.AvatarURL,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	if user.DeletedAt.Valid {
		resp.DeletedAt = user.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return resp, nil
}
