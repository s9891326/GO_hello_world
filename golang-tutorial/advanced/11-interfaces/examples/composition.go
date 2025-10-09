package main

import "fmt"

// 基礎接口
type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// 組合接口
type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// 或者更簡潔的方式
type ReadWriteCloser2 interface {
	ReadWriter
	Closer
}

// 實現組合接口的類型
type Buffer struct {
	data []byte
	pos  int
}

func (b *Buffer) Read(p []byte) (int, error) {
	if b.pos >= len(b.data) {
		return 0, fmt.Errorf("EOF")
	}
	
	n := copy(p, b.data[b.pos:])
	b.pos += n
	fmt.Printf("📖 讀取 %d 字節: %s\n", n, string(p[:n]))
	return n, nil
}

func (b *Buffer) Write(p []byte) (int, error) {
	b.data = append(b.data, p...)
	fmt.Printf("✍️ 寫入 %d 字節: %s\n", len(p), string(p))
	return len(p), nil
}

func (b *Buffer) Close() error {
	fmt.Println("🔒 關閉緩衝區")
	b.data = nil
	b.pos = 0
	return nil
}

func (b *Buffer) Reset() {
	b.data = b.data[:0]
	b.pos = 0
	fmt.Println("🔄 重置緩衝區")
}

func (b *Buffer) String() string {
	return fmt.Sprintf("Buffer(len=%d, pos=%d): %s", len(b.data), b.pos, string(b.data))
}

func demonstrateInterfaceComposition() {
	fmt.Println("\n--- 接口組合演示 ---")
	
	buffer := &Buffer{}
	
	// 作為 Writer 使用
	fmt.Println("📝 使用 Writer 接口:")
	var w Writer = buffer
	w.Write([]byte("Hello, "))
	w.Write([]byte("Interface! "))
	w.Write([]byte("Composition is powerful."))
	
	fmt.Printf("緩衝區狀態: %s\n", buffer.String())
	
	// 作為 Reader 使用
	fmt.Println("\n📚 使用 Reader 接口:")
	var r Reader = buffer
	data := make([]byte, 8)
	r.Read(data)
	
	data2 := make([]byte, 12)
	r.Read(data2)
	
	data3 := make([]byte, 20)
	r.Read(data3)
	
	// 嘗試讀取超出範圍
	data4 := make([]byte, 10)
	n, err := r.Read(data4)
	if err != nil {
		fmt.Printf("❌ 讀取錯誤: %v (讀取了 %d 字節)\n", err, n)
	}
	
	// 重置緩衝區
	buffer.Reset()
	
	// 作為組合接口使用
	fmt.Println("\n🔧 使用組合接口:")
	var rwc ReadWriteCloser = buffer
	
	rwc.Write([]byte("New data after reset"))
	
	readData := make([]byte, 10)
	rwc.Read(readData)
	
	rwc.Close()
	
	// 演示接口的靈活性
	fmt.Println("\n🎭 接口靈活性演示:")
	processData(buffer)
}

// 接受不同接口的函數
func processData(rw ReadWriter) {
	fmt.Println("🔄 處理數據...")
	
	// 寫入數據
	rw.Write([]byte("Processing data..."))
	
	// 如果同時實現了 Reader，則讀取數據
	if buf, ok := rw.(*Buffer); ok {
		buf.Reset() // 重置以便讀取
		rw.Write([]byte("Test data for processing"))
		
		data := make([]byte, 15)
		n, _ := rw.Read(data)
		fmt.Printf("📖 處理讀取到的數據: %s (%d bytes)\n", string(data[:n]), n)
	}
}

// 高級接口組合示例
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

type ReadSeeker interface {
	Reader
	Seeker
}

type WriteSeeker interface {
	Writer
	Seeker
}

type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}

// 實現完整的讀寫查找接口
type AdvancedBuffer struct {
	data []byte
	pos  int64
}

func (ab *AdvancedBuffer) Read(p []byte) (int, error) {
	if ab.pos >= int64(len(ab.data)) {
		return 0, fmt.Errorf("EOF")
	}
	
	n := copy(p, ab.data[ab.pos:])
	ab.pos += int64(n)
	return n, nil
}

func (ab *AdvancedBuffer) Write(p []byte) (int, error) {
	// 如果位置在末尾，則追加
	if ab.pos >= int64(len(ab.data)) {
		ab.data = append(ab.data, p...)
	} else {
		// 否則覆寫
		for i, b := range p {
			if int64(i)+ab.pos < int64(len(ab.data)) {
				ab.data[ab.pos+int64(i)] = b
			} else {
				ab.data = append(ab.data, b)
			}
		}
	}
	ab.pos += int64(len(p))
	return len(p), nil
}

func (ab *AdvancedBuffer) Seek(offset int64, whence int) (int64, error) {
	var newPos int64
	
	switch whence {
	case 0: // 相對於開始
		newPos = offset
	case 1: // 相對於當前位置
		newPos = ab.pos + offset
	case 2: // 相對於末尾
		newPos = int64(len(ab.data)) + offset
	default:
		return ab.pos, fmt.Errorf("無效的 whence 值")
	}
	
	if newPos < 0 {
		return ab.pos, fmt.Errorf("負數位置")
	}
	
	ab.pos = newPos
	return ab.pos, nil
}

func (ab *AdvancedBuffer) String() string {
	return fmt.Sprintf("AdvancedBuffer(len=%d, pos=%d)", len(ab.data), ab.pos)
}

func demonstrateAdvancedComposition() {
	fmt.Println("\n--- 高級接口組合演示 ---")
	
	buf := &AdvancedBuffer{}
	
	// 使用 ReadWriteSeeker 接口
	var rws ReadWriteSeeker = buf
	
	// 寫入一些數據
	rws.Write([]byte("Hello, Advanced Interface!"))
	fmt.Printf("寫入後: %s\n", buf.String())
	
	// 尋找到開始位置
	pos, _ := rws.Seek(0, 0)
	fmt.Printf("尋找到位置 %d\n", pos)
	
	// 讀取數據
	data := make([]byte, 5)
	rws.Read(data)
	fmt.Printf("讀取數據: %s\n", string(data))
	
	// 尋找到特定位置並覆寫
	rws.Seek(7, 0)
	rws.Write([]byte("World"))
	
	// 讀取全部內容
	rws.Seek(0, 0)
	fullData := make([]byte, len(buf.data))
	rws.Read(fullData)
	fmt.Printf("完整內容: %s\n", string(fullData))
}