// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TipoPagoDestino.avsc
 */
package AltaOrdenEnvioEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type TipoPagoDestino int32

const (
	TipoPagoDestinoUndefined TipoPagoDestino = 0
	TipoPagoDestinoP         TipoPagoDestino = 1
	TipoPagoDestinoD         TipoPagoDestino = 2
)

func (e TipoPagoDestino) String() string {
	switch e {
	case TipoPagoDestinoUndefined:
		return "undefined"
	case TipoPagoDestinoP:
		return "P"
	case TipoPagoDestinoD:
		return "D"
	}
	return "unknown"
}

func writeTipoPagoDestino(r TipoPagoDestino, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewTipoPagoDestinoValue(raw string) (r TipoPagoDestino, err error) {
	switch raw {
	case "undefined":
		return TipoPagoDestinoUndefined, nil
	case "P":
		return TipoPagoDestinoP, nil
	case "D":
		return TipoPagoDestinoD, nil
	}

	return -1, fmt.Errorf("invalid value for TipoPagoDestino: '%s'", raw)

}

func (b TipoPagoDestino) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *TipoPagoDestino) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewTipoPagoDestinoValue(stringVal)
	*b = val
	return err
}

type TipoPagoDestinoWrapper struct {
	Target *TipoPagoDestino
}

func (b TipoPagoDestinoWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b TipoPagoDestinoWrapper) SetInt(v int32) {
	*(b.Target) = TipoPagoDestino(v)
}

func (b TipoPagoDestinoWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b TipoPagoDestinoWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b TipoPagoDestinoWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b TipoPagoDestinoWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b TipoPagoDestinoWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b TipoPagoDestinoWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b TipoPagoDestinoWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b TipoPagoDestinoWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b TipoPagoDestinoWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b TipoPagoDestinoWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b TipoPagoDestinoWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b TipoPagoDestinoWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b TipoPagoDestinoWrapper) Finalize() {}
