// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NovedadContrato.avsc
 */
package FacturacionClienteEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NovedadContrato struct {
	SistemaDestino string `json:"SistemaDestino"`

	CodigoDeContratoInterno string `json:"CodigoDeContratoInterno"`

	Descripcion string `json:"Descripcion"`

	TipoDeServicio string `json:"TipoDeServicio"`

	CodigoDeDireccion string `json:"CodigoDeDireccion"`

	CodigoDeClienteInterno string `json:"CodigoDeClienteInterno"`

	Fecha int64 `json:"Fecha"`

	VigenciaDesde int64 `json:"VigenciaDesde"`

	VigenciaHasta int64 `json:"VigenciaHasta"`
}

const NovedadContratoAvroCRC64Fingerprint = "\x9f\x99(w\xbc\f^F"

func NewNovedadContrato() NovedadContrato {
	r := NovedadContrato{}
	return r
}

func DeserializeNovedadContrato(r io.Reader) (NovedadContrato, error) {
	t := NewNovedadContrato()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNovedadContratoFromSchema(r io.Reader, schema string) (NovedadContrato, error) {
	t := NewNovedadContrato()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNovedadContrato(r NovedadContrato, w io.Writer) error {
	var err error
	err = vm.WriteString(r.SistemaDestino, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoDeContratoInterno, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Descripcion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoDeServicio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoDeDireccion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoDeClienteInterno, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Fecha, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.VigenciaDesde, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.VigenciaHasta, w)
	if err != nil {
		return err
	}
	return err
}

func (r NovedadContrato) Serialize(w io.Writer) error {
	return writeNovedadContrato(r, w)
}

func (r NovedadContrato) Schema() string {
	return "{\"fields\":[{\"name\":\"SistemaDestino\",\"type\":\"string\"},{\"name\":\"CodigoDeContratoInterno\",\"type\":\"string\"},{\"name\":\"Descripcion\",\"type\":\"string\"},{\"name\":\"TipoDeServicio\",\"type\":\"string\"},{\"name\":\"CodigoDeDireccion\",\"type\":\"string\"},{\"name\":\"CodigoDeClienteInterno\",\"type\":\"string\"},{\"name\":\"Fecha\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"VigenciaDesde\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"VigenciaHasta\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"Andreani.FacturacionCliente.Events.Record.NovedadContrato\",\"type\":\"record\"}"
}

func (r NovedadContrato) SchemaName() string {
	return "Andreani.FacturacionCliente.Events.Record.NovedadContrato"
}

func (_ NovedadContrato) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NovedadContrato) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NovedadContrato) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NovedadContrato) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NovedadContrato) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NovedadContrato) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NovedadContrato) SetString(v string)   { panic("Unsupported operation") }
func (_ NovedadContrato) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NovedadContrato) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.SistemaDestino}

		return w

	case 1:
		w := types.String{Target: &r.CodigoDeContratoInterno}

		return w

	case 2:
		w := types.String{Target: &r.Descripcion}

		return w

	case 3:
		w := types.String{Target: &r.TipoDeServicio}

		return w

	case 4:
		w := types.String{Target: &r.CodigoDeDireccion}

		return w

	case 5:
		w := types.String{Target: &r.CodigoDeClienteInterno}

		return w

	case 6:
		w := types.Long{Target: &r.Fecha}

		return w

	case 7:
		w := types.Long{Target: &r.VigenciaDesde}

		return w

	case 8:
		w := types.Long{Target: &r.VigenciaHasta}

		return w

	}
	panic("Unknown field index")
}

func (r *NovedadContrato) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NovedadContrato) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ NovedadContrato) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NovedadContrato) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NovedadContrato) HintSize(int)                     { panic("Unsupported operation") }
func (_ NovedadContrato) Finalize()                        {}

func (_ NovedadContrato) AvroCRC64Fingerprint() []byte {
	return []byte(NovedadContratoAvroCRC64Fingerprint)
}

func (r NovedadContrato) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["SistemaDestino"], err = json.Marshal(r.SistemaDestino)
	if err != nil {
		return nil, err
	}
	output["CodigoDeContratoInterno"], err = json.Marshal(r.CodigoDeContratoInterno)
	if err != nil {
		return nil, err
	}
	output["Descripcion"], err = json.Marshal(r.Descripcion)
	if err != nil {
		return nil, err
	}
	output["TipoDeServicio"], err = json.Marshal(r.TipoDeServicio)
	if err != nil {
		return nil, err
	}
	output["CodigoDeDireccion"], err = json.Marshal(r.CodigoDeDireccion)
	if err != nil {
		return nil, err
	}
	output["CodigoDeClienteInterno"], err = json.Marshal(r.CodigoDeClienteInterno)
	if err != nil {
		return nil, err
	}
	output["Fecha"], err = json.Marshal(r.Fecha)
	if err != nil {
		return nil, err
	}
	output["VigenciaDesde"], err = json.Marshal(r.VigenciaDesde)
	if err != nil {
		return nil, err
	}
	output["VigenciaHasta"], err = json.Marshal(r.VigenciaHasta)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NovedadContrato) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["SistemaDestino"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.SistemaDestino); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SistemaDestino")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoDeContratoInterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeContratoInterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoDeContratoInterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Descripcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Descripcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Descripcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeServicio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeServicio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoDeServicio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoDeDireccion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeDireccion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoDeDireccion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodigoDeClienteInterno"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodigoDeClienteInterno); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodigoDeClienteInterno")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Fecha"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Fecha); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Fecha")
	}
	val = func() json.RawMessage {
		if v, ok := fields["VigenciaDesde"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VigenciaDesde); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for VigenciaDesde")
	}
	val = func() json.RawMessage {
		if v, ok := fields["VigenciaHasta"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.VigenciaHasta); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for VigenciaHasta")
	}
	return nil
}
