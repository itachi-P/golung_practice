package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	var n, _ = strconv.Atoi(sc.Text())
	fmt.Println(n * 12) //ダース計算
}