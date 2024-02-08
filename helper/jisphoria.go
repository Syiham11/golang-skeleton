package helper

func ProrderJishphoria(category, first_name string) string {

	htmlString := `<html lang="en">
	<head>
	  <meta charset="UTF-8" />
	  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
	  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
	  <title>Document</title>
	  <style>
		.bgEmail {
		  background-image: url("https://storage.googleapis.com/st-core/public/temp/1662359676-bgEmail.png");
		  background-repeat: no-repeat;
		  width: 600px;
		  height: 1024px;
		  background-position: center;
		  margin: auto;
		}
		body {
		  margin: 0;
		}
		.wrapperLogo {
		  padding-top: 40px;
		}
		.logo {
		  display: block;
		  margin: 0 auto;
		}
		.banner {
		  width: 450px;
		  height: 400px;
		}
		.wrapperBanner {
		  margin: auto;
		  padding-top: 28px;
		}
		.wrapperDescription {
		  height: auto;
		  background: white;
		  width: 450px;
		  margin: auto;
		  border-radius: 0px 0px 30px 30px;
		  box-sizing: border-box;
		  padding: 16px;
		}
		.textTitle {
		  font-size: 18px;
		  font-weight: 700;
		  color: black;
		}
		.textDescription {
		  font-size: 12px;
		  padding-top: 20px;
		}
		.wrapperSocmed {
		  text-align: center;
		  margin-top: 16px;
		  gap: 10px;
		}
	  </style>
	</head>
	<body>
	  <div class="bgEmail">
		<div class="wrapperLogo">
		  <img
			class="logo"
			src="https://storage.googleapis.com/st-core/public/temp/1662359906-logoEmail.png"
			alt="image logo"
		  />
		</div>
		<div class="wrapperBanner">
		  <img
			class="logo banner"
			src="https://storage.googleapis.com/st-core/public/temp/1662360084-bannerEmail.png"
			alt="image banner"
		  />
		  <div class="wrapperDescription">
			<div class="textTitle">The Ticket(s) is in your hand!</div>
			<div class="textDescription">
			  Congratulations, <strong>` + first_name + `</strong> You’re successfully
			  registered to buy the pre-ordered ticket! You’re only allowed to buy
			  this category ticket.<br />
			  <br />
			  Category: <strong>` + category + `</strong> <br />
			  <br />
			  Your pre-ordered ticket will be available on
			  <strong>HardcodeDate</strong>. You can buy your ticket either at
			  <strong>Eventori vTicket</strong> or <strong>Gotix/Loket!</strong>
			  <br /><br />
			  Thank you!
			</div>
		  </div>
		  <div class="wrapperSocmed">
			<img
			  src="https://storage.googleapis.com/st-core/public/temp/1662361398-iconIG.png"
			  alt="logo instagram"
			/>
			<img
			  src="https://storage.googleapis.com/st-core/public/temp/1662361131-iconFB.png"
			  alt="logo facebook"
			/>
			<img
			  src="https://storage.googleapis.com/st-core/public/temp/1662361436-iconTiktok.png"
			  alt="logo tiktok"
			/>
		  </div>
		</div>
	  </div>
	</body>
  </html>`

	return htmlString
}
