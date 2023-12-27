package helper

import (
	"strconv"
	"strings"
	"time"
)

func GenerateInvoice(tipe, max int) string {
	//1. Prepaid : INV/20230912/PR/123 = langsung bayar di bedakan berdasarakn jenis user dia member
	//2. Postpaid : INV/20230912/PS/123 = pasti invoice di bedakan berdasarakn jenis user
	inv := ""
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	removeChar := strings.Replace(date, "-", "", -1)
	if tipe == 1 {
		inv = "INV/" + removeChar + "/PR/" + strconv.Itoa(max)
	}
	if tipe == 2 {
		inv = "INV/" + removeChar + "/PS/" + strconv.Itoa(max)
	}

	return inv
}
