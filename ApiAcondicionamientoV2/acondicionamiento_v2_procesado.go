// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     AcondicionamientoV2Procesado.avsc
 */
package ApiAcondicionamientoV2Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type AcondicionamientoV2Procesado struct {
	Estado string `json:"estado"`

	Descripcion *UnionNullString `json:"descripcion"`

	IdTransaccion string `json:"idTransaccion"`

	Almacen *UnionNullString `json:"almacen"`

	Instancia *UnionNullString `json:"instancia"`

	Propietario *UnionNullString `json:"propietario"`

	NumeroOrdenExterna *UnionNullString `json:"numeroOrdenExterna"`

	AcondicionamientoEnFrio *UnionNullAcondicionamientoEnFrio `json:"acondicionamientoEnFrio"`

	Pedido *UnionNullPedidoAcondi `json:"pedido"`

	Abastecimiento *UnionNullAbastecimientoAcondi `json:"abastecimiento"`
}

const AcondicionamientoV2ProcesadoAvroCRC64Fingerprint = "\xe6 \xe9[E\x97ny"

func NewAcondicionamientoV2Procesado() AcondicionamientoV2Procesado {
	r := AcondicionamientoV2Procesado{}
	r.Descripcion = nil
	r.Almacen = nil
	r.Instancia = nil
	r.Propietario = nil
	r.NumeroOrdenExterna = nil
	r.AcondicionamientoEnFrio = nil
	r.Pedido = nil
	r.Abastecimiento = nil
	return r
}

func DeserializeAcondicionamientoV2Procesado(r io.Reader) (AcondicionamientoV2Procesado, error) {
	t := NewAcondicionamientoV2Procesado()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeAcondicionamientoV2ProcesadoFromSchema(r io.Reader, schema string) (AcondicionamientoV2Procesado, error) {
	t := NewAcondicionamientoV2Procesado()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeAcondicionamientoV2Procesado(r AcondicionamientoV2Procesado, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Estado, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.IdTransaccion, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Almacen, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Instancia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Propietario, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroOrdenExterna, w)
	if err != nil {
		return err
	}
	err = writeUnionNullAcondicionamientoEnFrio(r.AcondicionamientoEnFrio, w)
	if err != nil {
		return err
	}
	err = writeUnionNullPedidoAcondi(r.Pedido, w)
	if err != nil {
		return err
	}
	err = writeUnionNullAbastecimientoAcondi(r.Abastecimiento, w)
	if err != nil {
		return err
	}
	return err
}

func (r AcondicionamientoV2Procesado) Serialize(w io.Writer) error {
	return writeAcondicionamientoV2Procesado(r, w)
}

func (r AcondicionamientoV2Procesado) Schema() string {
	return "{\"fields\":[{\"name\":\"estado\",\"type\":\"string\"},{\"default\":null,\"name\":\"descripcion\",\"type\":[\"null\",\"string\"]},{\"name\":\"idTransaccion\",\"type\":\"string\"},{\"default\":null,\"name\":\"almacen\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"instancia\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"propietario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"numeroOrdenExterna\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"acondicionamientoEnFrio\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"fechaAcondi\",\"type\":[\"null\",\"string\"]}],\"name\":\"AcondicionamientoEnFrio\",\"namespace\":\"Andreani.ApiAcondicionamientoV2.Events.Common\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"pedido\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]},{\"name\":\"lineas\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"almacenCliente\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"codigo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cantidad\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoLote\",\"type\":[\"null\",\"string\"]}],\"name\":\"LineaPedido\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"PedidoAcondi\",\"namespace\":\"Andreani.ApiAcondicionamientoV2.Events.Common\",\"type\":\"record\"}]},{\"default\":null,\"name\":\"abastecimiento\",\"type\":[\"null\",{\"fields\":[{\"default\":null,\"name\":\"contratoWarehouse\",\"type\":[\"null\",\"string\"]},{\"name\":\"lineas\",\"type\":{\"items\":{\"fields\":[{\"default\":null,\"name\":\"codigo\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cantidadPedida\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"almacenWMS\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"tipoAcondicionamiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteDeFabricante\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"loteSecundario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fechaDeVencimiento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoLote\",\"type\":[\"null\",\"string\"]}],\"name\":\"LineaAbastecimiento\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"AbastecimientoAcondi\",\"namespace\":\"Andreani.ApiAcondicionamientoV2.Events.Common\",\"type\":\"record\"}]}],\"name\":\"Andreani.ApiAcondicionamientoV2.Events.Record.AcondicionamientoV2Procesado\",\"type\":\"record\"}"
}

func (r AcondicionamientoV2Procesado) SchemaName() string {
	return "Andreani.ApiAcondicionamientoV2.Events.Record.AcondicionamientoV2Procesado"
}

func (_ AcondicionamientoV2Procesado) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ AcondicionamientoV2Procesado) SetInt(v int32)       { panic("Unsupported operation") }
func (_ AcondicionamientoV2Procesado) SetLong(v int64)      { panic("Unsupported operation") }
func (_ AcondicionamientoV2Procesado) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ AcondicionamientoV2Procesado) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ AcondicionamientoV2Procesado) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ AcondicionamientoV2Procesado) SetString(v string)   { panic("Unsupported operation") }
func (_ AcondicionamientoV2Procesado) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AcondicionamientoV2Procesado) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Estado}

		return w

	case 1:
		r.Descripcion = NewUnionNullString()

		return r.Descripcion
	case 2:
		w := types.String{Target: &r.IdTransaccion}

		return w

	case 3:
		r.Almacen = NewUnionNullString()

		return r.Almacen
	case 4:
		r.Instancia = NewUnionNullString()

		return r.Instancia
	case 5:
		r.Propietario = NewUnionNullString()

		return r.Propietario
	case 6:
		r.NumeroOrdenExterna = NewUnionNullString()

		return r.NumeroOrdenExterna
	case 7:
		r.AcondicionamientoEnFrio = NewUnionNullAcondicionamientoEnFrio()

		return r.AcondicionamientoEnFrio
	case 8:
		r.Pedido = NewUnionNullPedidoAcondi()

		return r.Pedido
	case 9:
		r.Abastecimiento = NewUnionNullAbastecimientoAcondi()

		return r.Abastecimiento
	}
	panic("Unknown field index")
}

func (r *AcondicionamientoV2Procesado) SetDefault(i int) {
	switch i {
	case 1:
		r.Descripcion = nil
		return
	case 3:
		r.Almacen = nil
		return
	case 4:
		r.Instancia = nil
		return
	case 5:
		r.Propietario = nil
		return
	case 6:
		r.NumeroOrdenExterna = nil
		return
	case 7:
		r.AcondicionamientoEnFrio = nil
		return
	case 8:
		r.Pedido = nil
		return
	case 9:
		r.Abastecimiento = nil
		return
	}
	panic("Unknown field index")
}

func (r *AcondicionamientoV2Procesado) NullField(i int) {
	switch i {
	case 1:
		r.Descripcion = nil
		return
	case 3:
		r.Almacen = nil
		return
	case 4:
		r.Instancia = nil
		return
	case 5:
		r.Propietario = nil
		return
	case 6:
		r.NumeroOrdenExterna = nil
		return
	case 7:
		r.AcondicionamientoEnFrio = nil
		return
	case 8:
		r.Pedido = nil
		return
	case 9:
		r.Abastecimiento = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ AcondicionamientoV2Procesado) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ AcondicionamientoV2Procesado) AppendArray() types.Field { panic("Unsupported operation") }
func (_ AcondicionamientoV2Procesado) HintSize(int)             { panic("Unsupported operation") }
func (_ AcondicionamientoV2Procesado) Finalize()                {}

func (_ AcondicionamientoV2Procesado) AvroCRC64Fingerprint() []byte {
	return []byte(AcondicionamientoV2ProcesadoAvroCRC64Fingerprint)
}

func (r AcondicionamientoV2Procesado) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["estado"], err = json.Marshal(r.Estado)
	if err != nil {
		return nil, err
	}
	output["descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["idTransaccion"], err = json.Marshal(r.IdTransaccion)
	if err != nil {
		return nil, err
	}
	output["almacen"], err = json.Marshal(r.Almacen)
	if err != nil {
		return nil, err
	}
	output["instancia"], err = json.Marshal(r.Instancia)
	if err != nil {
		return nil, err
	}
	output["propietario"], err = json.Marshal(r.Propietario)
	if err != nil {
		return nil, err
	}
	output["numeroOrdenExterna"], err = json.Marshal(r.NumeroOrdenExterna)
	if err != nil {
		return nil, err
	}
	output["acondicionamientoEnFrio"], err = json.Marshal(r.AcondicionamientoEnFrio)
	if err != nil {
		return nil, err
	}
	output["pedido"], err = json.Marshal(r.Pedido)
	if err != nil {
		return nil, err
	}
	output["abastecimiento"], err = json.Marshal(r.Abastecimiento)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *AcondicionamientoV2Procesado) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["estado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for estado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["descripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Descripcion); err != nil {
			return err
		}
	} else {
		r.Descripcion = NewUnionNullString()

		r.Descripcion = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["idTransaccion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IdTransaccion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for idTransaccion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["almacen"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Almacen); err != nil {
			return err
		}
	} else {
		r.Almacen = NewUnionNullString()

		r.Almacen = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["instancia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Instancia); err != nil {
			return err
		}
	} else {
		r.Instancia = NewUnionNullString()

		r.Instancia = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["propietario"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Propietario); err != nil {
			return err
		}
	} else {
		r.Propietario = NewUnionNullString()

		r.Propietario = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroOrdenExterna"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroOrdenExterna); err != nil {
			return err
		}
	} else {
		r.NumeroOrdenExterna = NewUnionNullString()

		r.NumeroOrdenExterna = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["acondicionamientoEnFrio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AcondicionamientoEnFrio); err != nil {
			return err
		}
	} else {
		r.AcondicionamientoEnFrio = NewUnionNullAcondicionamientoEnFrio()

		r.AcondicionamientoEnFrio = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["pedido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Pedido); err != nil {
			return err
		}
	} else {
		r.Pedido = NewUnionNullPedidoAcondi()

		r.Pedido = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["abastecimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Abastecimiento); err != nil {
			return err
		}
	} else {
		r.Abastecimiento = NewUnionNullAbastecimientoAcondi()

		r.Abastecimiento = nil
	}
	return nil
}
