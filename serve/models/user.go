/**
 * @Author: mariomang
 * @Date: 2023-12-05 15:25:11
 * @File: domain/user.go
**/

package domain

import "time"

type User struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Nickname string    `json:"nickname"`
	Email    string    `json:"email"`
	Birthday string    `json:"birthday"`
	CreateAt time.Time `json:"create_at"`
}

func (u *User) Exists() bool {
	return u != nil && u.Username != ""
}
