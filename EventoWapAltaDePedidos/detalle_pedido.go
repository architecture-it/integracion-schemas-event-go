// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AltaDePedidosSolicitada.avsc
 */
package avro

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

	NumeroPedido string `json:"numeroPedido"`

	UnidadMedida string `json:"unidadMedida"`

	LineaExterna *UnionNullString `json:"lineaExterna"`

	Unidades float64 `json:"unidades"`

	OtrosDatos *UnionNullListaDePropiedades `json:"otrosDatos"`
}

const DetallePedidoAvroCRC64Fingerprint = "\xd0[\xfd\xba \x06\xdcN"

func NewDetallePedido() DetallePedido {
	r := DetallePedido{}
	r.Articulo = NewArticulo()

	r.LineaExterna = nil
	r.OtrosDatos = nil
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
	err = vm.WriteString(r.NumeroPedido, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.UnidadMedida, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LineaExterna, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Unidades, w)
	if err != nil {
		return err
	}
	err = writeUnionNullListaDePropiedades(r.OtrosDatos, w)
	if err != nil {
		return err
	}
	return err
}

func (r DetallePedido) Serialize(w io.Writer) error {
	return writeDetallePedido(r, w)
}

func (r DetallePedido) Schema() string {
	return "{\"fields\":[{\"name\":\"articulo\",\"type\":{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"name\":\"propietario\",\"type\":\"string\"},{\"default\":null,\"name\":\"lote\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"name\":\"fechaDeVencimiento\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"namespace\":\"Andreani.Wap.Events.Record\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"namespace\":\"Andreani.Wap.Events.Record\",\"type\":\"record\"}]}],\"name\":\"LoteArticulo\",\"namespace\":\"Andreani.Wap.Events.Record\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"Andreani.Wap.Events.Record.ListaDePropiedades\"]}],\"name\":\"Articulo\",\"namespace\":\"Andreani.Wap.Events.Record\",\"type\":\"record\"}},{\"name\":\"numeroPedido\",\"type\":\"string\"},{\"name\":\"unidadMedida\",\"type\":\"string\"},{\"default\":null,\"name\":\"lineaExterna\",\"type\":[\"null\",\"string\"]},{\"name\":\"unidades\",\"type\":\"double\"},{\"default\":null,\"name\":\"otrosDatos\",\"type\":[\"null\",\"Andreani.Wap.Events.Record.ListaDePropiedades\"]}],\"name\":\"Andreani.Wap.Events.Record.DetallePedido\",\"type\":\"record\"}"
}

func (r DetallePedido) SchemaName() string {
	return "Andreani.Wap.Events.Record.DetallePedido"
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
		w := types.String{Target: &r.NumeroPedido}

		return w

	case 2:
		w := types.String{Target: &r.UnidadMedida}

		return w

	case 3:
		r.LineaExterna = NewUnionNullString()

		return r.LineaExterna
	case 4:
		w := types.Double{Target: &r.Unidades}

		return w

	case 5:
		r.OtrosDatos = NewUnionNullListaDePropiedades()

		return r.OtrosDatos
	}
	panic("Unknown field index")
}

func (r *DetallePedido) SetDefault(i int) {
	switch i {
	case 3:
		r.LineaExterna = nil
		return
	case 5:
		r.OtrosDatos = nil
		return
	}
	panic("Unknown field index")
}

func (r *DetallePedido) NullField(i int) {
	switch i {
	case 3:
		r.LineaExterna = nil
		return
	case 5:
		r.OtrosDatos = nil
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
	output["numeroPedido"], err = json.Marshal(r.NumeroPedido)
	if err != nil {
		return nil, err
	}
	output["unidadMedida"], err = json.Marshal(r.UnidadMedida)
	if err != nil {
		return nil, err
	}
	output["lineaExterna"], err = json.Marshal(r.LineaExterna)
	if err != nil {
		return nil, err
	}
	output["unidades"], err = json.Marshal(r.Unidades)
	if err != nil {
		return nil, err
	}
	output["otrosDatos"], err = json.Marshal(r.OtrosDatos)
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
		if v, ok := fields["numeroPedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroPedido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroPedido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["unidadMedida"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.UnidadMedida); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for unidadMedida")
	}
	val = func() json.RawMessage {
		if v, ok := fields["lineaExterna"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LineaExterna); err != nil {
			return err
		}
	} else {
		r.LineaExterna = NewUnionNullString()

		r.LineaExterna = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["unidades"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Unidades); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for unidades")
	}
	val = func() json.RawMessage {
		if v, ok := fields["otrosDatos"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OtrosDatos); err != nil {
			return err
		}
	} else {
		r.OtrosDatos = NewUnionNullListaDePropiedades()

		r.OtrosDatos = nil
	}
	return nil
}
