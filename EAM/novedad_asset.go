// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NovedadEAMPowerApp.avsc
 */
package EAMEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NovedadAsset struct {
	Nueva_planta *UnionNullString `json:"nueva_planta"`

	Ceco *UnionNullString `json:"ceco"`

	FechaNovedad *UnionNullString `json:"fechaNovedad"`

	UsuarioNovedad *UnionNullString `json:"usuarioNovedad"`

	Clase *UnionNullString `json:"clase"`

	Marca *UnionNullString `json:"marca"`

	Modelo *UnionNullString `json:"modelo"`

	Nro_serie *UnionNullString `json:"nro_serie"`

	Razon_social *UnionNullString `json:"razon_social"`

	Propulsion *UnionNullString `json:"propulsion"`
}

const NovedadAssetAvroCRC64Fingerprint = "s\xef\x89yԃ\xe2t"

func NewNovedadAsset() NovedadAsset {
	r := NovedadAsset{}
	r.Nueva_planta = nil
	r.Ceco = nil
	r.FechaNovedad = nil
	r.UsuarioNovedad = nil
	r.Clase = nil
	r.Marca = nil
	r.Modelo = nil
	r.Nro_serie = nil
	r.Razon_social = nil
	r.Propulsion = nil
	return r
}

func DeserializeNovedadAsset(r io.Reader) (NovedadAsset, error) {
	t := NewNovedadAsset()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNovedadAssetFromSchema(r io.Reader, schema string) (NovedadAsset, error) {
	t := NewNovedadAsset()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNovedadAsset(r NovedadAsset, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.Nueva_planta, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Ceco, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaNovedad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.UsuarioNovedad, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Clase, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Marca, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Modelo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Nro_serie, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Razon_social, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Propulsion, w)
	if err != nil {
		return err
	}
	return err
}

func (r NovedadAsset) Serialize(w io.Writer) error {
	return writeNovedadAsset(r, w)
}

func (r NovedadAsset) Schema() string {
	return "{\"fields\":[{\"default\":null,\"name\":\"nueva_planta\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ceco\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaNovedad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"usuarioNovedad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"clase\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"marca\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"modelo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nro_serie\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"razon_social\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"propulsion\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.EAM.Events.Sharepoint.NovedadAsset\",\"type\":\"record\"}"
}

func (r NovedadAsset) SchemaName() string {
	return "Andreani.EAM.Events.Sharepoint.NovedadAsset"
}

func (_ NovedadAsset) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NovedadAsset) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NovedadAsset) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NovedadAsset) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NovedadAsset) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NovedadAsset) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NovedadAsset) SetString(v string)   { panic("Unsupported operation") }
func (_ NovedadAsset) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NovedadAsset) Get(i int) types.Field {
	switch i {
	case 0:
		r.Nueva_planta = NewUnionNullString()

		return r.Nueva_planta
	case 1:
		r.Ceco = NewUnionNullString()

		return r.Ceco
	case 2:
		r.FechaNovedad = NewUnionNullString()

		return r.FechaNovedad
	case 3:
		r.UsuarioNovedad = NewUnionNullString()

		return r.UsuarioNovedad
	case 4:
		r.Clase = NewUnionNullString()

		return r.Clase
	case 5:
		r.Marca = NewUnionNullString()

		return r.Marca
	case 6:
		r.Modelo = NewUnionNullString()

		return r.Modelo
	case 7:
		r.Nro_serie = NewUnionNullString()

		return r.Nro_serie
	case 8:
		r.Razon_social = NewUnionNullString()

		return r.Razon_social
	case 9:
		r.Propulsion = NewUnionNullString()

		return r.Propulsion
	}
	panic("Unknown field index")
}

func (r *NovedadAsset) SetDefault(i int) {
	switch i {
	case 0:
		r.Nueva_planta = nil
		return
	case 1:
		r.Ceco = nil
		return
	case 2:
		r.FechaNovedad = nil
		return
	case 3:
		r.UsuarioNovedad = nil
		return
	case 4:
		r.Clase = nil
		return
	case 5:
		r.Marca = nil
		return
	case 6:
		r.Modelo = nil
		return
	case 7:
		r.Nro_serie = nil
		return
	case 8:
		r.Razon_social = nil
		return
	case 9:
		r.Propulsion = nil
		return
	}
	panic("Unknown field index")
}

func (r *NovedadAsset) NullField(i int) {
	switch i {
	case 0:
		r.Nueva_planta = nil
		return
	case 1:
		r.Ceco = nil
		return
	case 2:
		r.FechaNovedad = nil
		return
	case 3:
		r.UsuarioNovedad = nil
		return
	case 4:
		r.Clase = nil
		return
	case 5:
		r.Marca = nil
		return
	case 6:
		r.Modelo = nil
		return
	case 7:
		r.Nro_serie = nil
		return
	case 8:
		r.Razon_social = nil
		return
	case 9:
		r.Propulsion = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ NovedadAsset) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NovedadAsset) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NovedadAsset) HintSize(int)                     { panic("Unsupported operation") }
func (_ NovedadAsset) Finalize()                        {}

func (_ NovedadAsset) AvroCRC64Fingerprint() []byte {
	return []byte(NovedadAssetAvroCRC64Fingerprint)
}

func (r NovedadAsset) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["nueva_planta"], err = json.Marshal(r.Nueva_planta)
	if err != nil {
		return nil, err
	}
	output["ceco"], err = json.Marshal(r.Ceco)
	if err != nil {
		return nil, err
	}
	output["fechaNovedad"], err = json.Marshal(r.FechaNovedad)
	if err != nil {
		return nil, err
	}
	output["usuarioNovedad"], err = json.Marshal(r.UsuarioNovedad)
	if err != nil {
		return nil, err
	}
	output["clase"], err = json.Marshal(r.Clase)
	if err != nil {
		return nil, err
	}
	output["marca"], err = json.Marshal(r.Marca)
	if err != nil {
		return nil, err
	}
	output["modelo"], err = json.Marshal(r.Modelo)
	if err != nil {
		return nil, err
	}
	output["nro_serie"], err = json.Marshal(r.Nro_serie)
	if err != nil {
		return nil, err
	}
	output["razon_social"], err = json.Marshal(r.Razon_social)
	if err != nil {
		return nil, err
	}
	output["propulsion"], err = json.Marshal(r.Propulsion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NovedadAsset) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["nueva_planta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nueva_planta); err != nil {
			return err
		}
	} else {
		r.Nueva_planta = NewUnionNullString()

		r.Nueva_planta = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ceco"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ceco); err != nil {
			return err
		}
	} else {
		r.Ceco = NewUnionNullString()

		r.Ceco = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fechaNovedad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaNovedad); err != nil {
			return err
		}
	} else {
		r.FechaNovedad = NewUnionNullString()

		r.FechaNovedad = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["usuarioNovedad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UsuarioNovedad); err != nil {
			return err
		}
	} else {
		r.UsuarioNovedad = NewUnionNullString()

		r.UsuarioNovedad = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["clase"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Clase); err != nil {
			return err
		}
	} else {
		r.Clase = NewUnionNullString()

		r.Clase = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["marca"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Marca); err != nil {
			return err
		}
	} else {
		r.Marca = NewUnionNullString()

		r.Marca = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["modelo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Modelo); err != nil {
			return err
		}
	} else {
		r.Modelo = NewUnionNullString()

		r.Modelo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["nro_serie"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nro_serie); err != nil {
			return err
		}
	} else {
		r.Nro_serie = NewUnionNullString()

		r.Nro_serie = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["razon_social"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Razon_social); err != nil {
			return err
		}
	} else {
		r.Razon_social = NewUnionNullString()

		r.Razon_social = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["propulsion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propulsion); err != nil {
			return err
		}
	} else {
		r.Propulsion = NewUnionNullString()

		r.Propulsion = nil
	}
	return nil
}
