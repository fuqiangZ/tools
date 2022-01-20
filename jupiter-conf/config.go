
package jupiterConf

import (
	"os"
	"github.com/BurntSushi/toml"
	"github.com/douyu/jupiter/pkg/conf"
	"github.com/douyu/jupiter/pkg/xlog"
)

// InitConf 兼容生产环境直接运行，或者单元测试
func InitConf(path string) error {
	if path == "" {
		path = "./config.toml"
	}
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		return conf.LoadFromReader(f, toml.Unmarshal)
	}
	// 调试场景，用指定目录的
	f, err = os.Open("../config.toml")
	if err == nil {
		defer f.Close()
		return conf.LoadFromReader(f, toml.Unmarshal)
	}
	xlog.Errorf("open config file error; %v", err)
	return err
}