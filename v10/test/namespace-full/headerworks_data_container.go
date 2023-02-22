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

func NewHeaderworksDataWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewHeaderworksData()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type HeaderworksDataReader struct {
	r io.Reader
	p *vm.Program
}

func NewHeaderworksDataReader(r io.Reader) (*HeaderworksDataReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewHeaderworksData()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &HeaderworksDataReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r HeaderworksDataReader) Read() (HeaderworksData, error) {
	t := NewHeaderworksData()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
