// 練習 1 解答：基本錯誤處理 - 計算器
package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Calculator struct{}

func (c *Calculator) Add(a, b string) (float64, error) {
	numA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, fmt.Errorf("無效的數字: %s", a)
	}
	
	numB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, fmt.Errorf("無效的數字: %s", b)
	}
	
	return numA + numB, nil
}

func (c *Calculator) Subtract(a, b string) (float64, error) {
	numA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, fmt.Errorf("無效的數字: %s", a)
	}
	
	numB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, fmt.Errorf("無效的數字: %s", b)
	}
	
	return numA - numB, nil
}

func (c *Calculator) Multiply(a, b string) (float64, error) {
	numA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, fmt.Errorf("無效的數字: %s", a)
	}
	
	numB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, fmt.Errorf("無效的數字: %s", b)
	}
	
	return numA * numB, nil
}

func (c *Calculator) Divide(a, b string) (float64, error) {
	numA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, fmt.Errorf("無效的數字: %s", a)
	}
	
	numB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, fmt.Errorf("無效的數字: %s", b)
	}
	
	if numB == 0 {
		return 0, errors.New("除數不能為零")
	}
	
	return numA / numB, nil
}

func main() {
	calc := &Calculator{}
	
	fmt.Println("=== 計算器測試 ===")
	
	// 測試加法
	fmt.Println("\n--- 加法測試 ---")
	testCases := []struct {
		a, b     string
		expected string
	}{
		{"10", "5", "15"},
		{"abc", "5", "錯誤"},
		{"10.5", "2.3", "12.8"},
	}
	
	for _, tc := range testCases {
		result, err := calc.Add(tc.a, tc.b)
		if err != nil {
			fmt.Printf("%s + %s = 錯誤: %v\n", tc.a, tc.b, err)
		} else {
			fmt.Printf("%s + %s = %.2f\n", tc.a, tc.b, result)
		}
	}
	
	// 測試除法
	fmt.Println("\n--- 除法測試 ---")
	divisionTests := []struct {
		a, b string
	}{
		{"10", "2"},
		{"10", "0"},
		{"abc", "5"},
	}
	
	for _, tc := range divisionTests {
		result, err := calc.Divide(tc.a, tc.b)
		if err != nil {
			fmt.Printf("%s ÷ %s = 錯誤: %v\n", tc.a, tc.b, err)
		} else {
			fmt.Printf("%s ÷ %s = %.2f\n", tc.a, tc.b, result)
		}
	}
	
	// 測試其他操作
	fmt.Println("\n--- 其他操作測試 ---")
	fmt.Print("減法: ")
	if result, err := calc.Subtract("10", "3"); err != nil {
		fmt.Printf("錯誤: %v\n", err)
	} else {
		fmt.Printf("10 - 3 = %.2f\n", result)
	}
	
	fmt.Print("乘法: ")
	if result, err := calc.Multiply("4", "5"); err != nil {
		fmt.Printf("錯誤: %v\n", err)
	} else {
		fmt.Printf("4 × 5 = %.2f\n", result)
	}
}