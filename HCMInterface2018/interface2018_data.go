// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListInterface2018.avsc
 */
package HCMInterface2018Events

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Interface2018Data struct {
	LegajoDelEmpleado string `json:"LegajoDelEmpleado"`

	Apellido string `json:"Apellido"`

	Nombres string `json:"Nombres"`

	FechaNacimiento string `json:"FechaNacimiento"`

	PaisDeNacimiento string `json:"PaisDeNacimiento"`

	Nacionalidad string `json:"Nacionalidad"`

	EstadoCivil string `json:"EstadoCivil"`

	Sexo string `json:"Sexo"`

	Parentesco string `json:"Parentesco"`

	Incapacitado string `json:"Incapacitado"`

	FechaDeDiscapacidad string `json:"FechaDeDiscapacidad"`

	TipoDocumento1 string `json:"TipoDocumento1"`

	CodPaisTipoDocumento string `json:"CodPaisTipoDocumento"`

	NroDocumento1 string `json:"NroDocumento1"`

	ObraSocial string `json:"ObraSocial"`

	PlanOS string `json:"PlanOS"`

	AvisarAnteEmergencia string `json:"AvisarAnteEmergencia"`

	PagaSalarioFamiliar string `json:"PagaSalarioFamiliar"`

	FechaInicioVinculo string `json:"FechaInicioVinculo"`

	FechaVto1 string `json:"FechaVto1"`

	Guarderia string `json:"Guarderia"`

	FechaVto2 string `json:"FechaVto2"`

	Adopcion string `json:"Adopcion"`

	Estudia string `json:"Estudia"`

	NivelDeEstudio string `json:"NivelDeEstudio"`

	FechaBaja string `json:"FechaBaja"`

	MotivoBaja string `json:"MotivoBaja"`

	BeneficiarioSeguroVida string `json:"BeneficiarioSeguroVida"`

	EntregoCertificado string `json:"EntregoCertificado"`
}

const Interface2018DataAvroCRC64Fingerprint = "$}$\xe6`h\x06\x1c"

func NewInterface2018Data() Interface2018Data {
	r := Interface2018Data{}
	return r
}

func DeserializeInterface2018Data(r io.Reader) (Interface2018Data, error) {
	t := NewInterface2018Data()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInterface2018DataFromSchema(r io.Reader, schema string) (Interface2018Data, error) {
	t := NewInterface2018Data()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInterface2018Data(r Interface2018Data, w io.Writer) error {
	var err error
	err = vm.WriteString(r.LegajoDelEmpleado, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Apellido, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Nombres, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaNacimiento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.PaisDeNacimiento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Nacionalidad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EstadoCivil, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Sexo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Parentesco, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Incapacitado, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaDeDiscapacidad, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.TipoDocumento1, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.CodPaisTipoDocumento, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NroDocumento1, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.ObraSocial, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.PlanOS, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.AvisarAnteEmergencia, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.PagaSalarioFamiliar, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaInicioVinculo, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaVto1, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Guarderia, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaVto2, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Adopcion, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Estudia, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NivelDeEstudio, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.FechaBaja, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.MotivoBaja, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.BeneficiarioSeguroVida, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.EntregoCertificado, w)
	if err != nil {
		return err
	}
	return err
}

func (r Interface2018Data) Serialize(w io.Writer) error {
	return writeInterface2018Data(r, w)
}

func (r Interface2018Data) Schema() string {
	return "{\"fields\":[{\"name\":\"LegajoDelEmpleado\",\"type\":\"string\"},{\"name\":\"Apellido\",\"type\":\"string\"},{\"name\":\"Nombres\",\"type\":\"string\"},{\"name\":\"FechaNacimiento\",\"type\":\"string\"},{\"name\":\"PaisDeNacimiento\",\"type\":\"string\"},{\"name\":\"Nacionalidad\",\"type\":\"string\"},{\"name\":\"EstadoCivil\",\"type\":\"string\"},{\"name\":\"Sexo\",\"type\":\"string\"},{\"name\":\"Parentesco\",\"type\":\"string\"},{\"name\":\"Incapacitado\",\"type\":\"string\"},{\"name\":\"FechaDeDiscapacidad\",\"type\":\"string\"},{\"name\":\"TipoDocumento1\",\"type\":\"string\"},{\"name\":\"CodPaisTipoDocumento\",\"type\":\"string\"},{\"name\":\"NroDocumento1\",\"type\":\"string\"},{\"name\":\"ObraSocial\",\"type\":\"string\"},{\"name\":\"PlanOS\",\"type\":\"string\"},{\"name\":\"AvisarAnteEmergencia\",\"type\":\"string\"},{\"name\":\"PagaSalarioFamiliar\",\"type\":\"string\"},{\"name\":\"FechaInicioVinculo\",\"type\":\"string\"},{\"name\":\"FechaVto1\",\"type\":\"string\"},{\"name\":\"Guarderia\",\"type\":\"string\"},{\"name\":\"FechaVto2\",\"type\":\"string\"},{\"name\":\"Adopcion\",\"type\":\"string\"},{\"name\":\"Estudia\",\"type\":\"string\"},{\"name\":\"NivelDeEstudio\",\"type\":\"string\"},{\"name\":\"FechaBaja\",\"type\":\"string\"},{\"name\":\"MotivoBaja\",\"type\":\"string\"},{\"name\":\"BeneficiarioSeguroVida\",\"type\":\"string\"},{\"name\":\"EntregoCertificado\",\"type\":\"string\"}],\"name\":\"Andreani.HCMInterface2018.Events.Record.Interface2018Data\",\"type\":\"record\"}"
}

func (r Interface2018Data) SchemaName() string {
	return "Andreani.HCMInterface2018.Events.Record.Interface2018Data"
}

func (_ Interface2018Data) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Interface2018Data) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Interface2018Data) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Interface2018Data) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Interface2018Data) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Interface2018Data) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Interface2018Data) SetString(v string)   { panic("Unsupported operation") }
func (_ Interface2018Data) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Interface2018Data) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.LegajoDelEmpleado}

		return w

	case 1:
		w := types.String{Target: &r.Apellido}

		return w

	case 2:
		w := types.String{Target: &r.Nombres}

		return w

	case 3:
		w := types.String{Target: &r.FechaNacimiento}

		return w

	case 4:
		w := types.String{Target: &r.PaisDeNacimiento}

		return w

	case 5:
		w := types.String{Target: &r.Nacionalidad}

		return w

	case 6:
		w := types.String{Target: &r.EstadoCivil}

		return w

	case 7:
		w := types.String{Target: &r.Sexo}

		return w

	case 8:
		w := types.String{Target: &r.Parentesco}

		return w

	case 9:
		w := types.String{Target: &r.Incapacitado}

		return w

	case 10:
		w := types.String{Target: &r.FechaDeDiscapacidad}

		return w

	case 11:
		w := types.String{Target: &r.TipoDocumento1}

		return w

	case 12:
		w := types.String{Target: &r.CodPaisTipoDocumento}

		return w

	case 13:
		w := types.String{Target: &r.NroDocumento1}

		return w

	case 14:
		w := types.String{Target: &r.ObraSocial}

		return w

	case 15:
		w := types.String{Target: &r.PlanOS}

		return w

	case 16:
		w := types.String{Target: &r.AvisarAnteEmergencia}

		return w

	case 17:
		w := types.String{Target: &r.PagaSalarioFamiliar}

		return w

	case 18:
		w := types.String{Target: &r.FechaInicioVinculo}

		return w

	case 19:
		w := types.String{Target: &r.FechaVto1}

		return w

	case 20:
		w := types.String{Target: &r.Guarderia}

		return w

	case 21:
		w := types.String{Target: &r.FechaVto2}

		return w

	case 22:
		w := types.String{Target: &r.Adopcion}

		return w

	case 23:
		w := types.String{Target: &r.Estudia}

		return w

	case 24:
		w := types.String{Target: &r.NivelDeEstudio}

		return w

	case 25:
		w := types.String{Target: &r.FechaBaja}

		return w

	case 26:
		w := types.String{Target: &r.MotivoBaja}

		return w

	case 27:
		w := types.String{Target: &r.BeneficiarioSeguroVida}

		return w

	case 28:
		w := types.String{Target: &r.EntregoCertificado}

		return w

	}
	panic("Unknown field index")
}

func (r *Interface2018Data) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Interface2018Data) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Interface2018Data) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Interface2018Data) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Interface2018Data) HintSize(int)                     { panic("Unsupported operation") }
func (_ Interface2018Data) Finalize()                        {}

func (_ Interface2018Data) AvroCRC64Fingerprint() []byte {
	return []byte(Interface2018DataAvroCRC64Fingerprint)
}

func (r Interface2018Data) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["LegajoDelEmpleado"], err = json.Marshal(r.LegajoDelEmpleado)
	if err != nil {
		return nil, err
	}
	output["Apellido"], err = json.Marshal(r.Apellido)
	if err != nil {
		return nil, err
	}
	output["Nombres"], err = json.Marshal(r.Nombres)
	if err != nil {
		return nil, err
	}
	output["FechaNacimiento"], err = json.Marshal(r.FechaNacimiento)
	if err != nil {
		return nil, err
	}
	output["PaisDeNacimiento"], err = json.Marshal(r.PaisDeNacimiento)
	if err != nil {
		return nil, err
	}
	output["Nacionalidad"], err = json.Marshal(r.Nacionalidad)
	if err != nil {
		return nil, err
	}
	output["EstadoCivil"], err = json.Marshal(r.EstadoCivil)
	if err != nil {
		return nil, err
	}
	output["Sexo"], err = json.Marshal(r.Sexo)
	if err != nil {
		return nil, err
	}
	output["Parentesco"], err = json.Marshal(r.Parentesco)
	if err != nil {
		return nil, err
	}
	output["Incapacitado"], err = json.Marshal(r.Incapacitado)
	if err != nil {
		return nil, err
	}
	output["FechaDeDiscapacidad"], err = json.Marshal(r.FechaDeDiscapacidad)
	if err != nil {
		return nil, err
	}
	output["TipoDocumento1"], err = json.Marshal(r.TipoDocumento1)
	if err != nil {
		return nil, err
	}
	output["CodPaisTipoDocumento"], err = json.Marshal(r.CodPaisTipoDocumento)
	if err != nil {
		return nil, err
	}
	output["NroDocumento1"], err = json.Marshal(r.NroDocumento1)
	if err != nil {
		return nil, err
	}
	output["ObraSocial"], err = json.Marshal(r.ObraSocial)
	if err != nil {
		return nil, err
	}
	output["PlanOS"], err = json.Marshal(r.PlanOS)
	if err != nil {
		return nil, err
	}
	output["AvisarAnteEmergencia"], err = json.Marshal(r.AvisarAnteEmergencia)
	if err != nil {
		return nil, err
	}
	output["PagaSalarioFamiliar"], err = json.Marshal(r.PagaSalarioFamiliar)
	if err != nil {
		return nil, err
	}
	output["FechaInicioVinculo"], err = json.Marshal(r.FechaInicioVinculo)
	if err != nil {
		return nil, err
	}
	output["FechaVto1"], err = json.Marshal(r.FechaVto1)
	if err != nil {
		return nil, err
	}
	output["Guarderia"], err = json.Marshal(r.Guarderia)
	if err != nil {
		return nil, err
	}
	output["FechaVto2"], err = json.Marshal(r.FechaVto2)
	if err != nil {
		return nil, err
	}
	output["Adopcion"], err = json.Marshal(r.Adopcion)
	if err != nil {
		return nil, err
	}
	output["Estudia"], err = json.Marshal(r.Estudia)
	if err != nil {
		return nil, err
	}
	output["NivelDeEstudio"], err = json.Marshal(r.NivelDeEstudio)
	if err != nil {
		return nil, err
	}
	output["FechaBaja"], err = json.Marshal(r.FechaBaja)
	if err != nil {
		return nil, err
	}
	output["MotivoBaja"], err = json.Marshal(r.MotivoBaja)
	if err != nil {
		return nil, err
	}
	output["BeneficiarioSeguroVida"], err = json.Marshal(r.BeneficiarioSeguroVida)
	if err != nil {
		return nil, err
	}
	output["EntregoCertificado"], err = json.Marshal(r.EntregoCertificado)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Interface2018Data) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["LegajoDelEmpleado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LegajoDelEmpleado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for LegajoDelEmpleado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Apellido"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Apellido); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Apellido")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Nombres"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nombres); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Nombres")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaNacimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaNacimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaNacimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PaisDeNacimiento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PaisDeNacimiento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PaisDeNacimiento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Nacionalidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Nacionalidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Nacionalidad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EstadoCivil"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EstadoCivil); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EstadoCivil")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Sexo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sexo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Sexo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Parentesco"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Parentesco); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Parentesco")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Incapacitado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Incapacitado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Incapacitado")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaDeDiscapacidad"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaDeDiscapacidad); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaDeDiscapacidad")
	}
	val = func() json.RawMessage {
		if v, ok := fields["TipoDocumento1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.TipoDocumento1); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for TipoDocumento1")
	}
	val = func() json.RawMessage {
		if v, ok := fields["CodPaisTipoDocumento"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CodPaisTipoDocumento); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for CodPaisTipoDocumento")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NroDocumento1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NroDocumento1); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NroDocumento1")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ObraSocial"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ObraSocial); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ObraSocial")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PlanOS"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PlanOS); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PlanOS")
	}
	val = func() json.RawMessage {
		if v, ok := fields["AvisarAnteEmergencia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.AvisarAnteEmergencia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AvisarAnteEmergencia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PagaSalarioFamiliar"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PagaSalarioFamiliar); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PagaSalarioFamiliar")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaInicioVinculo"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaInicioVinculo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaInicioVinculo")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaVto1"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaVto1); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaVto1")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Guarderia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Guarderia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Guarderia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaVto2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaVto2); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaVto2")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Adopcion"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Adopcion); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Adopcion")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Estudia"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Estudia); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Estudia")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NivelDeEstudio"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NivelDeEstudio); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NivelDeEstudio")
	}
	val = func() json.RawMessage {
		if v, ok := fields["FechaBaja"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FechaBaja); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for FechaBaja")
	}
	val = func() json.RawMessage {
		if v, ok := fields["MotivoBaja"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MotivoBaja); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for MotivoBaja")
	}
	val = func() json.RawMessage {
		if v, ok := fields["BeneficiarioSeguroVida"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BeneficiarioSeguroVida); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for BeneficiarioSeguroVida")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EntregoCertificado"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EntregoCertificado); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for EntregoCertificado")
	}
	return nil
}
