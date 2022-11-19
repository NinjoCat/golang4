/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес(индекс(ровно 6 цифр), город, улица, дом, квартира)

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

реализовать несколько методов у типа "Накладная"

createReader == NewReader
*/
package main

import (
	"fmt"
	"log"
	"strconv"
)

type Order struct {
	Tittle    string
	Amount    string
	BuyerName string
	Phone     string
	Address
}

type Address struct {
	PostalCode string
	City       string
	Street     string
	House      string
	Apartment  string
}

func newAddress(postalCode, city, street, house, apartment string) *Address {
	address := &Address{
		PostalCode: postalCode,
		City:       city,
		Street:     street,
		House:      house,
		Apartment:  apartment,
	}

	return address
}

func newOrder(title, buyerName, phone, amount string, address Address) *Order {
	order := &Order{
		Tittle:    title,
		BuyerName: buyerName,
		Phone:     phone,
		Amount:    amount,
		Address:   address,
	}

	return order
}

func (o Order) printInfo() {
	fmt.Println("________________________________________________________________")
	fmt.Println("Информация о заказе")
	fmt.Println("Название заказа: ", o.Tittle)
	fmt.Println("Имя покупателя: ", o.BuyerName)
	fmt.Println("Телефон: ", o.Phone)
	fmt.Println("Количество : ", o.Amount)
	fmt.Println("Адрес : ", o.City, ", ", o.Street, ", ", o.House, ", ", o.Apartment, ", ", o.PostalCode)

}

func (o *Order) changeAmount(amount int) {
	o.Amount = strconv.Itoa(amount)
}

func (o *Order) changeAddress(postalCode, city, street, house, apartment string) {
	o.PostalCode = postalCode
	o.City = city
	o.Street = street
	o.House = house
	o.Apartment = apartment

}

const (
	digits    = "0123456789"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
)

type scanString string

func main() {
	scanTitle := scanData("Введите напименование товара")
	scanTitle.checkTitle()

	scanName := scanData("Введите имя покупателя")
	scanName.checkLetters()

	scanPhone := scanData("Введите телефон покупателя")
	scanPhone.checkLength(10)
	scanPhone.checkDigits()

	scanAmount := scanData("Введите количество")
	scanAmount.checkDigits()

	scanPostCode := scanData("Введите индекс")
	scanPostCode.checkLength(6)
	scanPhone.checkDigits()

	scanCity := scanData("Введите город")

	scanStreet := scanData("Введите улицу")

	scanHouse := scanData("Введите дом")

	scanApartment := scanData("Введите кв")

	address := newAddress(scanPostCode.toStr(), scanCity.toStr(), scanStreet.toStr(), scanHouse.toStr(), scanApartment.toStr())
	order := newOrder(scanTitle.toStr(), scanName.toStr(), scanPhone.toStr(), scanAmount.toStr(), *address)
	order.printInfo()

	order.changeAmount(15)
	order.changeAddress("999888", "state100", "moscow", "street12387128", "sdf")
	order.printInfo()

}

func scanData(message string) scanString {
	var data string
	fmt.Println(message)
	_, err := fmt.Scanf("%s", &data)
	if err != nil {
		log.Fatal("Ошибка: " + err.Error())
	}

	return scanString(data)
}

func (str *scanString) checkTitle() {
	if len(*str) < 1 || len(*str) > 100 {
		log.Fatal("Ошибка: Длина минимум 1, максимум 100")
	}
}

func (str *scanString) checkDigits() {
	digitsMap := fillLettersMap(digits)
	for _, k := range *str {
		b := string(k)
		_, digitsMapCheck := digitsMap[b]
		if !digitsMapCheck {
			log.Fatal("Ошибка: Только числа")
		}
	}
}

func (str *scanString) checkLetters() {
	lowercaseMap := fillLettersMap(lowercase)
	for _, k := range *str {
		b := string(k)
		_, digitsMapCheck := lowercaseMap[b]
		if !digitsMapCheck {
			log.Fatal("Ошибка: Только буквы в нижнем регистре")
		}
	}
}

func (str *scanString) checkLength(l int) {
	if len(*str) != l {
		log.Fatal("Ошибка: Только " + strconv.Itoa(l) + " цифр")
	}
}

func (str *scanString) toStr() string {
	return string(*str)
}

func fillLettersMap(str string) map[string]string {
	var letters map[string]string
	letters = make(map[string]string)

	for _, k := range str {
		letters[string(k)] = string(k)
	}

	return letters
}
