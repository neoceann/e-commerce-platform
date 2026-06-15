package service

import "auth-service/internal/repository/db"

func RegisterInfoFromServiceToDB(r *RegisterRequest) *db.CreateUserParams {
	return &db.CreateUserParams{
		Email:     r.Email,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Phone:     r.Phone,
	}
}
