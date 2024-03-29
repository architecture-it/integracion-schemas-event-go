// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ReembolsoRendidoACentral.avsc
 */
package TrazasEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ComponenteDeCobro struct {
	Forma FormaDeCobroV2 `json:"forma"`

	Banco *UnionNullBanco `json:"banco"`

	Monto float64 `json:"monto"`

	NumeroDocumento *UnionNullString `json:"numeroDocumento"`

	Fecha *UnionNullString `json:"fecha"`
}

const ComponenteDeCobroAvroCRC64Fingerprint = "\x96)=\x91\"6\"p"

func NewComponenteDeCobro() ComponenteDeCobro {
	r := ComponenteDeCobro{}
	r.Banco = nil
	r.NumeroDocumento = nil
	r.Fecha = nil
	return r
}

func DeserializeComponenteDeCobro(r io.Reader) (ComponenteDeCobro, error) {
	t := NewComponenteDeCobro()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeComponenteDeCobroFromSchema(r io.Reader, schema string) (ComponenteDeCobro, error) {
	t := NewComponenteDeCobro()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeComponenteDeCobro(r ComponenteDeCobro, w io.Writer) error {
	var err error
	err = writeFormaDeCobroV2(r.Forma, w)
	if err != nil {
		return err
	}
	err = writeUnionNullBanco(r.Banco, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Monto, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.NumeroDocumento, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Fecha, w)
	if err != nil {
		return err
	}
	return err
}

func (r ComponenteDeCobro) Serialize(w io.Writer) error {
	return writeComponenteDeCobro(r, w)
}

func (r ComponenteDeCobro) Schema() string {
	return "{\"fields\":[{\"name\":\"forma\",\"type\":{\"name\":\"FormaDeCobroV2\",\"symbols\":[\"chequeDeTerceros\",\"boletaDeDepositoEnCtaCliente\",\"ajusteDeSaldo\",\"retencionesImpositivas\",\"chequeNominativo\",\"descuentoFinanciero\",\"notaCredito\",\"efectivo\",\"planillaRendicionExcCasa\",\"enviadoEnPManualAnteriormente\",\"ajusteImporteAnterior\",\"notaCreditoPami\",\"notaCreditoAConfirmarDevolucion\",\"enviaDiferenciaEnProximaLiquidacion\",\"pagoEfectivoCobrado\",\"pagoEfectivoChequePropio\",\"descuentoDevolucionMercaderia\",\"comprobanteTransferenciaACliente\",\"impuestoAlCheque\",\"saldoAFavor\",\"agregaPagoOtraFactura\",\"enviadoEnLiquidacionAnterior\",\"cartaCompromisoExcCasa\",\"visaDebitoCredito\",\"autorizacionAdjunta\",\"reciboOficialAndreani\",\"descuentoNotaAdjunta\",\"pagoFacilComprobante\",\"chequeElectronico\",\"mercadoPago\",\"ajustePorPendiente\",\"cbu\",\"cajaUnificadaExcCasa\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"banco\",\"type\":[\"null\",{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"}],\"name\":\"Banco\",\"type\":\"record\"}]},{\"name\":\"monto\",\"type\":\"double\"},{\"default\":null,\"name\":\"numeroDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fecha\",\"type\":[\"null\",\"string\"]}],\"name\":\"Integracion.Esquemas.Referencias.ComponenteDeCobro\",\"type\":\"record\"}"
}

func (r ComponenteDeCobro) SchemaName() string {
	return "Integracion.Esquemas.Referencias.ComponenteDeCobro"
}

func (_ ComponenteDeCobro) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ComponenteDeCobro) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ComponenteDeCobro) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ComponenteDeCobro) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ComponenteDeCobro) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ComponenteDeCobro) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ComponenteDeCobro) SetString(v string)   { panic("Unsupported operation") }
func (_ ComponenteDeCobro) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ComponenteDeCobro) Get(i int) types.Field {
	switch i {
	case 0:
		w := FormaDeCobroV2Wrapper{Target: &r.Forma}

		return w

	case 1:
		r.Banco = NewUnionNullBanco()

		return r.Banco
	case 2:
		w := types.Double{Target: &r.Monto}

		return w

	case 3:
		r.NumeroDocumento = NewUnionNullString()

		return r.NumeroDocumento
	case 4:
		r.Fecha = NewUnionNullString()

		return r.Fecha
	}
	panic("Unknown field index")
}

func (r *ComponenteDeCobro) SetDefault(i int) {
	switch i {
	case 1:
		r.Banco = nil
		return
	case 3:
		r.NumeroDocumento = nil
		return
	case 4:
		r.Fecha = nil
		return
	}
	panic("Unknown field index")
}

func (r *ComponenteDeCobro) NullField(i int) {
	switch i {
	case 1:
		r.Banco = nil
		return
	case 3:
		r.NumeroDocumento = nil
		return
	case 4:
		r.Fecha = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ ComponenteDeCobro) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ComponenteDeCobro) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ComponenteDeCobro) HintSize(int)                     { panic("Unsupported operation") }
func (_ ComponenteDeCobro) Finalize()                        {}

func (_ ComponenteDeCobro) AvroCRC64Fingerprint() []byte {
	return []byte(ComponenteDeCobroAvroCRC64Fingerprint)
}

func (r ComponenteDeCobro) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["forma"], err = json.Marshal(r.Forma)
	if err != nil {
		return nil, err
	}
	output["banco"], err = json.Marshal(r.Banco)
	if err != nil {
		return nil, err
	}
	output["monto"], err = json.Marshal(r.Monto)
	if err != nil {
		return nil, err
	}
	output["numeroDocumento"], err = json.Marshal(r.NumeroDocumento)
	if err != nil {
		return nil, err
	}
	output["fecha"], err = json.Marshal(r.Fecha)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ComponenteDeCobro) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["forma"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Forma); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for forma")
	}
	val = func() json.RawMessage {
		if v, ok := fields["banco"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Banco); err != nil {
			return err
		}
	} else {
		r.Banco = NewUnionNullBanco()

		r.Banco = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["monto"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Monto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for monto")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDocumento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDocumento); err != nil {
			return err
		}
	} else {
		r.NumeroDocumento = NewUnionNullString()

		r.NumeroDocumento = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["fecha"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Fecha); err != nil {
			return err
		}
	} else {
		r.Fecha = NewUnionNullString()

		r.Fecha = nil
	}
	return nil
}
