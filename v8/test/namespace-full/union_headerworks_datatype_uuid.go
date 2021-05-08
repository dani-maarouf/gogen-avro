// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v8/compiler"
	"github.com/actgardner/gogen-avro/v8/vm"
	"github.com/actgardner/gogen-avro/v8/vm/types"
)

type UnionHeaderworksDatatypeUUIDTypeEnum int

const (
	UnionHeaderworksDatatypeUUIDTypeEnumHeaderworksDatatypeUUID UnionHeaderworksDatatypeUUIDTypeEnum = 1
)

type UnionHeaderworksDatatypeUUID struct {
	Null                    *types.NullVal
	HeaderworksDatatypeUUID HeaderworksDatatypeUUID
	UnionType               UnionHeaderworksDatatypeUUIDTypeEnum
}

func writeUnionHeaderworksDatatypeUUID(r *UnionHeaderworksDatatypeUUID, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionHeaderworksDatatypeUUIDTypeEnumHeaderworksDatatypeUUID:
		return writeHeaderworksDatatypeUUID(r.HeaderworksDatatypeUUID, w)
	}
	return fmt.Errorf("invalid value for *UnionHeaderworksDatatypeUUID")
}

func NewUnionHeaderworksDatatypeUUID() *UnionHeaderworksDatatypeUUID {
	return &UnionHeaderworksDatatypeUUID{}
}

func (r *UnionHeaderworksDatatypeUUID) Serialize(w io.Writer) error {
	return writeUnionHeaderworksDatatypeUUID(r, w)
}

func DeserializeUnionHeaderworksDatatypeUUID(r io.Reader) (*UnionHeaderworksDatatypeUUID, error) {
	t := NewUnionHeaderworksDatatypeUUID()
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

func DeserializeUnionHeaderworksDatatypeUUIDFromSchema(r io.Reader, schema string) (*UnionHeaderworksDatatypeUUID, error) {
	t := NewUnionHeaderworksDatatypeUUID()
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

func (r *UnionHeaderworksDatatypeUUID) Schema() string {
	return "[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]"
}

func (_ *UnionHeaderworksDatatypeUUID) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionHeaderworksDatatypeUUID) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionHeaderworksDatatypeUUID) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionHeaderworksDatatypeUUID) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionHeaderworksDatatypeUUID) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionHeaderworksDatatypeUUID) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionHeaderworksDatatypeUUID) SetLong(v int64) {

	r.UnionType = (UnionHeaderworksDatatypeUUIDTypeEnum)(v)
}

func (r *UnionHeaderworksDatatypeUUID) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.HeaderworksDatatypeUUID = NewHeaderworksDatatypeUUID()
		return &types.Record{Target: (&r.HeaderworksDatatypeUUID)}
	}
	panic("Unknown field index")
}
func (_ *UnionHeaderworksDatatypeUUID) NullField(i int)  { panic("Unsupported operation") }
func (_ *UnionHeaderworksDatatypeUUID) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionHeaderworksDatatypeUUID) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ *UnionHeaderworksDatatypeUUID) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionHeaderworksDatatypeUUID) Finalize()                {}

func (r *UnionHeaderworksDatatypeUUID) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionHeaderworksDatatypeUUIDTypeEnumHeaderworksDatatypeUUID:
		return json.Marshal(map[string]interface{}{"headerworks.datatype.UUID": r.HeaderworksDatatypeUUID})
	}
	return nil, fmt.Errorf("invalid value for *UnionHeaderworksDatatypeUUID")
}

func (r *UnionHeaderworksDatatypeUUID) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["headerworks.datatype.UUID"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.HeaderworksDatatypeUUID)
	}
	return fmt.Errorf("invalid value for *UnionHeaderworksDatatypeUUID")
}
