package auth

// AccessTokenHeaderModel 请求头的token
type AccessTokenHeaderModel struct {
	Token	string	`header:"token" binding:"required"`
}
