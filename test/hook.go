package test

import "context"

type BeforeExecuteEvent interface {
	Context() context.Context
	SetContext(ctx context.Context)
	Table() string
	SetTable(table string)
	Op() Op
	// SQL return sql string
	SQL() string
	// Fields return select/update/insert fields
	Fields() []string
	// Args return sql args
	Args() []interface{}
}

type AfterExecuteEvent interface {
	Context() context.Context
	Table() string
	Op() Op
	// SQL return sql string
	SQL() string
	// Fields return select/update/insert fields
	Fields() []string
	// Args return sql args
	Args() []interface{}
	// Value return sql execute result
	Value() interface{}
}

type event interface {
	BeforeExecuteEvent
	AfterExecuteEvent
	SetValue(value interface{})
}

type Hook interface {
	Before(event BeforeExecuteEvent) error
	After(event AfterExecuteEvent) error
}

type HookUnimplemented struct{}

func (HookUnimplemented) Before(_ BeforeExecuteEvent) error {
	return nil
}

func (HookUnimplemented) After(_ AfterExecuteEvent) error {
	return nil
}
