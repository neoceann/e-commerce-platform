package service

import "errors"

var ErrUserAlreadyExists = errors.New("user already exists")
var ErrUserNotFound = errors.New("user not found")
var ErrTokenInvalid = errors.New("token is invalid or expired")
var ErrCredentialsInvalid = errors.New("invalid credentials")
var ErrInvalidOldPwd = errors.New("Invalid old password")
