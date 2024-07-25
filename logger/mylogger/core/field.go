package core

import "strconv"

type FieldType uint8

const (
	UnknownType FieldType = iota
	BinaryType
	BoolType
	ByteStringType
	Complex128Type
	Complex64Type
	DurationType
	Float64Type
	Float32Type
	Int64Type
	Int32Type
	Int16Type
	Int8Type
	StringType
	TimeFullType
	Uint64Type
	Uint32Type
	Uint16Type
	Uint8Type
)

type Field struct {
	Key     string
	Type    FieldType
	Integer int64
	String  string
}

func Uint64(key string, val uint64) Field {
	return Field{
		Key:     key,
		Type:    Uint64Type,
		Integer: int64(val),
		String:  strconv.FormatUint(val, 10),
	}
}