// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     CambioDeStockEstado.avsc
 */
package WarehouseStockEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type CambioDeStockEstado int32

const (
	CambioDeStockEstadoSolicitado CambioDeStockEstado = 0
	CambioDeStockEstadoAceptado   CambioDeStockEstado = 1
	CambioDeStockEstadoRechazado  CambioDeStockEstado = 2
	CambioDeStockEstadoCancelado  CambioDeStockEstado = 3
	CambioDeStockEstadoAjuste     CambioDeStockEstado = 4
)

func (e CambioDeStockEstado) String() string {
	switch e {
	case CambioDeStockEstadoSolicitado:
		return "Solicitado"
	case CambioDeStockEstadoAceptado:
		return "Aceptado"
	case CambioDeStockEstadoRechazado:
		return "Rechazado"
	case CambioDeStockEstadoCancelado:
		return "Cancelado"
	case CambioDeStockEstadoAjuste:
		return "Ajuste"
	}
	return "unknown"
}

func writeCambioDeStockEstado(r CambioDeStockEstado, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewCambioDeStockEstadoValue(raw string) (r CambioDeStockEstado, err error) {
	switch raw {
	case "Solicitado":
		return CambioDeStockEstadoSolicitado, nil
	case "Aceptado":
		return CambioDeStockEstadoAceptado, nil
	case "Rechazado":
		return CambioDeStockEstadoRechazado, nil
	case "Cancelado":
		return CambioDeStockEstadoCancelado, nil
	case "Ajuste":
		return CambioDeStockEstadoAjuste, nil
	}

	return -1, fmt.Errorf("invalid value for CambioDeStockEstado: '%s'", raw)

}

func (b CambioDeStockEstado) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *CambioDeStockEstado) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewCambioDeStockEstadoValue(stringVal)
	*b = val
	return err
}

type CambioDeStockEstadoWrapper struct {
	Target *CambioDeStockEstado
}

func (b CambioDeStockEstadoWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b CambioDeStockEstadoWrapper) SetInt(v int32) {
	*(b.Target) = CambioDeStockEstado(v)
}

func (b CambioDeStockEstadoWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b CambioDeStockEstadoWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b CambioDeStockEstadoWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b CambioDeStockEstadoWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b CambioDeStockEstadoWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b CambioDeStockEstadoWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b CambioDeStockEstadoWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b CambioDeStockEstadoWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b CambioDeStockEstadoWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b CambioDeStockEstadoWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b CambioDeStockEstadoWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b CambioDeStockEstadoWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b CambioDeStockEstadoWrapper) Finalize() {}
