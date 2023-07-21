package do

type Field struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Command struct {
	Name   string   `json:"name"` // 路径
	Desc   string   `json:"desc"`
	Fields []*Field `json:"fields"`
}

type PreExecCommand struct {
	CommandBaseInfo   *Command `json:"base_req_params"`
	Project           string   `json:"project"`
	CommandServerName string   `json:"command_server_name"`
	Env               string   `json:"env"`
	User              string   `json:"user"`
	Date              string   `json:"date"`
}

type PostExecCommand struct {
	RequestInfo *PreExecCommand `json:"request_info"`
	ExecResCode int             `json:"exec_res_code,omitempty"`
	ExecRes     string          `json:"exec_res,omitempty"`
}
