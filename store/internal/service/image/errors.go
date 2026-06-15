package service

import "errors"

var ErrInvalidImageData = errors.New("invalid image data")

var ErrInvalidID = errors.New("invalid UUID")

var ErrImageNotFound = errors.New("image not found")
