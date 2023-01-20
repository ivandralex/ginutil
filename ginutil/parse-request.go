package ginutil

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//TODO: move checks to validatior

const defaultLimit = "20"
const maxLimit = 100
const defaultOffset = "0"

const defaultThemeID = "0"
const defaultCategoryID = "0"
const defaultPropValue = ""

//ParseURLParamAsUint parses request for uint URL param
func ParseURLParamAsUint(ctx *gin.Context, key string) (uint, error) {
	paramStr := ctx.Param(key)
	return parseUint(paramStr)
}

//ParseURLParamAsUint parses request for uint URL param
func QueryParamAsBool(ctx *gin.Context, key string) (bool, error) {
	paramStr := ctx.Query(key)
	return strconv.ParseBool(paramStr)
}

//ParsePaginationParams parses request for limit and offset params
func ParsePaginationParams(ctx *gin.Context) (uint, uint, error) {
	limitParam := ctx.DefaultQuery("limit", defaultLimit)

	limit, err := parseUint(limitParam)

	if err != nil {
		return 0, 0, fmt.Errorf("invalid limit value (expected positive integer)")
	}

	if limit > maxLimit {
		return 0, 0, fmt.Errorf("max value for limit is 100")
	}

	offsetParam := ctx.DefaultQuery("offset", defaultOffset)

	offset, err := parseUint(offsetParam)

	if err != nil {
		return 0, 0, err
	}

	return limit, offset, nil
}

//ParseUintList parses request for uint list param
func ParseUintList(ctx *gin.Context, key string) ([]uint, error) {
	listStr := ctx.DefaultQuery(key, "")

	strList := strings.Split(listStr, ",")

	intList := []uint{}

	for _, str := range strList {

		id, err := parseUint(str)

		if err != nil {
			return nil, err
		}

		intList = append(intList, id)
	}

	return intList, nil
}

//ParseBodyParamAsUint parses request body for uint param
func ParseBodyParamAsUint(ctx *gin.Context, key string) (uint, error) {
	paramStr := ctx.DefaultPostForm(key, "")
	return parseUint(paramStr)
}

//ParseQueryParamAsUint parses request query for uint param
func ParseQueryParamAsUint(ctx *gin.Context, key string) (uint, error) {
	paramStr := ctx.Query(key)
	return parseUint(paramStr)
}

//ParseQueryParamAsUint parses request query for uint param
func ParseQueryParamAsFloat(ctx *gin.Context, key string) (float64, error) {
	paramStr := ctx.Query(key)
	return parseFloat(paramStr)
}

func parseUint(value string) (uint, error) {
	param, err := strconv.Atoi(value)

	if err != nil || param < 0 {
		return 0, fmt.Errorf("expected string representation of uint")
	}

	return uint(param), nil
}

func parseFloat(value string) (float64, error) {
	param, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return 0, fmt.Errorf("expected string representation of float32")
	}

	return param, nil
}

//ParseBodyParamAsBool parses request body for bool param
func ParseBodyParamAsBool(ctx *gin.Context, key string) (bool, error) {
	paramStr := ctx.DefaultPostForm(key, "")
	return strconv.ParseBool(paramStr)
}
