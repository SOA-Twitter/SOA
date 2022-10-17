package data

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"regexp"
)

// Defining the main struct for our API
type Product struct {
	ID          int     `json:"id"`                       //specifies that in the incoming Body the field to map to this will be called "id"
	Name        string  `json:"name" validate:"required"` //there are some integrated validation, for eg. this specifies that a value for name must be provided, otherwise it will not be valid
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"` //the tag "sku" is there so we can add custom validation
	CreatedOn   string  `json:"createdOn"`
	UpdatedOn   string  `json:"updatedOn"`
	DeletedOn   string  `json:"deletedOn"`
}

type Products []*Product

//Function to validate the incoming object from front.
//NOTE: if the tag "sku" is not present in the struct anotations we will get an error
func (p *Product) Validate() error {
	validate := validator.New()

	err := validate.RegisterValidation("sku", validateSKU)
	if err != nil {
		return err
	}

	return validate.Struct(p)
}

//We use a regex to validate a custom look of sku-s
//For eg: abc-abc-abc
func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile("[a-z]+-[a-z]+-[a-z]")
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

//Functions to encode and decode products to json and from json.
//If we inject an interface we achieve dependancy injection, meaning that anything that implements this interface can be passed down
//For us it will be ResponseWriter, but it also may be a file writer or something similar.
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
