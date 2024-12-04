package config

type Redis struct {
	Addr         string `yaml:"addr"`         // 地址
	Port         string `yaml:"port"`         // 端口
	Password     string `yaml:"password"`     // 密码（如果没有密码则为空）
	DB           int    `yaml:"db"`           // 数据库编号
	PoolSize     int    `yaml:"pool_size"`    // 连接池大小
	MinIdleConns int    `yaml:"minIdleConns"` // 最小空闲连接数
}
