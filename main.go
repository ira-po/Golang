package main

import "fmt"

func main() {
	var our_money = 23.00
	var the_price_of_one_apple = 5.99
	var the_price_of_one_pear = 7.00

	fmt.Println("Загальна сума грошей у гривнях:", our_money)
	fmt.Println("Вартість одного яблука становитиме:", the_price_of_one_apple)
	fmt.Println("Вартість однієї груші становитиме:", the_price_of_one_pear)

	var how_much_money_we_must_spend = ((the_price_of_one_apple) * 9) + ((the_price_of_one_pear) * 8)
	fmt.Println("Скільки грошей треба витратити аби купити 9 яблук та 8 груш?", how_much_money_we_must_spend)

	var how_many_apples_we_can_buy = (our_money) / (the_price_of_one_apple)
	fmt.Println("Скільки яблук ми можемо купити:", int(how_many_apples_we_can_buy))

	var how_many_pears_we_can_buy = (our_money) / (the_price_of_one_pear)
	fmt.Println("Скільки груш ми можемо купити:", int(how_many_pears_we_can_buy))

	can_buy_both_things := (the_price_of_one_apple*2 + the_price_of_one_pear*2) <= our_money
	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?", can_buy_both_things)
}
