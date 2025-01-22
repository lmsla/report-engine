package services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"report-backend-golang/entities"
	"report-backend-golang/log"
	"sort"
	"time"
)

// request struct
type ElasticsearchRequest struct {
	Aggs            map[string]Aggregation `json:"aggs"`
	Size            int                    `json:"size"`
	Fields          []Field                `json:"fields"`
	ScriptFields    struct{}               `json:"script_fields"`
	StoredFields    []string               `json:"stored_fields"`
	RuntimeMappings struct{}               `json:"runtime_mappings"`
	Source          Source                 `json:"_source"`
	Query           Query                  `json:"query"`
}

// aggregation struct
type Aggregation struct {
	Terms Terms                  `json:"terms"`
	Aggs  map[string]Aggregation `json:"aggs,omitempty"`
}

// Terms struct
type Terms struct {
	Field     string            `json:"field"`
	Order     map[string]string `json:"order"`
	Size      int               `json:"size"`
	ShardSize int               `json:"shard_size"`
}

// Field struct
type Field struct {
	Field  string `json:"field"`
	Format string `json:"format"`
}

// _source struct
type Source struct {
	Excludes []string `json:"excludes"`
}

// Query struct
type Query struct {
	Bool struct {
		Must    []interface{} `json:"must"`
		Filter  []Filter      `json:"filter"`
		Should  []interface{} `json:"should"`
		MustNot []interface{} `json:"must_not"`
	} `json:"bool"`
}

// Filter struct
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
func buildAggregationRequest(fields []Column, timefromStr, nowStr string) ElasticsearchRequest {
	// 最外層的 aggregation
	aggregations := make(map[string]Aggregation)
	currentAgg := aggregations

	// 動態產生 aggregate terms
	for i, field := range fields {
		// name aggregations
		aggName := fmt.Sprintf("agg_%d", i+2)
		newAgg := Aggregation{
			Terms: Terms{
				Field:     field.Name,
				Order:     map[string]string{"_count": "desc"},
				Size:      10,
				ShardSize: field.Size,
			},
		}

		// 將新的 aggregation 加到目前層的 aggregation 中
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

	// Build  Complete Elasticsearch request struct
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
								Gte:    timefromStr,
								Lte:    nowStr,
							},
						},
					},
				},
			},
		},
	}

}

func EsTableQuery(table_columns []entities.Column, instance entities.Instance, data_view, timefrom, now string) (result string) {
	var fields []Column
	var Columns []Column

	for _, data := range table_columns {
		var column Column
		column.Name = data.Name
		column.Order = data.Order
		column.Size = data.Size
		Columns = append(Columns, column)
		sort.Sort(sort.Reverse(ColumnSlice(Columns))) // 按照 order 排序
	}
	for _, data := range Columns {
		fields = append(fields, Column{Name: data.Name, Order: data.Order, Size: data.Size})
	}

	nowStr := now + "T00:00:00.000+08:00"
	timefromStr := timefrom + "T00:00:00.000+08:00"

	// nowStr:= "2024-05-13T15:30:00.000+08:00"
	// timefromStr := "2024-05-13T15:40:00.000+08:00"

	//// 動態指定一或多個欄位作為 aggregation 的 terms
	//// create request body
	//// fields sample := []string{"sourceAddress.keyword", "Method.keyword", "App.keyword", "User.keyword"}
	reqBody := buildAggregationRequest(fields, timefromStr, nowStr)

	//// reqbody to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Logrecord("ERROR", "Error marshaling JSON: "+ err.Error())
		return
	}

	//// create HTTP request
	url := fmt.Sprintf("%s/%s/_async_search", instance.EsUrl, data_view)
	// url := "https://10.99.1.93:9200/logstash-l7_network*/_async_search" // change to your Elasticsearch URL
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Logrecord("ERROR", "Creating request error at EsTableQuery stage: "+ err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	// Elasticsearch 帳號＆密碼
	req.SetBasicAuth(instance.User, instance.Password)

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
		log.Logrecord("ERROR", "Sending request error at EsTableQuery stage: "+ err.Error())
		return
	}
	defer resp.Body.Close()

	// read & print resp.Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Logrecord("ERROR", "Reading response body error at EsTableQuery stage: "+ err.Error())
		return
	}

	response := string(body)

	return response
}

func DataDealing(table entities.Table, timefrom, now string) (json_data string, data_len int) {

	jsonData := EsTableQuery(table.Columns, table.Instance, table.DataView, timefrom, now)

	// 將JSON字符串解碼為Go的map類型
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		log.Logrecord("ERROR", "Parsing JSON error at DataDealing stage: "+ err.Error())
		return
	}

	// 準備存儲結果的切片
	var finalResults []map[string]interface{}

	// 從 response 開始處理
	if response, ok := result["response"].(map[string]interface{}); ok {
		if aggs, ok := response["aggregations"].(map[string]interface{}); ok {
			for _, v := range aggs {
				if agg, ok := v.(map[string]interface{}); ok {
					flattenAggs(agg, &finalResults, map[string]string{})
				}
			}
		}
	}

	// 將最終結果轉換為 JSON 格式
	jsonOutput, err := json.MarshalIndent(finalResults, "", "  ")
	if err != nil {
		log.Logrecord("ERROR", "Converting to JSON error at DataDealing stage: "+ err.Error())
		// os.Exit(1)
	}

	// 輸出 JSON 結果
	// fmt.Println("yaya",string(jsonOutput))

	return string(jsonOutput), len(finalResults)
}

// 遞迴處理 JSON 的函數
func flattenAggs(data map[string]interface{}, result *[]map[string]interface{}, levels map[string]string) {
	if buckets, ok := data["buckets"].([]interface{}); ok {
		// 遍歷所有的 bucket
		for _, bucket := range buckets {
			bucketMap := bucket.(map[string]interface{})

			// 提取 key 和 doc_count
			key := bucketMap["key"].(string)
			docCount := int(bucketMap["doc_count"].(float64))

			// 複製當前層級的聚合信息
			newLevels := make(map[string]string)
			for k, v := range levels {
				newLevels[k] = v
			}

			// 確定這一層的聚合名稱
			levelName := fmt.Sprintf("%d", len(newLevels)+2)
			newLevels[levelName] = key

			// 檢查是否有下一層的聚合
			foundNext := false
			for k, v := range bucketMap {
				if subAgg, ok := v.(map[string]interface{}); ok && k[:4] == "agg_" {
					foundNext = true
					flattenAggs(subAgg, result, newLevels)
				}
			}

			// 如果沒有找到下一層的聚合，則添加結果到最終列表
			if !foundNext {
				aggResult := make(map[string]interface{})
				for k, v := range newLevels {
					aggResult[k] = v
				}
				aggResult["Count"] = docCount
				*result = append(*result, aggResult)
			}
		}
	}
}
