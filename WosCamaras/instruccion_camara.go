// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     InstruccionCamara.avsc
 */
package WosCamarasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type InstruccionCamara struct {
	Camara Camara `json:"Camara"`

	Grabacion Grabacion `json:"Grabacion"`

	Accion AccionesCamara `json:"Accion"`
}

const InstruccionCamaraAvroCRC64Fingerprint = "\xee\x01\xeb(F\xd9\xf8!"

func NewInstruccionCamara() InstruccionCamara {
	r := InstruccionCamara{}
	r.Camara = NewCamara()

	r.Grabacion = NewGrabacion()

	return r
}

func DeserializeInstruccionCamara(r io.Reader) (InstruccionCamara, error) {
	t := NewInstruccionCamara()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInstruccionCamaraFromSchema(r io.Reader, schema string) (InstruccionCamara, error) {
	t := NewInstruccionCamara()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInstruccionCamara(r InstruccionCamara, w io.Writer) error {
	var err error
	err = writeCamara(r.Camara, w)
	if err != nil {
		return err
	}
	err = writeGrabacion(r.Grabacion, w)
	if err != nil {
		return err
	}
	err = writeAccionesCamara(r.Accion, w)
	if err != nil {
		return err
	}
	return err
}

func (r InstruccionCamara) Serialize(w io.Writer) error {
	return writeInstruccionCamara(r, w)
}

func (r InstruccionCamara) Schema() string {
	return "{\"fields\":[{\"name\":\"Camara\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":[\"null\",\"string\"]},{\"name\":\"Ip\",\"type\":[\"null\",\"string\"]},{\"name\":\"Configuracion\",\"type\":{\"fields\":[{\"name\":\"Extension\",\"type\":[\"null\",\"string\"]},{\"name\":\"Codec\",\"type\":[\"null\",\"string\"]},{\"name\":\"Fps\",\"type\":[\"null\",\"int\"]},{\"name\":\"ResolucionVideoWidth\",\"type\":[\"null\",\"int\"]},{\"name\":\"ResolucionVideoHeight\",\"type\":[\"null\",\"int\"]}],\"name\":\"ConfiguracionCamara\",\"type\":\"record\"}}],\"name\":\"Camara\",\"type\":\"record\"}},{\"name\":\"Grabacion\",\"type\":{\"fields\":[{\"name\":\"Nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"CarpetaBase\",\"type\":[\"null\",\"string\"]},{\"name\":\"TiempoMaximoDeGrabacion\",\"type\":[\"null\",\"int\"]},{\"name\":\"TiempoPorParte\",\"type\":[\"null\",\"int\"]},{\"name\":\"TamanoPorParte\",\"type\":[\"null\",\"float\"]}],\"name\":\"Grabacion\",\"type\":\"record\"}},{\"name\":\"Accion\",\"type\":{\"name\":\"AccionesCamara\",\"symbols\":[\"IniciarGrabacion\",\"DetenerGrabacion\"],\"type\":\"enum\"}}],\"name\":\"Andreani.WosCamaras.Events.Record.InstruccionCamara\",\"type\":\"record\"}"
}

func (r InstruccionCamara) SchemaName() string {
	return "Andreani.WosCamaras.Events.Record.InstruccionCamara"
}

func (_ InstruccionCamara) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ InstruccionCamara) SetInt(v int32)       { panic("Unsupported operation") }
func (_ InstruccionCamara) SetLong(v int64)      { panic("Unsupported operation") }
func (_ InstruccionCamara) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ InstruccionCamara) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ InstruccionCamara) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ InstruccionCamara) SetString(v string)   { panic("Unsupported operation") }
func (_ InstruccionCamara) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *InstruccionCamara) Get(i int) types.Field {
	switch i {
	case 0:
		r.Camara = NewCamara()

		w := types.Record{Target: &r.Camara}

		return w

	case 1:
		r.Grabacion = NewGrabacion()

		w := types.Record{Target: &r.Grabacion}

		return w

	case 2:
		w := AccionesCamaraWrapper{Target: &r.Accion}

		return w

	}
	panic("Unknown field index")
}

func (r *InstruccionCamara) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *InstruccionCamara) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ InstruccionCamara) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ InstruccionCamara) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ InstruccionCamara) HintSize(int)                     { panic("Unsupported operation") }
func (_ InstruccionCamara) Finalize()                        {}

func (_ InstruccionCamara) AvroCRC64Fingerprint() []byte {
	return []byte(InstruccionCamaraAvroCRC64Fingerprint)
}

func (r InstruccionCamara) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Camara"], err = json.Marshal(r.Camara)
	if err != nil {
		return nil, err
	}
	output["Grabacion"], err = json.Marshal(r.Grabacion)
	if err != nil {
		return nil, err
	}
	output["Accion"], err = json.Marshal(r.Accion)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *InstruccionCamara) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Camara"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Camara); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Camara")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Grabacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Grabacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Grabacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Accion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Accion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Accion")
	}
	return nil
}