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
	Id int32 `json:"Id"`

	Inventario string `json:"Inventario"`

	Categoria string `json:"Categoria"`

	Planta string `json:"Planta"`

	Descripcion string `json:"Descripcion"`

	FueraDeServicio bool `json:"fueraDeServicio"`
}

const AssetAvroCRC64Fingerprint = "\xac\xe3+\x8eq\xfb(\xe8"

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
	err = vm.WriteString(r.Inventario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Categoria, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Planta, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.FueraDeServicio, w)
	if err != nil {
		return err
	}
	return err
}

func (r Asset) Serialize(w io.Writer) error {
	return writeAsset(r, w)
}

func (r Asset) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"int\"},{\"name\":\"Inventario\",\"type\":\"string\"},{\"name\":\"Categoria\",\"type\":\"string\"},{\"name\":\"Planta\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"fueraDeServicio\",\"type\":\"boolean\"}],\"name\":\"Andreani.EAM.Events.Sharepoint.Asset\",\"type\":\"record\"}"
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
		w := types.String{Target: &r.Inventario}

		return w

	case 2:
		w := types.String{Target: &r.Categoria}

		return w

	case 3:
		w := types.String{Target: &r.Planta}

		return w

	case 4:
		w := types.String{Target: &r.Descripcion}

		return w

	case 5:
		w := types.Boolean{Target: &r.FueraDeServicio}

		return w

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
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["Inventario"], err = json.Marshal(r.Inventario)
	if err != nil {
		return nil, err
	}
	output["Categoria"], err = json.Marshal(r.Categoria)
	if err != nil {
		return nil, err
	}
	output["Planta"], err = json.Marshal(r.Planta)
	if err != nil {
		return nil, err
	}
	output["Descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["fueraDeServicio"], err = json.Marshal(r.FueraDeServicio)
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
		if v, ok := fields["Id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Inventario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Inventario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Inventario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Categoria"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Categoria); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Categoria")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Planta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Planta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Planta")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Descripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Descripcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Descripcion")
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
	return nil
}
