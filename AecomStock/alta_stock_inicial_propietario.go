// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AltaStockInicialPropietario.avsc
 */
package AecomStockEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type AltaStockInicialPropietario struct {
	Identificacion Identificacion `json:"Identificacion"`
}

const AltaStockInicialPropietarioAvroCRC64Fingerprint = "ݥée\xd1sm"

func NewAltaStockInicialPropietario() AltaStockInicialPropietario {
	r := AltaStockInicialPropietario{}
	r.Identificacion = NewIdentificacion()

	return r
}

func DeserializeAltaStockInicialPropietario(r io.Reader) (AltaStockInicialPropietario, error) {
	t := NewAltaStockInicialPropietario()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAltaStockInicialPropietarioFromSchema(r io.Reader, schema string) (AltaStockInicialPropietario, error) {
	t := NewAltaStockInicialPropietario()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAltaStockInicialPropietario(r AltaStockInicialPropietario, w io.Writer) error {
	var err error
	err = writeIdentificacion(r.Identificacion, w)
	if err != nil {
		return err
	}
	return err
}

func (r AltaStockInicialPropietario) Serialize(w io.Writer) error {
	return writeAltaStockInicialPropietario(r, w)
}

func (r AltaStockInicialPropietario) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.AecomStock.Events.StockCommon\",\"type\":\"record\"}}],\"name\":\"Andreani.AecomStock.Events.Record.AltaStockInicialPropietario\",\"type\":\"record\"}"
}

func (r AltaStockInicialPropietario) SchemaName() string {
	return "Andreani.AecomStock.Events.Record.AltaStockInicialPropietario"
}

func (_ AltaStockInicialPropietario) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AltaStockInicialPropietario) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AltaStockInicialPropietario) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AltaStockInicialPropietario) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AltaStockInicialPropietario) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AltaStockInicialPropietario) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AltaStockInicialPropietario) SetString(v string)   { panic("Unsupported operation") }
func (_ AltaStockInicialPropietario) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AltaStockInicialPropietario) Get(i int) types.Field {
	switch i {
	case 0:
		r.Identificacion = NewIdentificacion()

		w := types.Record{Target: &r.Identificacion}

		return w

	}
	panic("Unknown field index")
}

func (r *AltaStockInicialPropietario) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *AltaStockInicialPropietario) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ AltaStockInicialPropietario) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ AltaStockInicialPropietario) AppendArray() types.Field { panic("Unsupported operation") }
func (_ AltaStockInicialPropietario) HintSize(int)             { panic("Unsupported operation") }
func (_ AltaStockInicialPropietario) Finalize()                {}

func (_ AltaStockInicialPropietario) AvroCRC64Fingerprint() []byte {
	return []byte(AltaStockInicialPropietarioAvroCRC64Fingerprint)
}

func (r AltaStockInicialPropietario) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Identificacion"], err = json.Marshal(r.Identificacion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AltaStockInicialPropietario) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Identificacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Identificacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Identificacion")
	}
	return nil
}