package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type commandListRes struct {
	Status string     `json:"status"`
	Msg    string     `json:"msg"`
	Data   []*Command `json:"data"`
}

type Field struct {
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Desc     string   `json:"desc"`
	Black    bool     `json:"black"`
	Choices  []string `json:"choices"`
	HelpText string   `json:"help_text"`
	Default  any      `json:"default"`
}
type Command struct {
	Name   string   `json:"name"` // 路径
	Desc   string   `json:"desc"`
	Fields []*Field `json:"fields"`
}

func queryCommandServerCommandList(addr string) ([]*Command, error) {
	// todo 请求指令对应的服务器commandlist
	url := addr + "/commandlist"
	if url != "" {
		resp := &commandListRes{}
		err := httpGet(url, resp)
		if err != nil {
			return nil, err
		}

		return resp.Data, nil
	}
	cmd1 := &Command{
		Name: "/get_player_data",
		Desc: "通过id获取玩家数据",
	}
	cmd1.Fields = append(cmd1.Fields, &Field{
		Name:     "player_id",
		Type:     "int64",
		Desc:     "玩家id",
		Black:    false,
		Choices:  nil,
		HelpText: "识别码",
		Default:  "-1",
	})
	cmd1.Fields = append(cmd1.Fields, &Field{
		Name:     "data_type",
		Type:     "int",
		Desc:     "数据类型，四种可选：base|bag|team|all",
		Black:    true,
		Choices:  []string{"base", "bag", "team", "all"},
		HelpText: "获取的数据类型",
		Default:  "base",
	})

	cmd2 := &Command{
		Name: "/set_player_data",
		Desc: "通过id设置玩家数据",
	}
	cmd2.Fields = append(cmd2.Fields, &Field{
		Name:     "player_id",
		Type:     "int64",
		Desc:     "玩家id",
		Black:    false,
		Choices:  nil,
		HelpText: "识别码",
		Default:  "-1",
	})
	cmd2.Fields = append(cmd2.Fields, &Field{
		Name:     "data_type",
		Type:     "int",
		Desc:     "数据类型，四种可选：base|bag|team|all",
		Black:    true,
		Choices:  []string{"base", "bag", "team", "all"},
		HelpText: "设置的数据类型",
		Default:  "base",
	})

	resp := make([]*Command, 0)
	resp = append(resp, cmd1)
	resp = append(resp, cmd2)

	return resp, nil
}

func execCommandServerCommand(addr, path, params string) (resp map[string]interface{}, code int, err error) {
	url := addr + path + "?" + params
	code = 200
	resp = make(map[string]interface{})
	err = httpGet(url, &resp)
	if err != nil {
		return nil, 300, err
	}
	return
	str := ""
	for i := 0; i < 100; i++ {
		str += strconv.Itoa(i)
	}
	return map[string]interface{}{"result": "ok", "msg": str}, 200, nil
	reqUrl := addr + path + params
	execRes, err := http.Get(reqUrl)
	if err != nil {

	}
	defer execRes.Body.Close()
	_, err = io.ReadAll(execRes.Body)
	if err != nil {

	}

	return nil, 200, nil
}

func httpGet(url string, respReceiver interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http get error:%v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body error:%v", err)
	}

	//fmt.Printf("exec %v resp %v\n", url, string(body))

	err = json.Unmarshal(body, respReceiver)
	if err != nil {
		return fmt.Errorf("json unmarshal error:%v", err)
	}

	return nil
}
