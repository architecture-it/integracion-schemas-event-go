// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     SeriesGeneradasRechazada.avsc
 */
package WosTrazabilidadEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type SeriesGeneradasRechazada struct {
	TareaSeriesGeneradas TareaConSeriesGeneradas `json:"TareaSeriesGeneradas"`

	RazonFalla *UnionNullString `json:"razonFalla"`
}

const SeriesGeneradasRechazadaAvroCRC64Fingerprint = "\x92X\xaa\x95\xcc\xf7\xa8\xe0"

func NewSeriesGeneradasRechazada() SeriesGeneradasRechazada {
	r := SeriesGeneradasRechazada{}
	r.TareaSeriesGeneradas = NewTareaConSeriesGeneradas()

	r.RazonFalla = nil
	return r
}

func DeserializeSeriesGeneradasRechazada(r io.Reader) (SeriesGeneradasRechazada, error) {
	t := NewSeriesGeneradasRechazada()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSeriesGeneradasRechazadaFromSchema(r io.Reader, schema string) (SeriesGeneradasRechazada, error) {
	t := NewSeriesGeneradasRechazada()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSeriesGeneradasRechazada(r SeriesGeneradasRechazada, w io.Writer) error {
	var err error
	err = writeTareaConSeriesGeneradas(r.TareaSeriesGeneradas, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.RazonFalla, w)
	if err != nil {
		return err
	}
	return err
}

func (r SeriesGeneradasRechazada) Serialize(w io.Writer) error {
	return writeSeriesGeneradasRechazada(r, w)
}

func (r SeriesGeneradasRechazada) Schema() string {
	return "{\"fields\":[{\"name\":\"TareaSeriesGeneradas\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"nroTarea\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"tipoDeTarea\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"almacen\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sku\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteInterno\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteCajita\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cantidad\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"gtin\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteFechaVencimiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"series\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"default\":null,\"name\":\"propietario\",\"type\":[\"null\",\"string\"]}],\"name\":\"TareaConSeriesGeneradas\",\"namespace\":\"Andreani.WosTrazabilidad.Events.WosTrazaTareaSeriesGeneradas\",\"type\":\"record\"}},{\"default\":null,\"name\":\"razonFalla\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.WosTrazabilidad.Events.Record.SeriesGeneradasRechazada\",\"type\":\"record\"}"
}

func (r SeriesGeneradasRechazada) SchemaName() string {
	return "Andreani.WosTrazabilidad.Events.Record.SeriesGeneradasRechazada"
}

func (_ SeriesGeneradasRechazada) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SeriesGeneradasRechazada) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SeriesGeneradasRechazada) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SeriesGeneradasRechazada) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SeriesGeneradasRechazada) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SeriesGeneradasRechazada) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SeriesGeneradasRechazada) SetString(v string)   { panic("Unsupported operation") }
func (_ SeriesGeneradasRechazada) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SeriesGeneradasRechazada) Get(i int) types.Field {
	switch i {
	case 0:
		r.TareaSeriesGeneradas = NewTareaConSeriesGeneradas()

		w := types.Record{Target: &r.TareaSeriesGeneradas}

		return w

	case 1:
		r.RazonFalla = NewUnionNullString()

		return r.RazonFalla
	}
	panic("Unknown field index")
}

func (r *SeriesGeneradasRechazada) SetDefault(i int) {
	switch i {
	case 1:
		r.RazonFalla = nil
		return
	}
	panic("Unknown field index")
}

func (r *SeriesGeneradasRechazada) NullField(i int) {
	switch i {
	case 1:
		r.RazonFalla = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ SeriesGeneradasRechazada) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SeriesGeneradasRechazada) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SeriesGeneradasRechazada) HintSize(int)                     { panic("Unsupported operation") }
func (_ SeriesGeneradasRechazada) Finalize()                        {}

func (_ SeriesGeneradasRechazada) AvroCRC64Fingerprint() []byte {
	return []byte(SeriesGeneradasRechazadaAvroCRC64Fingerprint)
}

func (r SeriesGeneradasRechazada) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["TareaSeriesGeneradas"], err = json.Marshal(r.TareaSeriesGeneradas)
	if err != nil {
		return nil, err
	}
	output["razonFalla"], err = json.Marshal(r.RazonFalla)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SeriesGeneradasRechazada) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["TareaSeriesGeneradas"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TareaSeriesGeneradas); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TareaSeriesGeneradas")
	}
	val = func() json.RawMessage {
		if v, ok := fields["razonFalla"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RazonFalla); err != nil {
			return err
		}
	} else {
		r.RazonFalla = NewUnionNullString()

		r.RazonFalla = nil
	}
	return nil
}