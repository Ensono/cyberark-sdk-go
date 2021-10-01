package models

type AuthReq struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	Newpassword       string `json:"newPassword,omitempty"`
	ConcurrentSession bool   `json:"concurrentSession,omitempty"`
}

type AuthToken string
