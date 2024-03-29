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

type UnionNullResponseTypeEnum int

const (
	UnionNullResponseTypeEnumResponse UnionNullResponseTypeEnum = 1
)

type UnionNullResponse struct {
	Null      *types.NullVal
	Response  Response
	UnionType UnionNullResponseTypeEnum
}

func writeUnionNullResponse(r *UnionNullResponse, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullResponseTypeEnumResponse:
		return writeResponse(r.Response, w)
	}
	return fmt.Errorf("invalid value for *UnionNullResponse")
}

func NewUnionNullResponse() *UnionNullResponse {
	return &UnionNullResponse{}
}

func (r *UnionNullResponse) Serialize(w io.Writer) error {
	return writeUnionNullResponse(r, w)
}

func DeserializeUnionNullResponse(r io.Reader) (*UnionNullResponse, error) {
	t := NewUnionNullResponse()
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

func DeserializeUnionNullResponseFromSchema(r io.Reader, schema string) (*UnionNullResponse, error) {
	t := NewUnionNullResponse()
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

func (r *UnionNullResponse) Schema() string {
	return "[\"null\",{\"fields\":[{\"default\":null,\"name\":\"estado\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"motivo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalDeDistribucion\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"ID\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"nomenclatura\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]}],\"name\":\"SucursalResponse\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"sucursalDeImposicion\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.SucursalResponse\"]},{\"default\":null,\"name\":\"sucursalDeRendicion\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.SucursalResponse\"]},{\"default\":null,\"name\":\"sucursalAbastecedora\",\"type\":[\"null\",\"Andreani.AltaOrdenEnvio.Events.Common.SucursalResponse\"]},{\"default\":null,\"name\":\"fechaCreacion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"zonaDeReparto\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroDePermisionaria\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"descripcionServicio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"etiquetaRemito\",\"type\":[\"null\",\"string\"]},{\"name\":\"Bultos\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"NumeroDeBulto\",\"type\":\"string\"},{\"name\":\"NumeroDeEnvio\",\"type\":\"string\"},{\"name\":\"Totalizador\",\"type\":\"string\"},{\"default\":null,\"name\":\"Linking\",\"type\":[\"null\",{\"items\":{\"fields\":[{\"name\":\"meta\",\"type\":\"string\"},{\"name\":\"contenido\",\"type\":\"string\"}],\"name\":\"Metadato\",\"type\":\"record\"},\"type\":\"array\"}]}],\"name\":\"BultoResponse\",\"type\":\"record\"},\"type\":\"array\"}]},{\"default\":null,\"name\":\"FechaEstimadaDeEntrega\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"huellaDeCarbono\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"gastoEnergetico\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"agrupadorDeBultos\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"etiquetasPorAgrupador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"etiquetasDocumentoDeCambio\",\"type\":[\"null\",\"string\"]}],\"name\":\"Response\",\"namespace\":\"Andreani.AltaOrdenEnvio.Events.Common\",\"type\":\"record\"}]"
}

func (_ *UnionNullResponse) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullResponse) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullResponse) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullResponse) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullResponse) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullResponse) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullResponse) SetLong(v int64) {

	r.UnionType = (UnionNullResponseTypeEnum)(v)
}

func (r *UnionNullResponse) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.Response = NewResponse()
		return &types.Record{Target: (&r.Response)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullResponse) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullResponse) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullResponse) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullResponse) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullResponse) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullResponse) Finalize()                        {}

func (r *UnionNullResponse) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullResponseTypeEnumResponse:
		return json.Marshal(map[string]interface{}{"Andreani.AltaOrdenEnvio.Events.Common.Response": r.Response})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullResponse")
}

func (r *UnionNullResponse) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["Andreani.AltaOrdenEnvio.Events.Common.Response"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Response)
	}
	return fmt.Errorf("invalid value for *UnionNullResponse")
}
