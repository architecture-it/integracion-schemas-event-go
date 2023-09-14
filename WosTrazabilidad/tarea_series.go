// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TareaSeries.avsc
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

type TareaSeries struct {
	TareaSeriesIngresadas TareaSeries `json:"TareaSeriesIngresadas"`
}

const TareaSeriesAvroCRC64Fingerprint = "\xecA\xd9\xceXO\xa8w"

func NewTareaSeries() TareaSeries {
	r := TareaSeries{}
	r.TareaSeriesIngresadas = NewTareaSeries()

	return r
}

func DeserializeTareaSeries(r io.Reader) (TareaSeries, error) {
	t := NewTareaSeries()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTareaSeriesFromSchema(r io.Reader, schema string) (TareaSeries, error) {
	t := NewTareaSeries()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTareaSeries(r TareaSeries, w io.Writer) error {
	var err error
	err = writeTareaSeries(r.TareaSeriesIngresadas, w)
	if err != nil {
		return err
	}
	return err
}

func (r TareaSeries) Serialize(w io.Writer) error {
	return writeTareaSeries(r, w)
}

func (r TareaSeries) Schema() string {
	return "{\"fields\":[{\"name\":\"TareaSeriesIngresadas\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"serialkey\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"nroTarea\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"tipoDeTarea\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"fecha\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"almacen\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"propietario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nroDocumentoInterno\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nroDocumentoExterno\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nroLPN\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nroLinea\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sku\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteCajita\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cantidad\",\"type\":[\"null\",\"int\"]},{\"name\":\"Series\",\"type\":{\"items\":{\"fields\":[{\"name\":\"serie\",\"type\":\"string\"},{\"default\":null,\"name\":\"estado\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"descripcionEstado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoTransaccion\",\"type\":[\"null\",\"string\"]}],\"name\":\"Serie\",\"namespace\":\"Andreani.WosTrazabilidad.Events.AnmatCommon\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"TareaSeries\",\"namespace\":\"Andreani.WosTrazabilidad.Events.WosTrazaTareaSeriesIngresadas\",\"type\":\"record\"}}],\"name\":\"Andreani.WosTrazabilidad.Events.Record.TareaSeries\",\"type\":\"record\"}"
}

func (r TareaSeries) SchemaName() string {
	return "Andreani.WosTrazabilidad.Events.Record.TareaSeries"
}

func (_ TareaSeries) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ TareaSeries) SetInt(v int32)       { panic("Unsupported operation") }
func (_ TareaSeries) SetLong(v int64)      { panic("Unsupported operation") }
func (_ TareaSeries) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ TareaSeries) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ TareaSeries) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ TareaSeries) SetString(v string)   { panic("Unsupported operation") }
func (_ TareaSeries) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *TareaSeries) Get(i int) types.Field {
	switch i {
	case 0:
		r.TareaSeriesIngresadas = NewTareaSeries()

		w := types.Record{Target: &r.TareaSeriesIngresadas}

		return w

	}
	panic("Unknown field index")
}

func (r *TareaSeries) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *TareaSeries) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ TareaSeries) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ TareaSeries) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ TareaSeries) HintSize(int)                     { panic("Unsupported operation") }
func (_ TareaSeries) Finalize()                        {}

func (_ TareaSeries) AvroCRC64Fingerprint() []byte {
	return []byte(TareaSeriesAvroCRC64Fingerprint)
}

func (r TareaSeries) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["TareaSeriesIngresadas"], err = json.Marshal(r.TareaSeriesIngresadas)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *TareaSeries) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["TareaSeriesIngresadas"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TareaSeriesIngresadas); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TareaSeriesIngresadas")
	}
	return nil
}
