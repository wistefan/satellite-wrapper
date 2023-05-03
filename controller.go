package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	g "wistefan/satellite-wrapper/generated"

	"github.com/fiware/VCVerifier/logging"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrorNoDID           = errors.New("no_did_param")
	ErrorInvalidPageSize = errors.New("invalid_page_size_param")
)

const (
	DEFAULT_PAGE_SIZE = 100
)

// A Route defines the parameters for an api endpoint
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc fiber.Handler
}

// Routes are a collection of defined api endpoints
type Routes []Route

// Router defines the required methods for retrieving api routes
type Router interface {
	Routes() Routes
}

type ApiController struct {
	service      g.DefaultApiServicer
	errorHandler ErrorHandler
}

func NewApiController(s g.DefaultApiServicer) Router {
	return &ApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}
}

func (a *ApiController) Routes() Routes {
	return Routes{
		{
			"GetIssuer",
			strings.ToUpper("Get"),
			"/v3/issuers/:did",
			a.GetIssuer,
		},
		{
			"GetIssuers",
			strings.ToUpper("Get"),
			"/v3/issuers",
			a.GetIssuers,
		},
	}
}

func (a *ApiController) GetIssuer(ctx *fiber.Ctx) (err error) {
	logging.Log().Warn("GET")
	didParam := ctx.Params("did")
	if didParam == "" {
		logging.Log().Warn("Received a query without a did.")
		return ErrorNoDID
	}
	resp, err := a.service.GetIssuer(ctx.Context(), didParam)
	if err != nil {
		return err
	}
	ctx.Status(http.StatusOK).JSON(resp)
	return
}

func (a *ApiController) GetIssuers(ctx *fiber.Ctx) (err error) {
	var pageSize = DEFAULT_PAGE_SIZE
	pageSizeParam := ctx.Query("page[size]")
	pageAfter := ctx.Query("page[after]")
	if pageSizeParam != "" {
		logging.Log().Debug("Received a query without a page size.")
		pageSize, err = strconv.Atoi(pageSizeParam)
		if err != nil {
			logging.Log().Debugf("Received an invalid pageSize %s", pageSizeParam)
			return ErrorInvalidPageSize
		}
	}
	resp, err := a.service.GetIssuers(ctx.Context(), float32(pageSize), pageAfter)
	if err != nil {
		return err
	}
	ctx.Status(http.StatusOK).JSON(resp)
	return
}
