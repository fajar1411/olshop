package helper

import (
	"net/http"
	"strings"
)

func PesanGagalHelper(msg string) (int, map[string]any) {
	var code int
	resp := map[string]interface{}{}
	if msg != "" {
		resp["message"] = msg
	}

	if strings.Contains(msg, "server") {
		code = http.StatusInternalServerError
	} else if strings.Contains(msg, "format") {
		code = http.StatusBadRequest
	} else if strings.Contains(msg, "not found") || strings.Contains(msg, "belum terdaftar") {
		code = http.StatusNotFound
	} else if strings.Contains(msg, "password") {
		code = http.StatusUnauthorized
	} else if strings.Contains(msg, "sudah terdaftar") {
		code = http.StatusConflict
	} else if strings.Contains(msg, "required") {
		code = http.StatusBadRequest
	}
	return code, resp
}
func PesanSuksesHelper(msg string) map[string]any {
	return map[string]any{
		"Status": "Berhasil",
		"MSG":    msg,
	}
}
func PesanDataBerhasilHelper(msg string, data any) map[string]any {
	return map[string]any{

		"data": data,
		"MSG":  msg,
	}
}
