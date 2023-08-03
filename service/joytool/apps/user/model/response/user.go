package response

type UserInSystem struct {
	UserName   string `json:"username"`
	IsAdmin    bool   `json:"is_admin"`
	IsAdminStr string `json:"is_admin_str"`
	IsMyInfo   bool   `json:"is_my_info"`
	Group      string `json:"group"`
	CreatedAt  string `json:"created_at"`
}
