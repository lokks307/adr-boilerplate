package action

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lokks307/adr-boilerplate/domain"
	"github.com/lokks307/adr-boilerplate/responder"
	"github.com/lokks307/go-util/djson"
)

func EchoHello(ctx echo.Context) error {
	return responder.TxtResponder.Response(ctx, http.StatusOK, "Hello ADR!")
}

func GetCustomerInfo(ctx echo.Context) error {
	cid := ctx.Param("customer_id")
	cidInt, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return responder.TxtResponder.Response(ctx, http.StatusBadRequest, "")
	}

	customer, err := domain.ReadCustomerByID(domainConn, cidInt)
	if err != nil {
		return responder.TxtResponder.Response(ctx, http.StatusNotFound, "")
	}

	res := djson.NewDJSON().Put(djson.Object{
		"id":      customer.CustomerId,
		"name":    customer.FirstName + " " + customer.LastName,
		"company": customer.Company.String,
		"address": customer.Address.String,
		"city":    customer.City.String,
		"State":   customer.State.String,
		"phone":   customer.Phone.String,
	})

	return responder.JsonResponder.Response(ctx, http.StatusOK, res)
}
