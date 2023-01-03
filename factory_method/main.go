package main

/*
	Go には、 クラスや継承などの OOP 機能がないため、 Go で古典的な Factory Method パターンを実装することは不可能。 ただし、 パターンの簡易版である単純ファクトリーなら実装できる。
*/

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
