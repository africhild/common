package controller

import (
	"errors"
	"fmt"
	"log"

	// "strings"

	"github.com/africhild/common/pkg/injection"
	"github.com/africhild/common/pkg/util"
	"github.com/gin-gonic/gin"

	// "github.com/go-playground/locales/en"
	// "github.com/go-playground/locales/es"
	// ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BaseController struct{}

const (
	defaultLanguage      = "en"
	defaultSingleMessage = false
)

var validate *validator.Validate

func (ctrl *BaseController) SQL(c *gin.Context) *gorm.DB {
	return injection.GetSQL(c)
}

func (ctrl *BaseController) KV(c *gin.Context) *redis.Client {
	return injection.GetKV(c)
}

func (ctrl *BaseController) User(c *gin.Context) (util.Object, error) {
	user := injection.GetUser(c)
	userId, ok := user["id"].(string)
	if !ok || userId == "" {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (ctrl *BaseController) Workspace(c *gin.Context) (util.Object, error) {
	workspace := injection.GetWorkspace(c)
	workspaceId, ok := workspace["id"].(string)
	if !ok || workspaceId == "" {
		return nil, errors.New("workspace not found")
	}

	return workspace, nil
}

func (ctrl *BaseController) Validate(c *gin.Context, payload any) (any, bool) {
	err := c.ShouldBindJSON(payload)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
			return nil, false
		}

		if data, ok := err.(validator.ValidationErrors); ok {
			log.Println(err)
			return data, false
		}

		log.Println(err)
		return nil, false
	}
	validate = validator.New()
	if err := validate.Struct(payload); err != nil {
		return err, false
	} else {
	fmt.Println("payload", payload)
	return nil, true
	}
}

func (ctrl *BaseController) ShouldBindJSON(c *gin.Context, payload any) (any, bool) {
	err := c.ShouldBindJSON(payload)
	if err != nil {
		return err, false
	}
	return nil, true
}

func (ctrl *BaseController) Success(c *gin.Context, message string, data any) {
	c.JSON(200, gin.H{
		"status":  true,
		"message": message,
		"data":    data,
	})
}

func (ctrl *BaseController) Error(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"status":  false,
		"message": message,
	})
}

func (ctrl *BaseController) ErrorWithData(c *gin.Context, message string, data any) {
	c.JSON(400, gin.H{
		"status":  false,
		"message": message,
		"data":    data,
	})
}

func (ctrl *BaseController) ErrorWithCode(c *gin.Context, message string, code int) {
	if code == 0 {
		code = 400
	}
	c.JSON(code, gin.H{
		"status":  false,
		"message": message,
	})
}

func (ctrl *BaseController) ErrorWithDataAndCode(c *gin.Context, message string, data any, code int) {
	if dataArr, ok := data.([]string); ok {
		for i, data := range dataArr {
			log.Println(i, data)
		}
	}
	c.JSON(code, gin.H{
		"status":  false,
		"message": message,
		"data":    data,
	})
}

var customMessages = map[string]string{
	"required": "%s is required.",
	"email":    "%s must be a valid email address.",
	"gte":      "%s must be greater than or equal to %s.",
	"lte":      "%s must be less than or equal to %s.",
}

// func (ctrl *BaseController) ValidationErorrs(c *gin.Context, message string,  data any, others ...any) []string {
// 	var singleMessage bool
// 	var language string
// 	if len(others) > 0 {
// 		if _language, ok := others[0].(string); ok {
// 			language = _language
// 		} else {
// 			language = "en"
// 		}
// 		if _singleMessage, ok := others[0].(bool); ok {
// 			singleMessage = _singleMessage
// 		} else {
// 			singleMessage = false
// 		}
// 	} else {
// 		singleMessage = false
// 		language = "en"
// 	}

// 	eng := en.New()
// 	spanish := es.New()
// 	uni := ut.New(eng, spanish)
// 	trans, _ := uni.GetTranslator(language)
// 	errMessages := []string{}
// 	if data != nil {
// 		validate = validator.New()
// 		// if validationErrors, ok := data.(validator.ValidationErrors); ok {
// 		if err := validate.Struct(data); err != nil {
// 			// Collect validation errors for the current user
// 			validationErrors := err.(validator.ValidationErrors)
// 			// userErrors := make(map[string]string)
// 			// for _, fieldErr := range validationErrors {
// 			// 	userErrors[fieldErr.Field()] = fieldErr.Tag()
// 			// }
// 			// errorsMap[i] = userErrors
// 			if singleMessage {
// 				errMessages = append(errMessages, validationErrors[0].Error())
// 			} else {
// 				for _, fieldErr := range validationErrors {
// 					// fmt.Println(fieldErr.Translate(trans))
// 					msg := customMessages[fieldErr.Tag()]
// 					if msg == "" {
// 						msg = "Validation failed on the %s tag."
// 					}
// 					var errMsg string
// 					// fieldName := strings.ToLower(fieldErr.Field())

// 					if param := fieldErr.Param(); param != "" {
// 						errMsg = fmt.Sprintf(msg, fieldErr.Field(), fieldErr.Param())
// 					} else {
// 						errMsg = fmt.Sprintf(msg, fieldErr.Field())
// 					}
// 					errMessages = append(errMessages, errMsg)
// 				}
// 			}
// 		} else {
// 			errMessages = append(errMessages, fmt.Sprintf("%v", data))
// 		}
// 	} else {
// 		errMessages = append(errMessages, message)
// 	}
// 	return errMessages
// }

func (ctrl *BaseController) ValidationErorrs(c *gin.Context, message string, data any, others ...any) []string {
	// Set default values
	singleMessage := defaultSingleMessage
	// language := defaultLanguage

	// Process optional parameters
	for _, param := range others {
		switch v := param.(type) {
		// case string:
		//     language = v
		case bool:
			singleMessage = v
		}
	}

	// Initialize translators
	// eng := en.New()
	// spanish := es.New()
	// uni := ut.New(eng, spanish)
	// trans, found := uni.GetTranslator(language)
	// if !found {
	//     trans, _ = uni.GetTranslator(defaultLanguage)
	// }

	errMessages := []string{}

	if data != nil {
		validate = validator.New()
		if err := validate.Struct(data); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			if validationErrors != nil {
				if singleMessage {
					// errMessages = append(errMessages, validationErrors[0].Translate(trans))
					errMessages = append(errMessages, validationErrors[0].Error())
				} else {
					for _, fieldErr := range validationErrors {
						msg := customMessages[fieldErr.Tag()]
						if msg == "" {
							msg = "Validation failed on the %s tag."
						}

						var errMsg string
						if param := fieldErr.Param(); param != "" {
							errMsg = fmt.Sprintf(msg, fieldErr.Field(), param)
						} else {
							errMsg = fmt.Sprintf(msg, fieldErr.Field())
						}
						errMessages = append(errMessages, errMsg)
					}
				}
			}
		}
	} else {
		errMessages = append(errMessages, message)
	}

	return errMessages
}
