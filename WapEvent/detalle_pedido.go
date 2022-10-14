// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoDeAlmacenSCE.avsc
 */
package WapEventEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type DetallePedido struct {
	Articulo Articulo `json:"articulo"`

	Numerodelinea *UnionNullString `json:"numerodelinea"`

	Tipoacondicionamientoescundario *UnionNullString `json:"tipoacondicionamientoescundario"`

	Admitepickingparcial *UnionNullString `json:"admitepickingparcial"`
}

const DetallePedidoAvroCRC64Fingerprint = "\x8e[)\xd8\x121\x0e\xb1"

func NewDetallePedido() DetallePedido {
	r := DetallePedido{}
	r.Articulo = NewArticulo()

	r.Admitepickingparcial = nil
	return r
}

func DeserializeDetallePedido(r io.Reader) (DetallePedido, error) {
	t := NewDetallePedido()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDetallePedidoFromSchema(r io.Reader, schema string) (DetallePedido, error) {
	t := NewDetallePedido()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDetallePedido(r DetallePedido, w io.Writer) error {
	var err error
	err = writeArticulo(r.Articulo, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Numerodelinea, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Tipoacondicionamientoescundario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Admitepickingparcial, w)
	if err != nil {
		return err
	}
	return err
}

func (r DetallePedido) Serialize(w io.Writer) error {
	return writeDetallePedido(r, w)
}

func (r DetallePedido) Schema() string {
	return "{\"fields\":[{\"name\":\"articulo\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"cantidad\",\"type\":\"double\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"name\":\"numeropedido\",\"type\":[\"null\",\"string\"]},{\"name\":\"unidadmedida\",\"type\":\"string\"},{\"name\":\"lineaexterna\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"datosadicionales\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"name\":\"lote\",\"type\":{\"fields\":[{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"string\"]}],\"name\":\"LoteArticulo\",\"type\":\"record\"}}],\"name\":\"Articulo\",\"type\":\"record\"}},{\"name\":\"numerodelinea\",\"type\":[\"null\",\"string\"]},{\"name\":\"tipoacondicionamientoescundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"admitepickingparcial\",\"type\":[\"null\",\"string\"]}],\"name\":\"Wap.Events.Record.DetallePedido\",\"type\":\"record\"}"
}

func (r DetallePedido) SchemaName() string {
	return "Wap.Events.Record.DetallePedido"
}

func (_ DetallePedido) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DetallePedido) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DetallePedido) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DetallePedido) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DetallePedido) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DetallePedido) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DetallePedido) SetString(v string)   { panic("Unsupported operation") }
func (_ DetallePedido) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DetallePedido) Get(i int) types.Field {
	switch i {
	case 0:
		r.Articulo = NewArticulo()

		w := types.Record{Target: &r.Articulo}

		return w

	case 1:
		r.Numerodelinea = NewUnionNullString()

		return r.Numerodelinea
	case 2:
		r.Tipoacondicionamientoescundario = NewUnionNullString()

		return r.Tipoacondicionamientoescundario
	case 3:
		r.Admitepickingparcial = NewUnionNullString()

		return r.Admitepickingparcial
	}
	panic("Unknown field index")
}

func (r *DetallePedido) SetDefault(i int) {
	switch i {
	case 3:
		r.Admitepickingparcial = nil
		return
	}
	panic("Unknown field index")
}

func (r *DetallePedido) NullField(i int) {
	switch i {
	case 1:
		r.Numerodelinea = nil
		return
	case 2:
		r.Tipoacondicionamientoescundario = nil
		return
	case 3:
		r.Admitepickingparcial = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ DetallePedido) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DetallePedido) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DetallePedido) HintSize(int)                     { panic("Unsupported operation") }
func (_ DetallePedido) Finalize()                        {}

func (_ DetallePedido) AvroCRC64Fingerprint() []byte {
	return []byte(DetallePedidoAvroCRC64Fingerprint)
}

func (r DetallePedido) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["articulo"], err = json.Marshal(r.Articulo)
	if err != nil {
		return nil, err
	}
	output["numerodelinea"], err = json.Marshal(r.Numerodelinea)
	if err != nil {
		return nil, err
	}
	output["tipoacondicionamientoescundario"], err = json.Marshal(r.Tipoacondicionamientoescundario)
	if err != nil {
		return nil, err
	}
	output["admitepickingparcial"], err = json.Marshal(r.Admitepickingparcial)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DetallePedido) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["articulo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Articulo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for articulo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numerodelinea"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Numerodelinea); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numerodelinea")
	}
	val = func() json.RawMessage {
		if v, ok := fields["tipoacondicionamientoescundario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Tipoacondicionamientoescundario); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for tipoacondicionamientoescundario")
	}
	val = func() json.RawMessage {
		if v, ok := fields["admitepickingparcial"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Admitepickingparcial); err != nil {
			return err
		}
	} else {
		r.Admitepickingparcial = NewUnionNullString()

		r.Admitepickingparcial = nil
	}
	return nil
}
