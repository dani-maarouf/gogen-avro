// Code generated by github.com/dani-maarouf/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
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

type UnionNullBodyworksDatatypeUUIDTypeEnum int

const (
	UnionNullBodyworksDatatypeUUIDTypeEnumBodyworksDatatypeUUID UnionNullBodyworksDatatypeUUIDTypeEnum = 1
)

type UnionNullBodyworksDatatypeUUID struct {
	Null                  *types.NullVal
	BodyworksDatatypeUUID BodyworksDatatypeUUID
	UnionType             UnionNullBodyworksDatatypeUUIDTypeEnum
}

func writeUnionNullBodyworksDatatypeUUID(r *UnionNullBodyworksDatatypeUUID, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullBodyworksDatatypeUUIDTypeEnumBodyworksDatatypeUUID:
		return writeBodyworksDatatypeUUID(r.BodyworksDatatypeUUID, w)
	}
	return fmt.Errorf("invalid value for *UnionNullBodyworksDatatypeUUID")
}

func NewUnionNullBodyworksDatatypeUUID() *UnionNullBodyworksDatatypeUUID {
	return &UnionNullBodyworksDatatypeUUID{}
}

func (r *UnionNullBodyworksDatatypeUUID) Serialize(w io.Writer) error {
	return writeUnionNullBodyworksDatatypeUUID(r, w)
}

func DeserializeUnionNullBodyworksDatatypeUUID(r io.Reader) (*UnionNullBodyworksDatatypeUUID, error) {
	t := NewUnionNullBodyworksDatatypeUUID()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNullBodyworksDatatypeUUIDFromSchema(r io.Reader, schema string) (*UnionNullBodyworksDatatypeUUID, error) {
	t := NewUnionNullBodyworksDatatypeUUID()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r *UnionNullBodyworksDatatypeUUID) Schema() string {
	return "[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"bodyworks.datatype\",\"type\":\"record\"}]"
}

func (_ *UnionNullBodyworksDatatypeUUID) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullBodyworksDatatypeUUID) SetLong(v int64) {

	r.UnionType = (UnionNullBodyworksDatatypeUUIDTypeEnum)(v)
}

func (r *UnionNullBodyworksDatatypeUUID) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.BodyworksDatatypeUUID = NewBodyworksDatatypeUUID()
		return &types.Record{Target: (&r.BodyworksDatatypeUUID)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullBodyworksDatatypeUUID) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) HintSize(i int)   { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionNullBodyworksDatatypeUUID) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullBodyworksDatatypeUUID) Finalize()                {}

func (r *UnionNullBodyworksDatatypeUUID) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullBodyworksDatatypeUUIDTypeEnumBodyworksDatatypeUUID:
		return json.Marshal(map[string]interface{}{"bodyworks.datatype.UUID": r.BodyworksDatatypeUUID})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullBodyworksDatatypeUUID")
}

func (r *UnionNullBodyworksDatatypeUUID) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["bodyworks.datatype.UUID"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.BodyworksDatatypeUUID)
	}
	return fmt.Errorf("invalid value for *UnionNullBodyworksDatatypeUUID")
}
