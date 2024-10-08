// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     EmployeeData.avsc
 */
package appnameEvents

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EmployeeData struct {
	DateOfBirth int64 `json:"DateOfBirth"`

	HiringDate int64 `json:"HiringDate"`

	ExternalId *UnionNullString `json:"ExternalId"`

	FamilyName *UnionNullString `json:"FamilyName"`

	FirstName *UnionNullString `json:"FirstName"`

	ImageUrl *UnionNullString `json:"ImageUrl"`

	LastName *UnionNullString `json:"LastName"`

	MiddleName *UnionNullString `json:"MiddleName"`

	IsActive bool `json:"IsActive"`

	Id int32 `json:"Id"`
}

const EmployeeDataAvroCRC64Fingerprint = "\x11\xc1\xc0ڋ\xec\xd89"

func NewEmployeeData() EmployeeData {
	r := EmployeeData{}
	r.ExternalId = nil
	r.FamilyName = nil
	r.FirstName = nil
	r.ImageUrl = nil
	r.LastName = nil
	r.MiddleName = nil
	return r
}

func DeserializeEmployeeData(r io.Reader) (EmployeeData, error) {
	t := NewEmployeeData()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEmployeeDataFromSchema(r io.Reader, schema string) (EmployeeData, error) {
	t := NewEmployeeData()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEmployeeData(r EmployeeData, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.DateOfBirth, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.HiringDate, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ExternalId, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FamilyName, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.FirstName, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.ImageUrl, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.LastName, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.MiddleName, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.IsActive, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Id, w)
	if err != nil {
		return err
	}
	return err
}

func (r EmployeeData) Serialize(w io.Writer) error {
	return writeEmployeeData(r, w)
}

func (r EmployeeData) Schema() string {
	return "{\"fields\":[{\"name\":\"DateOfBirth\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"name\":\"HiringDate\",\"type\":{\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}},{\"default\":null,\"name\":\"ExternalId\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FamilyName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"FirstName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"ImageUrl\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"LastName\",\"type\":[\"null\",\"string\"]},{\"default\":null,\"name\":\"MiddleName\",\"type\":[\"null\",\"string\"]},{\"name\":\"IsActive\",\"type\":\"boolean\"},{\"name\":\"Id\",\"type\":\"int\"}],\"name\":\"Andreani.RHpro.Events.Common.EmployeeData\",\"type\":\"record\"}"
}

func (r EmployeeData) SchemaName() string {
	return "Andreani.RHpro.Events.Common.EmployeeData"
}

func (_ EmployeeData) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EmployeeData) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EmployeeData) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EmployeeData) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EmployeeData) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EmployeeData) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EmployeeData) SetString(v string)   { panic("Unsupported operation") }
func (_ EmployeeData) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EmployeeData) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.DateOfBirth}

		return w

	case 1:
		w := types.Long{Target: &r.HiringDate}

		return w

	case 2:
		r.ExternalId = NewUnionNullString()

		return r.ExternalId
	case 3:
		r.FamilyName = NewUnionNullString()

		return r.FamilyName
	case 4:
		r.FirstName = NewUnionNullString()

		return r.FirstName
	case 5:
		r.ImageUrl = NewUnionNullString()

		return r.ImageUrl
	case 6:
		r.LastName = NewUnionNullString()

		return r.LastName
	case 7:
		r.MiddleName = NewUnionNullString()

		return r.MiddleName
	case 8:
		w := types.Boolean{Target: &r.IsActive}

		return w

	case 9:
		w := types.Int{Target: &r.Id}

		return w

	}
	panic("Unknown field index")
}

func (r *EmployeeData) SetDefault(i int) {
	switch i {
	case 2:
		r.ExternalId = nil
		return
	case 3:
		r.FamilyName = nil
		return
	case 4:
		r.FirstName = nil
		return
	case 5:
		r.ImageUrl = nil
		return
	case 6:
		r.LastName = nil
		return
	case 7:
		r.MiddleName = nil
		return
	}
	panic("Unknown field index")
}

func (r *EmployeeData) NullField(i int) {
	switch i {
	case 2:
		r.ExternalId = nil
		return
	case 3:
		r.FamilyName = nil
		return
	case 4:
		r.FirstName = nil
		return
	case 5:
		r.ImageUrl = nil
		return
	case 6:
		r.LastName = nil
		return
	case 7:
		r.MiddleName = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ EmployeeData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EmployeeData) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EmployeeData) HintSize(int)                     { panic("Unsupported operation") }
func (_ EmployeeData) Finalize()                        {}

func (_ EmployeeData) AvroCRC64Fingerprint() []byte {
	return []byte(EmployeeDataAvroCRC64Fingerprint)
}

func (r EmployeeData) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["DateOfBirth"], err = json.Marshal(r.DateOfBirth)
	if err != nil {
		return nil, err
	}
	output["HiringDate"], err = json.Marshal(r.HiringDate)
	if err != nil {
		return nil, err
	}
	output["ExternalId"], err = json.Marshal(r.ExternalId)
	if err != nil {
		return nil, err
	}
	output["FamilyName"], err = json.Marshal(r.FamilyName)
	if err != nil {
		return nil, err
	}
	output["FirstName"], err = json.Marshal(r.FirstName)
	if err != nil {
		return nil, err
	}
	output["ImageUrl"], err = json.Marshal(r.ImageUrl)
	if err != nil {
		return nil, err
	}
	output["LastName"], err = json.Marshal(r.LastName)
	if err != nil {
		return nil, err
	}
	output["MiddleName"], err = json.Marshal(r.MiddleName)
	if err != nil {
		return nil, err
	}
	output["IsActive"], err = json.Marshal(r.IsActive)
	if err != nil {
		return nil, err
	}
	output["Id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EmployeeData) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["DateOfBirth"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DateOfBirth); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DateOfBirth")
	}
	val = func() json.RawMessage {
		if v, ok := fields["HiringDate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.HiringDate); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for HiringDate")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ExternalId"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ExternalId); err != nil {
			return err
		}
	} else {
		r.ExternalId = NewUnionNullString()

		r.ExternalId = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["FamilyName"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FamilyName); err != nil {
			return err
		}
	} else {
		r.FamilyName = NewUnionNullString()

		r.FamilyName = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["FirstName"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.FirstName); err != nil {
			return err
		}
	} else {
		r.FirstName = NewUnionNullString()

		r.FirstName = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["ImageUrl"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.ImageUrl); err != nil {
			return err
		}
	} else {
		r.ImageUrl = NewUnionNullString()

		r.ImageUrl = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["LastName"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.LastName); err != nil {
			return err
		}
	} else {
		r.LastName = NewUnionNullString()

		r.LastName = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["MiddleName"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.MiddleName); err != nil {
			return err
		}
	} else {
		r.MiddleName = NewUnionNullString()

		r.MiddleName = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["IsActive"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IsActive); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IsActive")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Id")
	}
	return nil
}
