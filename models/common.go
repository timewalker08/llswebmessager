package models

import (
    "fmt"
)

const (
    FriendStatusNormal = 1
    FriendStatusHide = 2
    FriendStatusDeleted = 3
)

var NormalFriendStatus *Friendstatus = &Friendstatus{Id: FriendStatusNormal}
var HideFriendStatus *Friendstatus = &Friendstatus{Id: FriendStatusHide}
var DeletedFriendStatus *Friendstatus = &Friendstatus{Id: FriendStatusDeleted}

var LoginSessionKey string = "LLSLoginSessionKey"

type WebApiResult struct {
    Code     int
	Msg      string
}

type CommonError struct {
    ErrorMsg    string
}

func (ce *CommonError) Error() string {
	strFormat := `Error message: %s`
	return fmt.Sprintf(strFormat, ce.ErrorMsg)
}

type AccountNotExistError struct {
    AccountName string
}

func (ae *AccountNotExistError) Error() string {
	return fmt.Sprintf("%s doesnot exist.", ae.AccountName)
}

type AddFriendFailedError struct {
    UserName   string
	FriendName string
}

func (this *AddFriendFailedError) Error() string {
	return fmt.Sprintf("Failed to add friend %s to %s.", this.FriendName, this.UserName)
}

type LoginError struct {
    UserName string
	Password string
}

func (this *LoginError) Error() string {
    return fmt.Sprintf("Login failed. Name of password is incorrect.")
}

type NameAlreadyUsedWhenRegisterError struct {
    UserName string
}

func (this *NameAlreadyUsedWhenRegisterError) Error() string {
    return fmt.Sprintf("User name %s has already been used. Please use another name.", this.UserName)
}

type InvalidUserInfoWhenRegister struct {
    UserName string
	Password string
	ConfirmPassword string
	ErrorMessage string
}

func (this *InvalidUserInfoWhenRegister) Error() string {
    return this.ErrorMessage
}