package helper

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GenerateInvoice(buyername, event, orderid, ticket, qty, promo, priceticket, venue, date string, tax, platform_fee, type_platform_fee int) error {
	price, _ := strconv.Atoi(priceticket)
	qtys, _ := strconv.Atoi(qty)
	promos, _ := strconv.Atoi(promo)

	tot := (price * qtys) - promos
	total := strconv.Itoa(tot)
	totawal := price * qtys
	totalharga := strconv.Itoa(totawal)
	hitungtax := (float64(price) * float64(qtys)) * (float64(tax) / 100)
	taxResult := int(hitungtax)

	platform_fee_sum := float64(0)

	if platform_fee != 0 && type_platform_fee == 1 {
		platform_fee_sum = (float64(price) * float64(qtys)) * (float64(platform_fee) / 100)
	}

	if platform_fee != 0 && type_platform_fee == 2 {
		platform_fee_sum = float64(platform_fee)
	}

	convert_platform_fee := int(platform_fee_sum)

	hitungtotal := (price * qtys) + int(hitungtax) + convert_platform_fee
	totalbelanja := strconv.Itoa(hitungtotal)
	hitungtagihan := hitungtotal - promos
	totaltagihan := strconv.Itoa(hitungtagihan)
	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		return err
	}
	htmlStr := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<link
		href="https://fonts.googleapis.com/css2?family=Poppins:wght@400;500;600;700&display=swap"
		rel="stylesheet"
		/>
		<title>Document</title>
		<style>
		body {
			font-family: "Poppins", serif;
			font-size: 12px;
		}
		.flex {
			display: -webkit-box;
		}
		.flex-col {

		}
		.justify-center {
			/* -webkit-justify-content: center;
			justify-content: center; */
			-webkit-box-pack: center; 
		}
		.justify-between {
			/* -webkit-justify-content: space-between;
			justify-content: space-between; */
			-webkit-box-pack: justify; 
		}
		.justify-start {
			/* -webkit-justify-content: flex-start;
			justify-content: flex-start; */
			-webkit-box-pack: start; 
		}
		.justify-end {
			/* -webkit-justify-content: flex-end;
			justify-content: flex-end; */
			-webkit-box-pack: end; 
		}
		.items-center {
			/* -webkit-align-items: center;
			align-items: center; */
			-webkit-box-align: center;
		}
		.items-end {
			-webkit-box-align: end;
		}
		.w-full {
			width: 100%;
		}
		.mx-auto {
			margin: auto;
		}
		.my-10 {
			margin-top: 40px;
			margin-bottom: 40px;
		}
		.m-5 {
			margin: 20px;
		}
		.bg-white {
			background-color: white;
		}
		.text-orange {
			color: #f6921e;
		}
		.font-bold {
			font-weight: 700;
		}
		.bg-image {
			background-image: url("https://storage.googleapis.com/st-core/public/temp/1659939985-image 14.png");
			background-size: contain;
			background-repeat: no-repeat;
			background-position: center;
		}
		.img-logo {
			background-image: url("https://storage.googleapis.com/st-core/public/temp/1659940266-Group.png");
			background-size: contain;
			background-repeat: no-repeat;
			background-position: left;
			width: 165px;
			height: 50px;
		}
		.gap-8 {
			gap: 32px;
		}
		.font-normal {
			font-weight: 500;
		}
		</style>
	</head>
	<body>
		<div style="width: 595px; height: 842px" class="mx-auto bg-white bg-image">
		<div class="m-5">
			<div class="flex justify-between my-10">
			<div class="img-logo"></div>
			<div style="margin-top: 12px;">
				<div class="font-bold">Invoice</div>
				<div class="text-orange font-normal">INV/220620/TKT/44150001239</div>
			</div>
			</div>
			<div class="flex justify-start font-bold">Diterbitkan Atas Nama</div>
			<div style="margin-top: 10px" class="flex">
			<div style="width: 200px" class="font-normal">Penjual</div>
			<div style="width: 300px" class="font-normal">: PT. Eventori Media Semesta</div>
			</div>
			<div style="margin-top: 20px" class="font-bold">Untuk</div>
			<div style="margin-top: 10px" class="flex">
			<div style="width: 200px" class="font-normal">Pembeli</div>
			<div style="width: 300px" class="font-normal">: ` + buyername + `</div>
			</div>
			<div style="margin-top: 10px" class="flex">
			<div style="width: 200px" class="font-normal">Tanggal Pembelian</div>
			<div style="width: 300px" class="font-normal">: ` + date + `</div>
			</div>
			<div style="margin-top: 10px" class="flex">
			<div style="width: 200px" class="font-normal">Event</div>
			<div style="width: 300px" class="font-normal">: ` + event + `</div>
			</div>
			<div
			style="border: 2px dashed #000000; margin: 10px 0px 10px 0px"
			></div>
			<div style="margin: 0px 5px 0px 5px" class="flex justify-between">
			<div class="font-bold" style="width: 155px;">Kategori Tiket</div>
			<div class="flex justify-between font-bold" style="width: 200px;">
				<div style="width: 100px;">Jumlah</div>
				<div style="width: 100px;">Harga Satuan</div>
			</div>
			<div class="font-bold">Total Harga</div>
			</div>
			<div
			style="border: 2px dashed #000000; margin: 10px 0px 10px 0px"
			></div>
			<div style="margin: 0px 5px 0px 5px" class="flex justify-between">
			<div class="text-orange font-bold" style="width: 155px">` + ticket + `</div>
			<div class="flex justify-between" style="width: 200px;">
				<div style="width: 100px;" class="font-normal">` + qty + `</div>
				<div style="width: 100px;" class="font-normal">Rp ` + priceticket + `</div>
			</div>
			<div class="font-normal">Rp ` + total + `</div>
			</div>
			<div
			style="
				margin: 0px 5px;
				width: 150px;
				font-size: 10px;
				padding-top: 5px;
			"
			class="font-normal"
			>
			<strong>Venue : </strong> ` + venue + `
			</div>
			<div
			style="border: 1px dashed #000000; margin: 10px 0px 10px 0px"
			></div>
			<div class="flex justify-end" style="gap: 50px">
			<div class="font-bold" style="font-size: 12px; width: 200px;">TOTAL HARGA TIKET (` + qty + ` TIKET)</div>
			<div class="font-bold" style="font-size: 12px; width: 100px; text-align: end;">Rp ` + totalharga + `</div>
			</div>
			<div class="flex justify-end" style="gap: 50px; margin-top: 10px;">
			<div style="width: 200px;" class="font-normal">Biaya Layanan</div>
			<div style="width: 100px; text-align: end;" class="font-normal">Rp ` + strconv.Itoa(convert_platform_fee) + `</div>
			</div>
			<div class="flex justify-end" style="gap: 50px; margin-top: 10px;">
			<div style="width: 200px;" class="font-normal">Pajak</div>
			<div style="width: 100px; text-align: end;" class="font-normal">Rp ` + strconv.Itoa(taxResult) + `</div>
			</div>
			<div class="flex justify-end">  
			<div
			style="border: 1px dashed #000000; margin: 10px 0px 10px 0px; width: 300px;"
			></div>
			</div>
			<div class="flex justify-end" style="gap: 50px">
			<div class="font-bold" style="font-size: 12px; width: 200px;">TOTAL BELANJA</div>
			<div class="font-bold" style="font-size: 12px; width: 100px; text-align: end;">Rp ` + totalbelanja + `</div>
			</div>
			<div class="flex justify-end" style="gap: 50px; margin-top: 10px;">
			<div style="width: 200px;" class="font-normal">PROMO</div>
			<div style="width: 100px; text-align: end;" class="font-normal">- Rp ` + promo + `</div>
			</div>
			<div class="flex justify-end">  
			<div
			style="border: 1px dashed #000000; margin: 10px 0px 10px 0px; width: 300px;"
			></div>
			</div>
			<div class="flex justify-end" style="gap: 50px">
			<div class="font-bold" style="font-size: 12px; width: 200px;">TOTAL TAGIHAN</div>
			<div class="font-bold" style="font-size: 12px; width: 100px; text-align: end;">Rp ` + totaltagihan + `</div>
			</div>
			<div class="flex justify-end">  
			<div
			style="border: 1px dashed #000000; margin: 10px 0px 10px 0px; width: 100%;"
			></div>   
			</div>
			</div>
			<div class="flex items-end" style="height: 200px; width: 350px;">
			<div class="font-normal">Invoice ini sah dan diproses oleh komputer
				Silahkan hubungi <strong class="text-orange">
				Eventori CS
				</strong>  apabila perlu bantuan</div>
			</div>
		</div>
		</div>
	</body>
	</html>`

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(htmlStr)))
	pdfg.MarginLeft.Set(5)
	pdfg.MarginRight.Set(5)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	fmt.Println(err)

	if err != nil {
		log.Fatal(err)
	}

	//Your Pdf Name
	err = pdfg.WriteFile("helper/storage/" + orderid + ".pdf")
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("Done")
	return nil

}
