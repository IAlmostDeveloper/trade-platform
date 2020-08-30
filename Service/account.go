package service

import (
	dbaccess "trade-platform/DBAccess"
	entities "trade-platform/Entities"
)

func Authenticate(authRequest entities.AuthRequestJson) (bool, entities.UserCredentials){
	user := dbaccess.FindUserByLoginAndPassword(authRequest)
	return user.Id !=0, user
}

func Register(regRequest entities.RegRequestJson) bool{
	if dbaccess.FindUserByLogin(regRequest.Login).Id==0{
		dbaccess.InsertUser(regRequest)
		return true
	}
	return false
}
