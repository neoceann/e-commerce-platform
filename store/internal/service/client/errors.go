package service

import "errors"

var ErrInvalidClientData = errors.New("invalid client data")

var ErrInvalidID = errors.New("invalid UUID")

var ErrClientNotFound = errors.New("client not found")

var ErrInvalidPagination = errors.New("invalid pagination data (valid limit: 1-100)")

var ErrInvalidAddrData = errors.New("invalid address data")