package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hello_world/hello-JF/httputil"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	uid    = "95c2fcbb-4cc4-47ff-a59e-a975c3d5a697"
	key    = "021669817e2442eba0a1b0b1eda5a145"
	secret = "9MeUpc1q2WQ5DJmzWw2tqXYmnmZXhJtCLgUezPzYUNTVyba3r0fVWmdcVq0CvxGElerzsFKJc5EN9QAJbxVFqlaEnXvkphVWS5ukWsS6dOTaIc6udPH7MkQ19l76TLc5mcR18k5Jbi5uawsqWnFR76EqfLlAz1luBsjyO6kxgtLJ3LbcupTmsuGnBeiw2SXYAL84oETO6Oxwgyds6SHshzhVbGyLc4jKy1zFElzPK09tgonPNKNIiBmv4XsNSfkq"
)

func genRandomString() string {
	now := time.Now()
	return strings.ReplaceAll(now.Format("20060102150405.000000000"), ".", "")
}

func generateSign(data map[string]any) string {
	// 基于HMAC（Hash-based Message Authentication Code）的签名
	// 1. 排序 keys
	var keys []string
	for k, v := range data {
		if v != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	// 2. 拼接 value
	var sb strings.Builder
	for _, k := range keys {
		sb.WriteString(fmt.Sprintf("%s=%s&", k, data[k]))
	}
	// 去掉最後一個 '&'
	sbStr := sb.String()
	sbStr = sbStr[:len(sbStr)-1]

	// 4. 生成签名：使用密钥和哈希算法（SHA-256）
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(sbStr))
	signature := hex.EncodeToString(h.Sum(nil))

	fmt.Println("拼接後的字符串: ", sbStr, "生成的簽名: ", signature)
	return signature
}

// 生成随机UUID
func generateUUID() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "-", err
	}
	return uuid.String(), nil
}

func createOrder(orderId string) {
	uuid, _ := generateUUID()

	// data := map[string]string{
	// 	"transaction_type":  "income", // income(收入)/expense(支出)
	// 	"amount":            "20",
	// 	"merchant_order_id": orderId,
	// 	"merchant_uid":      uid,
	// 	"currency":          "cny",
	// 	"payment":           "a2a",
	// 	"response_type":     "json",
	// 	"country":           "China",
	// 	"accounting_type":   "receive", //recharge(充值)/receive(收款)/pay(付款)/settle(结算)
	// 	// "notify_url":        "https://d94d77e7046b.ngrok-free.app/notify",
	// 	"from":           "test866_cny",
	// 	"from_name":      "test866_china",
	// 	"from_bank_name": "中国建设银行",
	// 	// "channel":           "payment_buildin_a2a",
	// }

	data := map[string]any{
		"transaction_type":  "income", // income(收入)/expense(支出)
		"amount":            "20",
		"merchant_order_id": orderId,
		"merchant_uid":      uid,
		"currency":          "mmk",
		"payment":           "a2a",
		"response_type":     "json",
		"country":           "China",
		"accounting_type":   "receive", //recharge(充值)/receive(收款)/pay(付款)/settle(结算)
		// "notify_url":        "https://d94d77e7046b.ngrok-free.app/notify",
		"from":           "test866_cny",
		"from_name":      "test866_china",
		"from_bank_name": "test",
		// "channel":           "payment_buildin_a2a",

		"order_uid": uuid,
	}

	sign := generateSign(data)

	header := map[string]string{
		"Content-Type": "application/json",
		"OC-Key":       key,
		"OC-Signature": sign,
	}

	fmt.Println("create order data:", data, " header: ", header)

	// resp, err := httputil.DoPostRequest("https://demo.qc168.info/octopus/api/order/new", data, header)
	resp, err := httputil.DoPostRequest("http://127.0.0.1:8011/new", data, header)
	fmt.Println("create order response:", resp, " err: ", err)
}

func expenseOrder(orderId string) {
	uuid, _ := generateUUID()

	// 支出訂單
	// data := map[string]string{
	// 	"transaction_type":  "expense", // income(收入)/expense(支出)
	// 	"amount":            "5",
	// 	"merchant_order_id": orderId,
	// 	"merchant_uid":      uid,
	// 	"currency":          "cny",
	// 	"payment":           "a2a",
	// 	"response_type":     "json",
	// 	"country":           "China",
	// 	"accounting_type":   "pay", //recharge(充值)/receive(收款)/pay(付款)/settle(结算)
	// }

	data := map[string]any{
		"transaction_type":  "expense", // income(收入)/expense(支出)
		"amount":            "5",
		"merchant_order_id": orderId,
		"merchant_uid":      uid,
		"currency":          "mmk",
		"payment":           "a2a",
		"response_type":     "json",
		"country":           "China",
		"accounting_type":   "pay", //recharge(充值)/receive(收款)/pay(付款)/settle(结算)
		"order_uid":         uuid,
	}

	sign := generateSign(data)

	header := map[string]string{
		"Content-Type": "application/json",
		"OC-Key":       key,
		"OC-Signature": sign,
	}

	fmt.Println("expense order data:", data, " header: ", header)

	// resp, err := httputil.DoPostRequest("https://demo.qc168.info/octopus/api/order/new", data, header)
	resp, err := httputil.DoPostRequest("http://127.0.0.1:8011/new", data, header)
	fmt.Println("expense order response:", resp, " err: ", err)
}

func queryOrder(orderId string) {
	data := map[string]any{
		"merchant_uid": orderId,
	}

	sign := generateSign(data)
	header := map[string]string{
		"Content-Type": "application/json",
		"OC-Key":       key,
		"OC-Signature": sign,
	}
	fmt.Println("query order data:", data, " header: ", header)

	resp, err := httputil.DoPostRequest("https://demo.qc168.info/octopus/api/order/get", data, header)
	fmt.Println("query order response:", resp, " err: ", err)
}

func createPayAccount() {
	// data := map[string]any{
	// 	"channel":  "test", // income(收入)/expense(支出)
	// 	"username": "test866",
	// 	"deviceid": "device_001",
	// 	"enable":   true,
	// 	"password": "123456",
	// }

	// fedmobile
	data := map[string]any{
		"channel":  "fedmobile",
		"username": "918333075014",
		"deviceid": "fed1",
		"enable":   true,
		"password": "8080",
		"proxy":    "http://GEMDL5nh:i3xXRpay@34.14.128.242:8084",
	}

	// data := map[string]any{
	// 	"channel":  "test", // income(收入)/expense(支出)
	// 	"username": "917569790564",
	// 	"deviceid": "device_002",
	// 	"enable":   true,
	// 	"password": "8080",
	// }

	header := map[string]string{
		"Content-Type": "application/json",
		// "OC-Key":       key,
		// "OC-Signature": sign,
	}

	resp, err := httputil.DoPostRequest("http://127.0.0.1:8011/create_pay_account", data, header)
	fmt.Println("create order response:", resp, " err: ", err)

}

func updatePayAccount() {
	data := map[string]any{
		"id":       89,
		"username": "917032483253",
		"deviceid": "device_001",
		"enabled":  true,
		"password": "2580",
		"status":   "online",
	}
	// data := map[string]string{
	// 	"channel":   "test", // income(收入)/expense(支出)
	// 	"username":  "test866",
	// 	"deviceid": "device_001",
	// 	"password":  "123456",
	// 	"enable":    "true",
	// 	"status":    "online",
	// 	// "si":        "123456",
	// }

	header := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := httputil.DoPostRequest("http://127.0.0.1:8011/update_pay_account", data, header)
	fmt.Println("update order response:", resp, " err: ", err)
}

func listTags() {
	panic("unimplemented")
}

func createTags() {
	data := map[string]any{
		"name":             "payment_myanmar_tags.name", // income(收入)/expense(支出)
		"target":           "test866",
		"type":             "account",
		"transaction_type": "123456",
	}

	header := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := httputil.DoPostRequest("http://127.0.0.1:8011/add_tag", data, header)
	fmt.Println("creeate tags response:", resp, " err: ", err)
}

func main() {
	// orderId := genRandomString()
	// fmt.Println("生成的訂單號:", orderId)

	// 創建帳戶
	// createPayAccount()

	// 開啟帳戶
	// updatePayAccount()

	// 創建tags
	// createTags()

	// listTags()

	// createOrder(orderId)
	// queryOrder(orderId)

	// expenseOrder(orderId)

	// expenseOrder(orderId)
	// queryOrder(orderId)
}
