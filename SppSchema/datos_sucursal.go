// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Geocerca.avsc
 */
package SppSchemaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DatosSucursal struct {
	Codigo *UnionNullString `json:"codigo"`

	Id *UnionNullString `json:"id"`

	Nombre *UnionNullString `json:"nombre"`
}

const DatosSucursalAvroCRC64Fingerprint = "\x9e\x17!\xf9\x92e\xcf\xd7"

func NewDatosSucursal() DatosSucursal {
	r := DatosSucursal{}
	return r
}

func DeserializeDatosSucursal(r io.Reader) (DatosSucursal, error) {
	t := NewDatosSucursal()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDatosSucursalFromSchema(r io.Reader, schema string) (DatosSucursal, error) {
	t := NewDatosSucursal()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDatosSucursal(r DatosSucursal, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Codigo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Id, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Nombre, w)
	if err != nil {
		return err
	}
	return err
}

func (r DatosSucursal) Serialize(w io.Writer) error {
	return writeDatosSucursal(r, w)
}

func (r DatosSucursal) Schema() string {
	return "{\"fields\":[{\"name\":\"codigo\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":[\"null\",\"string\"]},{\"name\":\"nombre\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.SppSchema.Events.DatosSucursal\",\"type\":\"record\"}"
}

func (r DatosSucursal) SchemaName() string {
	return "Andreani.SppSchema.Events.DatosSucursal"
}

func (_ DatosSucursal) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DatosSucursal) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DatosSucursal) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DatosSucursal) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DatosSucursal) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DatosSucursal) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DatosSucursal) SetString(v string)   { panic("Unsupported operation") }
func (_ DatosSucursal) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DatosSucursal) Get(i int) types.Field {
	switch i {
	case 0:
		r.Codigo = NewUnionNullString()

		return r.Codigo
	case 1:
		r.Id = NewUnionNullString()

		return r.Id
	case 2:
		r.Nombre = NewUnionNullString()

		return r.Nombre
	}
	panic("Unknown field index")
}

func (r *DatosSucursal) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *DatosSucursal) NullField(i int) {
	switch i {
	case 0:
		r.Codigo = nil
		return
	case 1:
		r.Id = nil
		return
	case 2:
		r.Nombre = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DatosSucursal) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DatosSucursal) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DatosSucursal) HintSize(int)                     { panic("Unsupported operation") }
func (_ DatosSucursal) Finalize()                        {}

func (_ DatosSucursal) AvroCRC64Fingerprint() []byte {
	return []byte(DatosSucursalAvroCRC64Fingerprint)
}

func (r DatosSucursal) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["codigo"], err = json.Marshal(r.Codigo)
	if err != nil {
		return nil, err
	}
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DatosSucursal) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["codigo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Codigo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
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
		return fmt.Errorf("no value specified for nombre")
	}
	return nil
}
