package response

type UserInSystem struct {
	UserName  string `json:"username"`
	IsAdmin   bool   `json:"is_admin"`
	Group     string `json:"group"`
	CreatedAt string `json:"created_at"`
}
