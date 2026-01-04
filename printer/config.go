package printer

import (
	"sync"

	"github.com/zerroi/pdf-to-printer/internal"
)

// Config 配置结构
type Config struct {
	SumatraPath string // SumatraPDF可执行文件路径（为空则自动检测）
}

var (
	globalConfig Config
	configMutex  sync.RWMutex
)

// SetConfig 设置全局配置
func SetConfig(config Config) error {
	configMutex.Lock()
	defer configMutex.Unlock()

	// 如果指定了路径，验证路径是否存在
	if config.SumatraPath != "" {
		if !internal.FileExists(config.SumatraPath) {
			return ErrSumatraNotFound
		}
	}

	globalConfig = config
	return nil
}

// GetConfig 获取当前配置
func GetConfig() Config {
	configMutex.RLock()
	defer configMutex.RUnlock()
	return globalConfig
}

// resetConfig 重置配置（用于测试）
func resetConfig() {
	configMutex.Lock()
	defer configMutex.Unlock()
	globalConfig = Config{}
}
