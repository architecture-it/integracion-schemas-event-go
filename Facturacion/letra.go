// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     VueltaNroLegal.avsc
 */
package FacturacionEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Letra int32

const (
	LetraA         Letra = 0
	LetraB         Letra = 1
	LetraUndefined Letra = 2
	LetraAll       Letra = 3
)

func (e Letra) String() string {
	switch e {
	case LetraA:
		return "A"
	case LetraB:
		return "B"
	case LetraUndefined:
		return "undefined"
	case LetraAll:
		return "all"
	}
	return "unknown"
}

func writeLetra(r Letra, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewLetraValue(raw string) (r Letra, err error) {
	switch raw {
	case "A":
		return LetraA, nil
	case "B":
		return LetraB, nil
	case "undefined":
		return LetraUndefined, nil
	case "all":
		return LetraAll, nil
	}

	return -1, fmt.Errorf("invalid value for Letra: '%s'", raw)

}

func (b Letra) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *Letra) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewLetraValue(stringVal)
	*b = val
	return err
}

type LetraWrapper struct {
	Target *Letra
}

func (b LetraWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b LetraWrapper) SetInt(v int32) {
	*(b.Target) = Letra(v)
}

func (b LetraWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b LetraWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b LetraWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b LetraWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b LetraWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b LetraWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b LetraWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b LetraWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b LetraWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b LetraWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b LetraWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b LetraWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b LetraWrapper) Finalize() {}
