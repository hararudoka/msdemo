package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/hararudoka/msdemo/internal/model"
)

type Service struct {
	URL   string
	Token string
}

// GetAllEmployers returns all employees' data
func (s Service) GetAllEmployers() (model.Employees, error) {
	url := "https://online.moysklad.ru/api/remap/1.2/entity/employee"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.Employees{}, err
	}
	req.Header.Set("Authorization", "Bearer "+s.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Employees{}, err
	}

	defer resp.Body.Close()

	jsonBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.Employees{}, err
	}

	var data model.Employees
	err = json.Unmarshal(jsonBody, &data)
	return data, err
}

// GetEmployer gets one employee
func (s Service) GetEmployer(id string) (model.Employee, error) {
	// TODO: add realization

	// 	curl -X GET
	//   "https://online.moysklad.ru/api/remap/1.2/entity/employee/7944ef04-f831-11e5-7a69-971500188b19"
	//   -H "Authorization: Basic <Credentials>"

	return model.Employee{}, nil
}

// CreateEmployer creates the employee with data. TODO: data can be extended
func (s Service) CreateEmployer(lastName string) error {
	// TODO: add realization

	// curl -X POST
	// "https://online.moysklad.ru/api/remap/1.2/entity/employee/"
	// -H "Authorization: Basic <Credentials>"
	// -H "Content-Type: application/json"
	//   -d '{
	//         "firstName": "Петр",
	//         "middleName": "Иванович",
	//         "lastName": "Мойскладкин",
	//         "inn": "222490425273",
	//         "position": "Директор",
	//         "phone": "+7(999)888-7766",
	//         "description": "Описание",
	//         "attributes": [
	//           {
	//             "meta": {
	//               "href": "https://online.moysklad.ru/api/remap/1.2/entity/employee/metadata/attributes/ed14b498-cae3-11e8-9dd2-f3a300000044",
	//               "type": "attributemetadata",
	//               "mediaType": "application/json"
	//             },
	//             "value": "Строковое доп поле"
	//           }
	//         ]
	//       }'

	return nil
}

// EditEmployer edits the employee
func (s Service) EditEmployer(id string) error {
	// TODO: add realization

	// curl -X PUT // надо изучить
	// "https://online.moysklad.ru/api/remap/1.2/entity/employee/<id>"
	// -H "Authorization: Basic <Credentials>"
	// -H "Content-Type: application/json"
	//   -d '{
	//         "firstName": "Петр",
	//         "middleName": "Иванович",
	//         "lastName": "Мойскладкин",
	//         "inn": "222490425273",
	//         "position": "Директор",
	//         "phone": "+7(999)888-7766",
	//         "description": "Описание",
	//         "attributes": [
	//           {
	//             "meta": {
	//               "href": "https://online.moysklad.ru/api/remap/1.2/entity/employee/metadata/attributes/ed14b498-cae3-11e8-9dd2-f3a300000044",
	//               "type": "attributemetadata",
	//               "mediaType": "application/json"
	//             },
	//             "value": "Строковое доп поле"
	//           }
	//         ]
	//       }'

	return nil
}

// DeleteEmployer deletes the employee
func (s Service) DeleteEmployer(id string) error {
	// TODO: add realization

	// curl -X DELETE // надо изучить
	// "https://online.moysklad.ru/api/remap/1.2/entity/employee/<id>"
	// -H "Authorization: Basic <Credentials>"

	return nil
}
