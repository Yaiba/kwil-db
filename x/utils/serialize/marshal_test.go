package serialize_test

import (
	"kwil/x/utils/serialize"
	"testing"
)

func Test_Marshal(t *testing.T) {
	intVal := 100
	intValBytes, err := serialize.MarshalType(intVal)
	if err != nil {
		t.Errorf("failed to marshal int: %v", err)
	}

	val, err := serialize.UnmarshalType(intValBytes)
	if err != nil {
		t.Errorf("failed to unmarshal int: %v", err)
	}

	if val.(int) != intVal {
		t.Errorf("expected %v, got %v", intVal, val)
	}

	// test string
	strVal := "hello world"
	strValBytes, err := serialize.MarshalType(strVal)
	if err != nil {
		t.Errorf("failed to marshal string: %v", err)
	}

	val, err = serialize.UnmarshalType(strValBytes)
	if err != nil {
		t.Errorf("failed to unmarshal string: %v", err)
	}

	if val.(string) != strVal {
		t.Errorf("expected %v, got %v", strVal, val)
	}

	// test bool
	boolVal := true
	boolValBytes, err := serialize.MarshalType(boolVal)
	if err != nil {
		t.Errorf("failed to marshal bool: %v", err)
	}

	val, err = serialize.UnmarshalType(boolValBytes)
	if err != nil {
		t.Errorf("failed to unmarshal bool: %v", err)
	}

	if val.(bool) != boolVal {
		t.Errorf("expected %v, got %v", boolVal, val)
	}
}

func Test_Serialize(t *testing.T) {
	struct1 := TestStruct{
		Val1: 100,
		Val2: "hello world",
		Val3: true,
		Val4: []byte("hello world"),
	}

	bts, err := serialize.Serialize(struct1)
	if err != nil {
		t.Errorf("failed to serialize struct: %v", err)
	}

	struct2, err := serialize.Deserialize[TestStruct](bts)
	if err != nil {
		t.Errorf("failed to deserialize struct: %v", err)
	}

	if struct1.Val1 != struct2.Val1 {
		t.Errorf("expected %v, got %v", struct1.Val1, struct2.Val1)
	}

	if struct1.Val2 != struct2.Val2 {
		t.Errorf("expected %v, got %v", struct1.Val2, struct2.Val2)
	}

	if struct1.Val3 != struct2.Val3 {
		t.Errorf("expected %v, got %v", struct1.Val3, struct2.Val3)
	}

	if string(struct1.Val4) != string(struct2.Val4) {
		t.Errorf("expected %v, got %v", string(struct1.Val4), string(struct2.Val4))
	}

	// convert to struct 2
	struct3, err := serialize.Convert[TestStruct, TestStruct2](&struct1)
	if err != nil {
		t.Errorf("failed to convert struct: %v", err)
	}

	if struct1.Val1 != struct3.Val1 {
		t.Errorf("expected %v, got %v", struct1.Val1, struct3.Val1)
	}

	if struct1.Val2 != struct3.Val2 {
		t.Errorf("expected %v, got %v", struct1.Val2, struct3.Val2)
	}

	if struct1.Val3 != struct3.Val3 {
		t.Errorf("expected %v, got %v", struct1.Val3, struct3.Val3)
	}

	if string(struct1.Val4) != string(struct3.Val4) {
		t.Errorf("expected %v, got %v", string(struct1.Val4), string(struct3.Val4))
	}
}

type TestStruct struct {
	Val1 int    `json:"val1"`
	Val2 string `json:"val2"`
	Val3 bool   `json:"val3"`
	Val4 []byte `json:"val4"`
}

type TestStruct2 struct {
	Val1 int    `json:"val1"`
	Val2 string `json:"val2"`
	Val3 bool   `json:"val3"`
	Val4 []byte `json:"val4"`
}
