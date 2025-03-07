// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NuevaFechaDeEntregaRepactada.avsc
 */
package TrazasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NuevaFechaDeEntregaRepactada struct {
	Traza Traza `json:"traza"`

	NuevaFechaPactadaDesde *UnionNullString `json:"nuevaFechaPactadaDesde"`

	NuevaFechaPactadaHasta *UnionNullString `json:"nuevaFechaPactadaHasta"`
}

const NuevaFechaDeEntregaRepactadaAvroCRC64Fingerprint = "\x9e\xabF)rN#\x01"

func NewNuevaFechaDeEntregaRepactada() NuevaFechaDeEntregaRepactada {
	r := NuevaFechaDeEntregaRepactada{}
	r.Traza = NewTraza()

	return r
}

func DeserializeNuevaFechaDeEntregaRepactada(r io.Reader) (NuevaFechaDeEntregaRepactada, error) {
	t := NewNuevaFechaDeEntregaRepactada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNuevaFechaDeEntregaRepactadaFromSchema(r io.Reader, schema string) (NuevaFechaDeEntregaRepactada, error) {
	t := NewNuevaFechaDeEntregaRepactada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNuevaFechaDeEntregaRepactada(r NuevaFechaDeEntregaRepactada, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NuevaFechaPactadaDesde, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NuevaFechaPactadaHasta, w)
	if err != nil {
		return err
	}
	return err
}

func (r NuevaFechaDeEntregaRepactada) Serialize(w io.Writer) error {
	return writeNuevaFechaDeEntregaRepactada(r, w)
}

func (r NuevaFechaDeEntregaRepactada) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"nuevaFechaPactadaDesde\",\"type\":[\"null\",\"string\"]},{\"name\":\"nuevaFechaPactadaHasta\",\"type\":[\"null\",\"string\"]}],\"name\":\"Integracion.Esquemas.Trazas.NuevaFechaDeEntregaRepactada\",\"type\":\"record\"}"
}

func (r NuevaFechaDeEntregaRepactada) SchemaName() string {
	return "Integracion.Esquemas.Trazas.NuevaFechaDeEntregaRepactada"
}

func (_ NuevaFechaDeEntregaRepactada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NuevaFechaDeEntregaRepactada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NuevaFechaDeEntregaRepactada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NuevaFechaDeEntregaRepactada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NuevaFechaDeEntregaRepactada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NuevaFechaDeEntregaRepactada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NuevaFechaDeEntregaRepactada) SetString(v string)   { panic("Unsupported operation") }
func (_ NuevaFechaDeEntregaRepactada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NuevaFechaDeEntregaRepactada) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		r.NuevaFechaPactadaDesde = NewUnionNullString()

		return r.NuevaFechaPactadaDesde
	case 2:
		r.NuevaFechaPactadaHasta = NewUnionNullString()

		return r.NuevaFechaPactadaHasta
	}
	panic("Unknown field index")
}

func (r *NuevaFechaDeEntregaRepactada) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NuevaFechaDeEntregaRepactada) NullField(i int) {
	switch i {
	case 1:
		r.NuevaFechaPactadaDesde = nil
		return
	case 2:
		r.NuevaFechaPactadaHasta = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ NuevaFechaDeEntregaRepactada) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ NuevaFechaDeEntregaRepactada) AppendArray() types.Field { panic("Unsupported operation") }
func (_ NuevaFechaDeEntregaRepactada) HintSize(int)             { panic("Unsupported operation") }
func (_ NuevaFechaDeEntregaRepactada) Finalize()                {}

func (_ NuevaFechaDeEntregaRepactada) AvroCRC64Fingerprint() []byte {
	return []byte(NuevaFechaDeEntregaRepactadaAvroCRC64Fingerprint)
}

func (r NuevaFechaDeEntregaRepactada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["nuevaFechaPactadaDesde"], err = json.Marshal(r.NuevaFechaPactadaDesde)
	if err != nil {
		return nil, err
	}
	output["nuevaFechaPactadaHasta"], err = json.Marshal(r.NuevaFechaPactadaHasta)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NuevaFechaDeEntregaRepactada) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["traza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Traza); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for traza")
	}
	val = func() json.RawMessage {
		if v, ok := fields["nuevaFechaPactadaDesde"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NuevaFechaPactadaDesde); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for nuevaFechaPactadaDesde")
	}
	val = func() json.RawMessage {
		if v, ok := fields["nuevaFechaPactadaHasta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NuevaFechaPactadaHasta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for nuevaFechaPactadaHasta")
	}
	return nil
}
