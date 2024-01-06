package go_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn, bookabledate" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

/*
CustomValidator
自定义验证器
*/
func CustomValidator() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
	router.GET("/bookable", getBookable)
	router.Run()
}

func getBookable(context *gin.Context) {
	var b Booking
	if err := context.ShouldBindWith(&b, binding.Query); err == nil {
		context.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
