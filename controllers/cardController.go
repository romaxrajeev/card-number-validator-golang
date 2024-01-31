package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romaxrajeev/gotp/pkg"
	"github.com/romaxrajeev/gotp/structs"
	"go.uber.org/zap"
)

var LOGGER *zap.Logger = pkg.DefaultLogger()

func VerifyCardNumber(c *gin.Context) {
	// Fetch data
	var cardDetails structs.CardDetails
	error := c.ShouldBindJSON(&cardDetails)
	if error != nil {
		c.JSON(http.StatusBadRequest, structs.Error{ErrorMessage: "Could not fetch data from the body", ErrorCode: 123})
	}
	// Convert the card number string to list of integers
	cardNumber, err := pkg.StringToListOfIntConverter(cardDetails.CardNumber)
	if err != nil {
		LOGGER.Error("Not able to convert string to list. See err: " + err.Error() + "\nData: " + cardDetails.CardNumber)
		c.JSON(http.StatusBadRequest, structs.Error{ErrorMessage: "Not converted to list"})
	}
	// Validate with Luhn's algorithm
	correct, err := pkg.LuhnValidation(&cardNumber)
	if err == nil && correct {
		LOGGER.Info("Card verification successful for data: " + cardDetails.CardNumber)
		c.JSON(http.StatusOK, structs.Response{Message: "Card Verification Successful!"})
	} else {
		LOGGER.Error("Card verification failed for data: " + cardDetails.CardNumber + "\nReason: " + err.Error())
		c.JSON(http.StatusBadRequest, structs.Error{ErrorMessage: "Card Verification Failed!"})
	}
}
