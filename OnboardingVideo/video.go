// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Video.avsc
 */
package OnboardingVideoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Video struct {
	Id string `json:"id"`

	Titulo string `json:"titulo"`

	Propietario string `json:"propietario"`

	EsPrivado bool `json:"esPrivado"`

	Categorias []Categoria `json:"Categorias"`
}

const VideoAvroCRC64Fingerprint = "\xa1\x16?\xa7\xac_\xce!"

func NewVideo() Video {
	r := Video{}
	r.Categorias = make([]Categoria, 0)

	return r
}

func DeserializeVideo(r io.Reader) (Video, error) {
	t := NewVideo()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeVideoFromSchema(r io.Reader, schema string) (Video, error) {
	t := NewVideo()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeVideo(r Video, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Titulo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.EsPrivado, w)
	if err != nil {
		return err
	}
	err = writeArrayCategoria(r.Categorias, w)
	if err != nil {
		return err
	}
	return err
}

func (r Video) Serialize(w io.Writer) error {
	return writeVideo(r, w)
}

func (r Video) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"titulo\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"esPrivado\",\"type\":\"boolean\"},{\"name\":\"Categorias\",\"type\":{\"items\":{\"fields\":[{\"name\":\"nombre\",\"type\":\"string\"}],\"name\":\"Categoria\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.OnboardingVideo.Events.Record.Video\",\"type\":\"record\"}"
}

func (r Video) SchemaName() string {
	return "Andreani.OnboardingVideo.Events.Record.Video"
}

func (_ Video) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Video) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Video) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Video) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Video) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Video) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Video) SetString(v string)   { panic("Unsupported operation") }
func (_ Video) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Video) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Titulo}

		return w

	case 2:
		w := types.String{Target: &r.Propietario}

		return w

	case 3:
		w := types.Boolean{Target: &r.EsPrivado}

		return w

	case 4:
		r.Categorias = make([]Categoria, 0)

		w := ArrayCategoriaWrapper{Target: &r.Categorias}

		return w

	}
	panic("Unknown field index")
}

func (r *Video) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Video) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Video) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Video) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Video) HintSize(int)                     { panic("Unsupported operation") }
func (_ Video) Finalize()                        {}

func (_ Video) AvroCRC64Fingerprint() []byte {
	return []byte(VideoAvroCRC64Fingerprint)
}

func (r Video) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["titulo"], err = json.Marshal(r.Titulo)
	if err != nil {
		return nil, err
	}
	output["propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["esPrivado"], err = json.Marshal(r.EsPrivado)
	if err != nil {
		return nil, err
	}
	output["Categorias"], err = json.Marshal(r.Categorias)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Video) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["titulo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Titulo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for titulo")
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
		if v, ok := fields["esPrivado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EsPrivado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for esPrivado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Categorias"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Categorias); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Categorias")
	}
	return nil
}