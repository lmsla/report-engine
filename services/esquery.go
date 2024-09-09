package services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	// "io/ioutil"
	"net/http"
	"time"
)

// 定义请求结构体
type ElasticsearchRequest struct {
	Aggs map[string]Aggregation `json:"aggs"`
	Size          int           `json:"size"`
	Fields        []Field       `json:"fields"`
	ScriptFields  struct{}      `json:"script_fields"`
	StoredFields  []string      `json:"stored_fields"`
	RuntimeMappings struct{}    `json:"runtime_mappings"`
	Source        Source        `json:"_source"`
	Query         Query         `json:"query"`
}

// 定义聚合结构体
type Aggregation struct {
	Terms Terms                 `json:"terms"`
	Aggs  map[string]Aggregation `json:"aggs,omitempty"`
}

// 定义 Terms 结构体
type Terms struct {
	Field     string            `json:"field"`
	Order     map[string]string `json:"order"`
	Size      int               `json:"size"`
	ShardSize int               `json:"shard_size"`
}

// 定义字段结构体
type Field struct {
	Field  string `json:"field"`
	Format string `json:"format"`
}

// 定义 _source 结构体
type Source struct {
	Excludes []string `json:"excludes"`
}

// 定义查询结构体
type Query struct {
	Bool struct {
		Must     []interface{} `json:"must"`
		Filter   []Filter      `json:"filter"`
		Should   []interface{} `json:"should"`
		MustNot  []interface{} `json:"must_not"`
	} `json:"bool"`
}

// 定义过滤器结构体
type Filter struct {
	Range struct {
		Timestamp struct {
			Format string `json:"format"`
			Gte    string `json:"gte"`
			Lte    string `json:"lte"`
		} `json:"@timestamp"`
	} `json:"range"`
}

// Build aggregation body
func buildAggregationRequest(fields []string) ElasticsearchRequest {
	// 创建顶层的聚合结构体
	aggregations := make(map[string]Aggregation)
	currentAgg := aggregations

	// 动态构建聚合的 terms
	for i, field := range fields {
		// name aggregations
		aggName := fmt.Sprintf("agg_%d", i+2)
		newAgg := Aggregation{
			Terms: Terms{
				Field:     field,
				Order:     map[string]string{"_count": "desc"},
				Size:      10000,
				ShardSize: 25,
			},
		}

		// 将新的聚合添加到当前层的聚合中
		currentAgg[aggName] = newAgg

		// 如果不是最后一个字段，则为下一个嵌套聚合创建子聚合
		if i < len(fields)-1 {
			// 从 map 中取出结构体以进行修改
			agg := currentAgg[aggName]
			agg.Aggs = make(map[string]Aggregation)
			currentAgg[aggName] = agg
			currentAgg = agg.Aggs
		}
	}

	// 构建完整的 Elasticsearch 请求结构
	return ElasticsearchRequest{
		Aggs: aggregations,
		Size: 0,
		Fields: []Field{
			{
				Field:  "@timestamp",
				Format: "date_time",
			},
		},
		StoredFields: []string{"*"},
		Source: Source{
			Excludes: []string{},
		},
		Query: Query{
			Bool: struct {
				Must    []interface{} `json:"must"`
				Filter  []Filter      `json:"filter"`
				Should  []interface{} `json:"should"`
				MustNot []interface{} `json:"must_not"`
			}{
				Filter: []Filter{
					{
						Range: struct {
							Timestamp struct {
								Format string `json:"format"`
								Gte    string `json:"gte"`
								Lte    string `json:"lte"`
							} `json:"@timestamp"`
						}{
							Timestamp: struct {
								Format string `json:"format"`
								Gte    string `json:"gte"`
								Lte    string `json:"lte"`
							}{
								Format: "strict_date_optional_time",
								Gte:    "2023-09-04T16:00:00.000Z",
								Lte:    "2024-09-05T03:06:10.701Z",
							},
						},
					},
				},
			},
		},
	}
}

func EsTableQuery() {
	// 動態指定一或多個欄位作為 aggregation 的 terms
	fields := []string{"sourceAddress.keyword", "Method.keyword","App.keyword"}
	// create request body
	reqBody := buildAggregationRequest(fields)

	// 将请求体编码为 JSON 格式
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	// fmt.Println(string(jsonData))

	// 创建 HTTP 请求
	url := "https://10.99.1.93:9200/logstash-l7_network*/_async_search" // 请替换为你的 Elasticsearch URL
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	// 替换为你的 Elasticsearch 用户名和密码
	req.SetBasicAuth("elastic", "12345678") 

	// 創建 HTTP 客户端，支援自簽憑證（如果需要）
	client := &http.Client{
		// 跳過驗證憑證
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 10 * time.Second,
	}

	// 發送 request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// read & print resp.Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response Body:", string(body))

	// response 
	fmt.Println("Response Status:", resp.Status)
}
