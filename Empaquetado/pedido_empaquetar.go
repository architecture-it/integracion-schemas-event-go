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

type PedidoEmpaquetar struct {
	Identificacion Identificacion `json:"Identificacion"`

	OrdenWh string `json:"OrdenWh"`

	OrdenCliente string `json:"OrdenCliente"`

	Mesa *UnionNullString `json:"Mesa"`

	Calle *UnionNullString `json:"Calle"`

	FechaExpedicion *UnionNullLong `json:"FechaExpedicion"`

	MarcaEspecial *UnionNullString `json:"MarcaEspecial"`

	ContenedoresDePreparacion []ContenedorPreparacion `json:"ContenedoresDePreparacion"`

	ContenedoresEmbalajeDeAlmacen []ContenedorEmbalaje `json:"ContenedoresEmbalajeDeAlmacen"`
}

const PedidoEmpaquetarAvroCRC64Fingerprint = "\xa9j#3\x0f\xacmt"

func NewPedidoEmpaquetar() PedidoEmpaquetar {
	r := PedidoEmpaquetar{}
	r.Identificacion = NewIdentificacion()

	r.Mesa = nil
	r.Calle = nil
	r.FechaExpedicion = nil
	r.MarcaEspecial = nil
	r.ContenedoresDePreparacion = make([]ContenedorPreparacion, 0)

	r.ContenedoresEmbalajeDeAlmacen = make([]ContenedorEmbalaje, 0)

	return r
}

func DeserializePedidoEmpaquetar(r io.Reader) (PedidoEmpaquetar, error) {
	t := NewPedidoEmpaquetar()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePedidoEmpaquetarFromSchema(r io.Reader, schema string) (PedidoEmpaquetar, error) {
	t := NewPedidoEmpaquetar()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePedidoEmpaquetar(r PedidoEmpaquetar, w io.Writer) error {
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
	err = writeUnionNullString(r.Mesa, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Calle, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.FechaExpedicion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.MarcaEspecial, w)
	if err != nil {
		return err
	}
	err = writeArrayContenedorPreparacion(r.ContenedoresDePreparacion, w)
	if err != nil {
		return err
	}
	err = writeArrayContenedorEmbalaje(r.ContenedoresEmbalajeDeAlmacen, w)
	if err != nil {
		return err
	}
	return err
}

func (r PedidoEmpaquetar) Serialize(w io.Writer) error {
	return writePedidoEmpaquetar(r, w)
}

func (r PedidoEmpaquetar) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"},{\"default\":null,\"name\":\"PlantaOperacionId\",\"type\":[\"null\",\"int\"]}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"}},{\"name\":\"OrdenWh\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"default\":null,\"name\":\"Mesa\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaExpedicion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"MarcaEspecial\",\"type\":[\"null\",\"string\"]},{\"name\":\"ContenedoresDePreparacion\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Codigo\",\"type\":\"string\"},{\"name\":\"Articulos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Sku\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"default\":null,\"name\":\"Ean\",\"type\":[\"null\",\"string\"]},{\"name\":\"NroLineaPedido\",\"type\":\"string\"},{\"name\":\"CantidadPedido\",\"type\":\"int\"},{\"name\":\"CantidadPickeada\",\"type\":\"int\"},{\"default\":null,\"name\":\"Lote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Serie\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Zona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoZona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionZona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Longitud\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Altura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Ancho\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Peso\",\"type\":[\"null\",\"float\"]}],\"name\":\"ArticuloContenedorPreparacion\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ContenedorPreparacion\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"ContenedoresEmbalajeDeAlmacen\",\"type\":{\"items\":{\"fields\":[{\"name\":\"ContenedorId\",\"type\":\"string\"},{\"name\":\"ContenedorDescripcion\",\"type\":\"string\"},{\"default\":null,\"name\":\"Longuitud\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Altura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Ancho\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Peso\",\"type\":[\"null\",\"float\"]}],\"name\":\"ContenedorEmbalaje\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Empaquetado.Events.Record.PedidoEmpaquetar\",\"type\":\"record\"}"
}

func (r PedidoEmpaquetar) SchemaName() string {
	return "Andreani.Empaquetado.Events.Record.PedidoEmpaquetar"
}

func (_ PedidoEmpaquetar) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ PedidoEmpaquetar) SetInt(v int32)       { panic("Unsupported operation") }
func (_ PedidoEmpaquetar) SetLong(v int64)      { panic("Unsupported operation") }
func (_ PedidoEmpaquetar) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ PedidoEmpaquetar) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ PedidoEmpaquetar) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ PedidoEmpaquetar) SetString(v string)   { panic("Unsupported operation") }
func (_ PedidoEmpaquetar) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PedidoEmpaquetar) Get(i int) types.Field {
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
		r.Mesa = NewUnionNullString()

		return r.Mesa
	case 4:
		r.Calle = NewUnionNullString()

		return r.Calle
	case 5:
		r.FechaExpedicion = NewUnionNullLong()

		return r.FechaExpedicion
	case 6:
		r.MarcaEspecial = NewUnionNullString()

		return r.MarcaEspecial
	case 7:
		r.ContenedoresDePreparacion = make([]ContenedorPreparacion, 0)

		w := ArrayContenedorPreparacionWrapper{Target: &r.ContenedoresDePreparacion}

		return w

	case 8:
		r.ContenedoresEmbalajeDeAlmacen = make([]ContenedorEmbalaje, 0)

		w := ArrayContenedorEmbalajeWrapper{Target: &r.ContenedoresEmbalajeDeAlmacen}

		return w

	}
	panic("Unknown field index")
}

func (r *PedidoEmpaquetar) SetDefault(i int) {
	switch i {
	case 3:
		r.Mesa = nil
		return
	case 4:
		r.Calle = nil
		return
	case 5:
		r.FechaExpedicion = nil
		return
	case 6:
		r.MarcaEspecial = nil
		return
	}
	panic("Unknown field index")
}

func (r *PedidoEmpaquetar) NullField(i int) {
	switch i {
	case 3:
		r.Mesa = nil
		return
	case 4:
		r.Calle = nil
		return
	case 5:
		r.FechaExpedicion = nil
		return
	case 6:
		r.MarcaEspecial = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ PedidoEmpaquetar) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ PedidoEmpaquetar) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ PedidoEmpaquetar) HintSize(int)                     { panic("Unsupported operation") }
func (_ PedidoEmpaquetar) Finalize()                        {}

func (_ PedidoEmpaquetar) AvroCRC64Fingerprint() []byte {
	return []byte(PedidoEmpaquetarAvroCRC64Fingerprint)
}

func (r PedidoEmpaquetar) MarshalJSON() ([]byte, error) {
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
	output["Mesa"], err = json.Marshal(r.Mesa)
	if err != nil {
		return nil, err
	}
	output["Calle"], err = json.Marshal(r.Calle)
	if err != nil {
		return nil, err
	}
	output["FechaExpedicion"], err = json.Marshal(r.FechaExpedicion)
	if err != nil {
		return nil, err
	}
	output["MarcaEspecial"], err = json.Marshal(r.MarcaEspecial)
	if err != nil {
		return nil, err
	}
	output["ContenedoresDePreparacion"], err = json.Marshal(r.ContenedoresDePreparacion)
	if err != nil {
		return nil, err
	}
	output["ContenedoresEmbalajeDeAlmacen"], err = json.Marshal(r.ContenedoresEmbalajeDeAlmacen)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *PedidoEmpaquetar) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["Mesa"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Mesa); err != nil {
			return err
		}
	} else {
		r.Mesa = NewUnionNullString()

		r.Mesa = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Calle"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Calle); err != nil {
			return err
		}
	} else {
		r.Calle = NewUnionNullString()

		r.Calle = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaExpedicion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaExpedicion); err != nil {
			return err
		}
	} else {
		r.FechaExpedicion = NewUnionNullLong()

		r.FechaExpedicion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["MarcaEspecial"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MarcaEspecial); err != nil {
			return err
		}
	} else {
		r.MarcaEspecial = NewUnionNullString()

		r.MarcaEspecial = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ContenedoresDePreparacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContenedoresDePreparacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ContenedoresDePreparacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ContenedoresEmbalajeDeAlmacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ContenedoresEmbalajeDeAlmacen); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ContenedoresEmbalajeDeAlmacen")
	}
	return nil
}