// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Visita.avsc
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

type TipoDeTelefono int32

const (
	TipoDeTelefonoTrabajo TipoDeTelefono = 0
	TipoDeTelefonoCelular TipoDeTelefono = 1
	TipoDeTelefonoCasa    TipoDeTelefono = 2
	TipoDeTelefonoOtro    TipoDeTelefono = 3
)

func (e TipoDeTelefono) String() string {
	switch e {
	case TipoDeTelefonoTrabajo:
		return "trabajo"
	case TipoDeTelefonoCelular:
		return "celular"
	case TipoDeTelefonoCasa:
		return "casa"
	case TipoDeTelefonoOtro:
		return "otro"
	}
	return "unknown"
}

func writeTipoDeTelefono(r TipoDeTelefono, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewTipoDeTelefonoValue(raw string) (r TipoDeTelefono, err error) {
	switch raw {
	case "trabajo":
		return TipoDeTelefonoTrabajo, nil
	case "celular":
		return TipoDeTelefonoCelular, nil
	case "casa":
		return TipoDeTelefonoCasa, nil
	case "otro":
		return TipoDeTelefonoOtro, nil
	}

	return -1, fmt.Errorf("invalid value for TipoDeTelefono: '%s'", raw)

}

func (b TipoDeTelefono) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *TipoDeTelefono) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewTipoDeTelefonoValue(stringVal)
	*b = val
	return err
}

type TipoDeTelefonoWrapper struct {
	Target *TipoDeTelefono
}

func (b TipoDeTelefonoWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b TipoDeTelefonoWrapper) SetInt(v int32) {
	*(b.Target) = TipoDeTelefono(v)
}

func (b TipoDeTelefonoWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b TipoDeTelefonoWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b TipoDeTelefonoWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b TipoDeTelefonoWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b TipoDeTelefonoWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b TipoDeTelefonoWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b TipoDeTelefonoWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b TipoDeTelefonoWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b TipoDeTelefonoWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b TipoDeTelefonoWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b TipoDeTelefonoWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b TipoDeTelefonoWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b TipoDeTelefonoWrapper) Finalize() {}
