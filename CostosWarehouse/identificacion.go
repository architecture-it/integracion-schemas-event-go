// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Identificacion.avsc
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

type Identificacion struct {
	Evento *UnionStringNull `json:"Evento"`

	SchemaAvro *UnionStringNull `json:"SchemaAvro"`

	DescripcionProceso *UnionStringNull `json:"DescripcionProceso"`

	Almacen *UnionStringNull `json:"Almacen"`

	Propietario *UnionStringNull `json:"Propietario"`

	Instancia *UnionStringNull `json:"Instancia"`

	Key1 *UnionStringNull `json:"Key1"`

	Key2 *UnionStringNull `json:"Key2"`

	Key3 *UnionStringNull `json:"Key3"`

	Key4 *UnionStringNull `json:"Key4"`

	Key5 *UnionStringNull `json:"Key5"`

	Key6 *UnionStringNull `json:"Key6"`

	FechaGeneracion int64 `json:"FechaGeneracion"`
}

const IdentificacionAvroCRC64Fingerprint = "\xfb\xe0d]\xbc<\v\xbc"

func NewIdentificacion() Identificacion {
	r := Identificacion{}
	r.Evento = NewUnionStringNull()

	r.SchemaAvro = NewUnionStringNull()

	r.DescripcionProceso = NewUnionStringNull()

	r.Almacen = NewUnionStringNull()

	r.Propietario = NewUnionStringNull()

	r.Instancia = NewUnionStringNull()

	r.Key1 = NewUnionStringNull()

	r.Key2 = NewUnionStringNull()

	r.Key3 = NewUnionStringNull()

	r.Key4 = NewUnionStringNull()

	r.Key5 = NewUnionStringNull()

	r.Key6 = NewUnionStringNull()

	return r
}

func DeserializeIdentificacion(r io.Reader) (Identificacion, error) {
	t := NewIdentificacion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeIdentificacionFromSchema(r io.Reader, schema string) (Identificacion, error) {
	t := NewIdentificacion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeIdentificacion(r Identificacion, w io.Writer) error {
	var err error
	err = writeUnionStringNull(r.Evento, w)
	if err != nil {
		return err
	}
	err = writeUnionStringNull(r.SchemaAvro, w)
	if err != nil {
		return err
	}
	err = writeUnionStringNull(r.DescripcionProceso, w)
	if err != nil {
		return err
	}
	err = writeUnionStringNull(r.Almacen, w)
	if err != nil {
		return err
	}
	err = writeUnionStringNull(r.Propietario, w)
	if err != nil {
		return err
	}
	err = writeUnionStringNull(r.Instancia, w)
	if err != nil {
		return err
	}
	err = writeUnionStringNull(r.Key1, w)
	if err != nil {
		return err
	}
	err = writeUnionStringNull(r.Key2, w)
	if err != nil {
		return err
	}
	err = writeUnionStringNull(r.Key3, w)
	if err != nil {
		return err
	}
	err = writeUnionStringNull(r.Key4, w)
	if err != nil {
		return err
	}
	err = writeUnionStringNull(r.Key5, w)
	if err != nil {
		return err
	}
	err = writeUnionStringNull(r.Key6, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.FechaGeneracion, w)
	if err != nil {
		return err
	}
	return err
}

func (r Identificacion) Serialize(w io.Writer) error {
	return writeIdentificacion(r, w)
}

func (r Identificacion) Schema() string {
	return "{\"fields\":[{\"name\":\"Evento\",\"type\":[\"string\",\"null\"]},{\"name\":\"SchemaAvro\",\"type\":[\"string\",\"null\"]},{\"name\":\"DescripcionProceso\",\"type\":[\"string\",\"null\"]},{\"name\":\"Almacen\",\"type\":[\"string\",\"null\"]},{\"name\":\"Propietario\",\"type\":[\"string\",\"null\"]},{\"name\":\"Instancia\",\"type\":[\"string\",\"null\"]},{\"name\":\"Key1\",\"type\":[\"string\",\"null\"]},{\"name\":\"Key2\",\"type\":[\"string\",\"null\"]},{\"name\":\"Key3\",\"type\":[\"string\",\"null\"]},{\"name\":\"Key4\",\"type\":[\"string\",\"null\"]},{\"name\":\"Key5\",\"type\":[\"string\",\"null\"]},{\"name\":\"Key6\",\"type\":[\"string\",\"null\"]},{\"name\":\"FechaGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"Andreani.CostosWarehouseCommon.Events.Common.Identificacion\",\"type\":\"record\"}"
}

func (r Identificacion) SchemaName() string {
	return "Andreani.CostosWarehouseCommon.Events.Common.Identificacion"
}

func (_ Identificacion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Identificacion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Identificacion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Identificacion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Identificacion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Identificacion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Identificacion) SetString(v string)   { panic("Unsupported operation") }
func (_ Identificacion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Identificacion) Get(i int) types.Field {
	switch i {
	case 0:
		r.Evento = NewUnionStringNull()

		return r.Evento
	case 1:
		r.SchemaAvro = NewUnionStringNull()

		return r.SchemaAvro
	case 2:
		r.DescripcionProceso = NewUnionStringNull()

		return r.DescripcionProceso
	case 3:
		r.Almacen = NewUnionStringNull()

		return r.Almacen
	case 4:
		r.Propietario = NewUnionStringNull()

		return r.Propietario
	case 5:
		r.Instancia = NewUnionStringNull()

		return r.Instancia
	case 6:
		r.Key1 = NewUnionStringNull()

		return r.Key1
	case 7:
		r.Key2 = NewUnionStringNull()

		return r.Key2
	case 8:
		r.Key3 = NewUnionStringNull()

		return r.Key3
	case 9:
		r.Key4 = NewUnionStringNull()

		return r.Key4
	case 10:
		r.Key5 = NewUnionStringNull()

		return r.Key5
	case 11:
		r.Key6 = NewUnionStringNull()

		return r.Key6
	case 12:
		w := types.Long{Target: &r.FechaGeneracion}

		return w

	}
	panic("Unknown field index")
}

func (r *Identificacion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Identificacion) NullField(i int) {
	switch i {
	case 0:
		r.Evento = nil
		return
	case 1:
		r.SchemaAvro = nil
		return
	case 2:
		r.DescripcionProceso = nil
		return
	case 3:
		r.Almacen = nil
		return
	case 4:
		r.Propietario = nil
		return
	case 5:
		r.Instancia = nil
		return
	case 6:
		r.Key1 = nil
		return
	case 7:
		r.Key2 = nil
		return
	case 8:
		r.Key3 = nil
		return
	case 9:
		r.Key4 = nil
		return
	case 10:
		r.Key5 = nil
		return
	case 11:
		r.Key6 = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Identificacion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Identificacion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Identificacion) HintSize(int)                     { panic("Unsupported operation") }
func (_ Identificacion) Finalize()                        {}

func (_ Identificacion) AvroCRC64Fingerprint() []byte {
	return []byte(IdentificacionAvroCRC64Fingerprint)
}

func (r Identificacion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Evento"], err = json.Marshal(r.Evento)
	if err != nil {
		return nil, err
	}
	output["SchemaAvro"], err = json.Marshal(r.SchemaAvro)
	if err != nil {
		return nil, err
	}
	output["DescripcionProceso"], err = json.Marshal(r.DescripcionProceso)
	if err != nil {
		return nil, err
	}
	output["Almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["Propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["Instancia"], err = json.Marshal(r.Instancia)
	if err != nil {
		return nil, err
	}
	output["Key1"], err = json.Marshal(r.Key1)
	if err != nil {
		return nil, err
	}
	output["Key2"], err = json.Marshal(r.Key2)
	if err != nil {
		return nil, err
	}
	output["Key3"], err = json.Marshal(r.Key3)
	if err != nil {
		return nil, err
	}
	output["Key4"], err = json.Marshal(r.Key4)
	if err != nil {
		return nil, err
	}
	output["Key5"], err = json.Marshal(r.Key5)
	if err != nil {
		return nil, err
	}
	output["Key6"], err = json.Marshal(r.Key6)
	if err != nil {
		return nil, err
	}
	output["FechaGeneracion"], err = json.Marshal(r.FechaGeneracion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Identificacion) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Evento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Evento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Evento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SchemaAvro"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SchemaAvro); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SchemaAvro")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DescripcionProceso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DescripcionProceso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DescripcionProceso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Almacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Almacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Almacen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Propietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propietario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Propietario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Instancia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Instancia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Instancia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Key1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Key1); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Key1")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Key2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Key2); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Key2")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Key3"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Key3); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Key3")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Key4"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Key4); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Key4")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Key5"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Key5); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Key5")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Key6"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Key6); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Key6")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaGeneracion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaGeneracion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaGeneracion")
	}
	return nil
}