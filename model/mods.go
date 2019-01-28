package model

import (
    "fmt"
    "net/url"
    "strings"
    "time"
)

const modTimeFormat = "2006-01-02T15:04:05.999999999Z"

type Mods []Mod

type Mod struct {
    Name            string
    Title           string
    Owner           string
    Summary         string
    Description     string
    Download        ModDownload
    FAQ             string
    GitHub          string
    HomePage        string
    Releases        []ModRelease
    Tag             Tag
    Created         time.Time
    Updated         time.Time
}

type ModDownload struct {
    Count           int64
    Url             *url.URL
}

type ModRelease struct {

}

type ModTime struct {
    time.Time
}

func (t *ModTime) UnmarshalJSON(b []byte) (err error) {
    s := strings.Trim(string(b), "\"")
    if s == "null" {
        t.Time = time.Time{}
        return
    }
    t.Time, err = time.Parse(modTimeFormat, s)
    return
}

func (t *ModTime) MarshalJSON() ([]byte, error) {
    if t.Time.UnixNano() == (time.Time{}).UnixNano() {
        return []byte("null"), nil
    }
    return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(modTimeFormat))), nil
}