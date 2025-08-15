package services

import (
	"fmt"
	"net/smtp"
	"report-backend-golang/global"
)

// SMTPConfig 儲存偵測到的 SMTP 配置
type SMTPConfig struct {
	Host       string
	Port       string
	Auth       bool
	AuthType   string
	DisableTLS bool
	StartTLS   bool
	SSLPort    bool
	Error      error
}

// SMTPDetector 自動偵測 SMTP 伺服器配置
type SMTPDetector struct {
	host     string
	port     string
	username string
	password string
}

// NewSMTPDetector 創建新的 SMTP 偵測器
func NewSMTPDetector(host, port, username, password string) *SMTPDetector {
	return &SMTPDetector{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

// DetectSMTPConfig 自動偵測 SMTP 伺服器配置
func (d *SMTPDetector) DetectSMTPConfig() *SMTPConfig {
	config := &SMTPConfig{
		Host: d.host,
		Port: d.port,
	}

	// 測試不同的連線方式
	tests := []struct {
		name     string
		testFunc func() error
	}{
		{"SSL with Auth", d.testSSLAuth},
		{"STARTTLS with Auth", d.testStartTLSAuth},
		{"Plain Auth", d.testPlainAuth},
		{"No Auth with TLS", d.testNoAuthTLS},
		{"No Auth No TLS", d.testNoAuthNoTLS},
	}

	for _, test := range tests {
		if err := test.testFunc(); err == nil {
			// 成功找到可用的連線方式
			d.setConfigFromTest(test.name, config)
			return config
		} else {
			fmt.Printf("Test %s failed: %v\n", test.name, err)
		}
	}

	config.Error = fmt.Errorf("no working SMTP configuration found for %s:%s", d.host, d.port)
	return config
}

// testSSLAuth 測試 SSL + 認證
func (d *SMTPDetector) testSSLAuth() error {
	addr := fmt.Sprintf("%s:%s", d.host, d.port)
	auth := smtp.PlainAuth("", d.username, d.password, d.host)

	// 嘗試 LoginAuth
	if err := smtp.SendMail(addr, auth, "test@example.com", []string{"test@example.com"}, []byte("test")); err == nil {
		return nil
	}

	// 嘗試 PlainAuth
	auth = smtp.PlainAuth("", d.username, d.password, d.host)
	return smtp.SendMail(addr, auth, "test@example.com", []string{"test@example.com"}, []byte("test"))
}

// testStartTLSAuth 測試 STARTTLS + 認證
func (d *SMTPDetector) testStartTLSAuth() error {
	addr := fmt.Sprintf("%s:%s", d.host, d.port)

	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Quit()

	// 嘗試 STARTTLS
	if err = c.StartTLS(nil); err != nil {
		return err
	}

	// 嘗試認證
	auth := smtp.PlainAuth("", d.username, d.password, d.host)
	if err = c.Auth(auth); err != nil {
		return err
	}

	return nil
}

// testPlainAuth 測試純認證 (無 TLS)
func (d *SMTPDetector) testPlainAuth() error {
	addr := fmt.Sprintf("%s:%s", d.host, d.port)

	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Quit()

	auth := smtp.PlainAuth("", d.username, d.password, d.host)
	return c.Auth(auth)
}

// testNoAuthTLS 測試無認證 + TLS
func (d *SMTPDetector) testNoAuthTLS() error {
	addr := fmt.Sprintf("%s:%s", d.host, d.port)
	return smtp.SendMail(addr, nil, "test@example.com", []string{"test@example.com"}, []byte("test"))
}

// testNoAuthNoTLS 測試無認證 + 無 TLS
func (d *SMTPDetector) testNoAuthNoTLS() error {
	addr := fmt.Sprintf("%s:%s", d.host, d.port)

	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Quit()

	if err = c.Mail("test@example.com"); err != nil {
		return err
	}

	return c.Rcpt("test@example.com")
}

// setConfigFromTest 根據測試結果設定配置
func (d *SMTPDetector) setConfigFromTest(testName string, config *SMTPConfig) {
	switch testName {
	case "SSL with Auth":
		config.Auth = true
		config.DisableTLS = false
		config.AuthType = "PlainAuth"
	case "STARTTLS with Auth":
		config.Auth = true
		config.DisableTLS = false
		config.StartTLS = true
		config.AuthType = "PlainAuth"
	case "Plain Auth":
		config.Auth = true
		config.DisableTLS = true
		config.AuthType = "PlainAuth"
	case "No Auth with TLS":
		config.Auth = false
		config.DisableTLS = false
	case "No Auth No TLS":
		config.Auth = false
		config.DisableTLS = true
	}
}

// DetectAndUpdateConfig 偵測並更新全域配置
func DetectAndUpdateConfig(host, port, username, password string) error {
	detector := NewSMTPDetector(host, port, username, password)
	config := detector.DetectSMTPConfig()

	if config.Error != nil {
		return config.Error
	}

	// 更新全域配置
	global.EnvConfig.Email.Host = config.Host
	global.EnvConfig.Email.Port = config.Port
	global.EnvConfig.Email.User = username
	global.EnvConfig.Email.Password = password
	global.EnvConfig.Email.Auth = config.Auth
	global.EnvConfig.Email.AuthType = config.AuthType
	global.EnvConfig.Email.DisableTLS = config.DisableTLS

	fmt.Printf("Auto-detected SMTP config: Auth=%v, AuthType=%s, DisableTLS=%v\n",
		config.Auth, config.AuthType, config.DisableTLS)

	return nil
}

// TestSMTPConnection 測試 SMTP 連線
func TestSMTPConnection(host, port, username, password string) error {
	detector := NewSMTPDetector(host, port, username, password)
	config := detector.DetectSMTPConfig()

	if config.Error != nil {
		return fmt.Errorf("SMTP connection test failed: %w", config.Error)
	}

	fmt.Printf("SMTP connection test successful!\n")
	fmt.Printf("Configuration: %+v\n", config)

	return nil
}
