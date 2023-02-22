// Code generated by github.com/dani-maarouf/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/dani-maarouf/gogen-avro/v10/compiler"
	"github.com/dani-maarouf/gogen-avro/v10/vm"
	"github.com/dani-maarouf/gogen-avro/v10/vm/types"
)

type UnionAliasedRecordNestedTestRecordTypeEnum int

const (
	UnionAliasedRecordNestedTestRecordTypeEnumAliasedRecord UnionAliasedRecordNestedTestRecordTypeEnum = 0

	UnionAliasedRecordNestedTestRecordTypeEnumNestedTestRecord UnionAliasedRecordNestedTestRecordTypeEnum = 1
)

type UnionAliasedRecordNestedTestRecord struct {
	AliasedRecord    AliasedRecord
	NestedTestRecord NestedTestRecord
	UnionType        UnionAliasedRecordNestedTestRecordTypeEnum
}

func writeUnionAliasedRecordNestedTestRecord(r UnionAliasedRecordNestedTestRecord, w io.Writer) error {

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionAliasedRecordNestedTestRecordTypeEnumAliasedRecord:
		return writeAliasedRecord(r.AliasedRecord, w)
	case UnionAliasedRecordNestedTestRecordTypeEnumNestedTestRecord:
		return writeNestedTestRecord(r.NestedTestRecord, w)
	}
	return fmt.Errorf("invalid value for UnionAliasedRecordNestedTestRecord")
}

func NewUnionAliasedRecordNestedTestRecord() UnionAliasedRecordNestedTestRecord {
	return UnionAliasedRecordNestedTestRecord{}
}

func (r UnionAliasedRecordNestedTestRecord) Serialize(w io.Writer) error {
	return writeUnionAliasedRecordNestedTestRecord(r, w)
}

func DeserializeUnionAliasedRecordNestedTestRecord(r io.Reader) (UnionAliasedRecordNestedTestRecord, error) {
	t := NewUnionAliasedRecordNestedTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionAliasedRecordNestedTestRecordFromSchema(r io.Reader, schema string) (UnionAliasedRecordNestedTestRecord, error) {
	t := NewUnionAliasedRecordNestedTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r UnionAliasedRecordNestedTestRecord) Schema() string {
	return "[{\"aliases\":[\"NestedRecord\"],\"fields\":[{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"name\":\"BytesField\",\"type\":\"bytes\"}],\"name\":\"aliasedRecord\",\"type\":\"record\"},{\"fields\":[{\"name\":\"OtherField\",\"type\":\"aliasedRecord\"}],\"name\":\"NestedTestRecord\",\"type\":\"record\"}]"
}

func (_ UnionAliasedRecordNestedTestRecord) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ UnionAliasedRecordNestedTestRecord) SetInt(v int32)      { panic("Unsupported operation") }
func (_ UnionAliasedRecordNestedTestRecord) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ UnionAliasedRecordNestedTestRecord) SetDouble(v float64) { panic("Unsupported operation") }
func (_ UnionAliasedRecordNestedTestRecord) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ UnionAliasedRecordNestedTestRecord) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionAliasedRecordNestedTestRecord) SetLong(v int64) {

	r.UnionType = (UnionAliasedRecordNestedTestRecordTypeEnum)(v)
}

func (r *UnionAliasedRecordNestedTestRecord) Get(i int) types.Field {

	switch i {
	case 0:
		r.AliasedRecord = NewAliasedRecord()
		return &types.Record{Target: (&r.AliasedRecord)}
	case 1:
		r.NestedTestRecord = NewNestedTestRecord()
		return &types.Record{Target: (&r.NestedTestRecord)}
	}
	panic("Unknown field index")
}
func (_ UnionAliasedRecordNestedTestRecord) NullField(i int)  { panic("Unsupported operation") }
func (_ UnionAliasedRecordNestedTestRecord) HintSize(i int)   { panic("Unsupported operation") }
func (_ UnionAliasedRecordNestedTestRecord) SetDefault(i int) { panic("Unsupported operation") }
func (_ UnionAliasedRecordNestedTestRecord) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ UnionAliasedRecordNestedTestRecord) AppendArray() types.Field { panic("Unsupported operation") }
func (_ UnionAliasedRecordNestedTestRecord) Finalize()                {}

func (r UnionAliasedRecordNestedTestRecord) MarshalJSON() ([]byte, error) {

	switch r.UnionType {
	case UnionAliasedRecordNestedTestRecordTypeEnumAliasedRecord:
		return json.Marshal(map[string]interface{}{"aliasedRecord": r.AliasedRecord})
	case UnionAliasedRecordNestedTestRecordTypeEnumNestedTestRecord:
		return json.Marshal(map[string]interface{}{"NestedTestRecord": r.NestedTestRecord})
	}
	return nil, fmt.Errorf("invalid value for UnionAliasedRecordNestedTestRecord")
}

func (r *UnionAliasedRecordNestedTestRecord) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["aliasedRecord"]; ok {
		r.UnionType = 0
		return json.Unmarshal([]byte(value), &r.AliasedRecord)
	}
	if value, ok := fields["NestedTestRecord"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.NestedTestRecord)
	}
	return fmt.Errorf("invalid value for UnionAliasedRecordNestedTestRecord")
}
