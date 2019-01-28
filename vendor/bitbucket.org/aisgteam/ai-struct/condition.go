package aistruct


type Condition string

func (c Condition) String() string {
    return string(c)
}

const (
    Equal           Condition = "="
    Greater         Condition = ">"
    Less            Condition = "<"
    NotEqual        Condition = "!" + Equal
    GreaterOrEqual  Condition = Greater + Equal
    LessOrEqual     Condition = Less + Equal

)