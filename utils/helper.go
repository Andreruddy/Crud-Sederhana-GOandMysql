package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
)

func GetMD5Hash(password string) string {
	hash := md5.Sum([]byte(password))
	fmt.Println(hash[:])
	return hex.EncodeToString(hash[:])
}

func Paginate(req *http.Request, limit int) (int, int) {
	keys := req.URL.Query()
	if keys.Get("page") == "" {
		return 1, 0
	}
	page, _ := strconv.Atoi(keys.Get("page"))
	if page < 1 {
		return 1, 0
	}
	begin := (limit * page) - limit
	return page, begin
}