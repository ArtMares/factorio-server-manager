package model

type AuthToken string

func (t AuthToken) IsEmpty() bool {
    return t == ""
}

func (t AuthToken) String() string {
    return string(t)
}