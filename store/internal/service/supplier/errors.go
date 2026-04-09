package service

import "errors"

var ErrInvalidSupplierData = errors.New("invalid supplier data")

var ErrInvalidID = errors.New("invalid UUID")

var ErrSupplierNotFound = errors.New("supplier not found")

var ErrInvalidAddrData = errors.New("invalid address data")