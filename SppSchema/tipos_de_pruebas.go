// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     TiposDePruebas.avsc
 */
package SppSchemaEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type TiposDePruebas int32

const (
	TiposDePruebasALSA TiposDePruebas = 0
	TiposDePruebasCASA TiposDePruebas = 1
	TiposDePruebasDMS  TiposDePruebas = 2
)

func (e TiposDePruebas) String() string {
	switch e {
	case TiposDePruebasALSA:
		return "ALSA"
	case TiposDePruebasCASA:
		return "CASA"
	case TiposDePruebasDMS:
		return "DMS"
	}
	return "unknown"
}

func writeTiposDePruebas(r TiposDePruebas, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewTiposDePruebasValue(raw string) (r TiposDePruebas, err error) {
	switch raw {
	case "ALSA":
		return TiposDePruebasALSA, nil
	case "CASA":
		return TiposDePruebasCASA, nil
	case "DMS":
		return TiposDePruebasDMS, nil
	}

	return -1, fmt.Errorf("invalid value for TiposDePruebas: '%s'", raw)

}

func (b TiposDePruebas) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *TiposDePruebas) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewTiposDePruebasValue(stringVal)
	*b = val
	return err
}

type TiposDePruebasWrapper struct {
	Target *TiposDePruebas
}

func (b TiposDePruebasWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b TiposDePruebasWrapper) SetInt(v int32) {
	*(b.Target) = TiposDePruebas(v)
}

func (b TiposDePruebasWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b TiposDePruebasWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b TiposDePruebasWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b TiposDePruebasWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b TiposDePruebasWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b TiposDePruebasWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b TiposDePruebasWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b TiposDePruebasWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b TiposDePruebasWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b TiposDePruebasWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b TiposDePruebasWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b TiposDePruebasWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b TiposDePruebasWrapper) Finalize() {}
