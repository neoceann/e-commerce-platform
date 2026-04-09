package service

import "errors"

var ErrInvalidCategoryData = errors.New("invalid category data")

var ErrInvalidID = errors.New("invalid UUID")

var ErrCategoryNotFound = errors.New("category not found")