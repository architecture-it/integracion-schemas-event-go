// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ContabilizacionCobroReembolso.avsc
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

type ContabilizacionCobroReembolso struct {
	Traza Traza `json:"traza"`

	NumeroDeEnvio string `json:"numeroDeEnvio"`

	Fecha string `json:"fecha"`

	ComponenteDeCobro ComponenteDeCobro `json:"componenteDeCobro"`
}

const ContabilizacionCobroReembolsoAvroCRC64Fingerprint = "\xd9\xe7\x146\xeaYv\x90"

func NewContabilizacionCobroReembolso() ContabilizacionCobroReembolso {
	r := ContabilizacionCobroReembolso{}
	r.Traza = NewTraza()

	r.ComponenteDeCobro = NewComponenteDeCobro()

	return r
}

func DeserializeContabilizacionCobroReembolso(r io.Reader) (ContabilizacionCobroReembolso, error) {
	t := NewContabilizacionCobroReembolso()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeContabilizacionCobroReembolsoFromSchema(r io.Reader, schema string) (ContabilizacionCobroReembolso, error) {
	t := NewContabilizacionCobroReembolso()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeContabilizacionCobroReembolso(r ContabilizacionCobroReembolso, w io.Writer) error {
	var err error
	err = writeTraza(r.Traza, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NumeroDeEnvio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Fecha, w)
	if err != nil {
		return err
	}
	err = writeComponenteDeCobro(r.ComponenteDeCobro, w)
	if err != nil {
		return err
	}
	return err
}

func (r ContabilizacionCobroReembolso) Serialize(w io.Writer) error {
	return writeContabilizacionCobroReembolso(r, w)
}

func (r ContabilizacionCobroReembolso) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"numeroDeEnvio\",\"type\":\"string\"},{\"name\":\"fecha\",\"type\":\"string\"},{\"name\":\"componenteDeCobro\",\"type\":{\"fields\":[{\"name\":\"forma\",\"type\":{\"name\":\"FormaDeCobroV2\",\"symbols\":[\"chequeDeTerceros\",\"boletaDeDepositoEnCtaCliente\",\"ajusteDeSaldo\",\"retencionesImpositivas\",\"chequeNominativo\",\"descuentoFinanciero\",\"notaCredito\",\"efectivo\",\"planillaRendicionExcCasa\",\"enviadoEnPManualAnteriormente\",\"ajusteImporteAnterior\",\"notaCreditoPami\",\"notaCreditoAConfirmarDevolucion\",\"enviaDiferenciaEnProximaLiquidacion\",\"pagoEfectivoCobrado\",\"pagoEfectivoChequePropio\",\"descuentoDevolucionMercaderia\",\"comprobanteTransferenciaACliente\",\"impuestoAlCheque\",\"saldoAFavor\",\"agregaPagoOtraFactura\",\"enviadoEnLiquidacionAnterior\",\"cartaCompromisoExcCasa\",\"visaDebitoCredito\",\"autorizacionAdjunta\",\"reciboOficialAndreani\",\"descuentoNotaAdjunta\",\"pagoFacilComprobante\",\"chequeElectronico\",\"mercadoPago\",\"ajustePorPendiente\",\"cbu\",\"cajaUnificadaExcCasa\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"banco\",\"type\":[\"null\",{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"}],\"name\":\"Banco\",\"type\":\"record\"}]},{\"name\":\"monto\",\"type\":\"double\"},{\"default\":null,\"name\":\"numeroDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fecha\",\"type\":[\"null\",\"string\"]}],\"name\":\"ComponenteDeCobro\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}}],\"name\":\"Integracion.Esquemas.Trazas.ContabilizacionCobroReembolso\",\"type\":\"record\"}"
}

func (r ContabilizacionCobroReembolso) SchemaName() string {
	return "Integracion.Esquemas.Trazas.ContabilizacionCobroReembolso"
}

func (_ ContabilizacionCobroReembolso) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ContabilizacionCobroReembolso) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ContabilizacionCobroReembolso) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ContabilizacionCobroReembolso) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ContabilizacionCobroReembolso) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ContabilizacionCobroReembolso) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ContabilizacionCobroReembolso) SetString(v string)   { panic("Unsupported operation") }
func (_ ContabilizacionCobroReembolso) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ContabilizacionCobroReembolso) Get(i int) types.Field {
	switch i {
	case 0:
		r.Traza = NewTraza()

		w := types.Record{Target: &r.Traza}

		return w

	case 1:
		w := types.String{Target: &r.NumeroDeEnvio}

		return w

	case 2:
		w := types.String{Target: &r.Fecha}

		return w

	case 3:
		r.ComponenteDeCobro = NewComponenteDeCobro()

		w := types.Record{Target: &r.ComponenteDeCobro}

		return w

	}
	panic("Unknown field index")
}

func (r *ContabilizacionCobroReembolso) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ContabilizacionCobroReembolso) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ContabilizacionCobroReembolso) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ ContabilizacionCobroReembolso) AppendArray() types.Field { panic("Unsupported operation") }
func (_ ContabilizacionCobroReembolso) HintSize(int)             { panic("Unsupported operation") }
func (_ ContabilizacionCobroReembolso) Finalize()                {}

func (_ ContabilizacionCobroReembolso) AvroCRC64Fingerprint() []byte {
	return []byte(ContabilizacionCobroReembolsoAvroCRC64Fingerprint)
}

func (r ContabilizacionCobroReembolso) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["traza"], err = json.Marshal(r.Traza)
	if err != nil {
		return nil, err
	}
	output["numeroDeEnvio"], err = json.Marshal(r.NumeroDeEnvio)
	if err != nil {
		return nil, err
	}
	output["fecha"], err = json.Marshal(r.Fecha)
	if err != nil {
		return nil, err
	}
	output["componenteDeCobro"], err = json.Marshal(r.ComponenteDeCobro)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ContabilizacionCobroReembolso) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["traza"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Traza); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for traza")
	}
	val = func() json.RawMessage {
		if v, ok := fields["numeroDeEnvio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NumeroDeEnvio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for numeroDeEnvio")
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
		return fmt.Errorf("no value specified for fecha")
	}
	val = func() json.RawMessage {
		if v, ok := fields["componenteDeCobro"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ComponenteDeCobro); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for componenteDeCobro")
	}
	return nil
}
