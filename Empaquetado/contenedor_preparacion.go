// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoEmpaquetar.avsc
 */
package EmpaquetadoEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ContenedorPreparacion struct {
	Codigo string `json:"Codigo"`

	Articulos []ArticuloContenedorPreparacion `json:"Articulos"`
}

const ContenedorPreparacionAvroCRC64Fingerprint = "\x98\v\xf1HYF\x96H"

func NewContenedorPreparacion() ContenedorPreparacion {
	r := ContenedorPreparacion{}
	r.Articulos = make([]ArticuloContenedorPreparacion, 0)

	return r
}

func DeserializeContenedorPreparacion(r io.Reader) (ContenedorPreparacion, error) {
	t := NewContenedorPreparacion()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeContenedorPreparacionFromSchema(r io.Reader, schema string) (ContenedorPreparacion, error) {
	t := NewContenedorPreparacion()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeContenedorPreparacion(r ContenedorPreparacion, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Codigo, w)
	if err != nil {
		return err
	}
	err = writeArrayArticuloContenedorPreparacion(r.Articulos, w)
	if err != nil {
		return err
	}
	return err
}

func (r ContenedorPreparacion) Serialize(w io.Writer) error {
	return writeContenedorPreparacion(r, w)
}

func (r ContenedorPreparacion) Schema() string {
	return "{\"fields\":[{\"name\":\"Codigo\",\"type\":\"string\"},{\"name\":\"Articulos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Sku\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"default\":null,\"name\":\"Ean\",\"type\":[\"null\",\"string\"]},{\"name\":\"NroLineaPedido\",\"type\":\"string\"},{\"name\":\"CantidadPedido\",\"type\":\"int\"},{\"name\":\"CantidadPickeada\",\"type\":\"int\"},{\"default\":null,\"name\":\"Lote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Serie\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Zona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoZona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionZona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Longitud\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Altura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Ancho\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Peso\",\"type\":[\"null\",\"float\"]}],\"name\":\"ArticuloContenedorPreparacion\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Empaquetado.Events.Common.ContenedorPreparacion\",\"type\":\"record\"}"
}

func (r ContenedorPreparacion) SchemaName() string {
	return "Andreani.Empaquetado.Events.Common.ContenedorPreparacion"
}

func (_ ContenedorPreparacion) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ContenedorPreparacion) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ContenedorPreparacion) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ContenedorPreparacion) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ContenedorPreparacion) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ContenedorPreparacion) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ContenedorPreparacion) SetString(v string)   { panic("Unsupported operation") }
func (_ ContenedorPreparacion) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ContenedorPreparacion) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Codigo}

		return w

	case 1:
		r.Articulos = make([]ArticuloContenedorPreparacion, 0)

		w := ArrayArticuloContenedorPreparacionWrapper{Target: &r.Articulos}

		return w

	}
	panic("Unknown field index")
}

func (r *ContenedorPreparacion) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ContenedorPreparacion) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ContenedorPreparacion) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ContenedorPreparacion) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ContenedorPreparacion) HintSize(int)                     { panic("Unsupported operation") }
func (_ ContenedorPreparacion) Finalize()                        {}

func (_ ContenedorPreparacion) AvroCRC64Fingerprint() []byte {
	return []byte(ContenedorPreparacionAvroCRC64Fingerprint)
}

func (r ContenedorPreparacion) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Codigo"], err = json.Marshal(r.Codigo)
	if err != nil {
		return nil, err
	}
	output["Articulos"], err = json.Marshal(r.Articulos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ContenedorPreparacion) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Codigo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Codigo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Codigo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Articulos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Articulos); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Articulos")
	}
	return nil
}
