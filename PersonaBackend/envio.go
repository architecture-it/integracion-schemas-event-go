// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EnviosEvent.avsc
 */
package PersonaBackendEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Envio struct {
	Id string `json:"Id"`

	TipoDeEnvio string `json:"TipoDeEnvio"`

	ModoDeEntrega string `json:"ModoDeEntrega"`

	NumeroInterno string `json:"NumeroInterno"`

	NumeroDeSeguimiento string `json:"NumeroDeSeguimiento"`

	NumeroDeContrato string `json:"NumeroDeContrato"`

	Estado string `json:"Estado"`

	RemitoFactura string `json:"RemitoFactura"`

	FechaCreacion int64 `json:"FechaCreacion"`

	Origen *UnionNullOrigen `json:"Origen"`

	Destino *UnionNullDestino `json:"Destino"`

	Remitente *UnionNullRemitente `json:"Remitente"`

	Destinatario *UnionNullDestinatario `json:"Destinatario"`

	Paquetes *UnionNullArrayPaquete `json:"Paquetes"`

	DatosConstancia *UnionNullDatosConstancia `json:"DatosConstancia"`

	Cliente *UnionNullCliente `json:"Cliente"`

	Cotizacion *UnionNullCotizacion `json:"Cotizacion"`

	Canal *UnionNullString `json:"Canal"`

	Referencia *UnionNullString `json:"Referencia"`
}

const EnvioAvroCRC64Fingerprint = "\xc1\xbf\xa8\x8f\xd1 ^\x15"

func NewEnvio() Envio {
	r := Envio{}
	r.Origen = nil
	r.Destino = nil
	r.Remitente = nil
	r.Destinatario = nil
	r.Paquetes = nil
	r.DatosConstancia = nil
	r.Cliente = nil
	r.Cotizacion = nil
	r.Canal = nil
	r.Referencia = nil
	return r
}

func DeserializeEnvio(r io.Reader) (Envio, error) {
	t := NewEnvio()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioFromSchema(r io.Reader, schema string) (Envio, error) {
	t := NewEnvio()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvio(r Envio, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoDeEnvio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ModoDeEntrega, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroInterno, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroDeSeguimiento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroDeContrato, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Estado, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.RemitoFactura, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.FechaCreacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullOrigen(r.Origen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDestino(r.Destino, w)
	if err != nil {
		return err
	}
	err = writeUnionNullRemitente(r.Remitente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDestinatario(r.Destinatario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullArrayPaquete(r.Paquetes, w)
	if err != nil {
		return err
	}
	err = writeUnionNullDatosConstancia(r.DatosConstancia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullCliente(r.Cliente, w)
	if err != nil {
		return err
	}
	err = writeUnionNullCotizacion(r.Cotizacion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Canal, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Referencia, w)
	if err != nil {
		return err
	}
	return err
}

func (r Envio) Serialize(w io.Writer) error {
	return writeEnvio(r, w)
}

func (r Envio) Schema() string {
	return "{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"TipoDeEnvio\",\"type\":\"string\"},{\"name\":\"ModoDeEntrega\",\"type\":\"string\"},{\"name\":\"NumeroInterno\",\"type\":\"string\"},{\"name\":\"NumeroDeSeguimiento\",\"type\":\"string\"},{\"name\":\"NumeroDeContrato\",\"type\":\"string\"},{\"name\":\"Estado\",\"type\":\"string\"},{\"name\":\"RemitoFactura\",\"type\":\"string\"},{\"name\":\"FechaCreacion\",\"type\":\"long\"},{\"default\":null,\"name\":\"Origen\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"Calle\",\"type\":\"string\"},{\"default\":null,\"name\":\"Numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Unidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"default\":null,\"name\":\"SucursalId\",\"type\":[\"null\",\"string\"]}],\"name\":\"Origen\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Destino\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"Calle\",\"type\":\"string\"},{\"default\":null,\"name\":\"Numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Unidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"name\":\"Provincia\",\"type\":\"string\"},{\"default\":null,\"name\":\"ObservacionesAdicionales\",\"type\":[\"null\",\"string\"]}],\"name\":\"Destino\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Remitente\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"},{\"default\":null,\"name\":\"Apellido\",\"type\":[\"null\",\"string\"]},{\"name\":\"Email\",\"type\":\"string\"},{\"name\":\"Telefono\",\"type\":\"string\"},{\"name\":\"Dni\",\"type\":\"string\"},{\"name\":\"Calle\",\"type\":\"string\"},{\"default\":null,\"name\":\"Numero\",\"type\":[\"null\",\"string\"]},{\"name\":\"Localidad\",\"type\":\"string\"},{\"name\":\"CodigoPostal\",\"type\":\"string\"},{\"default\":null,\"name\":\"Piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Unidad\",\"type\":[\"null\",\"string\"]},{\"name\":\"Pais\",\"type\":\"string\"},{\"name\":\"Region\",\"type\":\"string\"}],\"name\":\"Remitente\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Destinatario\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Apellido\",\"type\":\"string\"},{\"name\":\"Email\",\"type\":\"string\"},{\"name\":\"Telefono\",\"type\":\"string\"},{\"name\":\"Dni\",\"type\":\"string\"}],\"name\":\"Destinatario\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Paquetes\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"Id\",\"type\":\"string\"},{\"name\":\"TipoId\",\"type\":\"string\"},{\"name\":\"Tipo\",\"type\":\"string\"},{\"name\":\"Alto\",\"type\":\"string\"},{\"name\":\"Ancho\",\"type\":\"string\"},{\"name\":\"Largo\",\"type\":\"string\"},{\"name\":\"Peso\",\"type\":\"string\"},{\"name\":\"ValorDeclarado\",\"type\":\"int\"},{\"default\":null,\"name\":\"NumeroDeBulto\",\"type\":[\"null\",\"string\"]}],\"name\":\"Paquete\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"DatosConstancia\",\"type\":[\"null\",{\"fields\":[{\"name\":\"Url\",\"type\":\"string\"},{\"name\":\"NumeroPermisionaria\",\"type\":\"string\"},{\"name\":\"SucursalDistribucion\",\"type\":\"string\"},{\"name\":\"SucursalDistribucionNomenclatura\",\"type\":\"string\"},{\"name\":\"SucursalDistribucionDescripcion\",\"type\":\"string\"},{\"name\":\"SucursalDistribucionId\",\"type\":\"string\"},{\"name\":\"SucursalRendicion\",\"type\":\"string\"},{\"name\":\"SucursalRendicionNomenclatura\",\"type\":\"string\"},{\"name\":\"SucursalRendicionDescripcion\",\"type\":\"string\"},{\"name\":\"SucursalRendicionId\",\"type\":\"string\"},{\"name\":\"CodigoSucursalCabecera\",\"type\":\"string\"},{\"name\":\"SucursalAbastecedoraDescripcion\",\"type\":\"string\"},{\"name\":\"SucursalAbastecedoraId\",\"type\":\"string\"},{\"name\":\"CodigoZonaReparto\",\"type\":\"string\"},{\"default\":null,\"name\":\"Region\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Provincia\",\"type\":[\"null\",\"string\"]}],\"name\":\"DatosConstancia\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Cliente\",\"type\":[\"null\",{\"fields\":[{\"name\":\"CodigoAndreani\",\"type\":\"string\"},{\"name\":\"Nombre\",\"type\":\"string\"},{\"name\":\"Logo\",\"type\":\"string\"}],\"name\":\"Cliente\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Cotizacion\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"IdMercadoPago\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Iva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"PesoAforado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SeguroDistribucionSinIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SeguroDistribucionConIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DistribucionSinIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"DistribucionConIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TarifaSinIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"TarifaConIva\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Descuento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Cupon\",\"type\":[\"null\",\"string\"]}],\"name\":\"Cotizacion\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"Canal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Referencia\",\"type\":[\"null\",\"string\"]}],\"name\":\"Andreani.PersonaBackend.Events.Common.Envio\",\"type\":\"record\"}"
}

func (r Envio) SchemaName() string {
	return "Andreani.PersonaBackend.Events.Common.Envio"
}

func (_ Envio) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Envio) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Envio) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Envio) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Envio) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Envio) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Envio) SetString(v string)   { panic("Unsupported operation") }
func (_ Envio) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Envio) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Id}

		return w

	case 1:
		w := types.String{Target: &r.TipoDeEnvio}

		return w

	case 2:
		w := types.String{Target: &r.ModoDeEntrega}

		return w

	case 3:
		w := types.String{Target: &r.NumeroInterno}

		return w

	case 4:
		w := types.String{Target: &r.NumeroDeSeguimiento}

		return w

	case 5:
		w := types.String{Target: &r.NumeroDeContrato}

		return w

	case 6:
		w := types.String{Target: &r.Estado}

		return w

	case 7:
		w := types.String{Target: &r.RemitoFactura}

		return w

	case 8:
		w := types.Long{Target: &r.FechaCreacion}

		return w

	case 9:
		r.Origen = NewUnionNullOrigen()

		return r.Origen
	case 10:
		r.Destino = NewUnionNullDestino()

		return r.Destino
	case 11:
		r.Remitente = NewUnionNullRemitente()

		return r.Remitente
	case 12:
		r.Destinatario = NewUnionNullDestinatario()

		return r.Destinatario
	case 13:
		r.Paquetes = NewUnionNullArrayPaquete()

		return r.Paquetes
	case 14:
		r.DatosConstancia = NewUnionNullDatosConstancia()

		return r.DatosConstancia
	case 15:
		r.Cliente = NewUnionNullCliente()

		return r.Cliente
	case 16:
		r.Cotizacion = NewUnionNullCotizacion()

		return r.Cotizacion
	case 17:
		r.Canal = NewUnionNullString()

		return r.Canal
	case 18:
		r.Referencia = NewUnionNullString()

		return r.Referencia
	}
	panic("Unknown field index")
}

func (r *Envio) SetDefault(i int) {
	switch i {
	case 9:
		r.Origen = nil
		return
	case 10:
		r.Destino = nil
		return
	case 11:
		r.Remitente = nil
		return
	case 12:
		r.Destinatario = nil
		return
	case 13:
		r.Paquetes = nil
		return
	case 14:
		r.DatosConstancia = nil
		return
	case 15:
		r.Cliente = nil
		return
	case 16:
		r.Cotizacion = nil
		return
	case 17:
		r.Canal = nil
		return
	case 18:
		r.Referencia = nil
		return
	}
	panic("Unknown field index")
}

func (r *Envio) NullField(i int) {
	switch i {
	case 9:
		r.Origen = nil
		return
	case 10:
		r.Destino = nil
		return
	case 11:
		r.Remitente = nil
		return
	case 12:
		r.Destinatario = nil
		return
	case 13:
		r.Paquetes = nil
		return
	case 14:
		r.DatosConstancia = nil
		return
	case 15:
		r.Cliente = nil
		return
	case 16:
		r.Cotizacion = nil
		return
	case 17:
		r.Canal = nil
		return
	case 18:
		r.Referencia = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ Envio) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Envio) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Envio) HintSize(int)                     { panic("Unsupported operation") }
func (_ Envio) Finalize()                        {}

func (_ Envio) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioAvroCRC64Fingerprint)
}

func (r Envio) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["TipoDeEnvio"], err = json.Marshal(r.TipoDeEnvio)
	if err != nil {
		return nil, err
	}
	output["ModoDeEntrega"], err = json.Marshal(r.ModoDeEntrega)
	if err != nil {
		return nil, err
	}
	output["NumeroInterno"], err = json.Marshal(r.NumeroInterno)
	if err != nil {
		return nil, err
	}
	output["NumeroDeSeguimiento"], err = json.Marshal(r.NumeroDeSeguimiento)
	if err != nil {
		return nil, err
	}
	output["NumeroDeContrato"], err = json.Marshal(r.NumeroDeContrato)
	if err != nil {
		return nil, err
	}
	output["Estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	output["RemitoFactura"], err = json.Marshal(r.RemitoFactura)
	if err != nil {
		return nil, err
	}
	output["FechaCreacion"], err = json.Marshal(r.FechaCreacion)
	if err != nil {
		return nil, err
	}
	output["Origen"], err = json.Marshal(r.Origen)
	if err != nil {
		return nil, err
	}
	output["Destino"], err = json.Marshal(r.Destino)
	if err != nil {
		return nil, err
	}
	output["Remitente"], err = json.Marshal(r.Remitente)
	if err != nil {
		return nil, err
	}
	output["Destinatario"], err = json.Marshal(r.Destinatario)
	if err != nil {
		return nil, err
	}
	output["Paquetes"], err = json.Marshal(r.Paquetes)
	if err != nil {
		return nil, err
	}
	output["DatosConstancia"], err = json.Marshal(r.DatosConstancia)
	if err != nil {
		return nil, err
	}
	output["Cliente"], err = json.Marshal(r.Cliente)
	if err != nil {
		return nil, err
	}
	output["Cotizacion"], err = json.Marshal(r.Cotizacion)
	if err != nil {
		return nil, err
	}
	output["Canal"], err = json.Marshal(r.Canal)
	if err != nil {
		return nil, err
	}
	output["Referencia"], err = json.Marshal(r.Referencia)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Envio) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeEnvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoDeEnvio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ModoDeEntrega"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ModoDeEntrega); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ModoDeEntrega")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroInterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroInterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NumeroInterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeSeguimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeSeguimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NumeroDeSeguimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NumeroDeContrato"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeContrato); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NumeroDeContrato")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Estado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["RemitoFactura"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RemitoFactura); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for RemitoFactura")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaCreacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaCreacion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaCreacion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Origen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Origen); err != nil {
			return err
		}
	} else {
		r.Origen = NewUnionNullOrigen()

		r.Origen = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Destino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Destino); err != nil {
			return err
		}
	} else {
		r.Destino = NewUnionNullDestino()

		r.Destino = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Remitente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Remitente); err != nil {
			return err
		}
	} else {
		r.Remitente = NewUnionNullRemitente()

		r.Remitente = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Destinatario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Destinatario); err != nil {
			return err
		}
	} else {
		r.Destinatario = NewUnionNullDestinatario()

		r.Destinatario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Paquetes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Paquetes); err != nil {
			return err
		}
	} else {
		r.Paquetes = NewUnionNullArrayPaquete()

		r.Paquetes = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["DatosConstancia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DatosConstancia); err != nil {
			return err
		}
	} else {
		r.DatosConstancia = NewUnionNullDatosConstancia()

		r.DatosConstancia = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cliente"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cliente); err != nil {
			return err
		}
	} else {
		r.Cliente = NewUnionNullCliente()

		r.Cliente = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Cotizacion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Cotizacion); err != nil {
			return err
		}
	} else {
		r.Cotizacion = NewUnionNullCotizacion()

		r.Cotizacion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Canal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Canal); err != nil {
			return err
		}
	} else {
		r.Canal = NewUnionNullString()

		r.Canal = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["Referencia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Referencia); err != nil {
			return err
		}
	} else {
		r.Referencia = NewUnionNullString()

		r.Referencia = nil
	}
	return nil
}
