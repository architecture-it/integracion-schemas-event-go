// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoDeTraza.avsc
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

type EventoDeTraza struct {
	CodigoDeEnvio string `json:"codigoDeEnvio"`

	Nombre *UnionNullString `json:"nombre"`

	Cuando string `json:"cuando"`

	CodigoDeContratoInterno string `json:"codigoDeContratoInterno"`

	EstadoDelEnvio *UnionNullString `json:"estadoDelEnvio"`

	CicloDelEnvio *UnionNullString `json:"cicloDelEnvio"`

	Operador *UnionNullString `json:"operador"`

	EstadoDeLaRendicion *UnionNullString `json:"estadoDeLaRendicion"`

	Comentario *UnionNullString `json:"comentario"`

	SucursalAsociadaAlEvento *UnionNullDatosSucursal `json:"sucursalAsociadaAlEvento"`
}

const EventoDeTrazaAvroCRC64Fingerprint = "z|&mX\x05a\xf0"

func NewEventoDeTraza() EventoDeTraza {
	r := EventoDeTraza{}
	return r
}

func DeserializeEventoDeTraza(r io.Reader) (EventoDeTraza, error) {
	t := NewEventoDeTraza()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoDeTrazaFromSchema(r io.Reader, schema string) (EventoDeTraza, error) {
	t := NewEventoDeTraza()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoDeTraza(r EventoDeTraza, w io.Writer) error {
	var err error
	err = vm.WriteString(r.CodigoDeEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Cuando, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoDeContratoInterno, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstadoDelEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CicloDelEnvio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Operador, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.EstadoDeLaRendicion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Comentario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosSucursal(r.SucursalAsociadaAlEvento, w)
	if err != nil {
		return err
	}
	return err
}

func (r EventoDeTraza) Serialize(w io.Writer) error {
	return writeEventoDeTraza(r, w)
}

func (r EventoDeTraza) Schema() string {
	return "{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":\"string\"},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":[\"null\",\"string\"]},{\"name\":\"nombre\",\"type\":[\"null\",\"string\"]}],\"name\":\"DatosSucursal\",\"type\":\"record\"}]}],\"name\":\"Andreani.SppSchema.Events.EventoDeTraza\",\"type\":\"record\"}"
}

func (r EventoDeTraza) SchemaName() string {
	return "Andreani.SppSchema.Events.EventoDeTraza"
}

func (_ EventoDeTraza) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoDeTraza) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoDeTraza) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoDeTraza) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoDeTraza) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoDeTraza) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoDeTraza) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoDeTraza) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoDeTraza) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.CodigoDeEnvio}

		return w

	case 1:
		r.Nombre = NewUnionNullString()

		return r.Nombre
	case 2:
		w := types.String{Target: &r.Cuando}

		return w

	case 3:
		w := types.String{Target: &r.CodigoDeContratoInterno}

		return w

	case 4:
		r.EstadoDelEnvio = NewUnionNullString()

		return r.EstadoDelEnvio
	case 5:
		r.CicloDelEnvio = NewUnionNullString()

		return r.CicloDelEnvio
	case 6:
		r.Operador = NewUnionNullString()

		return r.Operador
	case 7:
		r.EstadoDeLaRendicion = NewUnionNullString()

		return r.EstadoDeLaRendicion
	case 8:
		r.Comentario = NewUnionNullString()

		return r.Comentario
	case 9:
		r.SucursalAsociadaAlEvento = NewUnionNullDatosSucursal()

		return r.SucursalAsociadaAlEvento
	}
	panic("Unknown field index")
}

func (r *EventoDeTraza) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoDeTraza) NullField(i int) {
	switch i {
	case 1:
		r.Nombre = nil
		return
	case 4:
		r.EstadoDelEnvio = nil
		return
	case 5:
		r.CicloDelEnvio = nil
		return
	case 6:
		r.Operador = nil
		return
	case 7:
		r.EstadoDeLaRendicion = nil
		return
	case 8:
		r.Comentario = nil
		return
	case 9:
		r.SucursalAsociadaAlEvento = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EventoDeTraza) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventoDeTraza) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventoDeTraza) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventoDeTraza) Finalize()                        {}

func (_ EventoDeTraza) AvroCRC64Fingerprint() []byte {
	return []byte(EventoDeTrazaAvroCRC64Fingerprint)
}

func (r EventoDeTraza) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["codigoDeEnvio"], err = json.Marshal(r.CodigoDeEnvio)
	if err != nil {
		return nil, err
	}
	output["nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["cuando"], err = json.Marshal(r.Cuando)
	if err != nil {
		return nil, err
	}
	output["codigoDeContratoInterno"], err = json.Marshal(r.CodigoDeContratoInterno)
	if err != nil {
		return nil, err
	}
	output["estadoDelEnvio"], err = json.Marshal(r.EstadoDelEnvio)
	if err != nil {
		return nil, err
	}
	output["cicloDelEnvio"], err = json.Marshal(r.CicloDelEnvio)
	if err != nil {
		return nil, err
	}
	output["operador"], err = json.Marshal(r.Operador)
	if err != nil {
		return nil, err
	}
	output["estadoDeLaRendicion"], err = json.Marshal(r.EstadoDeLaRendicion)
	if err != nil {
		return nil, err
	}
	output["comentario"], err = json.Marshal(r.Comentario)
	if err != nil {
		return nil, err
	}
	output["sucursalAsociadaAlEvento"], err = json.Marshal(r.SucursalAsociadaAlEvento)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EventoDeTraza) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["codigoDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeEnvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigoDeEnvio")
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
	val = func() json.RawMessage {
		if v, ok := fields["cuando"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cuando); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cuando")
	}
	val = func() json.RawMessage {
		if v, ok := fields["codigoDeContratoInterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeContratoInterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for codigoDeContratoInterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["estadoDelEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoDelEnvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for estadoDelEnvio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["cicloDelEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CicloDelEnvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for cicloDelEnvio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["operador"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Operador); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for operador")
	}
	val = func() json.RawMessage {
		if v, ok := fields["estadoDeLaRendicion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoDeLaRendicion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for estadoDeLaRendicion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["comentario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Comentario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for comentario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["sucursalAsociadaAlEvento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SucursalAsociadaAlEvento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for sucursalAsociadaAlEvento")
	}
	return nil
}
