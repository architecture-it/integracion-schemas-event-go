// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ContenedorExpedido.avsc
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

type ContenedorExpedido struct {
	Traza TrazaContendor `json:"traza"`

	Origen DatosSucursal `json:"origen"`

	Destino DatosSucursal `json:"destino"`

	Viaje Viaje `json:"viaje"`
}

const ContenedorExpedidoAvroCRC64Fingerprint = "d\xaa\x03B7)\x1b\x11"

func NewContenedorExpedido() ContenedorExpedido {
	r := ContenedorExpedido{}
	r.Traza = NewTrazaContendor()

	r.Origen = NewDatosSucursal()

	r.Destino = NewDatosSucursal()

	r.Viaje = NewViaje()

	return r
}

func DeserializeContenedorExpedido(r io.Reader) (ContenedorExpedido, error) {
	t := NewContenedorExpedido()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeContenedorExpedidoFromSchema(r io.Reader, schema string) (ContenedorExpedido, error) {
	t := NewContenedorExpedido()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeContenedorExpedido(r ContenedorExpedido, w io.Writer) error {
	var err error
	err = writeTrazaContendor(r.Traza, w)
	if err != nil {
		return err
	}
	err = writeDatosSucursal(r.Origen, w)
	if err != nil {
		return err
	}
	err = writeDatosSucursal(r.Destino, w)
	if err != nil {
		return err
	}
	err = writeViaje(r.Viaje, w)
	if err != nil {
		return err
	}
	return err
}

func (r ContenedorExpedido) Serialize(w io.Writer) error {
	return writeContenedorExpedido(r, w)
}

func (r ContenedorExpedido) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"numero\",\"type\":\"string\"},{\"name\":\"tipo\",\"type\":\"string\"},{\"name\":\"ciclo\",\"type\":\"string\"},{\"name\":\"estado\",\"type\":\"string\"},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"TrazaContendor\",\"namespace\":\"Integracion.Esquemas.Contenedor.Referencias\",\"type\":\"record\"}},{\"name\":\"origen\",\"type\":\"Integracion.Esquemas.Referencias.DatosSucursal\"},{\"name\":\"destino\",\"type\":\"Integracion.Esquemas.Referencias.DatosSucursal\"},{\"name\":\"viaje\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"}],\"name\":\"Viaje\",\"namespace\":\"Integracion.Esquemas.Contenedor.Referencias\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Contenedor.Trazas.ContenedorExpedido\",\"type\":\"record\"}"
}

func (r ContenedorExpedido) SchemaName() string {
	return "Integracion.Esquemas.Contenedor.Trazas.ContenedorExpedido"
}

func (_ ContenedorExpedido) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ContenedorExpedido) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ContenedorExpedido) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ContenedorExpedido) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ContenedorExpedido) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ContenedorExpedido) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ContenedorExpedido) SetString(v string)   { panic("Unsupported operation") }
func (_ ContenedorExpedido) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ContenedorExpedido) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTrazaContendor()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		r.Origen = NewDatosSucursal()

		w := types.Record{Target: &r.Origen}

		return w

	case 2:
		r.Destino = NewDatosSucursal()

		w := types.Record{Target: &r.Destino}

		return w

	case 3:
		r.Viaje = NewViaje()

		w := types.Record{Target: &r.Viaje}

		return w

	}
	panic("Unknown field index")
}

func (r *ContenedorExpedido) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ContenedorExpedido) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ContenedorExpedido) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ContenedorExpedido) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ContenedorExpedido) HintSize(int)                     { panic("Unsupported operation") }
func (_ ContenedorExpedido) Finalize()                        {}

func (_ ContenedorExpedido) AvroCRC64Fingerprint() []byte {
	return []byte(ContenedorExpedidoAvroCRC64Fingerprint)
}

func (r ContenedorExpedido) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["origen"], err = json.Marshal(r.Origen)
	if err != nil {
		return nil, err
	}
	output["destino"], err = json.Marshal(r.Destino)
	if err != nil {
		return nil, err
	}
	output["viaje"], err = json.Marshal(r.Viaje)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ContenedorExpedido) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["origen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Origen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for origen")
	}
	val = func() json.RawMessage {
		if v, ok := fields["destino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Destino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for destino")
	}
	val = func() json.RawMessage {
		if v, ok := fields["viaje"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Viaje); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for viaje")
	}
	return nil
}
