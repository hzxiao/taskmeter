package config

import (
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"strings"
	"fmt"
)

type Config struct {
	Name string
}

func Init(cfgName string) error {
	c := &Config{
		Name: cfgName,
	}
	err := c.initConfig()
	if err != nil {
		return err
	}

	c.initLog()
	c.Print()
	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")     // 设置配置文件格式为YAML
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix("TASKMETER") // 读取环境变量的前缀为TASKMETER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}

func (c *Config) Print()  {
	for k, v := range viper.AllSettings() {
		fmt.Printf("%v = %v\n", k, v)
	}
}
func GetString(key string) string {
	return viper.GetString(key)
}
