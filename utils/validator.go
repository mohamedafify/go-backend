package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func SetupValidations(server *gin.Engine) {
	// used to get the json name in structs to return
	// a useful error when a field is missing
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

func BindBody(c *gin.Context, obj any) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		serr, ok := err.(*json.SyntaxError)
		if ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "Invalid JSON"})
			log.Println(serr.Error())
			return serr
		}
		if errsArr := parseError(err); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errsArr})
			log.Println(err)
			return err
		}
		return err
	}
	return nil
}

func parseError(err error) []string {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		errorMessages := make([]string, len(validationErrs))
		for i, e := range validationErrs {
			switch e.Tag() {
			case "required":
				errorMessages[i] = fmt.Sprintf("%s field is required", e.Field())
			}
		}
		return errorMessages
	} else if marshallingErr, ok := err.(*json.UnmarshalTypeError); ok {
		return []string{fmt.Sprintf("%s field must be a %s", marshallingErr.Field, marshallingErr.Type.String())}
	}
	return nil
}
