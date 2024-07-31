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

	CodigoEstadoAcondi *UnionNullString `json:"CodigoEstadoAcondi"`

	DescripcionEstadoAcondi *UnionNullString `json:"DescripcionEstadoAcondi"`

	CodigoEstadoTraza *UnionNullString `json:"CodigoEstadoTraza"`

	DescripcionEstadoTraza *UnionNullString `json:"DescripcionEstadoTraza"`

	DescripcionErrorTraza *UnionNullString `json:"DescripcionErrorTraza"`

	DescripcionErrorAcondi *UnionNullString `json:"DescripcionErrorAcondi"`

	ValePsicotropico *UnionNullString `json:"ValePsicotropico"`

	CiudadDestino *UnionNullString `json:"CiudadDestino"`

	ProvinciaDestino *UnionNullString `json:"ProvinciaDestino"`

	CodigoPostalDestino *UnionNullString `json:"CodigoPostalDestino"`

	ContenedoresDePreparacion []ContenedorPreparacion `json:"ContenedoresDePreparacion"`

	ContenedoresEmbalajeDeAlmacen []ContenedorEmbalaje `json:"ContenedoresEmbalajeDeAlmacen"`

	AtributosLote []AtributosDeLote `json:"AtributosLote"`

	ValidacionAtributosLote []ValidacionAtributosDeLote `json:"ValidacionAtributosLote"`
}

const PedidoEmpaquetarAvroCRC64Fingerprint = "\xa8t\xe2ΪE\xfb\xce"

func NewPedidoEmpaquetar() PedidoEmpaquetar {
	r := PedidoEmpaquetar{}
	r.Identificacion = NewIdentificacion()

	r.Mesa = nil
	r.Calle = nil
	r.FechaExpedicion = nil
	r.MarcaEspecial = nil
	r.CodigoEstadoAcondi = nil
	r.DescripcionEstadoAcondi = nil
	r.CodigoEstadoTraza = nil
	r.DescripcionEstadoTraza = nil
	r.DescripcionErrorTraza = nil
	r.DescripcionErrorAcondi = nil
	r.ValePsicotropico = nil
	r.CiudadDestino = nil
	r.ProvinciaDestino = nil
	r.CodigoPostalDestino = nil
	r.ContenedoresDePreparacion = make([]ContenedorPreparacion, 0)

	r.ContenedoresEmbalajeDeAlmacen = make([]ContenedorEmbalaje, 0)

	r.AtributosLote = make([]AtributosDeLote, 0)

	r.ValidacionAtributosLote = make([]ValidacionAtributosDeLote, 0)

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
	err = writeUnionNullString(r.CodigoEstadoAcondi, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DescripcionEstadoAcondi, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoEstadoTraza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DescripcionEstadoTraza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DescripcionErrorTraza, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.DescripcionErrorAcondi, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ValePsicotropico, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CiudadDestino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ProvinciaDestino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.CodigoPostalDestino, w)
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
	err = writeArrayAtributosDeLote(r.AtributosLote, w)
	if err != nil {
		return err
	}
	err = writeArrayValidacionAtributosDeLote(r.ValidacionAtributosLote, w)
	if err != nil {
		return err
	}
	return err
}

func (r PedidoEmpaquetar) Serialize(w io.Writer) error {
	return writePedidoEmpaquetar(r, w)
}

func (r PedidoEmpaquetar) Schema() string {
	return "{\"fields\":[{\"name\":\"Identificacion\",\"type\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"Evento\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Proceso\",\"type\":\"string\"},{\"name\":\"FechaHoraGeneracion\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"SistemaOrigen\",\"type\":\"string\"},{\"name\":\"Almacen\",\"type\":\"string\"},{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Instancia\",\"type\":\"string\"},{\"default\":null,\"name\":\"PlantaOperacionId\",\"type\":[\"null\",\"int\"]}],\"name\":\"Identificacion\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"}},{\"name\":\"OrdenWh\",\"type\":\"string\"},{\"name\":\"OrdenCliente\",\"type\":\"string\"},{\"default\":null,\"name\":\"Mesa\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FechaExpedicion\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]},{\"default\":null,\"name\":\"MarcaEspecial\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoEstadoAcondi\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionEstadoAcondi\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoEstadoTraza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionEstadoTraza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionErrorTraza\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionErrorAcondi\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ValePsicotropico\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CiudadDestino\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ProvinciaDestino\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoPostalDestino\",\"type\":[\"null\",\"string\"]},{\"name\":\"ContenedoresDePreparacion\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Codigo\",\"type\":\"string\"},{\"name\":\"Articulos\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Sku\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"default\":null,\"name\":\"Ean\",\"type\":[\"null\",\"string\"]},{\"name\":\"NroLineaPedido\",\"type\":\"string\"},{\"name\":\"CantidadPedido\",\"type\":\"int\"},{\"name\":\"CantidadPickeada\",\"type\":\"int\"},{\"default\":null,\"name\":\"Mesa\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Lote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Serie\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Zona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoZona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DescripcionZona\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Longitud\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Altura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Ancho\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Peso\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"InstruccionesEmbalaje\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UDM\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PickearTodos\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"ValidaLoteWos\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ValidaSerieWos\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FormatoSerie\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ValidaSerieEnSalida\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"IngresaSerieEnEmpaquetado\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"SerieDirigida\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PickDetailKey\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"CodigoValidacionSalidaLote\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"UsaDataMatrix\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"UbicacionOrigen\",\"type\":[\"null\",\"string\"]}],\"name\":\"ArticuloContenedorPreparacion\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"ContenedorPreparacion\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"ContenedoresEmbalajeDeAlmacen\",\"type\":{\"items\":{\"fields\":[{\"name\":\"ContenedorId\",\"type\":\"string\"},{\"name\":\"ContenedorDescripcion\",\"type\":\"string\"},{\"default\":null,\"name\":\"Longuitud\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Altura\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Ancho\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"Peso\",\"type\":[\"null\",\"float\"]},{\"name\":\"EsRetornable\",\"type\":\"boolean\"}],\"name\":\"ContenedorEmbalaje\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"AtributosLote\",\"type\":{\"items\":{\"fields\":[{\"name\":\"Propietario\",\"type\":\"string\"},{\"name\":\"Sku\",\"type\":\"string\"},{\"name\":\"Lote\",\"type\":\"string\"},{\"default\":null,\"name\":\"Atributo1\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo2\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo3\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo4\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo5\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo6\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo7\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo8\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo9\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo10\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo11\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Atributo12\",\"type\":[\"null\",\"string\"]}],\"name\":\"AtributosDeLote\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}},{\"name\":\"ValidacionAtributosLote\",\"type\":{\"items\":{\"fields\":[{\"name\":\"CodigoValidacionSalidaLote\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"default\":null,\"name\":\"Atributo1Display\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo2Display\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo3Display\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo4Display\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo5Display\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo6Display\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo7Display\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo8Display\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo9Display\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo10Display\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo11Display\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo12Display\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo1Required\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo2Required\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo3Required\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo4Required\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo5Required\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo6Required\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo7Required\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo8Required\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo9Required\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo10Required\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo11Required\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo12Required\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo1Choose\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo2Choose\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo3Choose\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo4Choose\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo5Choose\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo6Choose\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo7Choose\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo8Choose\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo9Choose\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo10Choose\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo11Choose\",\"type\":[\"null\",\"boolean\"]},{\"default\":null,\"name\":\"Atributo12Choose\",\"type\":[\"null\",\"boolean\"]}],\"name\":\"ValidacionAtributosDeLote\",\"namespace\":\"Andreani.Empaquetado.Events.Common\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Andreani.Empaquetado.Events.Record.PedidoEmpaquetar\",\"type\":\"record\"}"
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
		r.CodigoEstadoAcondi = NewUnionNullString()

		return r.CodigoEstadoAcondi
	case 8:
		r.DescripcionEstadoAcondi = NewUnionNullString()

		return r.DescripcionEstadoAcondi
	case 9:
		r.CodigoEstadoTraza = NewUnionNullString()

		return r.CodigoEstadoTraza
	case 10:
		r.DescripcionEstadoTraza = NewUnionNullString()

		return r.DescripcionEstadoTraza
	case 11:
		r.DescripcionErrorTraza = NewUnionNullString()

		return r.DescripcionErrorTraza
	case 12:
		r.DescripcionErrorAcondi = NewUnionNullString()

		return r.DescripcionErrorAcondi
	case 13:
		r.ValePsicotropico = NewUnionNullString()

		return r.ValePsicotropico
	case 14:
		r.CiudadDestino = NewUnionNullString()

		return r.CiudadDestino
	case 15:
		r.ProvinciaDestino = NewUnionNullString()

		return r.ProvinciaDestino
	case 16:
		r.CodigoPostalDestino = NewUnionNullString()

		return r.CodigoPostalDestino
	case 17:
		r.ContenedoresDePreparacion = make([]ContenedorPreparacion, 0)

		w := ArrayContenedorPreparacionWrapper{Target: &r.ContenedoresDePreparacion}

		return w

	case 18:
		r.ContenedoresEmbalajeDeAlmacen = make([]ContenedorEmbalaje, 0)

		w := ArrayContenedorEmbalajeWrapper{Target: &r.ContenedoresEmbalajeDeAlmacen}

		return w

	case 19:
		r.AtributosLote = make([]AtributosDeLote, 0)

		w := ArrayAtributosDeLoteWrapper{Target: &r.AtributosLote}

		return w

	case 20:
		r.ValidacionAtributosLote = make([]ValidacionAtributosDeLote, 0)

		w := ArrayValidacionAtributosDeLoteWrapper{Target: &r.ValidacionAtributosLote}

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
	case 7:
		r.CodigoEstadoAcondi = nil
		return
	case 8:
		r.DescripcionEstadoAcondi = nil
		return
	case 9:
		r.CodigoEstadoTraza = nil
		return
	case 10:
		r.DescripcionEstadoTraza = nil
		return
	case 11:
		r.DescripcionErrorTraza = nil
		return
	case 12:
		r.DescripcionErrorAcondi = nil
		return
	case 13:
		r.ValePsicotropico = nil
		return
	case 14:
		r.CiudadDestino = nil
		return
	case 15:
		r.ProvinciaDestino = nil
		return
	case 16:
		r.CodigoPostalDestino = nil
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
	case 7:
		r.CodigoEstadoAcondi = nil
		return
	case 8:
		r.DescripcionEstadoAcondi = nil
		return
	case 9:
		r.CodigoEstadoTraza = nil
		return
	case 10:
		r.DescripcionEstadoTraza = nil
		return
	case 11:
		r.DescripcionErrorTraza = nil
		return
	case 12:
		r.DescripcionErrorAcondi = nil
		return
	case 13:
		r.ValePsicotropico = nil
		return
	case 14:
		r.CiudadDestino = nil
		return
	case 15:
		r.ProvinciaDestino = nil
		return
	case 16:
		r.CodigoPostalDestino = nil
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
	output["CodigoEstadoAcondi"], err = json.Marshal(r.CodigoEstadoAcondi)
	if err != nil {
		return nil, err
	}
	output["DescripcionEstadoAcondi"], err = json.Marshal(r.DescripcionEstadoAcondi)
	if err != nil {
		return nil, err
	}
	output["CodigoEstadoTraza"], err = json.Marshal(r.CodigoEstadoTraza)
	if err != nil {
		return nil, err
	}
	output["DescripcionEstadoTraza"], err = json.Marshal(r.DescripcionEstadoTraza)
	if err != nil {
		return nil, err
	}
	output["DescripcionErrorTraza"], err = json.Marshal(r.DescripcionErrorTraza)
	if err != nil {
		return nil, err
	}
	output["DescripcionErrorAcondi"], err = json.Marshal(r.DescripcionErrorAcondi)
	if err != nil {
		return nil, err
	}
	output["ValePsicotropico"], err = json.Marshal(r.ValePsicotropico)
	if err != nil {
		return nil, err
	}
	output["CiudadDestino"], err = json.Marshal(r.CiudadDestino)
	if err != nil {
		return nil, err
	}
	output["ProvinciaDestino"], err = json.Marshal(r.ProvinciaDestino)
	if err != nil {
		return nil, err
	}
	output["CodigoPostalDestino"], err = json.Marshal(r.CodigoPostalDestino)
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
	output["AtributosLote"], err = json.Marshal(r.AtributosLote)
	if err != nil {
		return nil, err
	}
	output["ValidacionAtributosLote"], err = json.Marshal(r.ValidacionAtributosLote)
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
		if v, ok := fields["CodigoEstadoAcondi"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoEstadoAcondi); err != nil {
			return err
		}
	} else {
		r.CodigoEstadoAcondi = NewUnionNullString()

		r.CodigoEstadoAcondi = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DescripcionEstadoAcondi"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DescripcionEstadoAcondi); err != nil {
			return err
		}
	} else {
		r.DescripcionEstadoAcondi = NewUnionNullString()

		r.DescripcionEstadoAcondi = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoEstadoTraza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoEstadoTraza); err != nil {
			return err
		}
	} else {
		r.CodigoEstadoTraza = NewUnionNullString()

		r.CodigoEstadoTraza = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DescripcionEstadoTraza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DescripcionEstadoTraza); err != nil {
			return err
		}
	} else {
		r.DescripcionEstadoTraza = NewUnionNullString()

		r.DescripcionEstadoTraza = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DescripcionErrorTraza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DescripcionErrorTraza); err != nil {
			return err
		}
	} else {
		r.DescripcionErrorTraza = NewUnionNullString()

		r.DescripcionErrorTraza = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DescripcionErrorAcondi"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DescripcionErrorAcondi); err != nil {
			return err
		}
	} else {
		r.DescripcionErrorAcondi = NewUnionNullString()

		r.DescripcionErrorAcondi = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ValePsicotropico"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ValePsicotropico); err != nil {
			return err
		}
	} else {
		r.ValePsicotropico = NewUnionNullString()

		r.ValePsicotropico = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CiudadDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CiudadDestino); err != nil {
			return err
		}
	} else {
		r.CiudadDestino = NewUnionNullString()

		r.CiudadDestino = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ProvinciaDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ProvinciaDestino); err != nil {
			return err
		}
	} else {
		r.ProvinciaDestino = NewUnionNullString()

		r.ProvinciaDestino = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoPostalDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoPostalDestino); err != nil {
			return err
		}
	} else {
		r.CodigoPostalDestino = NewUnionNullString()

		r.CodigoPostalDestino = nil
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
	val = func() json.RawMessage {
		if v, ok := fields["AtributosLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AtributosLote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AtributosLote")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ValidacionAtributosLote"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ValidacionAtributosLote); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ValidacionAtributosLote")
	}
	return nil
}
