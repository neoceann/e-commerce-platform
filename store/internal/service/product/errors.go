package service

import "errors"

var ErrInvalidProductData = errors.New("invalid product data")

var ErrInvalidID = errors.New("invalid UUID")

var ErrProductNotFound = errors.New("product not found")

var ErrIncreaseFailed = errors.New("failed to increase product stock")

var ErrDecreaseFailed = errors.New("failed to decrease product stock")

var ErrBadPrice = errors.New("invalid price")

var ErrBadStock = errors.New("invalid stock")