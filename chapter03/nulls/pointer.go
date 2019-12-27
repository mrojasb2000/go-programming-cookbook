package nulls

import (
	"encoding/json"
	"fmt"
)

// ExamplePointer is the same, but
// uses a *Int
type ExamplePointer struct {
	Age  *int   `json:"age,omitempty"`
	Name string `json:"name"`
}

// PointerEncoding shows methods for
// dealing with nil/omitted values
func PointerEncoding() error {
	// note that no age = nil age
	e := ExamplePointer{}

	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Pointer Unamrshal, no age: %+v", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Pointer Marshal, no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Pointer Unamrshal, with age = 0: %+v", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Pointer Marshal, with age = 0:", string(value))

	return nil
}
