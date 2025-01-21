// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     SharepointAsset.avsc
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

type Asset struct {
	Id int32 `json:"id"`

	Tipo_objeto string `json:"tipo_objeto"`

	Descripcion string `json:"descripcion"`

	Clase string `json:"clase"`

	Codigo_costo string `json:"codigo_costo"`

	Estado string `json:"estado"`

	Fecha_alta string `json:"fecha_alta"`

	Organizacion string `json:"organizacion"`

	Fabricante string `json:"fabricante"`

	Modelo string `json:"modelo"`

	Nro_serie string `json:"nro_serie"`

	Propietario string `json:"propietario"`

	FueraDeServicio bool `json:"fueraDeServicio"`

	Propulsion string `json:"propulsion"`

	Cod_eam *UnionNullString `json:"cod_eam"`
}

const AssetAvroCRC64Fingerprint = "\xe8\xc3E\"\xf6$\xd5+"

func NewAsset() Asset {
	r := Asset{}
	return r
}

func DeserializeAsset(r io.Reader) (Asset, error) {
	t := NewAsset()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAssetFromSchema(r io.Reader, schema string) (Asset, error) {
	t := NewAsset()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAsset(r Asset, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Tipo_objeto, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Clase, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Codigo_costo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Estado, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Fecha_alta, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Organizacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Fabricante, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Modelo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Nro_serie, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.FueraDeServicio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Propulsion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Cod_eam, w)
	if err != nil {
		return err
	}
	return err
}

func (r Asset) Serialize(w io.Writer) error {
	return writeAsset(r, w)
}

func (r Asset) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"int\"},{\"name\":\"tipo_objeto\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"clase\",\"type\":\"string\"},{\"name\":\"codigo_costo\",\"type\":\"string\"},{\"name\":\"estado\",\"type\":\"string\"},{\"name\":\"fecha_alta\",\"type\":\"string\"},{\"name\":\"organizacion\",\"type\":\"string\"},{\"name\":\"fabricante\",\"type\":\"string\"},{\"name\":\"modelo\",\"type\":\"string\"},{\"name\":\"nro_serie\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"fueraDeServicio\",\"type\":\"boolean\"},{\"name\":\"propulsion\",\"type\":\"string\"},{\"name\":\"cod_eam\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.EAM.Events.Sharepoint.Asset\",\"type\":\"record\"}"
}

func (r Asset) SchemaName() string {
	return "Andreani.EAM.Events.Sharepoint.Asset"
}

func (_ Asset) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Asset) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Asset) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Asset) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Asset) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Asset) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Asset) SetString(v string)   { panic("Unsupported operation") }
func (_ Asset) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Asset) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Tipo_objeto}

		return w

	case 2:
		w := types.String{Target: &r.Descripcion}

		return w

	case 3:
		w := types.String{Target: &r.Clase}

		return w

	case 4:
		w := types.String{Target: &r.Codigo_costo}

		return w

	case 5:
		w := types.String{Target: &r.Estado}

		return w

	case 6:
		w := types.String{Target: &r.Fecha_alta}

		return w

	case 7:
		w := types.String{Target: &r.Organizacion}

		return w

	case 8:
		w := types.String{Target: &r.Fabricante}

		return w

	case 9:
		w := types.String{Target: &r.Modelo}

		return w

	case 10:
		w := types.String{Target: &r.Nro_serie}

		return w

	case 11:
		w := types.String{Target: &r.Propietario}

		return w

	case 12:
		w := types.Boolean{Target: &r.FueraDeServicio}

		return w

	case 13:
		w := types.String{Target: &r.Propulsion}

		return w

	case 14:
		r.Cod_eam = NewUnionNullString()

		return r.Cod_eam
	}
	panic("Unknown field index")
}

func (r *Asset) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Asset) NullField(i int) {
	switch i {
	case 14:
		r.Cod_eam = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Asset) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Asset) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Asset) HintSize(int)                     { panic("Unsupported operation") }
func (_ Asset) Finalize()                        {}

func (_ Asset) AvroCRC64Fingerprint() []byte {
	return []byte(AssetAvroCRC64Fingerprint)
}

func (r Asset) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["tipo_objeto"], err = json.Marshal(r.Tipo_objeto)
	if err != nil {
		return nil, err
	}
	output["descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["clase"], err = json.Marshal(r.Clase)
	if err != nil {
		return nil, err
	}
	output["codigo_costo"], err = json.Marshal(r.Codigo_costo)
	if err != nil {
		return nil, err
	}
	output["estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	output["fecha_alta"], err = json.Marshal(r.Fecha_alta)
	if err != nil {
		return nil, err
	}
	output["organizacion"], err = json.Marshal(r.Organizacion)
	if err != nil {
		return nil, err
	}
	output["fabricante"], err = json.Marshal(r.Fabricante)
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
	output["propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["fueraDeServicio"], err = json.Marshal(r.FueraDeServicio)
	if err != nil {
		return nil, err
	}
	output["propulsion"], err = json.Marshal(r.Propulsion)
	if err != nil {
		return nil, err
	}
	output["cod_eam"], err = json.Marshal(r.Cod_eam)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Asset) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		if v, ok := fields["tipo_objeto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tipo_objeto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipo_objeto")
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
		return fmt.Errorf("no value specified for descripcion")
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
		return fmt.Errorf("no value specified for clase")
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigo_costo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Codigo_costo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigo_costo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["estado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for estado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fecha_alta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Fecha_alta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fecha_alta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["organizacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Organizacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for organizacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fabricante"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Fabricante); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fabricante")
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
		return fmt.Errorf("no value specified for modelo")
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
		return fmt.Errorf("no value specified for nro_serie")
	}
	val = func() json.RawMessage {
		if v, ok := fields["propietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propietario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for propietario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["fueraDeServicio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FueraDeServicio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for fueraDeServicio")
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
		return fmt.Errorf("no value specified for propulsion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cod_eam"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cod_eam); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cod_eam")
	}
	return nil
}
