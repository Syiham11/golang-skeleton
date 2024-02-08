package helper

import (
	"strconv"
)

func Invoicetempalte(buyername, event, orderid, ticket, qty, promo, priceticket string, tax, platform_fee, type_platform_fee int) string {
	price, _ := strconv.Atoi(priceticket)
	qtys, _ := strconv.Atoi(qty)
	promos, _ := strconv.Atoi(promo)

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

	tot := (price * qtys) + taxResult + convert_platform_fee - promos
	total := strconv.Itoa(tot)
	htmlString := `<html>
	<head>
	  <meta name="viewport" content="width=device-width,initial-scale=1.0"/>
	  <meta charset="utf-8"/>
	  <title>Invoice</title>
	  <style>
	  body {
		font-family: 'San Francisco', Arial, Helvetica, san-serif;
		font-size: 14px;
		margin: 0;
	  }
	
	  * {
		box-sizing: border-box;
	  }
	
	  table {
		border-spacing: 0;
		border-collapse: collapse;
		width: 100%;
		max-width: 600px;
	  }
	
	  .invoice-table-wrapper {
		width: 100%;
		height: 100%;
		max-width: 650px;
		border-spacing: 0;
		border-collapse: collapse;
		margin: 0 auto;
		background: #F2F2F2;
	  }
	
	  .invoice-table {
		width: 100%;
		height: 100%;
		max-width: 600px;
		border-spacing: 0;
		border-collapse: collapse;
		border: 1px solid #E6E7EB;
		margin: 0 auto;
	  }
	
	  .logo {
		margin: 0;
		padding: 20px;
		text-align: left;
	  }
	
	  .invoice-header-table {
		border-spacing: 0;
		border-radius: 10px 10px 0 0;
		border-collapse: collapse;
		width: 100%;
		border: none;
		height: 64px;
		background-color: #F9AE3C;
	  }
	
	  .customer-name {
		padding: 24px 0px 0px 16px;
	  }
	
	  .amount-title {
		font-size: 11px;
		color: #AAA;
		margin: 10px 0;
	  }
	
	  .amount-paid {
		font-weight: 600;
		font-size: 32px;
		display: inline-block;
		padding: 0 10px 10px 0;
		margin: 0;
		vertical-align: -webkit-baseline-middle;
	  }
	
	  .payment-type {
		display: inline-block;
		border-radius: 4px;
		padding: 5px 10px;
		color: white;
		font-size: 11px;
		letter-spacing: 0.3px;
		background: #39B351;
		margin: 0;
	  }
	
	  .payment-type.payment-type-go {
		background: #5CA5DA;
	  }
	
	  .payment-tip {
		color: #777;
		padding: 10px 0 25px;
		border-bottom: 1px solid #E6E7EB;
		margin: 0 0 10px;
		font-size: 13px;
	  }
	
	  .order-wrapper {
		font-size: 14px;
		font-weight: 600;
		margin: 0;
		padding: 0 20px;
	
	  }
	
	  .promotion_banner {
		width: 100%;
		padding: 0 25px;
	  }
	
	  .order-wrapper span {
		color: #AAA;
		font-size: 11px;
		text-transform: uppercase;
		display: block;
		padding-bottom: 7px;
	  }
	
	  .pickup-time {
		color: #178FFC;
		text-transform: uppercase;
		font-size: 11px;
		vertical-align: text-top;
		padding-left: 8px;
	  }
	
	  .pickup-time strong {
		color: #333;
	  }
	
	  .drop-time {
		color: #EF8F03;
		text-transform: uppercase;
		font-size: 11px;
		vertical-align: text-top;
		padding-left: 8px;
	  }
	
	  .drop-time strong {
		color: #333;
	  }
	
	  .location-details {
		padding: 0 22px;
		margin: 10px 6px;
		font-size: 13px;
		line-height: 1.4;
	  }
	
	  .location-title {
		font-weight: bold;
		text-transform: uppercase;
		display: block;
		padding-bottom: 10px;
		font-size: 14px;
		line-height: 1;
	  }
	
	  .driver-phone-number {
		color: #777777;
		font-size: 10px;
	  }
	
	  .order-rating-text {
		background: #EEE;
		color: #999;
		font-size: 10px;
		display: inline-block;
		border-radius: 5px;
		padding: 5px;
		margin-top: 0px;
	  }
	
	  .order-summary-title {
		font-size: 11px;
		font-weight: 600;
		color: #999;
		padding: 26px 0px 5px;
		border-top: 1px solid #E6E7EB;
	  }
	
	  .order-summary-details {
		font-size: 13px;
		color: #777777;
		padding-bottom: 18px;
	  }
	
	  .order-summary-value {
		font-size: 14px;
		text-align: right;
		font-weight: 500;
		padding-bottom: 18px;
	  }
	
	  .total-price {
		font-size: 12px;
		font-weight: bold;
		padding: 15px 0;
	  }
	
	  .total-price-value {
		font-size: 16px;
		font-weight: bold;
		text-align: right;
		padding: 15px 0;
	  }

	  .item:not(:last-child)::after { content: ',' }
	
	  .paid-with-text {
		font-size: 11px;
		color: #777;
		padding: 0 10px 10px 0;
	  }
	
	  .help-text {
		font-size: 11px;
		color: #999;
		text-align: -webkit-right;
		padding: 5px 25px;
		font-weight: 300;
	  }
	
	  .address {
		font-size: 11px;
		color: #999;
		text-align: -webkit-center;
		padding: 5px 0;
	  }
	
	  .service_logo{
		width: 25%;
	  }
	
	  .brand_logo{
		margin-top: 25px;
	  }
	
	  .social_site_img{
		text-decoration: none
	  }
	
	  @media only screen and (max-width: 500px) {
		td.full {
		  display: block !important;
		  width: 100% !important;
		}
	
		.driver-details {
		  padding: 10px !important;
		}
	
		.order-wrapper {
		  font-size: 12px !important;
		  padding: 0 0 0 20px !important;
		  vertical-align: -webkit-baseline-middle !important;
		}
	  }
	
	  @media only screen and (max-width: 640px) {
		.promotion_banner {
		  padding: 0 20px;
		}
	  }
	
	  @media screen and (min-width: 500px) {
		.card-table {
		  width: 50% !important;
		}
	
		.user-info {
		  padding: 10px 0 0 10px !important;
		}
	  }
	
	</style>
	
	</head>
	<body style="font-family: 'San Francisco', Arial, Helvetica, san-serif,serif;
		font-size:14px;
		margin:0;
		box-sizing: border-box;
		color: black">
	<table style="width: 100%;
		height: 100%;
		max-width: 650px;
		border-spacing: 0;
		border-collapse: collapse;
		margin: 0 auto;
		background: #F2F2F2;" align="center">
	  <tbody>
	  <tr>
		<td style="padding:20px;">
		  <table style="width: 100%;
			  height: 100%;
			  max-width: 600px;
			  border-spacing: 0;
			  border-collapse: collapse;
			  border: 1px solid #E6E7EB;
			  margin: 0 auto;" align="center">
			<tbody>
			<tr>
	  <td>
		<table width="100%" align="center"
			   style="border-spacing: 0; border-collapse: collapse; width:100%; height: 100%;">
		  <thead>
		  <tr  style='background-image: url(https://web-dev.eventori.id/bgetiket.png);
		  height: 70px;
		  background-position: center;
		  background-repeat: no-repeat;
		  background-size: cover;'>
			<th width="100%">
			  <h2 style="margin:0; padding:20px; text-align: center;">
				<img alt="Eventori" src="https://storage.googleapis.com/st-core/public/temp/1660028139-logo-putih 1 (1).png" style="width:200px; height: 100%; object-fit: contain;" class="service_logo">
			  </h2>
			</th>
			
		  </tr>
		  </thead>
		</table>
	  </td>
	</tr>
	
	<tr bgcolor="white">
	  <td style="padding:32px 32px 10px 20px;">
		<table style="border-spacing: 0; border-collapse: collapse; width:100%;">
		  <tbody>
		  <tr>
			<span style="font-size:18px; color: #1c1d1d; font-weight:bold; padding-left: 20px;">
			  
			  Hallo ` + buyername + `,
			</span>
		  </tr>
		  <tr>
			<p style="color: #494A48; font-size: 13px; font-weight: 300; padding-left: 20px;">
			  Selamat Pembayaran Anda Berhasil, Silahkan melihat bukti pembayaran Anda di attachment
			</p>
		  </tr>
		  </tbody>
		</table>
	  </td>
	</tr>
	
	<tr bgcolor="white">
	  <td style="padding-left: 20px;">
		<table width="100%" align="center" style="border-spacing: 0; border-collapse: collapse; padding-bottom: 8px;">
		  <tbody>
		  <tr>
			<td width="90%">
			  <div style="margin: 0; padding: 0 20px;"
				  class="order-wrapper">
				<span style="color:#727272; font-size: 13px; text-transform: uppercase; display: block; padding-bottom: 4px; font-weight: 300;">
				  DETAIL PEMESANAN
				</span>
				<div style="color: #1c1d1d; font-size: 10px;  height: 20px; line-height: 10px; margin-bottom: 10px;">
	  			` + event + `
				</div>
			  </div>
			  <div style="color: #727272;height: 16px; font-weight: bold; line-height: 16px; font-size: 10px; text-transform: uppercase; display: block; padding-left:20px; padding-bottom: 7px; margin-top: 20px; margin-bottom: 10px;">
				No.Pemesanan: ` + orderid + `
			  </div>
			  <div style="margin: 0; padding: 12px 0 0 0; font-weight: 300;">
			  </div>
			</td>
			<td width="10%" style="padding:0 20px;">
			  <img src="https://ops-service-production.s3.amazonaws.com/assets/order-details.png"
				   alt="Order Details"
				   style="vertical-align: -webkit-baseline-middle; float:right;"/>
			</td>
		  </tr>
		  </tbody>
		</table>
	  </td>
	</tr>
	
	
	<tr bgcolor="white">
	  <td style="padding:0 20px;">
		<h2 style="font-size: 11px; font-weight: 600; color: #999; padding: 7px 0 5px; border-top: 1px solid #E6E7EB;">
		</h2>
	  </td>
	</tr>
	
	<tr bgcolor="white">
	  <td style="padding:0 20px;">
		<table width="100%" align="center" style="border-spacing: 0; border-collapse: collapse;">
		  <tbody>
				
				<tr className='item'>
				<td style="font-size: 11px; color: #494a4a; padding-left:20px; padding-bottom: 12px; font-weight: 300;">
				` + ticket + ` - (` + qty + ` Qty)
				</td>
				<td style="font-size: 11px; color: #494a4a; text-align: right; font-weight: 300; padding-right: 20px; padding-bottom: 12px;">
				  ` + priceticket + ` 
				</td>
			  </tr>
			  
			  <tr className='item'>
				<td style="font-size: 11px; color: #494a4a; padding-left:20px; padding-bottom: 12px; font-weight: 300;">
					Biaya Layanan
				</td>
				<td style="font-size: 11px; color: #494a4a; text-align: right; font-weight: 300; padding-right: 20px; padding-bottom: 12px;">
				  ` + strconv.Itoa(convert_platform_fee) + ` 
				</td>
			  </tr>
			  
			  <tr className='item'>
				<td style="font-size: 11px; color: #494a4a; padding-left:20px; padding-bottom: 12px; font-weight: 300;">
					Pajak
				</td>
				<td style="font-size: 11px; color: #494a4a; text-align: right; font-weight: 300; padding-right: 20px; padding-bottom: 12px;">
				  ` + strconv.Itoa(taxResult) + ` 
				</td>
			  </tr>

			  <tr className='item'>
				<td style="font-size: 11px; color: #494a4a; padding-left:20px; padding-bottom: 12px; font-weight: 300;">
					Promo
				</td>
				<td style="font-size: 11px; color: #494a4a; text-align: right; font-weight: 300; padding-right: 20px; padding-bottom: 12px;">
				  ` + promo + ` 
				</td>
			  </tr>
			
		  <tr style="border-bottom: 1px dashed #cacccf;">
			<td style="color:#1c1d1d;font-size: 11px; font-weight: bold; padding-left:20px; padding-top: 20px; padding-bottom:16px;">
			  TOTAL PEMBAYARAN
			</td>
			<td style="color:#1c1d1d;font-size: 11px; font-weight: bold; padding-top: 20px; text-align: right; padding-right: 20px; padding-bottom:16px;">
			  ` + total + `
			</td>
		  </tr>
		  <br> <br> <br>
	
	
		  
		  </tbody>
		</table>
	  </td>
	</tr>
	
	
			</tbody>
		  </table>
		</td>
	  </tr>
	
	  <tr>
	  <td>
		<p style="font-size: 11px; color: #999; text-align: right; padding: 5px 25px; font-weight: 300;">
		  <span style="font-weight:500;">Butuh bantuan? 
			<a target="_blank" href="https://eventori.id/contact-us">Kunjungi halaman</a> Bantuan di aplikasi eventori</span>
		</p>
	  </td>
	</tr>
	
	<tr bgcolor="#E8E8E8">
	  <td style="padding:0 20px;">
		<h2 style="text-align:center;margin: 0;padding: 10px 0 3px 0;">
			<img alt="Eventori" src="https://storage.googleapis.com/st-core/public/temp/1659940266-Group.png" style="height: 50px; width:200px; object-fit: contain;" class="service_logo">
		</h2>
		<p style=" font-size: 11px; color: #999; text-align: center; padding: 5px 0;">
			Nifarro Park, ITS Tower Jl. Raya Pasar Minggu Jakarta Selatan 12510 DKI Jakarta, Indonesia
		</p>
		<h2 style="margin: 0;text-align: center;padding: 8px 0 8px 0; background-color: black;">
			  <a target="_blank" style="text-decoration: none;" href= "https://www.facebook.com/eventori.id">
				<img alt="Facebook" class="social_site_img" width="30px" height="20px" style="object-fit: contain;" src="https://storage.googleapis.com/st-core/public/temp/1660019617-Icon awesome-facebook-f.png" />
			  </a>
			  <a target="_blank" style="text-decoration: none;" href= "https://www.instagram.com/eventori.id/">
				<img alt="Instagram" class="social_site_img"  width="30px" height="20px" style="object-fit: contain;" src="https://storage.googleapis.com/st-core/public/temp/1660019632-Icon awesome-instagram.png" />
			  </a>
			  <a target="_blank" style="text-decoration: none;" href= "https://www.youtube.com/channel/UChtIWkfidv0MzIFcAO3Oimg">
				<img alt="Youtube" class="social_site_img"  width="30px" height="20px" style="object-fit: contain;" src="https://storage.googleapis.com/st-core/public/temp/1660019598-Icon awesome-youtube.png" />
			  </a>
			  
		</h2>
	  </td>
	</tr>
	
	  </tbody>
	</table>
	</body>
	</html>`

	return htmlString
}
