// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
 */
package Wapv2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Distribuidor struct {
	Nombre *UnionNullString `json:"nombre"`

	Remito *UnionNullString `json:"remito"`

	Etiqueta *UnionNullString `json:"etiqueta"`

	Numerodeenvio *UnionNullString `json:"numerodeenvio"`

	Datosadicionales *UnionNullString `json:"datosadicionales"`
}

const DistribuidorAvroCRC64Fingerprint = "(]I\xaf\xb6&r\xd7"

func NewDistribuidor() Distribuidor {
	r := Distribuidor{}
	r.Nombre = nil
	r.Remito = nil
	r.Etiqueta = nil
	r.Numerodeenvio = nil
	r.Datosadicionales = nil
	return r
}

func DeserializeDistribuidor(r io.Reader) (Distribuidor, error) {
	t := NewDistribuidor()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDistribuidorFromSchema(r io.Reader, schema string) (Distribuidor, error) {
	t := NewDistribuidor()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDistribuidor(r Distribuidor, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Remito, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Etiqueta, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Numerodeenvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Datosadicionales, w)
	if err != nil {
		return err
	}
	return err
}

func (r Distribuidor) Serialize(w io.Writer) error {
	return writeDistribuidor(r, w)
}

func (r Distribuidor) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"remito\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"etiqueta\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numerodeenvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosadicionales\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.Wapv2.Events.Record.Distribuidor\",\"type\":\"record\"}"
}

func (r Distribuidor) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.Distribuidor"
}

func (_ Distribuidor) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Distribuidor) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Distribuidor) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Distribuidor) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Distribuidor) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Distribuidor) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Distribuidor) SetString(v string)   { panic("Unsupported operation") }
func (_ Distribuidor) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Distribuidor) Get(i int) types.Field {
	switch i {
	case 0:
		r.Nombre = NewUnionNullString()

		return r.Nombre
	case 1:
		r.Remito = NewUnionNullString()

		return r.Remito
	case 2:
		r.Etiqueta = NewUnionNullString()

		return r.Etiqueta
	case 3:
		r.Numerodeenvio = NewUnionNullString()

		return r.Numerodeenvio
	case 4:
		r.Datosadicionales = NewUnionNullString()

		return r.Datosadicionales
	}
	panic("Unknown field index")
}

func (r *Distribuidor) SetDefault(i int) {
	switch i {
	case 0:
		r.Nombre = nil
		return
	case 1:
		r.Remito = nil
		return
	case 2:
		r.Etiqueta = nil
		return
	case 3:
		r.Numerodeenvio = nil
		return
	case 4:
		r.Datosadicionales = nil
		return
	}
	panic("Unknown field index")
}

func (r *Distribuidor) NullField(i int) {
	switch i {
	case 0:
		r.Nombre = nil
		return
	case 1:
		r.Remito = nil
		return
	case 2:
		r.Etiqueta = nil
		return
	case 3:
		r.Numerodeenvio = nil
		return
	case 4:
		r.Datosadicionales = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Distribuidor) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Distribuidor) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Distribuidor) HintSize(int)                     { panic("Unsupported operation") }
func (_ Distribuidor) Finalize()                        {}

func (_ Distribuidor) AvroCRC64Fingerprint() []byte {
	return []byte(DistribuidorAvroCRC64Fingerprint)
}

func (r Distribuidor) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["remito"], err = json.Marshal(r.Remito)
	if err != nil {
		return nil, err
	}
	output["etiqueta"], err = json.Marshal(r.Etiqueta)
	if err != nil {
		return nil, err
	}
	output["numerodeenvio"], err = json.Marshal(r.Numerodeenvio)
	if err != nil {
		return nil, err
	}
	output["datosadicionales"], err = json.Marshal(r.Datosadicionales)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Distribuidor) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["nombre"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nombre); err != nil {
			return err
		}
	} else {
		r.Nombre = NewUnionNullString()

		r.Nombre = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["remito"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Remito); err != nil {
			return err
		}
	} else {
		r.Remito = NewUnionNullString()

		r.Remito = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["etiqueta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Etiqueta); err != nil {
			return err
		}
	} else {
		r.Etiqueta = NewUnionNullString()

		r.Etiqueta = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["numerodeenvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Numerodeenvio); err != nil {
			return err
		}
	} else {
		r.Numerodeenvio = NewUnionNullString()

		r.Numerodeenvio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["datosadicionales"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Datosadicionales); err != nil {
			return err
		}
	} else {
		r.Datosadicionales = NewUnionNullString()

		r.Datosadicionales = nil
	}
	return nil
}
