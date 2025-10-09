package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 日誌中間件
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// 創建響應記錄器
		recorder := &ResponseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			size:          0,
		}
		
		// 記錄請求開始
		log.Printf("開始 %s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		
		// 執行下一個處理器
		next.ServeHTTP(recorder, r)
		
		// 記錄請求完成
		duration := time.Since(start)
		log.Printf("完成 %s %s %d %d bytes %v", 
			r.Method, r.URL.Path, recorder.statusCode, recorder.size, duration)
	})
}

// 響應記錄器
type ResponseRecorder struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func (rr *ResponseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}

func (rr *ResponseRecorder) Write(b []byte) (int, error) {
	size, err := rr.ResponseWriter.Write(b)
	rr.size += size
	return size, err
}

// 認證中間件
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 檢查 Authorization 頭
		auth := r.Header.Get("Authorization")
		
		// 簡單的 Bearer token 驗證
		if !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "未授權：缺少 Bearer token", http.StatusUnauthorized)
			return
		}
		
		token := strings.TrimPrefix(auth, "Bearer ")
		if !isValidToken(token) {
			http.Error(w, "未授權：無效的 token", http.StatusUnauthorized)
			return
		}
		
		// 將用戶信息添加到上下文（這裡簡化處理）
		log.Printf("用戶認證成功: token=%s", token)
		
		next.ServeHTTP(w, r)
	})
}

// 模擬token驗證
func isValidToken(token string) bool {
	validTokens := []string{"abc123", "def456", "ghi789"}
	for _, valid := range validTokens {
		if token == valid {
			return true
		}
	}
	return false
}

// CORS中間件
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 設置CORS頭
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "86400")
		
		// 處理預檢請求
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

// 限流中間件
type RateLimiter struct {
	visitors map[string]*Visitor
	mutex    sync.RWMutex
	rate     int           // 每分鐘允許的請求數
	cleanup  time.Duration // 清理間隔
}

type Visitor struct {
	requests []time.Time
	mutex    sync.Mutex
}

func NewRateLimiter(rate int) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*Visitor),
		rate:     rate,
		cleanup:  time.Minute,
	}
	
	// 啟動清理協程
	go rl.cleanupLoop()
	
	return rl
}

func (rl *RateLimiter) cleanupLoop() {
	ticker := time.NewTicker(rl.cleanup)
	defer ticker.Stop()
	
	for range ticker.C {
		rl.mutex.Lock()
		now := time.Now()
		
		for ip, visitor := range rl.visitors {
			visitor.mutex.Lock()
			
			// 移除一分鐘前的請求記錄
			var validRequests []time.Time
			for _, reqTime := range visitor.requests {
				if now.Sub(reqTime) < time.Minute {
					validRequests = append(validRequests, reqTime)
				}
			}
			
			if len(validRequests) == 0 {
				delete(rl.visitors, ip)
			} else {
				visitor.requests = validRequests
			}
			
			visitor.mutex.Unlock()
		}
		
		rl.mutex.Unlock()
	}
}

func (rl *RateLimiter) Allow(ip string) bool {
	rl.mutex.Lock()
	visitor, exists := rl.visitors[ip]
	if !exists {
		visitor = &Visitor{
			requests: make([]time.Time, 0),
		}
		rl.visitors[ip] = visitor
	}
	rl.mutex.Unlock()
	
	visitor.mutex.Lock()
	defer visitor.mutex.Unlock()
	
	now := time.Now()
	
	// 移除一分鐘前的請求
	var validRequests []time.Time
	for _, reqTime := range visitor.requests {
		if now.Sub(reqTime) < time.Minute {
			validRequests = append(validRequests, reqTime)
		}
	}
	visitor.requests = validRequests
	
	// 檢查是否超過限制
	if len(visitor.requests) >= rl.rate {
		return false
	}
	
	// 記錄當前請求
	visitor.requests = append(visitor.requests, now)
	return true
}

func RateLimitMiddleware(limiter *RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := getClientIP(r)
			
			if !limiter.Allow(ip) {
				http.Error(w, "請求過於頻繁，請稍後再試", http.StatusTooManyRequests)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	}
}

// 獲取客戶端IP
func getClientIP(r *http.Request) string {
	// 檢查 X-Forwarded-For 頭
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}
	
	// 檢查 X-Real-IP 頭
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}
	
	// 使用 RemoteAddr
	ip := r.RemoteAddr
	if idx := strings.LastIndex(ip, ":"); idx != -1 {
		ip = ip[:idx]
	}
	
	return ip
}

// 請求大小限制中間件
func RequestSizeLimitMiddleware(maxBytes int64) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ContentLength > maxBytes {
				http.Error(w, "請求體過大", http.StatusRequestEntityTooLarge)
				return
			}
			
			// 限制請求體大小
			r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
			
			next.ServeHTTP(w, r)
		})
	}
}

// 超時中間件
func TimeoutMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.TimeoutHandler(next, timeout, "請求超時")
	}
}

// 安全頭中間件
func SecurityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 設置安全相關的HTTP頭
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		
		next.ServeHTTP(w, r)
	})
}

// 恢復中間件（處理panic）
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic 恢復: %v", err)
				http.Error(w, "內部服務器錯誤", http.StatusInternalServerError)
			}
		}()
		
		next.ServeHTTP(w, r)
	})
}

// 中間件鏈
type MiddlewareChain struct {
	middlewares []func(http.Handler) http.Handler
}

func NewMiddlewareChain() *MiddlewareChain {
	return &MiddlewareChain{
		middlewares: make([]func(http.Handler) http.Handler, 0),
	}
}

func (mc *MiddlewareChain) Use(middleware func(http.Handler) http.Handler) *MiddlewareChain {
	mc.middlewares = append(mc.middlewares, middleware)
	return mc
}

func (mc *MiddlewareChain) Then(handler http.Handler) http.Handler {
	// 反向應用中間件
	for i := len(mc.middlewares) - 1; i >= 0; i-- {
		handler = mc.middlewares[i](handler)
	}
	return handler
}

// 示例處理器
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "歡迎來到中間件演示頁面！\n")
	fmt.Fprintf(w, "請求時間: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "這是受保護的頁面，需要認證才能訪問。\n")
	fmt.Fprintf(w, "認證成功！時間: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"message": "API 響應",
		"time":    time.Now().Unix(),
		"method":  r.Method,
		"path":    r.URL.Path,
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	// 故意觸發panic來測試恢復中間件
	panic("這是一個測試panic")
}

func main() {
	// 創建限流器
	rateLimiter := NewRateLimiter(10) // 每分鐘10個請求
	
	// 創建中間件鏈
	chain := NewMiddlewareChain().
		Use(RecoveryMiddleware).
		Use(LoggingMiddleware).
		Use(CORSMiddleware).
		Use(SecurityHeadersMiddleware).
		Use(RateLimitMiddleware(rateLimiter)).
		Use(RequestSizeLimitMiddleware(1024*1024)). // 1MB限制
		Use(TimeoutMiddleware(30*time.Second))
	
	// 受保護的路由需要額外的認證中間件
	protectedChain := NewMiddlewareChain().
		Use(RecoveryMiddleware).
		Use(LoggingMiddleware).
		Use(AuthMiddleware).
		Use(CORSMiddleware).
		Use(SecurityHeadersMiddleware)
	
	// 路由設置
	mux := http.NewServeMux()
	
	// 公開路由
	mux.Handle("/", chain.Then(http.HandlerFunc(homeHandler)))
	mux.Handle("/api", chain.Then(http.HandlerFunc(apiHandler)))
	mux.Handle("/panic", chain.Then(http.HandlerFunc(panicHandler)))
	
	// 受保護路由
	mux.Handle("/protected", protectedChain.Then(http.HandlerFunc(protectedHandler)))
	
	// 創建服務器
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	fmt.Println("中間件演示服務器啟動在 http://localhost:8080")
	fmt.Println("可用路由:")
	fmt.Println("  GET  /           - 首頁（公開）")
	fmt.Println("  GET  /api        - API端點（公開）")
	fmt.Println("  GET  /panic      - Panic測試（公開）")
	fmt.Println("  GET  /protected  - 受保護頁面（需要認證）")
	fmt.Println("")
	fmt.Println("測試認證路由，請在請求頭中添加:")
	fmt.Println("  Authorization: Bearer abc123")
	fmt.Println("")
	fmt.Println("按 Ctrl+C 停止服務器")
	
	log.Fatal(server.ListenAndServe())
}