package utils

import (
	"fmt"
	"reflect"
	"strings"
)

type Validator struct {
	Entity interface{}
	DTO    interface{}
	Error  []map[string]interface{}
}

// CheckValidator : body io.ReadCloser, entities interface{}, request interface{}
func CheckValidator(entities interface{}, request interface{}) []map[string]interface{} {
	validate := Validator{
		DTO:    request,
		Entity: entities,
	}
	validate.Validator()

	if len(validate.Error) > 0 {
		return validate.Error
	}
	return nil
}

func (validate *Validator) Validator() *Validator {
	value := validate.newType()
	numFields := value.NumField()
	for i := 0; i < numFields; i++ {
		string1 := strings.SplitN(value.Field(i).Tag.Get("validator"), ",", 2)[0]
		string2 := strings.Split(string1, "|")
		if (string2)[0] != "" {
			for j := 0; j < len(string2); j++ {
				validate.checkValidate(value.Field(i).Tag.Get("json"), string2[j], value.Field(i))
			}
		}
	}
	return validate
}

func (validate *Validator) checkValidate(name string, values interface{}, field reflect.StructField) {
	message1 := strings.SplitN(field.Tag.Get("validator_message"), ",", 2)[0]
	message := strings.Split(message1, "|")[0]
	message2 := strings.Split(message1, "|")

	datas := convertStructToMap(validate.DTO)

	switch values {
	case "string":
		error := validate.checkString(datas[name])
		msg := validate.customMessage(values, message, message2)

		if !error {
			validate.Error = append(validate.Error, map[string]interface{}{
				"input":   name,
				"message": msg,
			})
		}
	case "required":
		error := validate.checkRequired(datas[name])
		msg := validate.customMessage(values, message, message2)

		if !error {
			validate.Error = append(validate.Error, map[string]interface{}{
				"input":   name,
				"message": msg,
			})
		}
	case "number":
		error := validate.checkNumber(datas[name])
		msg := validate.customMessage(values, message, message2)

		if !error {
			validate.Error = append(validate.Error, map[string]interface{}{
				"input":   name,
				"message": msg,
			})
		}
		// case "exists":
		// 	error := validate.checkUnique(datas[name], name)
		// 	msg := validate.customMessage(values, message, message2)

		// 	if !error {
		// 		validate.Error = append(validate.Error, map[string]interface{}{
		// 			"input":   name,
		// 			"message": msg,
		// 		})
		// 	}
	}
}

func (validate *Validator) customMessage(values interface{}, message string, message2 []string) string {
	var msg string
	if len(message2) > 0 {
		for i := 0; i < len(message2); i++ {
			check2 := strings.Contains(message2[i], fmt.Sprintf("%s", values))
			if check2 {
				msg = strings.Split(message2[i], fmt.Sprintf("%s:", values))[1]
			}
		}
	}
	if msg == "" {
		msg = fmt.Sprintf("is %s", values)
	}
	return msg
}

func (validate *Validator) checkString(value interface{}) bool {
	xType := fmt.Sprintf("%T", value)
	if value != nil {
		if xType != "string" {
			return false
		}
	}
	return true
}

func (validate *Validator) checkNumber(value interface{}) bool {
	xType := fmt.Sprintf("%T", value)
	return xType == "number"
}

func (validate *Validator) checkRequired(value interface{}) bool {
	if value == "" {
		return false
	}
	if value == nil {
		return false
	}

	return true
}

// func (validate *Validator) checkUnique(val interface{}, dest ...interface{}) bool {
// 	datas := validate.Entity
// 	var number int64
// 	validate.DB.Model(datas).
// 		Where(fmt.Sprintf("%s = ?", dest[0]), val).
// 		Count(&number)

// 	return number <= 0
// }

// func (validate *Validator) checkUniqueSkipId(val interface{}, dest ...interface{}) bool {
// 	datas := validate.Entity
// 	var number int64
// 	validate.DB.Model(datas).
// 		Where(fmt.Sprintf("%s = ?", dest[0]), val).
// 		Count(&number)

// 	return number <= 0
// }

// function utils

func (validate *Validator) newType() reflect.Type {
	value := reflect.New(reflect.TypeOf(validate.DTO)).Interface()
	v := reflect.ValueOf(value)
	i := reflect.Indirect(v)
	s := i.Type()
	return s
}

func convertStructToMap(st interface{}) map[string]interface{} {

	reqRules := make(map[string]interface{})

	v := reflect.ValueOf(st)
	t := reflect.TypeOf(st)

	for i := 0; i < v.NumField(); i++ {
		key := strings.ToLower(t.Field(i).Name)
		typ := v.FieldByName(t.Field(i).Name).Kind().String()
		structTag := t.Field(i).Tag.Get("json")
		jsonName := strings.TrimSpace(strings.Split(structTag, ",")[0])
		value := v.FieldByName(t.Field(i).Name)

		// if jsonName is not empty use it for the key
		if jsonName != "" && jsonName != "-" {
			key = jsonName
		}

		if typ == "string" {
			if !(value.String() == "" && strings.Contains(structTag, "omitempty")) {
				// fmt.Println(key, value)
				// fmt.Println(key, value.String())
				reqRules[key] = value.String()
			}
		} else if typ == "int" {
			reqRules[key] = value.Int()
		} else {
			reqRules[key] = value.Interface()
		}

	}

	return reqRules
}
