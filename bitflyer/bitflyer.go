package bitflyer

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// bitflyer lightning APIを利用するためのURL
// 末尾に各API指定のURLを付け加え、リクエストを送ることで使用できる。
const baseURL = "https://api.bitflyer.com/v1/"

// クライアントが持っているAPI認証のための情報
type APIClient struct {
	key        string
	secret     string
	httpClient *http.Client
}

func New(key, secret string) *APIClient {
	apiClient := &APIClient{key, secret, &http.Client{}}
	return apiClient
}

// リクエストするデータ
func (api APIClient) header(method, endpoint string, body []byte) map[string]string {
	// タイムスタンプ（stringに変換）
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	// ログに出力
	log.Println(timestamp)
	// データをまとめる（bitflyerAPI側の指示）
	message := timestamp + method + endpoint + string(body)

	// HMAC署名
	mac := hmac.New(sha256.New, []byte(api.secret))
	mac.Write([]byte(message))
	sign := hex.EncodeToString(mac.Sum(nil))

	return map[string]string{
		"ACCESS-KEY":       api.key,
		"ACCESS-TIMESTAMP": timestamp,
		"SCCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}
}

func (api *APIClient) doRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
	// 有効なアドレスか解析（bitflyerAPI基準URL）
	baseURL, err := url.Parse(baseURL)
	if err != nil {
		return
	}
	// 有効なアドレスか解析（BalanceAPI利用時URL）
	apiURL, err := url.Parse(urlPath)
	if err != nil {
		return
	}
	// 特定のbitflyerAPIを利用するためのリクエストURL（エンドポイント）
	endpoint := baseURL.ResolveReference(apiURL).String()
	// エンドポイントをログ出力
	log.Printf("action=doRequest endpoint=%s", endpoint)
	// httpリクエストを送信
	// GETの場合：method, endpoint, query
	// POSTの場合：method, endpoint, bytes.NewBuffer(data)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
// RawQueryに変換（例：q1=foo&q2=bar）
	req.URL.RawQuery = q.Encode()

	for key, value := range api.header(method, req.URL.RequestURI(), data) {
		reqHeader.Add(key, value)

		resq, err != api.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resq.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
}


// BalanceAPIで取得したデータを保管する構造体
type Balance struct {
	CurrentCode string `json:"currency_code"` // 通貨コード
	Amount float64 `json:"amount"`						// 所持してる金額
	Available float64 `json:"available"`			// 利用する金額
}

// GetBalanceAPIにアクセスするための関数
func (api *APIClient) GetBalance() ([]Balance, error){
	url := "me/getbalance"
	// GetBalanceAPIにリクエストを送る。（処理はdoRequest関数）
	resp, err := api.doRequest("GET", url, map[string]string{}, nil)
	log.Printf("url=%s resp=%s", url, string(resp))
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	var balance []Balance
	// レスポンスされたjson形式の値をGo objectに変換して構造体に保管
	err = json.Unmarshal(resp, &balance)
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	return balance, nil
}