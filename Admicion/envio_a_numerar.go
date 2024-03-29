// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     ListaEnviosANumerar.avsc
 */
package AdmicionEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EnvioANumerar struct {
	PrePieceHeader int64 `json:"PrePieceHeader"`

	PrePiece int32 `json:"PrePiece"`

	PostalStampNumber *UnionNullString `json:"PostalStampNumber"`

	InternalContrac string `json:"internalContrac"`

	BranchID int32 `json:"branchID"`

	OperatorId int32 `json:"OperatorId"`

	Published bool `json:"Published"`
}

const EnvioANumerarAvroCRC64Fingerprint = "Is\x8c\xd0\xcf\x1d\x86\r"

func NewEnvioANumerar() EnvioANumerar {
	r := EnvioANumerar{}
	r.PostalStampNumber = nil
	return r
}

func DeserializeEnvioANumerar(r io.Reader) (EnvioANumerar, error) {
	t := NewEnvioANumerar()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnvioANumerarFromSchema(r io.Reader, schema string) (EnvioANumerar, error) {
	t := NewEnvioANumerar()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnvioANumerar(r EnvioANumerar, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.PrePieceHeader, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.PrePiece, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.PostalStampNumber, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.InternalContrac, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.BranchID, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.OperatorId, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Published, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnvioANumerar) Serialize(w io.Writer) error {
	return writeEnvioANumerar(r, w)
}

func (r EnvioANumerar) Schema() string {
	return "{\"fields\":[{\"name\":\"PrePieceHeader\",\"type\":\"long\"},{\"name\":\"PrePiece\",\"type\":\"int\"},{\"default\":null,\"name\":\"PostalStampNumber\",\"type\":[\"null\",\"string\"]},{\"name\":\"internalContrac\",\"type\":\"string\"},{\"name\":\"branchID\",\"type\":\"int\"},{\"name\":\"OperatorId\",\"type\":\"int\"},{\"name\":\"Published\",\"type\":\"boolean\"}],\"name\":\"Andreani.Admicion.Events.Record.EnvioANumerar\",\"type\":\"record\"}"
}

func (r EnvioANumerar) SchemaName() string {
	return "Andreani.Admicion.Events.Record.EnvioANumerar"
}

func (_ EnvioANumerar) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnvioANumerar) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnvioANumerar) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnvioANumerar) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnvioANumerar) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnvioANumerar) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnvioANumerar) SetString(v string)   { panic("Unsupported operation") }
func (_ EnvioANumerar) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnvioANumerar) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.PrePieceHeader}

		return w

	case 1:
		w := types.Int{Target: &r.PrePiece}

		return w

	case 2:
		r.PostalStampNumber = NewUnionNullString()

		return r.PostalStampNumber
	case 3:
		w := types.String{Target: &r.InternalContrac}

		return w

	case 4:
		w := types.Int{Target: &r.BranchID}

		return w

	case 5:
		w := types.Int{Target: &r.OperatorId}

		return w

	case 6:
		w := types.Boolean{Target: &r.Published}

		return w

	}
	panic("Unknown field index")
}

func (r *EnvioANumerar) SetDefault(i int) {
	switch i {
	case 2:
		r.PostalStampNumber = nil
		return
	}
	panic("Unknown field index")
}

func (r *EnvioANumerar) NullField(i int) {
	switch i {
	case 2:
		r.PostalStampNumber = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EnvioANumerar) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EnvioANumerar) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EnvioANumerar) HintSize(int)                     { panic("Unsupported operation") }
func (_ EnvioANumerar) Finalize()                        {}

func (_ EnvioANumerar) AvroCRC64Fingerprint() []byte {
	return []byte(EnvioANumerarAvroCRC64Fingerprint)
}

func (r EnvioANumerar) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["PrePieceHeader"], err = json.Marshal(r.PrePieceHeader)
	if err != nil {
		return nil, err
	}
	output["PrePiece"], err = json.Marshal(r.PrePiece)
	if err != nil {
		return nil, err
	}
	output["PostalStampNumber"], err = json.Marshal(r.PostalStampNumber)
	if err != nil {
		return nil, err
	}
	output["internalContrac"], err = json.Marshal(r.InternalContrac)
	if err != nil {
		return nil, err
	}
	output["branchID"], err = json.Marshal(r.BranchID)
	if err != nil {
		return nil, err
	}
	output["OperatorId"], err = json.Marshal(r.OperatorId)
	if err != nil {
		return nil, err
	}
	output["Published"], err = json.Marshal(r.Published)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnvioANumerar) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["PrePieceHeader"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PrePieceHeader); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PrePieceHeader")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PrePiece"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PrePiece); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for PrePiece")
	}
	val = func() json.RawMessage {
		if v, ok := fields["PostalStampNumber"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.PostalStampNumber); err != nil {
			return err
		}
	} else {
		r.PostalStampNumber = NewUnionNullString()

		r.PostalStampNumber = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["internalContrac"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.InternalContrac); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for internalContrac")
	}
	val = func() json.RawMessage {
		if v, ok := fields["branchID"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BranchID); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for branchID")
	}
	val = func() json.RawMessage {
		if v, ok := fields["OperatorId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.OperatorId); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for OperatorId")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Published"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Published); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Published")
	}
	return nil
}
