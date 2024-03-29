// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TipoDeEntrega.avsc
 */
package TrazasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type TipoDeEntrega int32

const (
	TipoDeEntregaUndefined      TipoDeEntrega = 0
	TipoDeEntregaDistribucion   TipoDeEntrega = 1
	TipoDeEntregaDeliveryWindow TipoDeEntrega = 2
)

func (e TipoDeEntrega) String() string {
	switch e {
	case TipoDeEntregaUndefined:
		return "undefined"
	case TipoDeEntregaDistribucion:
		return "distribucion"
	case TipoDeEntregaDeliveryWindow:
		return "deliveryWindow"
	}
	return "unknown"
}

func writeTipoDeEntrega(r TipoDeEntrega, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewTipoDeEntregaValue(raw string) (r TipoDeEntrega, err error) {
	switch raw {
	case "undefined":
		return TipoDeEntregaUndefined, nil
	case "distribucion":
		return TipoDeEntregaDistribucion, nil
	case "deliveryWindow":
		return TipoDeEntregaDeliveryWindow, nil
	}

	return -1, fmt.Errorf("invalid value for TipoDeEntrega: '%s'", raw)

}

func (b TipoDeEntrega) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *TipoDeEntrega) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewTipoDeEntregaValue(stringVal)
	*b = val
	return err
}

type TipoDeEntregaWrapper struct {
	Target *TipoDeEntrega
}

func (b TipoDeEntregaWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b TipoDeEntregaWrapper) SetInt(v int32) {
	*(b.Target) = TipoDeEntrega(v)
}

func (b TipoDeEntregaWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b TipoDeEntregaWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b TipoDeEntregaWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b TipoDeEntregaWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b TipoDeEntregaWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b TipoDeEntregaWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b TipoDeEntregaWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b TipoDeEntregaWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b TipoDeEntregaWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b TipoDeEntregaWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b TipoDeEntregaWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b TipoDeEntregaWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b TipoDeEntregaWrapper) Finalize() {}
