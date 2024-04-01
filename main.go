package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/machinebox/graphql"
)

const (
	secret    string = "asdasdasda12312"
	qubeToken string = "sqp_a202070d6d59e0fece8cb2806470d30260b04b87"
)

func main() {
	asd
	for i := 0; i < 3; i++ {
		fmt.Println("idx:", i)
		// go myCart()
		// go addItemToCart()
		go customerCheckPostalCode()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("done")
}

func myCart() error {
	now := time.Now()
	defer func() {
		fmt.Println("duration:", time.Since(now))
	}()

	jsonData := map[string]string{
		"query": `
		{
			myCart {
				items {
					id
					item {
						id
						name
					}
					quantity
					subtotal
				}
			}
		}`,
	}
	b, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("error 1:", err)
		return err
	}
	req, err := http.NewRequest("POST", "http://localhost:4002", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("error 2:", err)
		return err
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("accessToken", "yDVy5D173ZiHs/swd1LTb1sdyEccrFTOIE1Cg9pPXuEOY9ZGUsD4oC2Ghht1/g2Y")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("error 3:", err)
		return err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 4:", err)
		return err
	}

	fmt.Println("Data:", string(data))

	return nil
}

func addItemToCart() error {
	now := time.Now()
	defer func() {
		fmt.Println("duration:", time.Since(now))
	}()

	query := `
	{
		addItemToCart(input: {
			orderType: "delivery"
			branchID: 20
			itemID: 19
			quantity: 3
			price: 1
			specifications: ""
			priceTierID: 0
		})
		{
			id
		}
	}`

	client := graphql.NewClient("http://localhost:4002")

	// req, err := http.NewRequest("POST", "http://localhost:4002", bytes.NewBuffer(b2))
	// if err != nil {
	// 	fmt.Println("error 2:", err)
	// 	return err
	// }

	request := graphql.NewRequest(query)
	request.Header.Add("accessToken", "asda;ksdbaisdpiabsd")
	request.Header.Add("content-type", "application/json")

	var response interface{}

	err := client.Run(context.Background(), request, &response)
	if err != nil {
		log.Fatal("err 1:", err)
	}

	fmt.Println("response:", response)

	// fmt.Println("b:", string(b))
	// req.Header.Add("content-type", "application/json")
	// req.Header.Add("accessToken", "asdnkabsbudiabu123123")
	// client := &http.Client{}
	// response, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("error 3:", err)
	// 	return err
	// }
	// defer response.Body.Close()

	// data, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println("error 4:", err)
	// 	return err
	// }

	// fmt.Println("Data:", string(data))

	return nil
}

func addItemToCart_Schema() string {
	return `mutation addItemToCart(
		$input: AddItemToCartInput!
	){
		addItemToCart(input: $input)
		{
			id
		}
	}`
}

type AddItemToCartInput struct {
	OrderType      string  `json:"orderType"`
	BranchID       int     `json:"branchID"`
	ItemID         int     `json:"itemID"`
	Quantity       int     `json:"quantity"`
	Price          float64 `json:"price"`
	Specifications string  `json:"specifications"`
	PriceTierID    int     `json:"priceTierID"`
}

type DataInput struct {
	AddItemToCartInput AddItemToCartInput `json:"addItemToCartInput"`
}

func customerCheckPostalCode() error {
	now := time.Now()
	defer func() {
		fmt.Println("duration:", time.Since(now))
	}()

	jsonData := map[string]string{
		"query": `
		{
			customerCheckPostalCode(
				locationID: 1,
				postalCode: "123123",
			) {
				id
			}
		}`,
	}
	b, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("error 1:", err)
		return err
	}
	req, err := http.NewRequest("POST", "http://localhost:4002", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("error 2:", err)
		return err
	}

	req.Header.Add("content-type", "application/json")
	// req.Header.Add("accessToken", "asdbasiubdasudbpiu2123")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("error 3:", err)
		return err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 4:", err)
		return err
	}

	fmt.Println("Data:", string(data))

	return nil
}
