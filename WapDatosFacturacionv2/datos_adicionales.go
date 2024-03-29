// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Pedido.avsc
 */
package WapDatosFacturacionv2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DatosAdicionales struct {
	Metadatos []Metadato `json:"metadatos"`
}

const DatosAdicionalesAvroCRC64Fingerprint = "H$\xbfI䣽h"

func NewDatosAdicionales() DatosAdicionales {
	r := DatosAdicionales{}
	r.Metadatos = make([]Metadato, 0)

	return r
}

func DeserializeDatosAdicionales(r io.Reader) (DatosAdicionales, error) {
	t := NewDatosAdicionales()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDatosAdicionalesFromSchema(r io.Reader, schema string) (DatosAdicionales, error) {
	t := NewDatosAdicionales()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDatosAdicionales(r DatosAdicionales, w io.Writer) error {
	var err error
	err = writeArrayMetadato(r.Metadatos, w)
	if err != nil {
		return err
	}
	return err
}

func (r DatosAdicionales) Serialize(w io.Writer) error {
	return writeDatosAdicionales(r, w)
}

func (r DatosAdicionales) Schema() string {
	return "{\"fields\":[{\"name\":\"metadatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.WapDatosFacturacionv2.Events.Record.DatosAdicionales\",\"type\":\"record\"}"
}

func (r DatosAdicionales) SchemaName() string {
	return "Andreani.WapDatosFacturacionv2.Events.Record.DatosAdicionales"
}

func (_ DatosAdicionales) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DatosAdicionales) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DatosAdicionales) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DatosAdicionales) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DatosAdicionales) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DatosAdicionales) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DatosAdicionales) SetString(v string)   { panic("Unsupported operation") }
func (_ DatosAdicionales) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DatosAdicionales) Get(i int) types.Field {
	switch i {
	case 0:
		r.Metadatos = make([]Metadato, 0)

		w := ArrayMetadatoWrapper{Target: &r.Metadatos}

		return w

	}
	panic("Unknown field index")
}

func (r *DatosAdicionales) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *DatosAdicionales) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ DatosAdicionales) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DatosAdicionales) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DatosAdicionales) HintSize(int)                     { panic("Unsupported operation") }
func (_ DatosAdicionales) Finalize()                        {}

func (_ DatosAdicionales) AvroCRC64Fingerprint() []byte {
	return []byte(DatosAdicionalesAvroCRC64Fingerprint)
}

func (r DatosAdicionales) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["metadatos"], err = json.Marshal(r.Metadatos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DatosAdicionales) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["metadatos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Metadatos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for metadatos")
	}
	return nil
}
