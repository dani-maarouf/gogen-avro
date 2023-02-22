// Code generated by github.com/dani-maarouf/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/dani-maarouf/gogen-avro/v10/compiler"
	"github.com/dani-maarouf/gogen-avro/v10/container"
	"github.com/dani-maarouf/gogen-avro/v10/vm"
)

func NewNestedTestRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewNestedTestRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type NestedTestRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewNestedTestRecordReader(r io.Reader) (*NestedTestRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewNestedTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &NestedTestRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r NestedTestRecordReader) Read() (NestedTestRecord, error) {
	t := NewNestedTestRecord()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
