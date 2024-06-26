// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CostosWarehouse.avsc
 */
package CostosWarehouseEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type CostosWarehouse struct {
	Identificacion Identificacion `json:"Identificacion"`

	Reporte Reporte `json:"Reporte"`
}

const CostosWarehouseAvroCRC64Fingerprint = "\xed\xd3w\xa8r\xa2>\x18"

func NewCostosWarehouse() CostosWarehouse {
	r := CostosWarehouse{}
	r.Identificacion = NewIdentificacion()

	r.Reporte = NewReporte()

	return r
}

func DeserializeCostosWarehouse(r io.Reader) (CostosWarehouse, error) {
	t := NewCostosWarehouse()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCostosWarehouseFromSchema(r io.Reader, schema string) (CostosWarehouse, error) {
	t := NewCostosWarehouse()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCostosWarehouse(r CostosWarehouse, w io.Writer) error {
	var err error
	err = writeIdentificacion(r.Identificacion, w)
	if err != nil {
		return err
	}
	err = writeReporte(r.Reporte, w)
	if err != nil {
		return err
	}
	return err
}

func (r CostosWarehouse) Serialize(w io.Writer) error {
	return writeCostosWarehouse(r, w)
}

func (r CostosWarehouse) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Evento\",\"type\":[\"string\",\"null\"]},{\"name\":\"SchemaAvro\",\"type\":[\"string\",\"null\"]},{\"name\":\"DescripcionProceso\",\"type\":[\"string\",\"null\"]},{\"name\":\"Almacen\",\"type\":[\"string\",\"null\"]},{\"name\":\"Propietario\",\"type\":[\"string\",\"null\"]},{\"name\":\"Instancia\",\"type\":[\"string\",\"null\"]},{\"name\":\"Key1\",\"type\":[\"string\",\"null\"]},{\"name\":\"Key2\",\"type\":[\"string\",\"null\"]},{\"name\":\"Key3\",\"type\":[\"string\",\"null\"]},{\"name\":\"Key4\",\"type\":[\"string\",\"null\"]},{\"name\":\"Key5\",\"type\":[\"string\",\"null\"]},{\"name\":\"Key6\",\"type\":[\"string\",\"null\"]},{\"name\":\"Key7\",\"type\":[{\"items\":\"string\",\"type\":\"array\"},\"null\"]},{\"name\":\"Key8\",\"type\":[{\"items\":\"string\",\"type\":\"array\"},\"null\"]},{\"name\":\"FechaGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.CostosWarehouse.Events.Common\",\"type\":\"record\"}},{\"name\":\"Reporte\",\"type\":{\"fields\":[{\"name\":\"Almacen\",\"type\":[\"null\",\"string\"]},{\"name\":\"Operacion\",\"type\":[\"null\",\"string\"]},{\"name\":\"Cliente\",\"type\":[\"null\",\"string\"]},{\"name\":\"Contrato\",\"type\":[\"null\",\"string\"]},{\"name\":\"UnidadDeVenta\",\"type\":[\"null\",\"string\"]},{\"name\":\"MacroProceso\",\"type\":[\"null\",\"string\"]},{\"name\":\"MicroProceso\",\"type\":[\"null\",\"string\"]},{\"name\":\"Zona\",\"type\":[\"null\",\"string\"]},{\"name\":\"Cantidad\",\"type\":[\"null\",\"int\"]},{\"name\":\"UnidadDeMedida\",\"type\":[\"null\",\"string\"]},{\"name\":\"Volumetria\",\"type\":[\"null\",\"float\"]},{\"name\":\"Peso\",\"type\":[\"null\",\"float\"]},{\"name\":\"UnidadFinales\",\"type\":[\"null\",\"int\"]},{\"name\":\"CostoUnitario\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"opcional\",\"type\":[\"null\",\"string\"]}],\"name\":\"Reporte\",\"namespace\":\"Andreani.CostosWarehouse.Events.Common\",\"type\":\"record\"}}],\"name\":\"Andreani.CostosWarehouse.Events.Record.CostosWarehouse\",\"type\":\"record\"}"
}

func (r CostosWarehouse) SchemaName() string {
	return "Andreani.CostosWarehouse.Events.Record.CostosWarehouse"
}

func (_ CostosWarehouse) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CostosWarehouse) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CostosWarehouse) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CostosWarehouse) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CostosWarehouse) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CostosWarehouse) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CostosWarehouse) SetString(v string)   { panic("Unsupported operation") }
func (_ CostosWarehouse) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CostosWarehouse) Get(i int) types.Field {
	switch i {
	case 0:
		r.Identificacion = NewIdentificacion()

		w := types.Record{Target: &r.Identificacion}

		return w

	case 1:
		r.Reporte = NewReporte()

		w := types.Record{Target: &r.Reporte}

		return w

	}
	panic("Unknown field index")
}

func (r *CostosWarehouse) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *CostosWarehouse) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ CostosWarehouse) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CostosWarehouse) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CostosWarehouse) HintSize(int)                     { panic("Unsupported operation") }
func (_ CostosWarehouse) Finalize()                        {}

func (_ CostosWarehouse) AvroCRC64Fingerprint() []byte {
	return []byte(CostosWarehouseAvroCRC64Fingerprint)
}

func (r CostosWarehouse) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Identificacion"], err = json.Marshal(r.Identificacion)
	if err != nil {
		return nil, err
	}
	output["Reporte"], err = json.Marshal(r.Reporte)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CostosWarehouse) UnmarshalJSON(data []byte) error {
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
	val = func() json.RawMessage {
		if v, ok := fields["Reporte"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Reporte); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Reporte")
	}
	return nil
}
