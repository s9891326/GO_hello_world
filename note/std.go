package note

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// 6.1 隨機數
func Random() {
	// rand.Seed(time.Now().UnixNano())  go version > 1.20的話就不用需要用rand.seed了
	fmt.Println(rand.Intn(10))
}

// 6.2 字符串類型轉換
func Strconv() {
	i1 := 123
	s1 := "gmail.com"
	s2 := fmt.Sprintf("%d@%s", i1, s1)
	fmt.Println("s2=", s2)

	var (
		i2 int
		s3 string
	)
	n, err := fmt.Sscanf(s2, "%d@%s", &i2, &s3)
	if err != nil {
		panic(err)
	}
	fmt.Println("成功解析", n, "個數據")
	fmt.Println("i2=", i2)
	fmt.Println("s3=", s3)


	s4 := strconv.FormatInt(123, 4)
	fmt.Println("s4=", s4)
	u1, err := strconv.ParseUint(s4, 4, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("u1=", u1)
}

// 6.3 string包常見函數
func PackageStrings()  {
	fmt.Println(strings.Contains("hello", "ll"))
	// 	true
	fmt.Println(strings.Index("hello", "o"))
	// 4
	fmt.Println(strings.Replace("hello", "l", "dd", -1))  // n: 替換次數 n < 0:不限制替換次數
	// heddddo

	fmt.Println(strings.Repeat("mia", 5))
	// miamiamiamiamia

	fmt.Println(strings.Fields("mia mia\n mia\t mia"))  // 按照空白(多個字符\n \t...)拆分成子串
	// 	[mia mia mia mia]

	fmt.Println(strings.Split("he-ll-o-world", "-"))
	// 	[he ll o world]

	fmt.Println(strings.SplitAfter("he-ll-o-world", "-"))
	// 	[he- ll- o- world]

	fmt.Println(strings.Trim("#*\nwww.www.ww&@#", "@#$%&\n*"))  // 清除指定的字串
	// www.www.ww
}

// 6.4 中文字符常見操作
func PackageUTF8()  {
	fmt.Println(utf8.RuneCountInString("hello,世紀"))
	// 8

	str := "hello,世界"
	// 因為是使用bytes進行切片的，所以一個中文在切片中需要3bytes，所以會是無效的字串
	// 要轉換成rune才會是Unicode 字元
	runes := []rune(str) // 轉換為 rune 陣列  
	fmt.Println(runes)  //  [104 101 108 108 111 44 19990 30028]
	fmt.Println(string(runes[:len(runes) - 1]))  // hello,世

	fmt.Println(str[: len(str) - 1])   // hello,世�
	fmt.Println(utf8.ValidString(str[:len(str) - 1]))  // false

}


// 6.5 時間常見操作
func PackageTime()  {
	fmt.Println("6.5.1 時段")
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Println()
	d1, err := time.ParseDuration("1000s")
	if err != nil {
		panic(err)
	}
	fmt.Println("d1=", d1)  // d1= 16m40s

	// 根據 layout 解析時間字串 value，回傳 time.Time（預設 UTC 時區）
	t1, err := time.Parse("2006年1月2日, 15點4分5秒", "2028年1月1日, 18點18分15秒")
	if err != nil {
		panic(err)
	}
	fmt.Println("t1=", t1)  // 2022-01-01 18:18:15 +0000 UTC
	
	// 計算 time.Now() 和 t 的時間差，回傳 time.Duration
	fmt.Println("now:", time.Now())
	fmt.Println(time.Since(t1))

	// time.After(d) <- chan 將在等待時段D後巷返回的chan上發送當前時段t
	// var c1 chan int = make(chan int)
	c1 := make(chan int)
	select {
	case <- c1:
		fmt.Println("收到")
	case <- time.After(time.Second):
		fmt.Println("過期")
	}

	fmt.Println("6.5.2 時區")
	l1, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		panic(err)
	}
	fmt.Println(l1.String())

	fmt.Println("6.5.5 時間")
	fmt.Println(time.Now().Format("2006/01/02 15:04:05"))
	t2, err := time.ParseInLocation("2006/01/02 15:04:05", "2028/03/07 15:07:50", l1)
	if err != nil {
		panic(err)
	}
	fmt.Println("t2=", t2)

	fmt.Println(t2.Location())
	fmt.Println(t2.Add(d1))
} 
