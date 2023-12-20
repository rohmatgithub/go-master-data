package validation_playground

import (
	"errors"
	"fmt"
	en_US "github.com/go-playground/locales/en"
	id_ID "github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	idtranslations "github.com/go-playground/validator/v10/translations/id"
	"strings"
	"testing"
)

type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func TestTranslation(t *testing.T) {

	// NOTE: ommitting allot of error checking for brevity

	en := en_US.New()
	id := id_ID.New()
	uni = ut.New(en, id)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	err := entranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return
	}

	err = validate.Var("e", "required,min=2")
	if err != nil {
		// translate all error at once
		var errs validator.ValidationErrors
		errors.As(err, &errs)

		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned, and you'll see something surprising
		// translations are i18n aware!!!!
		for _, fieldError := range errs {
			//result[jsonTag] = fieldError.Translate(v.getTranslator(contextModel.AuthAccessTokenModel.Locale))
			fmt.Println("-->>", strings.Replace(fieldError.Translate(trans), fieldError.Field()+" ", "", 1))
			fmt.Println(fieldError.Translate(trans))
		}
	}

	translateAll(trans)
	translateIndividual(trans)
	translateOverride(trans) // yep you can specify your own in whatever locale you want!

	fmt.Printf("\n\n\n")

	trans, _ = uni.GetTranslator("id")

	//validate = validator.New()
	err = idtranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
		return
	}

	translateAll(trans)
	translateIndividual(trans)
	translateOverride(trans) // yep you can specify your own in whatever locale you want!
}

func translateAll(trans ut.Translator) {

	type User struct {
		//Username string `validate:"required"`
		//Tagline  string `validate:"required,lt=10"`
		//Tagline2 string `validate:"required,gt=1"`
		//Email    string `validate:"required,email"`
		Name string `validate:"gte=2,eq=10"`
	}

	user := User{
		//Username: "Joeybloggs",
		//Tagline:  "This tagline is way too long.",
		//Tagline2: "1",
		//Email:    "email_test@mail.com",
		Name: "test dulu kali",
	}

	err := validate.Struct(user)
	if err != nil {

		// translate all error at once
		errs := err.(validator.ValidationErrors)

		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware!!!!
		// eg. '10 characters' vs '1 character'
		fmt.Println(errs.Translate(trans))
	}
}

func translateIndividual(trans ut.Translator) {

	type User struct {
		Username string `validate:"required"`
	}

	var user User

	err := validate.Struct(user)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
		}
	}
}

func translateOverride(trans ut.Translator) {

	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	type User struct {
		Username string `validate:"required"`
	}

	var user User

	err := validate.Struct(user)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
		}
	}
}
