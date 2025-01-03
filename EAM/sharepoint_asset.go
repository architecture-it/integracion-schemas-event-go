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

type SharepointAsset struct {
	Asset Asset `json:"Asset"`
}

const SharepointAssetAvroCRC64Fingerprint = "\xfco|\xb3\U000df8a3"

func NewSharepointAsset() SharepointAsset {
	r := SharepointAsset{}
	r.Asset = NewAsset()

	return r
}

func DeserializeSharepointAsset(r io.Reader) (SharepointAsset, error) {
	t := NewSharepointAsset()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeSharepointAssetFromSchema(r io.Reader, schema string) (SharepointAsset, error) {
	t := NewSharepointAsset()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeSharepointAsset(r SharepointAsset, w io.Writer) error {
	var err error
	err = writeAsset(r.Asset, w)
	if err != nil {
		return err
	}
	return err
}

func (r SharepointAsset) Serialize(w io.Writer) error {
	return writeSharepointAsset(r, w)
}

func (r SharepointAsset) Schema() string {
	return "{\"fields\":[{\"name\":\"Asset\",\"type\":{\"fields\":[{\"name\":\"id\",\"type\":\"int\"},{\"name\":\"tipo_objeto\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"},{\"name\":\"clase\",\"type\":\"string\"},{\"name\":\"codigo_costo\",\"type\":\"string\"},{\"name\":\"estado\",\"type\":\"string\"},{\"name\":\"fecha_alta\",\"type\":\"string\"},{\"name\":\"organizacion\",\"type\":\"string\"},{\"name\":\"fabricante\",\"type\":\"string\"},{\"name\":\"modelo\",\"type\":\"string\"},{\"name\":\"nro_serie\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"fueraDeServicio\",\"type\":\"boolean\"},{\"name\":\"cod_eam\",\"type\":\"string\"}],\"name\":\"Asset\",\"namespace\":\"Andreani.EAM.Events.Sharepoint\",\"type\":\"record\"}}],\"name\":\"Andreani.EAM.Events.Record.SharepointAsset\",\"type\":\"record\"}"
}

func (r SharepointAsset) SchemaName() string {
	return "Andreani.EAM.Events.Record.SharepointAsset"
}

func (_ SharepointAsset) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ SharepointAsset) SetInt(v int32)       { panic("Unsupported operation") }
func (_ SharepointAsset) SetLong(v int64)      { panic("Unsupported operation") }
func (_ SharepointAsset) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ SharepointAsset) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ SharepointAsset) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ SharepointAsset) SetString(v string)   { panic("Unsupported operation") }
func (_ SharepointAsset) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *SharepointAsset) Get(i int) types.Field {
	switch i {
	case 0:
		r.Asset = NewAsset()

		w := types.Record{Target: &r.Asset}

		return w

	}
	panic("Unknown field index")
}

func (r *SharepointAsset) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *SharepointAsset) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ SharepointAsset) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ SharepointAsset) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ SharepointAsset) HintSize(int)                     { panic("Unsupported operation") }
func (_ SharepointAsset) Finalize()                        {}

func (_ SharepointAsset) AvroCRC64Fingerprint() []byte {
	return []byte(SharepointAssetAvroCRC64Fingerprint)
}

func (r SharepointAsset) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Asset"], err = json.Marshal(r.Asset)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *SharepointAsset) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Asset"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Asset); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Asset")
	}
	return nil
}
