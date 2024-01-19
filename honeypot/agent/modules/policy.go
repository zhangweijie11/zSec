package modules

import (
	"github.com/zhangweijie11/zSec/honeypot/agent/models"
	"github.com/zhangweijie11/zSec/honeypot/agent/vars"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

// read policy data from local yaml file
func ReadPolicyFromYaml() (data *models.PolicyData, err error) {
	data = new(models.PolicyData)
	var content []byte

	content, err = ioutil.ReadFile(filepath.Join(vars.CurDir, "conf", "policy.yaml"))
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, data)
	return data, err

}

func LoadPolicy() (*models.PolicyData, error) {
	var err error
	vars.PolicyData, err = ReadPolicyFromYaml()
	if err != nil {
		return nil, err
	}

	if len(vars.PolicyData.Service) > 0 {
		for _, service := range vars.Services {
			vars.Services = append(vars.Services, service)
		}
	}

	if len(vars.PolicyData.Policy) > 0 {
		vars.HoneypotPolicy = vars.PolicyData.Policy[0]
	}

	return vars.PolicyData, err
}
