package controllers

// func Mercadopago() {
// 	response, mercadopagoErr, err := mercadopago.CreatePayment(mercadopago.PaymentRequest{
// 		ExternalReference: "seu-id-interno-0001",
// 		Items: []mercadopago.Item{
// 			{
// 				Title:     "Pagamendo mensalidade PagueTry",
// 				Quantity:  1,
// 				UnitPrice: 50,
// 			},
// 		},
// 		Payer: mercadopago.Payer{
// 			Identification: mercadopago.PayerIdentification{
// 				Type:   "CPF",
// 				Number: "12345678909",
// 			},
// 			Name:    "Eduardo",
// 			Surname: "Mior",
// 			Email:   "eduardo-mior@hotmail.com",
// 		},
// 		NotificationURL:   "https://localhost/webhook/mercadopago",
// 	}, "seu-access-token")

// 	if err != nil {
// 			ctx.JSON(400, errors.New("add pedido error"))
// 			return
// 	} else if mercadopagoErr != nil {
// 		// Erro retornado do MercadoPago
// 	} else {
// 		// Sucesso!
// 	}
// }
