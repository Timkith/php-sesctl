package parser

type SessionValue interface {
	isSessionValue()
}

type SessionEntry struct {
	Key   string
	Value SessionValue
}

type StringValue struct{ value string }

func (StringValue) isSessionValue() {}

type IntValue struct{ value int64 }

func (IntValue) isSessionValue() {}

type ArrayValue struct{ value []SessionEntry }

func (ArrayValue) isSessionValue() {}

type BoolValue struct{ value bool }

func (BoolValue) isSessionValue() {}

type DoubleValue struct{ value float64 }

func (DoubleValue) isSessionValue() {}

type NullValue struct{}

func (NullValue) isSessionValue() {}
