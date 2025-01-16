// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EventoDeRespuesta.avsc
 */
package TrazasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EventoDeRespuesta struct {
	Traza Traza `json:"traza"`

	EsRespuestaDe string `json:"esRespuestaDe"`

	Motivo string `json:"motivo"`

	Submotivo string `json:"submotivo"`
}

const EventoDeRespuestaAvroCRC64Fingerprint = "\xbd!\xc5\x00\xbc\x0fn\r"

func NewEventoDeRespuesta() EventoDeRespuesta {
	r := EventoDeRespuesta{}
	r.Traza = NewTraza()

	return r
}

func DeserializeEventoDeRespuesta(r io.Reader) (EventoDeRespuesta, error) {
	t := NewEventoDeRespuesta()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEventoDeRespuestaFromSchema(r io.Reader, schema string) (EventoDeRespuesta, error) {
	t := NewEventoDeRespuesta()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEventoDeRespuesta(r EventoDeRespuesta, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EsRespuestaDe, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Motivo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Submotivo, w)
	if err != nil {
		return err
	}
	return err
}

func (r EventoDeRespuesta) Serialize(w io.Writer) error {
	return writeEventoDeRespuesta(r, w)
}

func (r EventoDeRespuesta) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"esRespuestaDe\",\"type\":\"string\"},{\"name\":\"motivo\",\"type\":\"string\"},{\"name\":\"submotivo\",\"type\":\"string\"}],\"name\":\"Integracion.Esquemas.Incidencias.EventoDeRespuesta\",\"type\":\"record\"}"
}

func (r EventoDeRespuesta) SchemaName() string {
	return "Integracion.Esquemas.Incidencias.EventoDeRespuesta"
}

func (_ EventoDeRespuesta) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EventoDeRespuesta) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EventoDeRespuesta) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EventoDeRespuesta) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EventoDeRespuesta) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EventoDeRespuesta) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EventoDeRespuesta) SetString(v string)   { panic("Unsupported operation") }
func (_ EventoDeRespuesta) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EventoDeRespuesta) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		w := types.String{Target: &r.EsRespuestaDe}

		return w

	case 2:
		w := types.String{Target: &r.Motivo}

		return w

	case 3:
		w := types.String{Target: &r.Submotivo}

		return w

	}
	panic("Unknown field index")
}

func (r *EventoDeRespuesta) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *EventoDeRespuesta) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EventoDeRespuesta) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EventoDeRespuesta) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EventoDeRespuesta) HintSize(int)                     { panic("Unsupported operation") }
func (_ EventoDeRespuesta) Finalize()                        {}

func (_ EventoDeRespuesta) AvroCRC64Fingerprint() []byte {
	return []byte(EventoDeRespuestaAvroCRC64Fingerprint)
}

func (r EventoDeRespuesta) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["esRespuestaDe"], err = json.Marshal(r.EsRespuestaDe)
	if err != nil {
		return nil, err
	}
	output["motivo"], err = json.Marshal(r.Motivo)
	if err != nil {
		return nil, err
	}
	output["submotivo"], err = json.Marshal(r.Submotivo)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EventoDeRespuesta) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["traza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Traza); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for traza")
	}
	val = func() json.RawMessage {
		if v, ok := fields["esRespuestaDe"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EsRespuestaDe); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for esRespuestaDe")
	}
	val = func() json.RawMessage {
		if v, ok := fields["motivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Motivo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for motivo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["submotivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Submotivo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for submotivo")
	}
	return nil
}
