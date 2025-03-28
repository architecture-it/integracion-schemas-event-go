// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     RecepcionReembolsoEnCentral.avsc
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

type RecepcionReembolsoEnCentral struct {
	Traza Traza `json:"traza"`

	NumeroDeEnvio string `json:"numeroDeEnvio"`

	Fecha string `json:"fecha"`

	ComponentesDeCobro []ComponenteDeCobro `json:"componentesDeCobro"`
}

const RecepcionReembolsoEnCentralAvroCRC64Fingerprint = "\x91\xfa\x14\"aZA/"

func NewRecepcionReembolsoEnCentral() RecepcionReembolsoEnCentral {
	r := RecepcionReembolsoEnCentral{}
	r.Traza = NewTraza()

	r.ComponentesDeCobro = make([]ComponenteDeCobro, 0)

	return r
}

func DeserializeRecepcionReembolsoEnCentral(r io.Reader) (RecepcionReembolsoEnCentral, error) {
	t := NewRecepcionReembolsoEnCentral()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRecepcionReembolsoEnCentralFromSchema(r io.Reader, schema string) (RecepcionReembolsoEnCentral, error) {
	t := NewRecepcionReembolsoEnCentral()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRecepcionReembolsoEnCentral(r RecepcionReembolsoEnCentral, w io.Writer) error {
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
	err = writeArrayComponenteDeCobro(r.ComponentesDeCobro, w)
	if err != nil {
		return err
	}
	return err
}

func (r RecepcionReembolsoEnCentral) Serialize(w io.Writer) error {
	return writeRecepcionReembolsoEnCentral(r, w)
}

func (r RecepcionReembolsoEnCentral) Schema() string {
	return "{\"fields\":[{\"name\":\"traza\",\"type\":{\"fields\":[{\"name\":\"codigoDeEnvio\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"cuando\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"codigoDeContratoInterno\",\"type\":\"string\"},{\"default\":null,\"name\":\"codigoDeContrato\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"cicloDelEnvio\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"operador\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"estadoDeLaRendicion\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"comentario\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"sucursalAsociadaAlEvento\",\"type\":[\"null\",{\"fields\":[{\"name\":\"codigo\",\"type\":\"string\"},{\"default\":null,\"name\":\"nombre\",\"type\":[\"null\",\"string\"]},{\"name\":\"id\",\"type\":\"string\"}],\"name\":\"DatosSucursal\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"}]}],\"name\":\"Traza\",\"namespace\":\"Integracion.Esquemas\",\"type\":\"record\"}},{\"name\":\"numeroDeEnvio\",\"type\":\"string\"},{\"name\":\"fecha\",\"type\":\"string\"},{\"name\":\"componentesDeCobro\",\"type\":{\"items\":{\"fields\":[{\"name\":\"forma\",\"type\":{\"name\":\"FormaDeCobroV2\",\"symbols\":[\"chequeDeTerceros\",\"boletaDeDepositoEnCtaCliente\",\"ajusteDeSaldo\",\"retencionesImpositivas\",\"chequeNominativo\",\"descuentoFinanciero\",\"notaCredito\",\"efectivo\",\"planillaRendicionExcCasa\",\"enviadoEnPManualAnteriormente\",\"ajusteImporteAnterior\",\"notaCreditoPami\",\"notaCreditoAConfirmarDevolucion\",\"enviaDiferenciaEnProximaLiquidacion\",\"pagoEfectivoCobrado\",\"pagoEfectivoChequePropio\",\"descuentoDevolucionMercaderia\",\"comprobanteTransferenciaACliente\",\"impuestoAlCheque\",\"saldoAFavor\",\"agregaPagoOtraFactura\",\"enviadoEnLiquidacionAnterior\",\"cartaCompromisoExcCasa\",\"visaDebitoCredito\",\"autorizacionAdjunta\",\"reciboOficialAndreani\",\"descuentoNotaAdjunta\",\"pagoFacilComprobante\",\"chequeElectronico\",\"mercadoPago\",\"ajustePorPendiente\",\"cbu\",\"cajaUnificadaExcCasa\"],\"type\":\"enum\"}},{\"default\":null,\"name\":\"banco\",\"type\":[\"null\",{\"fields\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"descripcion\",\"type\":\"string\"}],\"name\":\"Banco\",\"type\":\"record\"}]},{\"name\":\"monto\",\"type\":\"double\"},{\"default\":null,\"name\":\"numeroDocumento\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"fecha\",\"type\":[\"null\",\"string\"]}],\"name\":\"ComponenteDeCobro\",\"namespace\":\"Integracion.Esquemas.Referencias\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Integracion.Esquemas.Trazas.RecepcionReembolsoEnCentral\",\"type\":\"record\"}"
}

func (r RecepcionReembolsoEnCentral) SchemaName() string {
	return "Integracion.Esquemas.Trazas.RecepcionReembolsoEnCentral"
}

func (_ RecepcionReembolsoEnCentral) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RecepcionReembolsoEnCentral) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RecepcionReembolsoEnCentral) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RecepcionReembolsoEnCentral) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RecepcionReembolsoEnCentral) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RecepcionReembolsoEnCentral) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RecepcionReembolsoEnCentral) SetString(v string)   { panic("Unsupported operation") }
func (_ RecepcionReembolsoEnCentral) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RecepcionReembolsoEnCentral) Get(i int) types.Field {
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
		r.ComponentesDeCobro = make([]ComponenteDeCobro, 0)

		w := ArrayComponenteDeCobroWrapper{Target: &r.ComponentesDeCobro}

		return w

	}
	panic("Unknown field index")
}

func (r *RecepcionReembolsoEnCentral) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *RecepcionReembolsoEnCentral) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ RecepcionReembolsoEnCentral) AppendMap(key string) types.Field {
	panic("Unsupported operation")
}
func (_ RecepcionReembolsoEnCentral) AppendArray() types.Field { panic("Unsupported operation") }
func (_ RecepcionReembolsoEnCentral) HintSize(int)             { panic("Unsupported operation") }
func (_ RecepcionReembolsoEnCentral) Finalize()                {}

func (_ RecepcionReembolsoEnCentral) AvroCRC64Fingerprint() []byte {
	return []byte(RecepcionReembolsoEnCentralAvroCRC64Fingerprint)
}

func (r RecepcionReembolsoEnCentral) MarshalJSON() ([]byte, error) {
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
	output["componentesDeCobro"], err = json.Marshal(r.ComponentesDeCobro)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RecepcionReembolsoEnCentral) UnmarshalJSON(data []byte) error {
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
		if v, ok := fields["componentesDeCobro"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ComponentesDeCobro); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for componentesDeCobro")
	}
	return nil
}
