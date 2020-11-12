// 加载配置文件，目前支持json
// author: ericlang
package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

func init() {
	fmt.Println("xxxxxxxxx")
}

// LoadTomlConfig 读取配置文件内容，并将内容读取至 ConfigInfo
func LoadTomlConfig(fileName string, configContent interface{}) (err error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("config file open failed, filename:", fileName, ", err:", err.Error())
		return err
	}
	defer file.Close()
	// 读取配置文件内容
	_, err = toml.DecodeReader(file, configContent)
	if err != nil {
		fmt.Println("config toml file decode failed, filename:", fileName, " err: ", err.Error())
		return err
	}
	return nil
}

func String(cfg interface{}) string {
	contentBytes, err := json.Marshal(cfg)
	if err != nil {
		fmt.Println("marshal json failed, check your data's format, err:", err.Error())
		return string("")
	}
	return string(contentBytes)
}
