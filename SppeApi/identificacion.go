// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     OrdenDeRetiroRechazada.avsc
 */
package SppeApiEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Identificacion struct {
	Id string `json:"Id"`

	Evento string `json:"Evento"`

	Nombre string `json:"Nombre"`

	Proceso string `json:"Proceso"`

	FechaHoraGeneracion int64 `json:"FechaHoraGeneracion"`

	SistemaOrigen string `json:"SistemaOrigen"`

	Almacen string `json:"Almacen"`

	Propietario string `json:"Propietario"`

	Instancia string `json:"Instancia"`
}

const IdentificacionAvroCRC64Fingerprint = "\xd2\xfeop\x11Y\xe1\xfb"

func NewIdentificacion() Identificacion {
	r := Identificacion{}
	return r
}

func DeserializeIdentificacion(r io.Reader) (Identificacion, error) {
	t := NewIdentificacion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeIdentificacionFromSchema(r io.Reader, schema string) (Identificacion, error) {
	t := NewIdentificacion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeIdentificacion(r Identificacion, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Evento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Proceso, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.FechaHoraGeneracion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.SistemaOrigen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Almacen, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Instancia, w)
	if err != nil {
		return err
	}
	return err
}

func (r Identificacion) Serialize(w io.Writer) error {
	return writeIdentificacion(r, w)
}

func (r Identificacion) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"}],\"name\":\"Andreani.SppeApi.Events.OrdenDeRetiroRechazadaCommon.Identificacion\",\"type\":\"record\"}"
}

func (r Identificacion) SchemaName() string {
	return "Andreani.SppeApi.Events.OrdenDeRetiroRechazadaCommon.Identificacion"
}

func (_ Identificacion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Identificacion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Identificacion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Identificacion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Identificacion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Identificacion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Identificacion) SetString(v string)   { panic("Unsupported operation") }
func (_ Identificacion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Identificacion) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Evento}

		return w

	case 2:
		w := types.String{Target: &r.Nombre}

		return w

	case 3:
		w := types.String{Target: &r.Proceso}

		return w

	case 4:
		w := types.Long{Target: &r.FechaHoraGeneracion}

		return w

	case 5:
		w := types.String{Target: &r.SistemaOrigen}

		return w

	case 6:
		w := types.String{Target: &r.Almacen}

		return w

	case 7:
		w := types.String{Target: &r.Propietario}

		return w

	case 8:
		w := types.String{Target: &r.Instancia}

		return w

	}
	panic("Unknown field index")
}

func (r *Identificacion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Identificacion) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Identificacion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Identificacion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Identificacion) HintSize(int)                     { panic("Unsupported operation") }
func (_ Identificacion) Finalize()                        {}

func (_ Identificacion) AvroCRC64Fingerprint() []byte {
	return []byte(IdentificacionAvroCRC64Fingerprint)
}

func (r Identificacion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["Evento"], err = json.Marshal(r.Evento)
	if err != nil {
		return nil, err
	}
	output["Nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["Proceso"], err = json.Marshal(r.Proceso)
	if err != nil {
		return nil, err
	}
	output["FechaHoraGeneracion"], err = json.Marshal(r.FechaHoraGeneracion)
	if err != nil {
		return nil, err
	}
	output["SistemaOrigen"], err = json.Marshal(r.SistemaOrigen)
	if err != nil {
		return nil, err
	}
	output["Almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["Propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["Instancia"], err = json.Marshal(r.Instancia)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Identificacion) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["Evento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Evento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Evento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Nombre"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nombre); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Nombre")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Proceso"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Proceso); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Proceso")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaHoraGeneracion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaHoraGeneracion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaHoraGeneracion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SistemaOrigen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SistemaOrigen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SistemaOrigen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Almacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Almacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Almacen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Propietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propietario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Propietario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Instancia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Instancia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Instancia")
	}
	return nil
}
