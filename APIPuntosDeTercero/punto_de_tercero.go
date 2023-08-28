// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PuntoDeTercero.avsc
 */
package APIPuntosDeTerceroEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type PuntoDeTercero struct {
	Id int32 `json:"id"`

	Nombre string `json:"nombre"`

	Telefono string `json:"telefono"`

	HorarioDeAtencion string `json:"horarioDeAtencion"`

	Observaciones string `json:"observaciones"`

	AdmiteEnvios int32 `json:"admiteEnvios"`

	EntregaEnvios int32 `json:"entregaEnvios"`

	Tipo string `json:"tipo"`

	Referencia string `json:"referencia"`

	Responsable Responsable `json:"responsable"`

	Ubicacion Ubicacion `json:"ubicacion"`

	Activo bool `json:"activo"`
}

const PuntoDeTerceroAvroCRC64Fingerprint = "\t\xb8w\x1c+\x8fӸ"

func NewPuntoDeTercero() PuntoDeTercero {
	r := PuntoDeTercero{}
	r.Responsable = NewResponsable()

	r.Ubicacion = NewUbicacion()

	return r
}

func DeserializePuntoDeTercero(r io.Reader) (PuntoDeTercero, error) {
	t := NewPuntoDeTercero()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePuntoDeTerceroFromSchema(r io.Reader, schema string) (PuntoDeTercero, error) {
	t := NewPuntoDeTercero()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePuntoDeTercero(r PuntoDeTercero, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Nombre, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Telefono, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.HorarioDeAtencion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Observaciones, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.AdmiteEnvios, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.EntregaEnvios, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Tipo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Referencia, w)
	if err != nil {
		return err
	}
	err = writeResponsable(r.Responsable, w)
	if err != nil {
		return err
	}
	err = writeUbicacion(r.Ubicacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Activo, w)
	if err != nil {
		return err
	}
	return err
}

func (r PuntoDeTercero) Serialize(w io.Writer) error {
	return writePuntoDeTercero(r, w)
}

func (r PuntoDeTercero) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":\"int\"},{\"name\":\"nombre\",\"type\":\"string\"},{\"name\":\"telefono\",\"type\":\"string\"},{\"name\":\"horarioDeAtencion\",\"type\":\"string\"},{\"name\":\"observaciones\",\"type\":\"string\"},{\"name\":\"admiteEnvios\",\"type\":\"int\"},{\"name\":\"entregaEnvios\",\"type\":\"int\"},{\"name\":\"tipo\",\"type\":\"string\"},{\"name\":\"referencia\",\"type\":\"string\"},{\"name\":\"responsable\",\"type\":{\"fields\":[{\"name\":\"id\",\"type\":\"int\"},{\"name\":\"puntoDeTerceroID\",\"type\":\"int\"},{\"name\":\"nombre\",\"type\":\"string\"},{\"name\":\"apellido\",\"type\":\"string\"},{\"name\":\"mail\",\"type\":\"string\"}],\"name\":\"Responsable\",\"type\":\"record\"}},{\"name\":\"ubicacion\",\"type\":{\"fields\":[{\"name\":\"id\",\"type\":\"int\"},{\"name\":\"puntoDeTerceroID\",\"type\":\"int\"},{\"name\":\"calle\",\"type\":\"string\"},{\"name\":\"numero\",\"type\":\"string\"},{\"name\":\"piso\",\"type\":\"string\"},{\"name\":\"departamento\",\"type\":\"string\"},{\"name\":\"codigoPostal\",\"type\":\"string\"},{\"name\":\"latitud\",\"type\":\"string\"},{\"name\":\"longitud\",\"type\":\"string\"}],\"name\":\"Ubicacion\",\"type\":\"record\"}},{\"name\":\"activo\",\"type\":\"boolean\"}],\"name\":\"Andreani.PuntoDeTercero.Events.Record.PuntoDeTercero\",\"type\":\"record\"}"
}

func (r PuntoDeTercero) SchemaName() string {
	return "Andreani.PuntoDeTercero.Events.Record.PuntoDeTercero"
}

func (_ PuntoDeTercero) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ PuntoDeTercero) SetInt(v int32)       { panic("Unsupported operation") }
func (_ PuntoDeTercero) SetLong(v int64)      { panic("Unsupported operation") }
func (_ PuntoDeTercero) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ PuntoDeTercero) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ PuntoDeTercero) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ PuntoDeTercero) SetString(v string)   { panic("Unsupported operation") }
func (_ PuntoDeTercero) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PuntoDeTercero) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.Nombre}

		return w

	case 2:
		w := types.String{Target: &r.Telefono}

		return w

	case 3:
		w := types.String{Target: &r.HorarioDeAtencion}

		return w

	case 4:
		w := types.String{Target: &r.Observaciones}

		return w

	case 5:
		w := types.Int{Target: &r.AdmiteEnvios}

		return w

	case 6:
		w := types.Int{Target: &r.EntregaEnvios}

		return w

	case 7:
		w := types.String{Target: &r.Tipo}

		return w

	case 8:
		w := types.String{Target: &r.Referencia}

		return w

	case 9:
		r.Responsable = NewResponsable()

		w := types.Record{Target: &r.Responsable}

		return w

	case 10:
		r.Ubicacion = NewUbicacion()

		w := types.Record{Target: &r.Ubicacion}

		return w

	case 11:
		w := types.Boolean{Target: &r.Activo}

		return w

	}
	panic("Unknown field index")
}

func (r *PuntoDeTercero) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *PuntoDeTercero) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ PuntoDeTercero) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ PuntoDeTercero) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ PuntoDeTercero) HintSize(int)                     { panic("Unsupported operation") }
func (_ PuntoDeTercero) Finalize()                        {}

func (_ PuntoDeTercero) AvroCRC64Fingerprint() []byte {
	return []byte(PuntoDeTerceroAvroCRC64Fingerprint)
}

func (r PuntoDeTercero) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["nombre"], err = json.Marshal(r.Nombre)
	if err != nil {
		return nil, err
	}
	output["telefono"], err = json.Marshal(r.Telefono)
	if err != nil {
		return nil, err
	}
	output["horarioDeAtencion"], err = json.Marshal(r.HorarioDeAtencion)
	if err != nil {
		return nil, err
	}
	output["observaciones"], err = json.Marshal(r.Observaciones)
	if err != nil {
		return nil, err
	}
	output["admiteEnvios"], err = json.Marshal(r.AdmiteEnvios)
	if err != nil {
		return nil, err
	}
	output["entregaEnvios"], err = json.Marshal(r.EntregaEnvios)
	if err != nil {
		return nil, err
	}
	output["tipo"], err = json.Marshal(r.Tipo)
	if err != nil {
		return nil, err
	}
	output["referencia"], err = json.Marshal(r.Referencia)
	if err != nil {
		return nil, err
	}
	output["responsable"], err = json.Marshal(r.Responsable)
	if err != nil {
		return nil, err
	}
	output["ubicacion"], err = json.Marshal(r.Ubicacion)
	if err != nil {
		return nil, err
	}
	output["activo"], err = json.Marshal(r.Activo)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *PuntoDeTercero) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["telefono"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Telefono); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for telefono")
	}
	val = func() json.RawMessage {
		if v, ok := fields["horarioDeAtencion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.HorarioDeAtencion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for horarioDeAtencion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["observaciones"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Observaciones); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for observaciones")
	}
	val = func() json.RawMessage {
		if v, ok := fields["admiteEnvios"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AdmiteEnvios); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for admiteEnvios")
	}
	val = func() json.RawMessage {
		if v, ok := fields["entregaEnvios"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EntregaEnvios); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for entregaEnvios")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tipo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["referencia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Referencia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for referencia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["responsable"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Responsable); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for responsable")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ubicacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ubicacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ubicacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["activo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Activo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for activo")
	}
	return nil
}
