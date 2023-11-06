package validation_playground

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"testing"
)

type user struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func TestSimple(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	us := user{
		FirstName: "",
		LastName:  "",
	}

	err := validate.Struct(us)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(us).FieldByName(err.Field())
			jsonTag := field.Tag.Get("json")
			fmt.Println("json :", jsonTag)
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println("tag -> " + err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
	}
}

type DbBackedUser struct {
	Name sql.NullString `validate:"required"`
	Age  sql.NullInt64  `validate:"required"`
}

func ValidateValuer(field reflect.Value) interface{} {

	if valuer, ok := field.Interface().(driver.Valuer); ok {

		val, err := valuer.Value()
		if err == nil {
			return val
		}
		// handle the error how you want
	}

	return nil
}

func TestCustom(t *testing.T) {
	validate := validator.New()

	// register all sql.Null* types to use the ValidateValuer CustomTypeFunc
	validate.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})

	// build object for validation
	x := DbBackedUser{Name: sql.NullString{String: "", Valid: true}, Age: sql.NullInt64{Int64: 0, Valid: false}}

	err := validate.Struct(x)

	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}

// MyStruct ..
type MyStruct struct {
	String string `validate:"is-awesome"`
}

func ValidateMyVal(fl validator.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("is-awesome", ValidateMyVal)
	if err != nil {
		return
	}

	s := MyStruct{String: "awesome"}

	err = validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}

	s.String = "not awesome"
	err = validate.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}
