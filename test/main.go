package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type Item struct {
	ItemName  string
	ProductID int
	Quantity  int
	ItemValue float64
	SKUCode   string
}

func parseItem(s string) (*Item, error) {
	item := &Item{}
	//pairs := strings.Split(s, ", ")
	//fmt.Println("pairs", s)
	re := regexp.MustCompile(`(\w+): ([^,]*)(, )?`)
	matches := re.FindAllStringSubmatch(s, -1)
	fmt.Println("lens", len(matches))
	for _, match := range matches {
		//fmt.Println("match", match)

		//fmt.Println("pair", strings.TrimSpace(pair))

		key := match[1]
		//fmt.Println("key", key)
		value := match[2]
		//fmt.Println("value", value)
		switch key {
		case "item_name":
			item.ItemName = value
		case "product_id":
			id, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			item.ProductID = id
		case "quantity":
			qty, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			item.Quantity = qty
		case "item_value":
			val, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, err
			}
			item.ItemValue = val
		case "sku_code":
			item.SKUCode = value
		default:
			return nil, fmt.Errorf("unknown key: %s", key)
		}
	}
	return item, nil
}

func main() {

	testString1 := "item_name: [Sample] Tiered Wire Basket, product_id: 97, quantity: 3, item_value: 119.950000, sku_code: TWB"
	//testString2 := "item_name: Chihuahua1234, product_id: 184211, quantity: 2, item_value: 12.5, sku_code: Chihuahua1234-LA|item_name: Before Before Before Before Before Before Before Before Before Before Before Before Before Before Before Before Before Before Choco nama After After After After After After After After After After Af, product_id: 184213, quantity: 2, item_value: 12.5, sku_code: 0|item_name: PlayStation 5, product_id: 184214, quantity: 10, item_value: 12.5, sku_code: 0|item_name: Forman Linda 7052, product_id: 93933, quantity: 2, item_value: 12.5, sku_code: 0|" +
	//	"item_name: Power Lines Overhead Sign, OSHA Caution Sign, product_id: 93954, quantity: 2, item_value: 12.5, sku_code: 0"
	//if strings.Contains(testString2, "|") {
	//	//arr := strings.Split(testString2, "|")
	//	//fmt.Println("The string contains |", arr)
	//	arr := strings.Split(testString2, "|")
	//	fmt.Println("arr", arr)
	//
	//	for _, s := range arr {
	//		item, err := parseItem(s)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		fmt.Println("item", item)
	//	}
	//}

	item, err := parseItem(testString1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("item", item.ItemName, item.ProductID, item.Quantity, item.ItemValue, item.SKUCode)

	success := 5

	newItem := &Item{
		ItemName:  "",
		ProductID: 0,
		Quantity:  0,
		ItemValue: 0,
		SKUCode:   "",
	}
	result, result2 := test(success, newItem)
	fmt.Println("aaaa", result)
	fmt.Println("bbbb", result2)

}

func test(success int, item *Item) (int, *Item) {
	success += 5

	item = nil

	return success, item
}
