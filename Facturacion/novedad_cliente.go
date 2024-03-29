// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     NovedadCliente.avsc
 */
package FacturacionEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type NovedadCliente struct {
	CodigoDeClienteInterno string `json:"CodigoDeClienteInterno"`

	Canal string `json:"Canal"`

	DepartamentoDeVentas string `json:"DepartamentoDeVentas"`

	CodigoDeDireccion string `json:"CodigoDeDireccion"`

	EjecutivoDeCuentas string `json:"EjecutivoDeCuentas"`

	EstaActivo int32 `json:"EstaActivo"`

	GerenteDeCuentas string `json:"GerenteDeCuentas"`

	RazonSocial string `json:"RazonSocial"`

	Segmento string `json:"Segmento"`

	TipoDeEntidadFiscal int32 `json:"TipoDeEntidadFiscal"`

	EjecutivoDeCuentasLogin string `json:"EjecutivoDeCuentasLogin"`

	NombreDeFantasia string `json:"NombreDeFantasia"`

	Fecha *UnionNullLong `json:"Fecha"`
}

const NovedadClienteAvroCRC64Fingerprint = "BD\xb5\xff\x8c&\x17\xf5"

func NewNovedadCliente() NovedadCliente {
	r := NovedadCliente{}
	return r
}

func DeserializeNovedadCliente(r io.Reader) (NovedadCliente, error) {
	t := NewNovedadCliente()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNovedadClienteFromSchema(r io.Reader, schema string) (NovedadCliente, error) {
	t := NewNovedadCliente()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNovedadCliente(r NovedadCliente, w io.Writer) error {
	var err error
	err = vm.WriteString(r.CodigoDeClienteInterno, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Canal, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.DepartamentoDeVentas, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodigoDeDireccion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EjecutivoDeCuentas, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.EstaActivo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.GerenteDeCuentas, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.RazonSocial, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Segmento, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.TipoDeEntidadFiscal, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EjecutivoDeCuentasLogin, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NombreDeFantasia, w)
	if err != nil {
		return err
	}
	err = writeUnionNullLong(r.Fecha, w)
	if err != nil {
		return err
	}
	return err
}

func (r NovedadCliente) Serialize(w io.Writer) error {
	return writeNovedadCliente(r, w)
}

func (r NovedadCliente) Schema() string {
	return "{\"fields\":[{\"name\":\"CodigoDeClienteInterno\",\"type\":\"string\"},{\"name\":\"Canal\",\"type\":\"string\"},{\"name\":\"DepartamentoDeVentas\",\"type\":\"string\"},{\"name\":\"CodigoDeDireccion\",\"type\":\"string\"},{\"name\":\"EjecutivoDeCuentas\",\"type\":\"string\"},{\"name\":\"EstaActivo\",\"type\":\"int\"},{\"name\":\"GerenteDeCuentas\",\"type\":\"string\"},{\"name\":\"RazonSocial\",\"type\":\"string\"},{\"name\":\"Segmento\",\"type\":\"string\"},{\"name\":\"TipoDeEntidadFiscal\",\"type\":\"int\"},{\"name\":\"EjecutivoDeCuentasLogin\",\"type\":\"string\"},{\"name\":\"NombreDeFantasia\",\"type\":\"string\"},{\"name\":\"Fecha\",\"type\":[\"null\",{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}]}],\"name\":\"Andreani.Facturacion.Events.Record.NovedadCliente\",\"type\":\"record\"}"
}

func (r NovedadCliente) SchemaName() string {
	return "Andreani.Facturacion.Events.Record.NovedadCliente"
}

func (_ NovedadCliente) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NovedadCliente) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NovedadCliente) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NovedadCliente) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NovedadCliente) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NovedadCliente) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NovedadCliente) SetString(v string)   { panic("Unsupported operation") }
func (_ NovedadCliente) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NovedadCliente) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.CodigoDeClienteInterno}

		return w

	case 1:
		w := types.String{Target: &r.Canal}

		return w

	case 2:
		w := types.String{Target: &r.DepartamentoDeVentas}

		return w

	case 3:
		w := types.String{Target: &r.CodigoDeDireccion}

		return w

	case 4:
		w := types.String{Target: &r.EjecutivoDeCuentas}

		return w

	case 5:
		w := types.Int{Target: &r.EstaActivo}

		return w

	case 6:
		w := types.String{Target: &r.GerenteDeCuentas}

		return w

	case 7:
		w := types.String{Target: &r.RazonSocial}

		return w

	case 8:
		w := types.String{Target: &r.Segmento}

		return w

	case 9:
		w := types.Int{Target: &r.TipoDeEntidadFiscal}

		return w

	case 10:
		w := types.String{Target: &r.EjecutivoDeCuentasLogin}

		return w

	case 11:
		w := types.String{Target: &r.NombreDeFantasia}

		return w

	case 12:
		r.Fecha = NewUnionNullLong()

		return r.Fecha
	}
	panic("Unknown field index")
}

func (r *NovedadCliente) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NovedadCliente) NullField(i int) {
	switch i {
	case 12:
		r.Fecha = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ NovedadCliente) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NovedadCliente) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NovedadCliente) HintSize(int)                     { panic("Unsupported operation") }
func (_ NovedadCliente) Finalize()                        {}

func (_ NovedadCliente) AvroCRC64Fingerprint() []byte {
	return []byte(NovedadClienteAvroCRC64Fingerprint)
}

func (r NovedadCliente) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["CodigoDeClienteInterno"], err = json.Marshal(r.CodigoDeClienteInterno)
	if err != nil {
		return nil, err
	}
	output["Canal"], err = json.Marshal(r.Canal)
	if err != nil {
		return nil, err
	}
	output["DepartamentoDeVentas"], err = json.Marshal(r.DepartamentoDeVentas)
	if err != nil {
		return nil, err
	}
	output["CodigoDeDireccion"], err = json.Marshal(r.CodigoDeDireccion)
	if err != nil {
		return nil, err
	}
	output["EjecutivoDeCuentas"], err = json.Marshal(r.EjecutivoDeCuentas)
	if err != nil {
		return nil, err
	}
	output["EstaActivo"], err = json.Marshal(r.EstaActivo)
	if err != nil {
		return nil, err
	}
	output["GerenteDeCuentas"], err = json.Marshal(r.GerenteDeCuentas)
	if err != nil {
		return nil, err
	}
	output["RazonSocial"], err = json.Marshal(r.RazonSocial)
	if err != nil {
		return nil, err
	}
	output["Segmento"], err = json.Marshal(r.Segmento)
	if err != nil {
		return nil, err
	}
	output["TipoDeEntidadFiscal"], err = json.Marshal(r.TipoDeEntidadFiscal)
	if err != nil {
		return nil, err
	}
	output["EjecutivoDeCuentasLogin"], err = json.Marshal(r.EjecutivoDeCuentasLogin)
	if err != nil {
		return nil, err
	}
	output["NombreDeFantasia"], err = json.Marshal(r.NombreDeFantasia)
	if err != nil {
		return nil, err
	}
	output["Fecha"], err = json.Marshal(r.Fecha)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NovedadCliente) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
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
		if v, ok := fields["Canal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Canal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Canal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DepartamentoDeVentas"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DepartamentoDeVentas); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DepartamentoDeVentas")
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
		if v, ok := fields["EjecutivoDeCuentas"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EjecutivoDeCuentas); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EjecutivoDeCuentas")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EstaActivo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstaActivo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EstaActivo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["GerenteDeCuentas"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.GerenteDeCuentas); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for GerenteDeCuentas")
	}
	val = func() json.RawMessage {
		if v, ok := fields["RazonSocial"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.RazonSocial); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for RazonSocial")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Segmento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Segmento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Segmento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDeEntidadFiscal"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDeEntidadFiscal); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoDeEntidadFiscal")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EjecutivoDeCuentasLogin"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EjecutivoDeCuentasLogin); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EjecutivoDeCuentasLogin")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NombreDeFantasia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NombreDeFantasia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NombreDeFantasia")
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
	return nil
}
