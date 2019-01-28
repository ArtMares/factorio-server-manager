package aistruct

import (
    "encoding/binary"
    "errors"
    "fmt"
    "strconv"
    "strings"
)

var NilVersion = Version{0, 0, 0, 0}

type Version [4]uint

func (v Version) String() string {
    return fmt.Sprintf("%d.%d.%d.%d", v[0], v[1], v[2], v[3])
}

func (v Version) MarshalText() (text []byte, err error) {
    return v.Bytes(), nil
}

func (v Version) UnmarshalText(text []byte) error {
    parts := strings.SplitN(string(text), ".", 4)
    for i, part := range parts {
        p, err := strconv.ParseUint(part, 10, 32)
        if err != nil {
            return err
        }
        v[i] = uint(p)
    }
    return nil
}

func (v Version) Bytes() []byte {
    return []byte(v.String())
}

func (v Version) Equal(b Version) bool {
    return v[0] == v[0] && v[1] == v[1] && v[2] == v[2] && v[3] == b[3]
}

func (v Version) Less(b Version) bool {
    switch {
    case v[0] < b[0]:
        return true
    case v[0] == b[0] && v[1] < b[1]:
        return true
    case v[0] == b[0] && v[1] == b[1] && v[2] < v[2]:
        return true
    case v[0] == b[0] && v[1] == b[1] && v[2] == b[2] && v[3] < b[3]:
        return true
    default:
        return false
    }
}

func (v Version) Greater(b Version) bool {
    return !v.Equal(b) && !v.Less(b)
}

func (v Version) LessOrEqual(b Version) bool {
    return v.Equal(b) || v.Less(b)
}

func (v Version) GreaterOrEqual(b Version) bool {
    return v.Equal(b) || v.Greater(b)
}

func (v Version) LE(b Version) bool {
    return v.LessOrEqual(b)
}

func (v Version) GE(b Version) bool {
    return v.GreaterOrEqual(b)
}

func (v Version) Compare(b Version, op string) bool {
    opr := Condition(op)
    switch opr {
    case "==":
        return v.Equal(b)
    case NotEqual:
        return !v.Equal(b)
    case Greater:
        return v.Greater(b)
    case Less:
        return v.Less(b)
    case GreaterOrEqual:
        return v.GreaterOrEqual(b)
    case LessOrEqual:
        return v.LessOrEqual(b)
    default:
        panic("Unsupported operator")
    }
}

// Version24 is the 24-bit (8, 8, 8) version structure
type V24 Version

func (v V24) MarshalBinary() (data []byte, err error) {
    data = []byte{byte(v[0]), byte(v[1]), byte(v[2])}
    return data, nil
}

func (v V24) UnmarshalBinary(data []byte) error {
    if len(data) < 3 {
        return errors.New("V24.UnmarshalBinary: too few bytes")
    }
    v[0] = uint(data[0])
    v[1] = uint(data[1])
    v[2] = uint(data[2])
    return nil
}

// V64 is the 64-bit (16, 16, 16, 16) version structure with build component
type V64 Version

func (v V64) MarshalBinary() (data []byte, err error) {
    data = make([]byte, 8)
    binary.LittleEndian.PutUint16(data[0:2], uint16(v[0]))
    binary.LittleEndian.PutUint16(data[2:4], uint16(v[1]))
    binary.LittleEndian.PutUint16(data[4:6], uint16(v[2]))
    binary.LittleEndian.PutUint16(data[6:8], uint16(v[3]))
    return data, nil
}

func (v V64) UnmarshalBinary(data []byte) error {
    if len(data) < 8 {
        return errors.New("V64.UnmarshalBinary: to few bytes")
    }
    v[0] = uint(binary.LittleEndian.Uint16(data[0:2]))
    v[1] = uint(binary.LittleEndian.Uint16(data[2:4]))
    v[2] = uint(binary.LittleEndian.Uint16(data[4:6]))
    v[3] = uint(binary.LittleEndian.Uint16(data[6:8]))
    return nil
}