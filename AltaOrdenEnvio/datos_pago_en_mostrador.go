// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OrdenDeEnvioSolicitada.avsc
 */
package AltaOrdenEnvioEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DatosPagoEnMostrador struct {
	MontoACobrar *UnionNullInt `json:"montoACobrar"`

	DocumentoTipo *UnionNullString `json:"documentoTipo"`

	DocumentoNumero *UnionNullString `json:"documentoNumero"`
}

const DatosPagoEnMostradorAvroCRC64Fingerprint = "\xa8\x11\xdb\xcc$\xf9\xf5\x1d"

func NewDatosPagoEnMostrador() DatosPagoEnMostrador {
	r := DatosPagoEnMostrador{}
	return r
}

func DeserializeDatosPagoEnMostrador(r io.Reader) (DatosPagoEnMostrador, error) {
	t := NewDatosPagoEnMostrador()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDatosPagoEnMostradorFromSchema(r io.Reader, schema string) (DatosPagoEnMostrador, error) {
	t := NewDatosPagoEnMostrador()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDatosPagoEnMostrador(r DatosPagoEnMostrador, w io.Writer) error {
	var err error
	err = writeUnionNullInt(r.MontoACobrar, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DocumentoTipo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DocumentoNumero, w)
	if err != nil {
		return err
	}
	return err
}

func (r DatosPagoEnMostrador) Serialize(w io.Writer) error {
	return writeDatosPagoEnMostrador(r, w)
}

func (r DatosPagoEnMostrador) Schema() string {
	return "{\"fields\":[{\"name\":\"montoACobrar\",\"type\":[\"null\",\"int\"]},{\"name\":\"documentoTipo\",\"type\":[\"null\",\"string\"]},{\"name\":\"documentoNumero\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.AltaOrdenEnvio.Events.Common.DatosPagoEnMostrador\",\"type\":\"record\"}"
}

func (r DatosPagoEnMostrador) SchemaName() string {
	return "Andreani.AltaOrdenEnvio.Events.Common.DatosPagoEnMostrador"
}

func (_ DatosPagoEnMostrador) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DatosPagoEnMostrador) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DatosPagoEnMostrador) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DatosPagoEnMostrador) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DatosPagoEnMostrador) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DatosPagoEnMostrador) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DatosPagoEnMostrador) SetString(v string)   { panic("Unsupported operation") }
func (_ DatosPagoEnMostrador) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DatosPagoEnMostrador) Get(i int) types.Field {
	switch i {
	case 0:
		r.MontoACobrar = NewUnionNullInt()

		return r.MontoACobrar
	case 1:
		r.DocumentoTipo = NewUnionNullString()

		return r.DocumentoTipo
	case 2:
		r.DocumentoNumero = NewUnionNullString()

		return r.DocumentoNumero
	}
	panic("Unknown field index")
}

func (r *DatosPagoEnMostrador) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *DatosPagoEnMostrador) NullField(i int) {
	switch i {
	case 0:
		r.MontoACobrar = nil
		return
	case 1:
		r.DocumentoTipo = nil
		return
	case 2:
		r.DocumentoNumero = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DatosPagoEnMostrador) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DatosPagoEnMostrador) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DatosPagoEnMostrador) HintSize(int)                     { panic("Unsupported operation") }
func (_ DatosPagoEnMostrador) Finalize()                        {}

func (_ DatosPagoEnMostrador) AvroCRC64Fingerprint() []byte {
	return []byte(DatosPagoEnMostradorAvroCRC64Fingerprint)
}

func (r DatosPagoEnMostrador) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["montoACobrar"], err = json.Marshal(r.MontoACobrar)
	if err != nil {
		return nil, err
	}
	output["documentoTipo"], err = json.Marshal(r.DocumentoTipo)
	if err != nil {
		return nil, err
	}
	output["documentoNumero"], err = json.Marshal(r.DocumentoNumero)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DatosPagoEnMostrador) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["montoACobrar"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MontoACobrar); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for montoACobrar")
	}
	val = func() json.RawMessage {
		if v, ok := fields["documentoTipo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DocumentoTipo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for documentoTipo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["documentoNumero"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DocumentoNumero); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for documentoNumero")
	}
	return nil
}
