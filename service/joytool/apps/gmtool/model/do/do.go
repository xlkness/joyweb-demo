package do

// Env 环境数据
type Env struct {
	Project   string `json:"project"`
	Name      string `json:"name"` // project内唯一
	Desc      string `json:"desc"`
	Index     int    `json:"index"` // 用于排序的索引
	CreatedAt string `json:"created_at"`
}

// CommandServer 指令服
type CommandServer struct {
	Project         string             `json:"project"`
	Name            string             `json:"name"` // project内唯一
	Desc            string             `json:"desc"`
	Addr            string             `json:"addr"`
	Envs            []string           `json:"envs"` // 环境标签，一个server可以挂钩多个env
	CreatedAt       string             `json:"created_at"`
	ExecHistoryList []*PostExecCommand `json:"exec_history_list"` // 保存最近执行的
}

// LikeExecCommand 收藏某个执行的指令
type LikeExecCommand struct {
	Project string          `json:"project"`
	User    string          `json:"user"`
	ID      int             `json:"id"`
	Command *PreExecCommand `json:"command"`
}

type Permission struct {
	Env           string `json:"env"`
	CommandServer string `json:"command_server"`
	Action        string `json:"action"` // write|read
}

type PermissionGroupData struct {
	Project     string        `json:"project"`
	Name        string        `json:"name"`
	Desc        string        `json:"desc"`
	Permissions []*Permission `json:"permissions"`
	CreatedAt   string        `json:"created_at"`
}
