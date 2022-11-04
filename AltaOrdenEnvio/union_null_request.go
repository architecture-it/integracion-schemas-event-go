// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TransactionEvent.avsc
 */
package AltaOrdenEnvioEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

type UnionNullRequestTypeEnum int

const (
	UnionNullRequestTypeEnumRequest UnionNullRequestTypeEnum = 1
)

type UnionNullRequest struct {
	Null      *types.NullVal
	Request   Request
	UnionType UnionNullRequestTypeEnum
}

func writeUnionNullRequest(r *UnionNullRequest, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullRequestTypeEnumRequest:
		return writeRequest(r.Request, w)
	}
	return fmt.Errorf("invalid value for *UnionNullRequest")
}

func NewUnionNullRequest() *UnionNullRequest {
	return &UnionNullRequest{}
}

func (r *UnionNullRequest) Serialize(w io.Writer) error {
	return writeUnionNullRequest(r, w)
}

func DeserializeUnionNullRequest(r io.Reader) (*UnionNullRequest, error) {
	t := NewUnionNullRequest()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNullRequestFromSchema(r io.Reader, schema string) (*UnionNullRequest, error) {
	t := NewUnionNullRequest()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r *UnionNullRequest) Schema() string {
	return "[\"null\",{\"fields\":[{\"name\":\"Bultos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"Kilos\",\"type\":\"double\"},{\"name\":\"LargoCm\",\"type\":\"double\"},{\"name\":\"AltoCm\",\"type\":\"double\"},{\"name\":\"AnchoCm\",\"type\":\"double\"},{\"name\":\"VolumenCm\",\"type\":\"double\"},{\"name\":\"ValorDeclaradoSinImpuestos\",\"type\":\"double\"},{\"name\":\"ValorDeclaradoConImpuestos\",\"type\":\"double\"},{\"name\":\"ValorDeclarado\",\"type\":\"double\"},{\"default\":null,\"name\":\"Descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Referencias\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"NumeroDeEnvio\",\"type\":[\"null\",\"string\"]}],\"name\":\"BultoRequest\",\"type\":\"record\"},\"type\":\"array\"}]},{\"name\":\"Contrato\",\"type\":\"string\"},{\"default\":null,\"name\":\"TipoServicio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"SucursalClienteID\",\"type\":[\"null\",\"int\"]},{\"name\":\"origen\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"postal\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"abreviaturaProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoDeDireccion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoCiudad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nombreProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO -3166-2 (https://es.wikipedia.org/wiki/ISO_3166-2:AR)\",\"name\":\"codigoISOProvincia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"doc\":\"segun ISO 3166-1 alpha-2 (https://es.wikipedia.org/wiki/ISO_3166-1#Tabla_de_c.C3.B3digos_alfa-2_asignados_o_reservados)\",\"name\":\"codigoISOPais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",{\"fields\":[{\"name\":\"metadatos\",\"type\":[\"null\",{\"items\":\"Andreani.AltaOrdenEnvio.Events.Common.Metadato\",\"type\":\"array\"}]}],\"name\":\"ListaDePropiedades\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"casillaDeCorreo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]}],\"name\":\"Direccion\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"datosSucursal\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"ID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nomenclatura\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"direccion\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"localidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"Region\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pais\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigoPostal\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"componentesDeDireccion\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDePropiedades\"]},{\"default\":null,\"name\":\"casillaDeCorreo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"calle\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"piso\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"departamento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"region\",\"type\":[\"null\",\"string\"]}],\"name\":\"DireccionPostal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"datosAdicionales\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDePropiedades\"]},{\"default\":null,\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"fields\":[{\"name\":\"listaDeTelefonos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"tipo\",\"type\":\"string\"},{\"name\":\"numero\",\"type\":\"string\"}],\"name\":\"Telefono\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"ListaDeTelefonos\",\"type\":\"record\"}]}],\"name\":\"Sucursal\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"GeoCoordenada\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"latitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"longitud\",\"type\":[\"null\",\"double\"]},{\"default\":null,\"name\":\"elevacion\",\"type\":[\"null\",\"double\"]}],\"name\":\"GeoCoordenada\",\"type\":\"record\"}]}],\"name\":\"LugarDeRetiroOEntrega\",\"type\":\"record\"}]},{\"name\":\"destino\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.LugarDeRetiroOEntrega\"]},{\"name\":\"idPedido\",\"type\":[\"null\",\"string\"]},{\"name\":\"remitente\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"nombreCompleto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"email\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"documentoTipo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"documentoNumero\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"telefonos\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.ListaDeTelefonos\"]}],\"name\":\"Persona\",\"type\":\"record\"}]},{\"name\":\"destinatario\",\"type\":[\"null\",{\"items\":\"Andreani.AltaOrdenEnvio.Events.Common.Persona\",\"type\":\"array\"}]},{\"name\":\"remito\",\"type\":[\"null\",{\"fields\":[{\"name\":\"complementarios\",\"type\":[\"null\",{\"items\":\"string\",\"type\":\"array\"}]},{\"name\":\"numeroDeRemito\",\"type\":[\"null\",\"string\"]}],\"name\":\"Remito\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"centroDeCosto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"productoAEntregar\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"productoARetirar\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoProducto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"categoriaDeFacturacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"pagoDestino\",\"type\":[\"null\",\"int\"]},{\"default\":null,\"name\":\"valorACobrar\",\"type\":[\"null\",\"float\"]},{\"default\":null,\"name\":\"codigoVerificadorDeEntrega\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaYHorario\",\"type\":[\"null\",{\"fields\":[{\"name\":\"fecha\",\"type\":[\"null\",\"string\"]},{\"name\":\"horaDesde\",\"type\":[\"null\",\"string\"]},{\"name\":\"horaHasta\",\"type\":[\"null\",\"string\"]}],\"name\":\"FechaYHorario\",\"type\":\"record\"}]}],\"name\":\"Request\",\"namespace\":\"Andreani.AltaOrdenEnvio.Events.Common\",\"type\":\"record\"}]"
}

func (_ *UnionNullRequest) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullRequest) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullRequest) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullRequest) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullRequest) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullRequest) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullRequest) SetLong(v int64) {

	r.UnionType = (UnionNullRequestTypeEnum)(v)
}

func (r *UnionNullRequest) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Request = NewRequest()
		return &types.Record{Target: (&r.Request)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullRequest) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullRequest) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullRequest) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullRequest) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullRequest) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullRequest) Finalize()                        {}

func (r *UnionNullRequest) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullRequestTypeEnumRequest:
		return json.Marshal(map[string]interface{}{"Andreani.AltaOrdenEnvio.Events.Common.Request": r.Request})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullRequest")
}

func (r *UnionNullRequest) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.AltaOrdenEnvio.Events.Common.Request"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Request)
	}
	return fmt.Errorf("invalid value for *UnionNullRequest")
}
