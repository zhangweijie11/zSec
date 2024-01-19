package vars

import (
	"github.com/zhangweijie11/zSec/honeypot/agent/models"
	"github.com/zhangweijie11/zSec/honeypot/agent/util"
)

var (
	HoneypotPolicy models.Policy
	PolicyData     *models.PolicyData

	Services = make([]models.BackendService, 0)
	CurDir   string
)

func init() {
	CurDir, _ = util.GetCurDir()
}
