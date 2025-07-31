package logic

import (
	"big_mall_api/internal/model"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
)

func (s *ServerLogic) CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.UserResponse, error) {
	// 密码加密
	hashedPassword := s.hashPassword(req.Password)

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Age:      req.Age,
		Status:   1,
	}

	// 保存到数据库
	err := s.mainMdb.DB.WithContext(ctx).Save(user).Error
	if err != nil {
		return nil, err
	}
	return s.userToResponse(user), nil
}

func (s *ServerLogic) GetUser(ctx context.Context, id uint) (*model.UserResponse, error) {
	// 先从缓存获取
	var user *model.User
	key := fmt.Sprintf("user:%d", id)
	data, err := s.mainRdb.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(data), &user)
	if err == nil {
		return s.userToResponse(user), nil
	}

	// 从数据库获取
	err = s.mainMdb.DB.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	//更新用户缓存
	value, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	s.mainRdb.Set(ctx, key, value, 0)

	return s.userToResponse(user), nil
}

func (s *ServerLogic) DeleteUser(ctx context.Context, id uint) error {
	if err := s.mainMdb.DB.WithContext(ctx).Delete(id).Error; err != nil {
		return err
	}

	// 删除缓存
	go func() {
		key := fmt.Sprintf("user:%d", id)
		_ = s.mainRdb.Del(ctx, key)
	}()

	return nil
}

func (s *ServerLogic) ListUsers(ctx context.Context, page, pageSize int) ([]*model.UserResponse, error) {
	offset := (page - 1) * pageSize
	var users []*model.User
	err := s.mainMdb.DB.WithContext(ctx).Offset(offset).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, err
	}

	responses := make([]*model.UserResponse, len(users))
	for i, user := range users {
		responses[i] = s.userToResponse(user)
	}

	return responses, nil
}

func (s *ServerLogic) hashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return fmt.Sprintf("%x", hash)
}

func (s *ServerLogic) userToResponse(user *model.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
