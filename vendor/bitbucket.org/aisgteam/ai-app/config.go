package aiapp

import (
    "fmt"
    "github.com/spf13/viper"
    "os"
    "path/filepath"
    "strings"
)

var (
    EnvPrefix  = "aiapp"
    ConfigType = "yaml"
    ConfigName = "config"
)

type ConfigEnv string

const (
    ConfigEnvProd   ConfigEnv = "prod"
    ConfigEnvTest   ConfigEnv = "test"
    ConfigEnvDev    ConfigEnv = "dev"
)

type RuntimeConfig struct{
    env     ConfigEnv
}

func loadRuntimeConfig() (*RuntimeConfig, error) {
    cfg := &RuntimeConfig{}


    EnvPrefix = strings.ToUpper(EnvPrefix)

    viper.SetEnvPrefix(EnvPrefix)
    viper.AutomaticEnv()
    viper.SetConfigType(ConfigType)
    viper.SetConfigName(ConfigName)
    replacer := strings.NewReplacer(".", "_")
    viper.SetEnvKeyReplacer(replacer)

    env := os.Getenv("ENV")
    env = strings.ToLower(env)
    switch env {
    case "prod":
        cfg.env = ConfigEnvProd
    case "product":
        cfg.env = ConfigEnvProd
    case "test":
        cfg.env = ConfigEnvTest
    case "stage":
        cfg.env = ConfigEnvTest
    default:
        cfg.env = ConfigEnvDev
    }
    prefix := fmt.Sprintf("%v/", cfg.env)
    appPath := filepath.Join(a.dir.LocalPath, "config", prefix)
    envPath := os.Getenv(fmt.Sprintf("%v_%v", EnvPrefix, "CONFIG_PATH"))
    if envPath != "" {
        sep := envPath[len(envPath)-1:]
        if sep != "/" {
            envPath += "/"
        }
        envPath += prefix
        viper.AddConfigPath(envPath)
    }
    viper.AddConfigPath(appPath)
    err := viper.ReadInConfig()
    if err != nil {
        return nil, Err("loadRuntimeConfig", err)
    }
    return cfg, nil
}

func (c *RuntimeConfig) GetString(name string, def string) string {
    v := viper.GetString(name)
    if v == "" {
        return def
    }
    return v
}

func (c *RuntimeConfig) GetStr(name, def string) string {
    return c.GetString(name, def)
}


func (c *RuntimeConfig) Str(name, def string) string {
    return c.GetString(name, def)
}

func (c *RuntimeConfig) GetInt(name string, def int) int {
    v := viper.GetInt(name)
    if v == 0 {
        return def
    }
    return v
}

func (c *RuntimeConfig) Int(name string, def int) int {
    return c.GetInt(name, def)
}

func (c *RuntimeConfig) GetInt64(name string, def int64) int64 {
    v := viper.GetInt64(name)
    if v == 0 {
        return def
    }
    return v
}

func (c *RuntimeConfig) Int64(name string, def int64) int64 {
    return c.GetInt64(name, def)
}

func (c *RuntimeConfig) GetBool(name string) bool {
    return viper.GetBool(name)
}

func (c *RuntimeConfig) Bool(name string) bool {
    return c.GetBool(name)
}

func (c *RuntimeConfig) GetFloat(name string, def float64) float64 {
    v := viper.GetFloat64(name)
    if v == 0 {
        return def
    }
    return v
}

func (c *RuntimeConfig) Float(name string, def float64) float64 {
    return c.GetFloat(name, def)
}

func (c *RuntimeConfig) GetBytes(name string, def []byte) []byte {
    return []byte(c.GetString(name, string(def)))
}

func (c *RuntimeConfig) Bytes(name string, def []byte) []byte {
    return c.GetBytes(name, def)
}
