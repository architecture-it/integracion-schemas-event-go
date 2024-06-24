// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CalculoTiempoEmbalaje.avsc
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

type CalculoTiempoEmbalaje struct {
	Identificacion Identificacion `json:"Identificacion"`

	OrdenWh string `json:"OrdenWh"`

	OrdenCliente string `json:"OrdenCliente"`

	Articulos []ArticuloContenedorPreparacion `json:"Articulos"`
}

const CalculoTiempoEmbalajeAvroCRC64Fingerprint = "7\xc6\xd6>\xfa\xc8\xd6\xfd"

func NewCalculoTiempoEmbalaje() CalculoTiempoEmbalaje {
	r := CalculoTiempoEmbalaje{}
	r.Identificacion = NewIdentificacion()

	r.Articulos = make([]ArticuloContenedorPreparacion, 0)

	return r
}

func DeserializeCalculoTiempoEmbalaje(r io.Reader) (CalculoTiempoEmbalaje, error) {
	t := NewCalculoTiempoEmbalaje()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeCalculoTiempoEmbalajeFromSchema(r io.Reader, schema string) (CalculoTiempoEmbalaje, error) {
	t := NewCalculoTiempoEmbalaje()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeCalculoTiempoEmbalaje(r CalculoTiempoEmbalaje, w io.Writer) error {
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
	err = writeArrayArticuloContenedorPreparacion(r.Articulos, w)
	if err != nil {
		return err
	}
	return err
}

func (r CalculoTiempoEmbalaje) Serialize(w io.Writer) error {
	return writeCalculoTiempoEmbalaje(r, w)
}

func (r CalculoTiempoEmbalaje) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"},{\"default\":null,\"name\":\"PlantaOperacionId\",\"type\":[\"null\",\"int\"]}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"}},{\"name\":\"OrdenWh\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"name\":\"Articulos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Sku\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"default\":null,\"name\":\"Ean\",\"type\":[\"null\",\"string\"]},{\"name\":\"NroLineaPedido\",\"type\":\"string\"},{\"name\":\"CantidadPedido\",\"type\":\"int\"},{\"name\":\"CantidadPickeada\",\"type\":\"int\"},{\"default\":null,\"name\":\"Lote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Serie\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Zona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoZona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionZona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Longitud\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Altura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Ancho\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Peso\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"InstruccionesEmbalaje\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UDM\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PickearTodos\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"ValidaLoteWos\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ValidaSerieWos\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FormatoSerie\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ValidaSerieEnSalida\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"IngresaSerieEnEmpaquetado\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"SerieDirigida\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PickDetailKey\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoValidacionSalidaLote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UsaDataMatrix\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"UbicacionOrigen\",\"type\":[\"null\",\"string\"]}],\"name\":\"ArticuloContenedorPreparacion\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Empaquetado.Events.Record.CalculoTiempoEmbalaje\",\"type\":\"record\"}"
}

func (r CalculoTiempoEmbalaje) SchemaName() string {
	return "Andreani.Empaquetado.Events.Record.CalculoTiempoEmbalaje"
}

func (_ CalculoTiempoEmbalaje) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ CalculoTiempoEmbalaje) SetInt(v int32)       { panic("Unsupported operation") }
func (_ CalculoTiempoEmbalaje) SetLong(v int64)      { panic("Unsupported operation") }
func (_ CalculoTiempoEmbalaje) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ CalculoTiempoEmbalaje) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ CalculoTiempoEmbalaje) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ CalculoTiempoEmbalaje) SetString(v string)   { panic("Unsupported operation") }
func (_ CalculoTiempoEmbalaje) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *CalculoTiempoEmbalaje) Get(i int) types.Field {
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
		r.Articulos = make([]ArticuloContenedorPreparacion, 0)

		w := ArrayArticuloContenedorPreparacionWrapper{Target: &r.Articulos}

		return w

	}
	panic("Unknown field index")
}

func (r *CalculoTiempoEmbalaje) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *CalculoTiempoEmbalaje) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ CalculoTiempoEmbalaje) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ CalculoTiempoEmbalaje) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ CalculoTiempoEmbalaje) HintSize(int)                     { panic("Unsupported operation") }
func (_ CalculoTiempoEmbalaje) Finalize()                        {}

func (_ CalculoTiempoEmbalaje) AvroCRC64Fingerprint() []byte {
	return []byte(CalculoTiempoEmbalajeAvroCRC64Fingerprint)
}

func (r CalculoTiempoEmbalaje) MarshalJSON() ([]byte, error) {
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
	output["Articulos"], err = json.Marshal(r.Articulos)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *CalculoTiempoEmbalaje) UnmarshalJSON(data []byte) error {
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
