package request

import "joytool/apps/gmtool/model/do"

type HomeView struct {
	Project string `json:"project"`
}

type CommandServerData struct {
	Project string   `json:"project"`
	Name    string   `json:"name"` // project内唯一
	Desc    string   `json:"desc"`
	Addr    string   `json:"addr"`
	Envs    []string `json:"envs"` // 环境标签，一个server可以挂钩多个env
}

type DeleteCommandServer struct {
	Project string `json:"project"`
	Name    string `json:"name"` // project内唯一
}

type EnvData struct {
	Project string `json:"project"`
	Name    string `json:"name"` // project内唯一
	Desc    string `json:"desc"`
	Index   int    `json:"index"` // 用于排序的索引
}

type DeleteEnv struct {
	Project string `json:"project"`
	Name    string `json:"name"` // project内唯一
}

type CommandList struct {
	Project string `json:"project"`
	Name    string `json:"name"` // project内唯一
}

type ExecCommand struct {
	Project           string      `json:"project"`
	CommandServerName string      `json:"command_server_name"`
	Env               string      `json:"env"`
	ExecBaseInfo      *do.Command `json:"base_info"`
}

type DeleteLikeCommand struct {
	Project string `json:"project"`
	ID      string `json:"id"`
}

type DeleteExecHistory struct {
	Project string `json:"project"`
	Name    string `json:"name"`  // project内唯一
	Index   int    `json:"index"` // 第几个索引处的历史
}

type PermissionGroupList struct {
	Project string `json:"project"`
}

type PermissionGroupData struct {
	Project     string           `json:"project"`
	Name        string           `json:"name"`
	Desc        string           `json:"desc"`
	Permissions []*do.Permission `json:"permissions"`
}

type DeletePermissionGroup struct {
	Project string `json:"project"`
	Name    string `json:"name"`
}
