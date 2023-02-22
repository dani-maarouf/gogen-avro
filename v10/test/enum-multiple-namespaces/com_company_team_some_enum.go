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

	"github.com/dani-maarouf/gogen-avro/v10/vm"
	"github.com/dani-maarouf/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ComCompanyTeamSomeEnum int32

const (
	ComCompanyTeamSomeEnumA ComCompanyTeamSomeEnum = 0
	ComCompanyTeamSomeEnumB ComCompanyTeamSomeEnum = 1
	ComCompanyTeamSomeEnumC ComCompanyTeamSomeEnum = 2
)

func (e ComCompanyTeamSomeEnum) String() string {
	switch e {
	case ComCompanyTeamSomeEnumA:
		return "A"
	case ComCompanyTeamSomeEnumB:
		return "B"
	case ComCompanyTeamSomeEnumC:
		return "C"
	}
	return "unknown"
}

func writeComCompanyTeamSomeEnum(r ComCompanyTeamSomeEnum, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewComCompanyTeamSomeEnumValue(raw string) (r ComCompanyTeamSomeEnum, err error) {
	switch raw {
	case "A":
		return ComCompanyTeamSomeEnumA, nil
	case "B":
		return ComCompanyTeamSomeEnumB, nil
	case "C":
		return ComCompanyTeamSomeEnumC, nil
	}

	return -1, fmt.Errorf("invalid value for ComCompanyTeamSomeEnum: '%s'", raw)

}

func (b ComCompanyTeamSomeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *ComCompanyTeamSomeEnum) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewComCompanyTeamSomeEnumValue(stringVal)
	*b = val
	return err
}

type ComCompanyTeamSomeEnumWrapper struct {
	Target *ComCompanyTeamSomeEnum
}

func (b ComCompanyTeamSomeEnumWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b ComCompanyTeamSomeEnumWrapper) SetInt(v int32) {
	*(b.Target) = ComCompanyTeamSomeEnum(v)
}

func (b ComCompanyTeamSomeEnumWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b ComCompanyTeamSomeEnumWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b ComCompanyTeamSomeEnumWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b ComCompanyTeamSomeEnumWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b ComCompanyTeamSomeEnumWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b ComCompanyTeamSomeEnumWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b ComCompanyTeamSomeEnumWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b ComCompanyTeamSomeEnumWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b ComCompanyTeamSomeEnumWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b ComCompanyTeamSomeEnumWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b ComCompanyTeamSomeEnumWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b ComCompanyTeamSomeEnumWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b ComCompanyTeamSomeEnumWrapper) Finalize() {}
