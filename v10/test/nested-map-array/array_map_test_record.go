// Code generated by github.com/dani-maarouf/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/dani-maarouf/gogen-avro/v10/vm"
	"github.com/dani-maarouf/gogen-avro/v10/vm/types"
)

func writeArrayMapTestRecord(r []MapTestRecord, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeMapTestRecord(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayMapTestRecordWrapper struct {
	Target *[]MapTestRecord
}

func (_ ArrayMapTestRecordWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayMapTestRecordWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayMapTestRecordWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayMapTestRecordWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayMapTestRecordWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayMapTestRecordWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayMapTestRecordWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayMapTestRecordWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayMapTestRecordWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayMapTestRecordWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayMapTestRecordWrapper) Finalize()                        {}
func (_ ArrayMapTestRecordWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayMapTestRecordWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]MapTestRecord, 0, s)
	}
}
func (r ArrayMapTestRecordWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayMapTestRecordWrapper) AppendArray() types.Field {
	var v MapTestRecord
	v = NewMapTestRecord()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
