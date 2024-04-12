package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("======File.csv swapping rows and columns======")
	// 檢查是否提供了檔案名稱(只能傳入一個檔案)
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go example.csv")
		return
	}

	// 獲取檔案名稱
	fileName := os.Args[1]

	// 打開 CSV 文件
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// 讀取 CSV 文件
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 將行列互換
	var transposed [][]string
	for i := 0; i < len(records[0]); i++ {
		var column []string
		for j := 0; j < len(records); j++ {
			column = append(column, records[j][i])
		}
		transposed = append(transposed, column)
	}

	// 打開同一個 CSV 文件以便寫入
	file, err = os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// 寫入處理後的數據到文件中
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range transposed {
		err := writer.Write(row)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	fmt.Printf("Swapping rows and columns file :'%s' successfully!\n", fileName)
}

