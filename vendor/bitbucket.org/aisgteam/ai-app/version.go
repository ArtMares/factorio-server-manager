package aiapp

import (
    "fmt"
    "regexp"
    "strconv"
)

type Version struct {
    Major   int64
    Minor   int64
    Patch   int64
}

func (v *Version) Parse(s string) {
    re := regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)`)
    match := re.FindStringSubmatch(s)
    if len(match)> 0 && len(match) > 1 {
        v.Major, _ = strconv.ParseInt(match[1], 10, 64)
        v.Minor, _ = strconv.ParseInt(match[2], 10, 64)
        v.Patch, _ = strconv.ParseInt(match[3], 10, 64)
    }
}

func (v Version) String() string {
    return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v Version) Bytes() []byte {
    return []byte(v.String())
}