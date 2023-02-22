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

var _ = fmt.Printf

type AliasRecord struct {
	A string `json:"a"`

	C string `json:"c"`
}

const AliasRecordAvroCRC64Fingerprint = "B~_@rz\xdb\xc7"

func NewAliasRecord() AliasRecord {
	r := AliasRecord{}
	return r
}

func DeserializeAliasRecord(r io.Reader) (AliasRecord, error) {
	t := NewAliasRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAliasRecordFromSchema(r io.Reader, schema string) (AliasRecord, error) {
	t := NewAliasRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAliasRecord(r AliasRecord, w io.Writer) error {
	var err error
	err = vm.WriteString(r.A, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.C, w)
	if err != nil {
		return err
	}
	return err
}

func (r AliasRecord) Serialize(w io.Writer) error {
	return writeAliasRecord(r, w)
}

func (r AliasRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"a\",\"type\":\"string\"},{\"name\":\"c\",\"type\":\"string\"}],\"name\":\"AliasRecord\",\"type\":\"record\"}"
}

func (r AliasRecord) SchemaName() string {
	return "AliasRecord"
}

func (_ AliasRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AliasRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AliasRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AliasRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AliasRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AliasRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AliasRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ AliasRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AliasRecord) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.A}

		return w

	case 1:
		w := types.String{Target: &r.C}

		return w

	}
	panic("Unknown field index")
}

func (r *AliasRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *AliasRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ AliasRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ AliasRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ AliasRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ AliasRecord) Finalize()                        {}

func (_ AliasRecord) AvroCRC64Fingerprint() []byte {
	return []byte(AliasRecordAvroCRC64Fingerprint)
}

func (r AliasRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["a"], err = json.Marshal(r.A)
	if err != nil {
		return nil, err
	}
	output["c"], err = json.Marshal(r.C)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AliasRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["a"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.A); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for a")
	}
	val = func() json.RawMessage {
		if v, ok := fields["c"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.C); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for c")
	}
	return nil
}
