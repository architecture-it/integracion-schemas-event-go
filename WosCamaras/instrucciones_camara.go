// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     InstruccionesCamara.avsc
 */
package WosCamarasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type InstruccionesCamara int32

const (
	InstruccionesCamaraIniciarGrabacion InstruccionesCamara = 0
	InstruccionesCamaraDetenerGrabacion InstruccionesCamara = 1
)

func (e InstruccionesCamara) String() string {
	switch e {
	case InstruccionesCamaraIniciarGrabacion:
		return "IniciarGrabacion"
	case InstruccionesCamaraDetenerGrabacion:
		return "DetenerGrabacion"
	}
	return "unknown"
}

func writeInstruccionesCamara(r InstruccionesCamara, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewInstruccionesCamaraValue(raw string) (r InstruccionesCamara, err error) {
	switch raw {
	case "IniciarGrabacion":
		return InstruccionesCamaraIniciarGrabacion, nil
	case "DetenerGrabacion":
		return InstruccionesCamaraDetenerGrabacion, nil
	}

	return -1, fmt.Errorf("invalid value for InstruccionesCamara: '%s'", raw)

}

func (b InstruccionesCamara) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *InstruccionesCamara) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewInstruccionesCamaraValue(stringVal)
	*b = val
	return err
}

type InstruccionesCamaraWrapper struct {
	Target *InstruccionesCamara
}

func (b InstruccionesCamaraWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b InstruccionesCamaraWrapper) SetInt(v int32) {
	*(b.Target) = InstruccionesCamara(v)
}

func (b InstruccionesCamaraWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b InstruccionesCamaraWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b InstruccionesCamaraWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b InstruccionesCamaraWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b InstruccionesCamaraWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b InstruccionesCamaraWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b InstruccionesCamaraWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b InstruccionesCamaraWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b InstruccionesCamaraWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b InstruccionesCamaraWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b InstruccionesCamaraWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b InstruccionesCamaraWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b InstruccionesCamaraWrapper) Finalize() {}