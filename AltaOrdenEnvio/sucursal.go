// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TransactionEvent.avsc
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

type Sucursal struct {
	ID *UnionNullString `json:"ID"`

	Nomenclatura *UnionNullString `json:"nomenclatura"`

	Descripcion *UnionNullString `json:"descripcion"`

	Direccion *UnionNullDireccionPostal `json:"direccion"`

	DatosAdicionales *UnionNullListaDePropiedades `json:"datosAdicionales"`

	ListaDeTelefonos *UnionNullListaDeTelefonos `json:"listaDeTelefonos"`
}

const SucursalAvroCRC64Fingerprint = "\x9dVHfDx\xb9\x03"

func NewSucursal() Sucursal {
	r := Sucursal{}
	r.ID = nil
	r.Nomenclatura = nil
	r.Descripcion = nil
	r.Direccion = nil
	r.DatosAdicionales = nil
	r.ListaDeTelefonos = nil
	return r
}

func DeserializeSucursal(r io.Reader) (Sucursal, error) {
	t := NewSucursal()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSucursalFromSchema(r io.Reader, schema string) (Sucursal, error) {
	t := NewSucursal()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSucursal(r Sucursal, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.ID, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Nomenclatura, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDireccionPostal(r.Direccion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullListaDePropiedades(r.DatosAdicionales, w)
	if err != nil {
		return err
	}
	err = writeUnionNullListaDeTelefonos(r.ListaDeTelefonos, w)
	if err != nil {
		return err
	}
	return err
}

func (r Sucursal) Serialize(w io.Writer) error {
	return writeSucursal(r, w)
}

func (r Sucursal) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"ID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nomenclatura\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"direccion\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"casillaDeCorreo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]}],\"name\":\"DireccionPostal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDePropiedades\"]},{\"default\":null,\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"tipo\",\"type\":\"string\"},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]}],\"name\":\"Andreani.AltaOrdenEnvio.Events.Common.Sucursal\",\"type\":\"record\"}"
}

func (r Sucursal) SchemaName() string {
	return "Andreani.AltaOrdenEnvio.Events.Common.Sucursal"
}

func (_ Sucursal) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Sucursal) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Sucursal) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Sucursal) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Sucursal) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Sucursal) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Sucursal) SetString(v string)   { panic("Unsupported operation") }
func (_ Sucursal) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Sucursal) Get(i int) types.Field {
	switch i {
	case 0:
		r.ID = NewUnionNullString()

		return r.ID
	case 1:
		r.Nomenclatura = NewUnionNullString()

		return r.Nomenclatura
	case 2:
		r.Descripcion = NewUnionNullString()

		return r.Descripcion
	case 3:
		r.Direccion = NewUnionNullDireccionPostal()

		return r.Direccion
	case 4:
		r.DatosAdicionales = NewUnionNullListaDePropiedades()

		return r.DatosAdicionales
	case 5:
		r.ListaDeTelefonos = NewUnionNullListaDeTelefonos()

		return r.ListaDeTelefonos
	}
	panic("Unknown field index")
}

func (r *Sucursal) SetDefault(i int) {
	switch i {
	case 0:
		r.ID = nil
		return
	case 1:
		r.Nomenclatura = nil
		return
	case 2:
		r.Descripcion = nil
		return
	case 3:
		r.Direccion = nil
		return
	case 4:
		r.DatosAdicionales = nil
		return
	case 5:
		r.ListaDeTelefonos = nil
		return
	}
	panic("Unknown field index")
}

func (r *Sucursal) NullField(i int) {
	switch i {
	case 0:
		r.ID = nil
		return
	case 1:
		r.Nomenclatura = nil
		return
	case 2:
		r.Descripcion = nil
		return
	case 3:
		r.Direccion = nil
		return
	case 4:
		r.DatosAdicionales = nil
		return
	case 5:
		r.ListaDeTelefonos = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Sucursal) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Sucursal) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Sucursal) HintSize(int)                     { panic("Unsupported operation") }
func (_ Sucursal) Finalize()                        {}

func (_ Sucursal) AvroCRC64Fingerprint() []byte {
	return []byte(SucursalAvroCRC64Fingerprint)
}

func (r Sucursal) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["ID"], err = json.Marshal(r.ID)
	if err != nil {
		return nil, err
	}
	output["nomenclatura"], err = json.Marshal(r.Nomenclatura)
	if err != nil {
		return nil, err
	}
	output["descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["direccion"], err = json.Marshal(r.Direccion)
	if err != nil {
		return nil, err
	}
	output["datosAdicionales"], err = json.Marshal(r.DatosAdicionales)
	if err != nil {
		return nil, err
	}
	output["listaDeTelefonos"], err = json.Marshal(r.ListaDeTelefonos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Sucursal) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["ID"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ID); err != nil {
			return err
		}
	} else {
		r.ID = NewUnionNullString()

		r.ID = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["nomenclatura"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nomenclatura); err != nil {
			return err
		}
	} else {
		r.Nomenclatura = NewUnionNullString()

		r.Nomenclatura = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["descripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Descripcion); err != nil {
			return err
		}
	} else {
		r.Descripcion = NewUnionNullString()

		r.Descripcion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["direccion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Direccion); err != nil {
			return err
		}
	} else {
		r.Direccion = NewUnionNullDireccionPostal()

		r.Direccion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["datosAdicionales"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DatosAdicionales); err != nil {
			return err
		}
	} else {
		r.DatosAdicionales = NewUnionNullListaDePropiedades()

		r.DatosAdicionales = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["listaDeTelefonos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ListaDeTelefonos); err != nil {
			return err
		}
	} else {
		r.ListaDeTelefonos = NewUnionNullListaDeTelefonos()

		r.ListaDeTelefonos = nil
	}
	return nil
}
