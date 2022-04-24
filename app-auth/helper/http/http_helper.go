package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
)

const (
	codeSuccess           int = 200
	codeInsertSucces      int = 201
	codeBadRequestError   int = 400
	codeUnauthorizedError int = 401
	codeError             int = 500
)

type HTTPHelper struct {
	Validate   *validator.Validate
	Translator ut.Translator
}

type ResponseHelper struct {
	Code        int
	Status      string
	Message     string
	MessageType string
	ResultData  interface{}
}

func (u *HTTPHelper) EmptyJsonMap() map[string]interface{} {
	return make(map[string]interface{})
}

// SET RESPONSE MESSAGE
func (u *HTTPHelper) SetResponse(c *gin.Context, code int, status string, message string, messageType string, data interface{}) ResponseHelper {
	return ResponseHelper{code, status, message, messageType, data}
}

// SEND ERROR MESSAGE
func (u *HTTPHelper) SendError(c *gin.Context, code int, message string, messageType string, data interface{}) error {
	res := u.SetResponse(c, codeError, `error`, message, messageType, data)

	return u.SendResponse(res, c, code)
}

// SEND DATA SUCCESS
func (u *HTTPHelper) SendSuccess(c *gin.Context, message string, data interface{}) error {
	res := u.SetResponse(c, codeSuccess, `ok`, message, `success`, data)

	return u.SendResponse(res, c, codeSuccess)
}

func (u *HTTPHelper) SendBadRequest(c *gin.Context, message string, data interface{}) error {
	res := u.SetResponse(c, codeBadRequestError, `error`, message, `badRequest`, data)

	return u.SendResponse(res, c, codeBadRequestError)
}

func (u *HTTPHelper) SendUnauthorizedError(c *gin.Context, message string, data interface{}) error {
	return u.SendError(c, codeUnauthorizedError, message, `unAuthorized`, data)
}

// SEND RESPONSE MESSAGE
func (u *HTTPHelper) SendResponse(res ResponseHelper, c *gin.Context, code int) error {
	if len(res.MessageType) == 0 {
		res.Message = `success`
	}

	var resCode int
	if res.Code != 200 {
		resCode = code
	} else {
		resCode = http.StatusOK
	}

	c.JSON(resCode, gin.H{
		"code":         resCode,
		"message_type": res.MessageType,
		"message":      res.Message,
		"result_data":  res.ResultData,
	})

	return nil
}

func SendHttpResponse(c *gin.Context, code int, messageType string, message string, resultData interface{}) error {
	if len(messageType) == 0 {
		message = `success`
	}

	var respCode int
	if code != 200 {
		respCode = code
	} else {
		respCode = http.StatusOK
	}

	c.JSON(code, gin.H{
		"code":         respCode,
		"message_type": messageType,
		"message":      message,
		"result_data":  resultData,
	})
	return nil
}
