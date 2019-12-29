package json

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type Company struct {
	Name 	string
}

type Person struct {
	index 		int
	Name 		string			`json:"name"`
	Age 		int8			`json:"age"`
	Employers	[]Company		`json:"employers"`
}

func FromJson(bytes []byte) (Person, error) {
	var result Person
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		//return nil, errors.New("invalid json input")
		//return result, fmt.Errorf("invalid json input")
		return result, errors.Wrap(err, "Error unmarshal json")
	}
	return result, nil
}

func ToJson(p Person) ([]byte, error) {
	var data []byte
	data, err := json.Marshal(p)
	if err != nil {
		return data, errors.Wrap(err, "Error marshal data into json")
	}
	return data, nil
}