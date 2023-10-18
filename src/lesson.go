package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type user struct {
	addr      string
	timestamp time.Time
}

type balance struct {
	user
	userBalance int
}

func (b *balance) Balances() string {
	return fmt.Sprintf("Your balance is: %v", b.userBalance)
}
func AccountGenerator(b user) string {
	addressWallet := sha256.Sum256([]byte(b.addr))
	hash := fmt.Sprintf("%x", addressWallet)
	return string("0x" + hash)
}

func (b balance) mint(amount int) func() int {
	return func() int {
		b.userBalance += amount
		return b.userBalance
	}

}

func main() {

	newAddress := AccountGenerator(user{addr: "0x2fa2131387afe"})
	ub := balance{user: user{addr: newAddress, timestamp: time.Now()}, userBalance: 100}
	var increaseAmount int
	fmt.Println("Enter amounts that you want to mint")
	fmt.Scan(&increaseAmount)
	FT := ub.mint(increaseAmount)
	result := FT() // Call the returned function to get the updated balance
	fmt.Printf("user balance is : %d\t\n  your address account is : %v\t\n and you have created your account at %v\t\n and your balance increase to : %v\t\n", ub.userBalance, ub.user.addr, ub.timestamp.Format(time.RFC3339), result)
}
