package controller

import (
	"net/http"
	"report-backend-golang/services"

	"github.com/gin-gonic/gin"
)

// TestSMTPConnection 測試 SMTP 連線
// @Summary 測試 SMTP 連線
// @Description 測試指定的 SMTP 伺服器連線並自動偵測配置
// @Tags SMTP
// @Accept json
// @Produce json
// @Param smtp body SMTPTestRequest true "SMTP 測試請求"
// @Success 200 {object} SMTPTestResponse "測試成功"
// @Failure 400 {object} ErrorResponse "請求格式錯誤"
// @Failure 500 {object} ErrorResponse "SMTP 連線失敗"
// @Router /test-smtp [post]
func TestSMTPConnection(c *gin.Context) {
	var req SMTPTestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request: " + err.Error(),
		})
		return
	}

	// 測試 SMTP 連線
	err := services.TestSMTPConnection(req.Host, req.Port, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "SMTP connection test failed: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SMTP connection test successful",
		"host":    req.Host,
		"port":    req.Port,
	})
}

// DetectSMTPConfig 自動偵測 SMTP 配置
// @Summary 自動偵測 SMTP 配置
// @Description 自動偵測 SMTP 伺服器的連線方式和認證類型
// @Tags SMTP
// @Accept json
// @Produce json
// @Param smtp body SMTPTestRequest true "SMTP 偵測請求"
// @Success 200 {object} SMTPTestResponse "偵測成功"
// @Failure 400 {object} ErrorResponse "請求格式錯誤"
// @Failure 500 {object} ErrorResponse "偵測失敗"
// @Router /detect-smtp [post]
func DetectSMTPConfig(c *gin.Context) {
	var req SMTPTestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request: " + err.Error(),
		})
		return
	}

	// 自動偵測 SMTP 配置
	err := services.DetectAndUpdateConfig(req.Host, req.Port, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "SMTP detection failed: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SMTP configuration detected and updated",
		"host":    req.Host,
		"port":    req.Port,
	})
}

// SMTPTestRequest SMTP 測試請求
type SMTPTestRequest struct {
	Host     string `json:"host" binding:"required"`
	Port     string `json:"port" binding:"required"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// SMTPTestResponse SMTP 測試回應
type SMTPTestResponse struct {
	Message string `json:"message"`
	Host    string `json:"host"`
	Port    string `json:"port"`
}

// SMTPConfigResponse SMTP 配置回應
type SMTPConfigResponse struct {
	Host       string `json:"host"`
	Port       string `json:"port"`
	Auth       bool   `json:"auth"`
	AuthType   string `json:"auth_type"`
	DisableTLS bool   `json:"disable_tls"`
	StartTLS   bool   `json:"start_tls"`
}

// ErrorResponse 錯誤回應
type ErrorResponse struct {
	Error string `json:"error"`
}
