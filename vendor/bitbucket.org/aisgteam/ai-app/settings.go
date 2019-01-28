package aiapp

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "runtime"
    "strings"
)

var (
    hasVendorName = true
    systemDirectories []string
    globalDirectory string
    cacheDirectory string
)

func init() {
    switch goos := runtime.GOOS; goos {
    case "darwin":
        systemDirectories = []string{"/Library/Application Support"}
        globalDirectory = os.Getenv("HOME") + "/Library/Application Support"
        cacheDirectory = os.Getenv("HOME") + "/Library/Caches"
    case "windows":
        systemDirectories = []string{os.Getenv("PROGRAMDATA")}
        globalDirectory = os.Getenv("APPDATA")
        cacheDirectory = os.Getenv("LOCALAPPDATA")
    default:
        if os.Getenv("XDG_CONFIG_HOME") != "" {
            globalDirectory = os.Getenv("XDG_CONFIG_HOME")
        } else {
            globalDirectory = filepath.Join(os.Getenv("HOME"), ".config")
        }
        if os.Getenv("XDG_CONFIG_DIRS") != "" {
            systemDirectories = strings.Split(os.Getenv("XDG_CONFIG_DIRS"), ":")
        } else {
            systemDirectories = []string{"/etc/xdg"}
        }
        if os.Getenv("XDG_CACHE_HOME") != "" {
            cacheDirectory = os.Getenv("XDG_CACHE_HOME")
        } else {
            cacheDirectory = filepath.Join(os.Getenv("HOME"), ".cache")
        }
    }
}

type DirType int

const (
    System DirType = iota
    Global
    All
    Existing
    Local
    Cache
)

type Dir struct {
    Path string
    Type DirType
}

func (d Dir) Open(fileName string) (*os.File, error) {
    return os.Open(filepath.Join(d.Path, fileName))
}

func (d Dir) Create(fileName string) (*os.File, error) {
    err := d.CreateParentDir(fileName)
    if err != nil {
        return nil, err
    }
    return os.Create(filepath.Join(d.Path, fileName))
}

func (d Dir) ReadFile(fileName string) ([]byte, error) {
    return ioutil.ReadFile(filepath.Join(d.Path, fileName))
}

func (d Dir) CreateParentDir(fileName string) error {
    return os.MkdirAll(filepath.Dir(filepath.Join(d.Path, fileName)), 0755)
}

func (d Dir) WriteFile(fileName string, data []byte) error {
    err := d.CreateParentDir(fileName)
    if err != nil {
        return err
    }
    return ioutil.WriteFile(filepath.Join(d.Path, fileName), data, 0644)
}

func (d Dir) MkdirAll() error {
    return os.MkdirAll(d.Path, 0755)
}

func (d Dir) Exists(fileName string) bool {
    _, err := os.Stat(filepath.Join(d.Path, fileName))
    return !os.IsNotExist(err)
}

type ConfigDir struct {
    VendorName      string
    ApplicationName string
    LocalPath       string
}

func loadDir(vendorName, applicationName string) ConfigDir {
    cfg := ConfigDir{
        VendorName:         vendorName,
        ApplicationName:    applicationName,
    }
    cfg.LocalPath, _ = filepath.Abs(".")

    appData := cfg.QueryFolders(Global)
    if appData != nil {

    }

    return cfg
}

func (c ConfigDir) joinPath(root string) string {
    if c.VendorName != "" && hasVendorName {
        return filepath.Join(root, c.VendorName, c.ApplicationName)
    }
    return filepath.Join(root, c.ApplicationName)
}

func (c ConfigDir) QueryFolders(configType DirType) []*Dir {
    if configType == Cache {
        return []*Dir{c.QueryCacheFolder()}
    }
    var result []*Dir
    if c.LocalPath != "" &&configType != System && configType != Global {
        result = append(result, &Dir{
            Path: c.LocalPath,
            Type: Local,
        })
    }
    if configType != System && configType != Local {
        result = append(result, &Dir{
            Path: c.joinPath(globalDirectory),
            Type: Global,
        })
    }
    if configType != Global && configType != Local {
        for _, root := range systemDirectories {
            result = append(result, &Dir{
                Path: c.joinPath(root),
                Type: System,
            })
        }
    }
    if configType != Existing {
        return result
    }
    var existing []*Dir
    for _, entry:= range result {
        if _, err := os.Stat(entry.Path); !os.IsNotExist(err) {
            existing = append(existing, entry)
        }
    }
    return existing
}

func (c ConfigDir) QueryFolderContainsFile(fileName string) *Dir {
    configs := c.QueryFolders(Existing)
    for _, config := range configs {
        if _, err := os.Stat(filepath.Join(config.Path, fileName)); !os.IsNotExist(err) {
            return config
        }
    }
    return nil
}

func (c ConfigDir) QueryCacheFolder() *Dir {
    return &Dir{
        Path: c.joinPath(cacheDirectory),
        Type: Cache,
    }
}

type RuntimeSettings struct {

}