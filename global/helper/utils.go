package helper

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"ariefsn/game-currency/global/models"

	"github.com/fatih/structs"
	"gorm.io/gorm"
)

func StatusText(code int) string {
	statusText := "Unknown Error"

	switch code {
	case 422:
		statusText = "Error Rendering Response"
	case 400:
		statusText = "Bad Request"
	case 401:
		statusText = "Unauthorized"
	case 403:
		statusText = "Forbidden"
	case 404:
		statusText = "Not Found"
	case 405:
		statusText = "Method Not Allowed"
	case 500:
		statusText = "Internal Server Error"
	case 502:
		statusText = "Bad Gateway"
	case 503:
		statusText = "Server Unavailable"
	}

	return statusText
}

func ParsePayload(r *http.Request, target interface{}) error {
	return json.NewDecoder(r.Body).Decode(&target)
}

func MapToCase(v map[string]interface{}, c string) map[string]interface{} {
	if c == "" {
		c = "c"
	}

	for key, val := range v {
		if reflect.TypeOf(val).Kind() == reflect.Map { // nested map
			val = MapToCase(val.(map[string]interface{}), c)
		}
		delete(v, key)
		newKey := key
		switch strings.ToLower(c) {
		case "l", "lower", "low":
			newKey = strings.ToLower(key)
		case "u", "upper", "up":
			newKey = strings.ToUpper(key)
		case "c", "camel", "cam":
			split := strings.Split(key, "")
			split[0] = strings.ToLower(split[0])
			newKey = strings.Join(split, "")
		}
		if strings.ToLower(key) == "id" {
			newKey = "id"
		}

		v[newKey] = val
	}
	return v
}

func HideFields(data interface{}, fields ...string) map[string]interface{} {
	m := MapToCase(structs.Map(data), "")

	for k := range m {
		for _, f := range fields {
			if f == k {
				delete(m, k)
				break
			}
		}
	}

	return m
}

func BuildWhereClause(wheres ...models.FilterModel) (string, []interface{}) {
	clause := []string{}
	values := []interface{}{}

	for _, w := range wheres {
		c, v := w.BuildClause()

		clause = append(clause, c)
		values = append(values, v)
	}

	return strings.Join(clause, " AND "), values
}

func CountAllRows(g *gorm.DB) int {
	var count int64
	g.Count(&count)

	return int(count)
}

func GetFilterField(r *http.Request, field string, opr models.FilterOpr) *models.FilterModel {
	val := r.URL.Query().Get(field)

	if opr == "" {
		opr = models.OprEq
	}

	if val != "" {
		return models.NewFilterModel(field, opr, val)
	}

	return nil
}

func AppendFilterField(r *http.Request, field string, opr models.FilterOpr, filters []models.FilterModel) []models.FilterModel {
	val := GetFilterField(r, field, opr)

	if val != nil {
		filters = append(filters, *val)
	}

	return filters
}

func GetSkipLimit(r *http.Request) (int, int) {
	s := r.URL.Query().Get("skip")
	l := r.URL.Query().Get("limit")

	var skip int = 0
	var limit int = 0

	if s != "" {
		skip, _ = strconv.Atoi(s)
	}

	if l != "" {
		limit, _ = strconv.Atoi(l)
	}

	return skip, limit
}

func GetOrders(r *http.Request) string {
	o := r.URL.Query().Get("orders")

	if o != "" {
		return strings.ReplaceAll(o, "*", " ")
	}

	return ""
}

func RandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func GetTimezone() *time.Location {
	tz := "Asia/Jakarta"

	if GetEnv().Db.Timezone != "Local" {
		tz = GetEnv().Db.Timezone
	}

	tz = strings.ReplaceAll(tz, "%2F", "/")

	loc, _ := time.LoadLocation(tz)

	return loc
}

func GetFullUrlQuery(r *http.Request) string {
	query := r.URL.RawQuery

	if query != "" {
		return "?" + query
	}

	return query
}

func AddContentDisposition(isDownload bool, rw http.ResponseWriter, fileName string) {
	if isDownload {
		contentDisposition := fmt.Sprintf("attachment; filename=%s", fileName)
		rw.Header().Set("Content-Disposition", contentDisposition)
	}
}
