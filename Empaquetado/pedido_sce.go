// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     PedidoSce.avsc
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

type PedidoSce struct {
	Identificacion Identificacion `json:"Identificacion"`

	OrdenWh string `json:"OrdenWh"`

	OrdenCliente string `json:"OrdenCliente"`

	BultosSce []BultoSce `json:"BultosSce"`
}

const PedidoSceAvroCRC64Fingerprint = "\x1d˓c\xf4\x05@\xfa"

func NewPedidoSce() PedidoSce {
	r := PedidoSce{}
	r.Identificacion = NewIdentificacion()

	r.BultosSce = make([]BultoSce, 0)

	return r
}

func DeserializePedidoSce(r io.Reader) (PedidoSce, error) {
	t := NewPedidoSce()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePedidoSceFromSchema(r io.Reader, schema string) (PedidoSce, error) {
	t := NewPedidoSce()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePedidoSce(r PedidoSce, w io.Writer) error {
	var err error
	err = writeIdentificacion(r.Identificacion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.OrdenWh, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.OrdenCliente, w)
	if err != nil {
		return err
	}
	err = writeArrayBultoSce(r.BultosSce, w)
	if err != nil {
		return err
	}
	return err
}

func (r PedidoSce) Serialize(w io.Writer) error {
	return writePedidoSce(r, w)
}

func (r PedidoSce) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"},{\"default\":null,\"name\":\"PlantaOperacionId\",\"type\":[\"null\",\"int\"]}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"}},{\"name\":\"OrdenWh\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"name\":\"BultosSce\",\"type\":{\"items\":{\"fields\":[{\"name\":\"CodigoEmbalaje\",\"type\":\"string\"},{\"name\":\"TipoContenedorEmbalaje\",\"type\":\"string\"},{\"name\":\"ContenedorPreparacion\",\"type\":\"string\"},{\"default\":null,\"name\":\"Longitud\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Altura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Ancho\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Peso\",\"type\":[\"null\",\"float\"]},{\"name\":\"ArticulosSce\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Sku\",\"type\":\"string\"},{\"name\":\"Cantidad\",\"type\":\"int\"},{\"name\":\"NroLinea\",\"type\":\"string\"},{\"name\":\"PickDetailKey\",\"type\":\"string\"},{\"name\":\"Lote\",\"type\":\"string\"},{\"default\":null,\"name\":\"ContenedorPreparacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo3\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo4\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo5\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo6\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo7\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo8\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo9\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo10\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo11\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo12\",\"type\":[\"null\",\"string\"]},{\"name\":\"InformaSeriesEnEmpaquetado\",\"type\":\"boolean\"},{\"default\":null,\"name\":\"Series\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]}],\"name\":\"ArticuloSce\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"BultoSce\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Empaquetado.Events.Record.PedidoSce\",\"type\":\"record\"}"
}

func (r PedidoSce) SchemaName() string {
	return "Andreani.Empaquetado.Events.Record.PedidoSce"
}

func (_ PedidoSce) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ PedidoSce) SetInt(v int32)       { panic("Unsupported operation") }
func (_ PedidoSce) SetLong(v int64)      { panic("Unsupported operation") }
func (_ PedidoSce) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ PedidoSce) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ PedidoSce) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ PedidoSce) SetString(v string)   { panic("Unsupported operation") }
func (_ PedidoSce) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PedidoSce) Get(i int) types.Field {
	switch i {
	case 0:
		r.Identificacion = NewIdentificacion()

		w := types.Record{Target: &r.Identificacion}

		return w

	case 1:
		w := types.String{Target: &r.OrdenWh}

		return w

	case 2:
		w := types.String{Target: &r.OrdenCliente}

		return w

	case 3:
		r.BultosSce = make([]BultoSce, 0)

		w := ArrayBultoSceWrapper{Target: &r.BultosSce}

		return w

	}
	panic("Unknown field index")
}

func (r *PedidoSce) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *PedidoSce) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ PedidoSce) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ PedidoSce) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ PedidoSce) HintSize(int)                     { panic("Unsupported operation") }
func (_ PedidoSce) Finalize()                        {}

func (_ PedidoSce) AvroCRC64Fingerprint() []byte {
	return []byte(PedidoSceAvroCRC64Fingerprint)
}

func (r PedidoSce) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Identificacion"], err = json.Marshal(r.Identificacion)
	if err != nil {
		return nil, err
	}
	output["OrdenWh"], err = json.Marshal(r.OrdenWh)
	if err != nil {
		return nil, err
	}
	output["OrdenCliente"], err = json.Marshal(r.OrdenCliente)
	if err != nil {
		return nil, err
	}
	output["BultosSce"], err = json.Marshal(r.BultosSce)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *PedidoSce) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Identificacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Identificacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Identificacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrdenWh"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenWh); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrdenWh")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OrdenCliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OrdenCliente); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OrdenCliente")
	}
	val = func() json.RawMessage {
		if v, ok := fields["BultosSce"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BultosSce); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for BultosSce")
	}
	return nil
}
