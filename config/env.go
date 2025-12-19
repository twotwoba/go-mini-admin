package config

import "os"

type Env string

const (
	Dev  Env = "dev"  // 开发
	Beta Env = "beta" // 测试
	Prod Env = "prod" // 生产
)

const (
	GMA_APP_ENV  = "GMA_APP_ENV"  // 应用环境变量 dev beta prod
	GMA_ENV_FILE = "GMA_ENV_FILE" // 配置文件路径 ./config.dev.yaml
)

func GetEnv() Env {
	env := os.Getenv(GMA_APP_ENV)
	switch env {
	case string(Beta):
		return Beta
	case string(Prod):
		return Prod
	default:
		return Dev
	}
}

func GetConfigPath() string {
	return os.Getenv(GMA_ENV_FILE)
}
