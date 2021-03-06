// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessExpire int64  `json:"accessExpire"`
	AccessToken  string `json:"accessToken"`
}

type RegisterRequest struct {
	Gender   int64  `json:"gender"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID     int64  `json:"id"`
	Gender int64  `json:"gender"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

type UserInfoRequest struct {
	ID int64 `json:"id"`
}

type UserInfoResponse struct {
	ID     int64  `json:"id"`
	Gender int64  `json:"gender"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}
