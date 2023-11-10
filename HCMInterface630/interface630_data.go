// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListInterface630.avsc
 */
package HCMInterface630Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Interface630Data struct {
	TipoDeEstructura string `json:"TipoDeEstructura"`

	Legajo int64 `json:"Legajo"`

	FechaDesde string `json:"FechaDesde"`

	Estructura *UnionNullString `json:"Estructura"`

	FechaHasta *UnionNullString `json:"FechaHasta"`

	TipoMotivo *UnionNullString `json:"TipoMotivo"`

	Motivo *UnionNullString `json:"Motivo"`

	TipoDeIL *UnionNullString `json:"TipoDeIL"`

	NumeroDeIL *UnionNullString `json:"NumeroDeIL"`

	NumeroDeExpediente *UnionNullString `json:"NumeroDeExpediente"`
}

const Interface630DataAvroCRC64Fingerprint = "\xf6\xce\xce\xdf\v\xfew4"

func NewInterface630Data() Interface630Data {
	r := Interface630Data{}
	r.FechaHasta = nil
	r.TipoMotivo = nil
	r.Motivo = nil
	r.TipoDeIL = nil
	r.NumeroDeIL = nil
	r.NumeroDeExpediente = nil
	return r
}

func DeserializeInterface630Data(r io.Reader) (Interface630Data, error) {
	t := NewInterface630Data()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInterface630DataFromSchema(r io.Reader, schema string) (Interface630Data, error) {
	t := NewInterface630Data()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInterface630Data(r Interface630Data, w io.Writer) error {
	var err error
	err = vm.WriteString(r.TipoDeEstructura, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Legajo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaDesde, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Estructura, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FechaHasta, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoMotivo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Motivo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.TipoDeIL, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeIL, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDeExpediente, w)
	if err != nil {
		return err
	}
	return err
}

func (r Interface630Data) Serialize(w io.Writer) error {
	return writeInterface630Data(r, w)
}

func (r Interface630Data) Schema() string {
	return "{\"fields\":[{\"name\":\"TipoDeEstructura\",\"type\":\"string\"},{\"name\":\"Legajo\",\"type\":\"long\"},{\"name\":\"FechaDesde\",\"type\":\"string\"},{\"name\":\"Estructura\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaHasta\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoMotivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Motivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TipoDeIL\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeIL\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"NumeroDeExpediente\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.HCMInterface630.Events.Record.Interface630Data\",\"type\":\"record\"}"
}

func (r Interface630Data) SchemaName() string {
	return "Andreani.HCMInterface630.Events.Record.Interface630Data"
}

func (_ Interface630Data) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Interface630Data) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Interface630Data) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Interface630Data) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Interface630Data) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Interface630Data) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Interface630Data) SetString(v string)   { panic("Unsupported operation") }
func (_ Interface630Data) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Interface630Data) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.TipoDeEstructura}

		return w

	case 1:
		w := types.Long{Target: &r.Legajo}

		return w

	case 2:
		w := types.String{Target: &r.FechaDesde}

		return w

	case 3:
		r.Estructura = NewUnionNullString()

		return r.Estructura
	case 4:
		r.FechaHasta = NewUnionNullString()

		return r.FechaHasta
	case 5:
		r.TipoMotivo = NewUnionNullString()

		return r.TipoMotivo
	case 6:
		r.Motivo = NewUnionNullString()

		return r.Motivo
	case 7:
		r.TipoDeIL = NewUnionNullString()

		return r.TipoDeIL
	case 8:
		r.NumeroDeIL = NewUnionNullString()

		return r.NumeroDeIL
	case 9:
		r.NumeroDeExpediente = NewUnionNullString()

		return r.NumeroDeExpediente
	}
	panic("Unknown field index")
}

func (r *Interface630Data) SetDefault(i int) {
	switch i {
	case 4:
		r.FechaHasta = nil
		return
	case 5:
		r.TipoMotivo = nil
		return
	case 6:
		r.Motivo = nil
		return
	case 7:
		r.TipoDeIL = nil
		return
	case 8:
		r.NumeroDeIL = nil
		return
	case 9:
		r.NumeroDeExpediente = nil
		return
	}
	panic("Unknown field index")
}

func (r *Interface630Data) NullField(i int) {
	switch i {
	case 3:
		r.Estructura = nil
		return
	case 4:
		r.FechaHasta = nil
		return
	case 5:
		r.TipoMotivo = nil
		return
	case 6:
		r.Motivo = nil
		return
	case 7:
		r.TipoDeIL = nil
		return
	case 8:
		r.NumeroDeIL = nil
		return
	case 9:
		r.NumeroDeExpediente = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Interface630Data) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Interface630Data) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Interface630Data) HintSize(int)                     { panic("Unsupported operation") }
func (_ Interface630Data) Finalize()                        {}

func (_ Interface630Data) AvroCRC64Fingerprint() []byte {
	return []byte(Interface630DataAvroCRC64Fingerprint)
}

func (r Interface630Data) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["TipoDeEstructura"], err = json.Marshal(r.TipoDeEstructura)
	if err != nil {
		return nil, err
	}
	output["Legajo"], err = json.Marshal(r.Legajo)
	if err != nil {
		return nil, err
	}
	output["FechaDesde"], err = json.Marshal(r.FechaDesde)
	if err != nil {
		return nil, err
	}
	output["Estructura"], err = json.Marshal(r.Estructura)
	if err != nil {
		return nil, err
	}
	output["FechaHasta"], err = json.Marshal(r.FechaHasta)
	if err != nil {
		return nil, err
	}
	output["TipoMotivo"], err = json.Marshal(r.TipoMotivo)
	if err != nil {
		return nil, err
	}
	output["Motivo"], err = json.Marshal(r.Motivo)
	if err != nil {
		return nil, err
	}
	output["TipoDeIL"], err = json.Marshal(r.TipoDeIL)
	if err != nil {
		return nil, err
	}
	output["NumeroDeIL"], err = json.Marshal(r.NumeroDeIL)
	if err != nil {
		return nil, err
	}
	output["NumeroDeExpediente"], err = json.Marshal(r.NumeroDeExpediente)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Interface630Data) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeEstructura"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeEstructura); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoDeEstructura")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Legajo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Legajo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Legajo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaDesde"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaDesde); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaDesde")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estructura"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estructura); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Estructura")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaHasta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaHasta); err != nil {
			return err
		}
	} else {
		r.FechaHasta = NewUnionNullString()

		r.FechaHasta = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoMotivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoMotivo); err != nil {
			return err
		}
	} else {
		r.TipoMotivo = NewUnionNullString()

		r.TipoMotivo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Motivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Motivo); err != nil {
			return err
		}
	} else {
		r.Motivo = NewUnionNullString()

		r.Motivo = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeIL"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeIL); err != nil {
			return err
		}
	} else {
		r.TipoDeIL = NewUnionNullString()

		r.TipoDeIL = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeIL"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeIL); err != nil {
			return err
		}
	} else {
		r.NumeroDeIL = NewUnionNullString()

		r.NumeroDeIL = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeExpediente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeExpediente); err != nil {
			return err
		}
	} else {
		r.NumeroDeExpediente = NewUnionNullString()

		r.NumeroDeExpediente = nil
	}
	return nil
}
