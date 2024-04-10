package main

import (
	"SensorQualityControl/log"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	sensors := log.ParseLog(reader)
	for _, s := range sensors {
		fmt.Println(s.Evaluate())
	}
}
