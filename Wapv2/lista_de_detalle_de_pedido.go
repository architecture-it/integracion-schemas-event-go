// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
 */
package Wapv2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ListaDeDetalleDePedido struct {
	DetallePedido []DetallePedido `json:"detallePedido"`
}

const ListaDeDetalleDePedidoAvroCRC64Fingerprint = "M\x03X\x9f\x81@)\xe3"

func NewListaDeDetalleDePedido() ListaDeDetalleDePedido {
	r := ListaDeDetalleDePedido{}
	r.DetallePedido = make([]DetallePedido, 0)

	return r
}

func DeserializeListaDeDetalleDePedido(r io.Reader) (ListaDeDetalleDePedido, error) {
	t := NewListaDeDetalleDePedido()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeListaDeDetalleDePedidoFromSchema(r io.Reader, schema string) (ListaDeDetalleDePedido, error) {
	t := NewListaDeDetalleDePedido()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeListaDeDetalleDePedido(r ListaDeDetalleDePedido, w io.Writer) error {
	var err error
	err = writeArrayDetallePedido(r.DetallePedido, w)
	if err != nil {
		return err
	}
	return err
}

func (r ListaDeDetalleDePedido) Serialize(w io.Writer) error {
	return writeListaDeDetalleDePedido(r, w)
}

func (r ListaDeDetalleDePedido) Schema() string {
	return "{\"fields\":[{\"name\":\"detallePedido\",\"type\":{\"items\":{\"fields\":[{\"name\":\"articulo\",\"type\":{\"fields\":[{\"name\":\"contratoWarehouse\",\"type\":\"string\"},{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"cantidad\",\"type\":\"double\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"numeropedido\",\"type\":[\"null\",\"string\"]},{\"name\":\"zonaConsumo\",\"type\":[\"null\",\"string\"]},{\"name\":\"unidadmedida\",\"type\":\"string\"},{\"default\":null,\"name\":\"datosadicionales\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"name\":\"lote\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"string\"]},{\"name\":\"estado\",\"type\":[\"null\",\"string\"]}],\"name\":\"LoteArticulo\",\"type\":\"record\"}}],\"name\":\"Articulo\",\"type\":\"record\"}},{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]},{\"name\":\"numerodelinea\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoacondicionamientoescundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"admitepickingparcial\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estatusOTdeAcondi\",\"type\":[\"null\",\"string\"]},{\"name\":\"diasVencimientoMinimo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tiposDeAcondicionamientoSecundario\",\"type\":[\"null\",{\"fields\":[{\"name\":\"tiposDeAcondicionamientoSecundario\",\"type\":{\"items\":{\"fields\":[{\"name\":\"tipoDeAcondi\",\"type\":\"int\"}],\"name\":\"TiposDeAcondicionamientoSecundario\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ListaDeTiposDeAcondicionamientoSecundario\",\"type\":\"record\"}]}],\"name\":\"DetallePedido\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Wapv2.Events.Record.ListaDeDetalleDePedido\",\"type\":\"record\"}"
}

func (r ListaDeDetalleDePedido) SchemaName() string {
	return "Andreani.Wapv2.Events.Record.ListaDeDetalleDePedido"
}

func (_ ListaDeDetalleDePedido) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ListaDeDetalleDePedido) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ListaDeDetalleDePedido) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ListaDeDetalleDePedido) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ListaDeDetalleDePedido) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ListaDeDetalleDePedido) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ListaDeDetalleDePedido) SetString(v string)   { panic("Unsupported operation") }
func (_ ListaDeDetalleDePedido) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ListaDeDetalleDePedido) Get(i int) types.Field {
	switch i {
	case 0:
		r.DetallePedido = make([]DetallePedido, 0)

		w := ArrayDetallePedidoWrapper{Target: &r.DetallePedido}

		return w

	}
	panic("Unknown field index")
}

func (r *ListaDeDetalleDePedido) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ListaDeDetalleDePedido) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ListaDeDetalleDePedido) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ListaDeDetalleDePedido) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ListaDeDetalleDePedido) HintSize(int)                     { panic("Unsupported operation") }
func (_ ListaDeDetalleDePedido) Finalize()                        {}

func (_ ListaDeDetalleDePedido) AvroCRC64Fingerprint() []byte {
	return []byte(ListaDeDetalleDePedidoAvroCRC64Fingerprint)
}

func (r ListaDeDetalleDePedido) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["detallePedido"], err = json.Marshal(r.DetallePedido)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ListaDeDetalleDePedido) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["detallePedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DetallePedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for detallePedido")
	}
	return nil
}
