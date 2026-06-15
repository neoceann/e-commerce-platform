package service

import "errors"

var ErrInvalidID = errors.New("invalid UUID")

var ErrAddrNotFound = errors.New("address not found")
