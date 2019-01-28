package wapp

import (
    "fmt"
    "time"

    "bitbucket.org/aisgteam/ai-struct"
    "github.com/satori/go.uuid"
)

var (
    a       *instance
    Name    string
    Version aistruct.Version
    Vendor  string
)

type instance struct {
    uid       uuid.UUID
    startTime time.Time
}

func New() *instance {
    a = &instance{}
    a.uid = uuid.NewV4()
    a.startTime = time.Now()
}

func FullName() string {
    if Version.String() != "" {
        return fmt.Sprintf("%s (%s)", Name, Version)
    }
    return Name
}

func UpTime() time.Duration {
    return time.Since(a.startTime)
}

func UUID() string {
    return a.uid.String()
}