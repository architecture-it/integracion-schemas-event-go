// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     DetallePreenvio.avsc
 */
package CalidadCertificadaEvents

import (
	"io"

	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

func writeArrayPersona(r []Persona, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writePersona(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayPersonaWrapper struct {
	Target *[]Persona
}

func (_ ArrayPersonaWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayPersonaWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayPersonaWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayPersonaWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayPersonaWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayPersonaWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayPersonaWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayPersonaWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayPersonaWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayPersonaWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayPersonaWrapper) Finalize()                        {}
func (_ ArrayPersonaWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayPersonaWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Persona, 0, s)
	}
}
func (r ArrayPersonaWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r ArrayPersonaWrapper) AppendArray() types.Field {
	var v Persona
	v = NewPersona()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
