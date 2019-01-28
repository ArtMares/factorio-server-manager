package aiapp

import (
    "fmt"
    "github.com/satori/go.uuid"
    "github.com/sirupsen/logrus"
    "os"
    "time"
)

var (
    a           *instance
    Vendor      string
    Name        string
    ShortName   string
    Ver         Version
)

type AppType int

const (
    Desktop AppType = iota
    Console
    Service
)

type instance struct {
    config    *RuntimeConfig
    dir       ConfigDir
    log       *logrus.Logger
    uid       uuid.UUID
    startTime time.Time
}

func Config() *RuntimeConfig {
    return a.config
}

func Logger() *logrus.Logger {
    return a.log
}

func AddLoggerHook(hook logrus.Hook) {
    a.log.AddHook(hook)
}

func UpTime() time.Duration {
    return time.Since(a.startTime)
}

func GetUUID() string {
    return a.uid.String()
}

func FullName() string {
    if Ver.String() != "" {
        return fmt.Sprintf("%s %s", Name, Ver)
    }

    return Name
}

func New() *instance {
    a = &instance{}

    a.log = logrus.StandardLogger()
    a.dir = loadDir(Vendor, ShortName)
    cfg, err := loadRuntimeConfig()
    if err != nil {
        a.log.Error("Cannot load configuration: ", err)
        panic(err)
    }
    a.config = cfg
    a.uid = uuid.NewV4()
    a.startTime = time.Now()

    return a
}

func init() {
    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.SetOutput(os.Stdout)
}
