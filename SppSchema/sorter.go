// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     Sorter.avsc
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

type Sorter int32

const (
	SorterMINI       Sorter = 0
	SorterVERTICAL   Sorter = 1
	SorterTRUXSORTER Sorter = 2
	SorterHORIZONTAL Sorter = 3
	SorterOTRO       Sorter = 4
)

func (e Sorter) String() string {
	switch e {
	case SorterMINI:
		return "MINI"
	case SorterVERTICAL:
		return "VERTICAL"
	case SorterTRUXSORTER:
		return "TRUXSORTER"
	case SorterHORIZONTAL:
		return "HORIZONTAL"
	case SorterOTRO:
		return "OTRO"
	}
	return "unknown"
}

func writeSorter(r Sorter, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewSorterValue(raw string) (r Sorter, err error) {
	switch raw {
	case "MINI":
		return SorterMINI, nil
	case "VERTICAL":
		return SorterVERTICAL, nil
	case "TRUXSORTER":
		return SorterTRUXSORTER, nil
	case "HORIZONTAL":
		return SorterHORIZONTAL, nil
	case "OTRO":
		return SorterOTRO, nil
	}

	return -1, fmt.Errorf("invalid value for Sorter: '%s'", raw)

}

func (b Sorter) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *Sorter) UnmarshalJSON(data []byte) error {
	var stringVal string
	err := json.Unmarshal(data, &stringVal)
	if err != nil {
		return err
	}
	val, err := NewSorterValue(stringVal)
	*b = val
	return err
}

type SorterWrapper struct {
	Target *Sorter
}

func (b SorterWrapper) SetBoolean(v bool) {
	panic("Unable to assign boolean to int field")
}

func (b SorterWrapper) SetInt(v int32) {
	*(b.Target) = Sorter(v)
}

func (b SorterWrapper) SetLong(v int64) {
	panic("Unable to assign long to int field")
}

func (b SorterWrapper) SetFloat(v float32) {
	panic("Unable to assign float to int field")
}

func (b SorterWrapper) SetUnionElem(v int64) {
	panic("Unable to assign union elem to int field")
}

func (b SorterWrapper) SetDouble(v float64) {
	panic("Unable to assign double to int field")
}

func (b SorterWrapper) SetBytes(v []byte) {
	panic("Unable to assign bytes to int field")
}

func (b SorterWrapper) SetString(v string) {
	panic("Unable to assign string to int field")
}

func (b SorterWrapper) Get(i int) types.Field {
	panic("Unable to get field from int field")
}

func (b SorterWrapper) SetDefault(i int) {
	panic("Unable to set default on int field")
}

func (b SorterWrapper) AppendMap(key string) types.Field {
	panic("Unable to append map key to from int field")
}

func (b SorterWrapper) AppendArray() types.Field {
	panic("Unable to append array element to from int field")
}

func (b SorterWrapper) NullField(int) {
	panic("Unable to null field in int field")
}

func (b SorterWrapper) HintSize(int) {
	panic("Unable to hint size in int field")
}

func (b SorterWrapper) Finalize() {}