package common

import ut "github.com/go-playground/universal-translator"

type ContextModel struct {
	LoggerModel          LoggerModel
	AuthAccessTokenModel AuthAccessTokenModel
	PermissionHave       string
	Translator           ut.Translator
}

type AuthAccessTokenModel struct {
	ResourceUserID int64  `json:"rid"`
	Scope          string `json:"scp"`
	ClientID       string `json:"cid"`
	Locale         string `json:"lang"`
}
