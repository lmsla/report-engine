package clients

import (
	// "context"
	"crypto/tls"
	"fmt"
	"io"
	"github.com/elastic/go-elasticsearch/v8"
	// "log"
	"net/http"
	"report-backend-golang/log"
	"report-backend-golang/models"
	"time"
)

// 測試 Elasticsearch 連線
func TestElasticsearch(instance models.Instance) (error, models.Response) {

	res_t := models.Response{}
	res_t.Success = false
	res_t.Body = nil

	esConfig := elasticsearch.Config{
		Addresses: []string{
			instance.EsUrl, // 替換為你的 Elasticsearch URL
		},
		Username: instance.User,     // 替換為你的 Elasticsearch 使用者名稱
		Password: instance.Password, // 替換為你的 Elasticsearch 密碼
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 忽略自簽證書
		},
	}

	es, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		log.Logrecord("ERROR", fmt.Sprintf("Error creating Elasticsearch client: %s", err.Error()))
		res_t.Msg = fmt.Sprintf("Error creating Elasticsearch client")
		return err, res_t
	}

	// 測試集群健康狀態˝
	res, err := es.Cluster.Health()
	if err != nil {
		log.Logrecord("ERROR", fmt.Sprintf("Error ES getting cluster health: %s", err.Error()))
		res_t.Msg = fmt.Sprintf("Error ES getting cluster health")
		return err, res_t
	}
	defer res.Body.Close()

	if res.IsError() {
		res_t.Msg = fmt.Sprintf("Elasticsearch response error")
		// res_t.Body = res.Body
		log.Logrecord("ERROR", fmt.Sprintf("Elasticsearch response error: %s", res.String()))
	} else {
		fmt.Println("Cluster Health Response:")
		fmt.Println(res.String())
		res_t.Msg = fmt.Sprintf("ES Cluster Health Check Succeed")
		res_t.Success = true
	}
	return err, res_t
}

// 測試 Kibana 連線
func TestKibana(instance models.Instance) (error, models.Response) {

	res_t := models.Response{}
	res_t.Success = false
	res_t.Body = nil

	// 自定義 HTTP 客戶端
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 忽略自簽證書
		},
	}
	// 建立請求
	req, err := http.NewRequest("GET", instance.URL+"/api/spaces/space", nil)
	if err != nil {
		log.Logrecord("ERROR", fmt.Sprintf("Error creating HTTP request: %s", err.Error()))
		res_t.Msg = fmt.Sprintf("Error creating HTTP request")
		return err, res_t
	}
	req.SetBasicAuth(instance.User, instance.Password)

	// 發送請求
	resp, err := client.Do(req)
	if err != nil {
		log.Logrecord("ERROR", fmt.Sprintf("Error connecting to Kibana: %s", err.Error()))
		res_t.Msg = fmt.Sprintf("Error connecting to Kibana")
		return err, res_t
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		res_t.Msg = "Kibana is reachable!"
		res_t.Success = true
	} else {
		log.Logrecord("ERROR", fmt.Sprintf("Kibana response error: %s", string(body)))
		res_t.Msg = fmt.Sprintf("Kibana response error")
	}
	return err, res_t
}
