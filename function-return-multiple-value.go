package main

import "fmt"

func getFriendName() (string, string, string) {
	return "Budi", "Anto", "Joko"
}

func main() {
	firstFriend, secondFriend, thirdFriend := getFriendName()
	fmt.Println(firstFriend)
	fmt.Println(secondFriend)
	fmt.Println(thirdFriend)

	myFirstFriend, _, _ := getFriendName()
	fmt.Println(myFirstFriend)
}
