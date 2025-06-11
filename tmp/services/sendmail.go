package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	log1 "log"
	"mime"
	"net/smtp"
	"report-backend-golang/global"
	"report-backend-golang/log"
	"report-backend-golang/tools"
	"strings"
	"time"
	"errors"
)

type Mail interface {
	Auth()
	Send(message Message) error
}

type SendMail struct {
	user     string
	password string
	host     string
	port     string
	auth     smtp.Auth
}

type Attachment struct {
	name        []string
	filepath    []string
	contentType string
	withFile    bool
}

type Message struct {
	from        string
	to          []string
	cc          []string
	bcc         []string
	subject     string
	body        string
	contentType string
	attachment  Attachment
}

func SendEmailBySchedule(nowtime int64, ScheduleID int) (err error) {

	defer func() {
		if err != nil {
			log.Logrecord("ERROR", "func SendEmailBySchedule error")
		}
	}()

	user := global.EnvConfig.Email.User
	password := global.EnvConfig.Email.Password
	host := global.EnvConfig.Email.Host
	port := global.EnvConfig.Email.Port
	report_path := global.EnvConfig.Files.ReportFile

	// GetReportByScheduleID(ScheduleID)

	report_data, err := GetReportByScheduleID(ScheduleID)
	if err != nil {
		fmt.Println(err)
	}
	// timefrom := tools.Timeconverter(report_data.TimeUnit,report_data.TimePeriod)
	var reportForSendList []string
	var nameList []string

	t := time.Unix(nowtime, 0)
	now := t.Format("2006-01-02")

	for _, report := range report_data {
		timefrom := tools.Timeconverter(nowtime, report.TimeUnit, report.TimePeriod, report.Alias)
		reportname := fmt.Sprintf("%s_%s_%s", report.Name, timefrom, now)
		nameList = append(nameList, reportname+".pdf")
		reportForSendList = append(reportForSendList, report_path+reportname)

	}
	var text string
	text = "Report Center 排程\n"
	for _, report := range report_data {
		timefrom := tools.Timeconverter(nowtime, report.TimeUnit, report.TimePeriod, report.Alias)
		text += fmt.Sprintf("%s，報表區間：%s ~ %s\n", report.Name, timefrom, now)
	}

	var reciver_list []string
	var cc_list []string
	var bcc_list []string
	schedule_data, err := GetScheduleBySheduleID(ScheduleID)
	if err != nil {
		fmt.Println(err)
	}

	for _, to := range schedule_data.To {
		reciver_list = append(reciver_list, to)

	}

	for _, cc := range schedule_data.CC {
		cc_list = append(cc_list, cc)

	}

	for _, bcc := range schedule_data.CC {
		bcc_list = append(bcc_list, bcc)

	}

	var mail Mail

	if user == "" {
		mail = &SendMail{host: host, port: port}
	} else {
		mail = &SendMail{user: user, password: password, host: host, port: port}
	}
	// fmt.Println("mail",mail)

	message := Message{from: global.EnvConfig.Email.Sender,
		to:          reciver_list,
		cc:          cc_list,
		bcc:         bcc_list,
		subject:     schedule_data.Name+"-"+ schedule_data.Subject,
		body:        text + schedule_data.Body,
		contentType: "text/plain;charset=utf-8",
		// attachment: Attachment{
		//     name:        "test.jpg",
		//     contentType: "image/jpg",
		//     withFile:    true,
		// },
		attachment: Attachment{
			filepath:    reportForSendList,
			name:        nameList,
			contentType: "application/octet-stream",
			withFile:    true,
		},
	}

	// mail.Send(message)
	err = newFunction(mail, message)
	if err != nil {
		log.Logrecord("排程", "信件送出失敗")
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		log.Logrecord("排程", "信件送出成功")
		fmt.Println("Send mail success!")
	}
	return err
}

func newFunction(mail Mail, message Message) error {
	err := mail.Send(message)
	return err
}

func (mail *SendMail) Auth() {
	// mail.auth = smtp.PlainAuth("", mail.user, mail.password, mail.host)
	mail.auth = LoginAuth(mail.user, mail.password)
}

func (mail SendMail) Send(message Message) error {
	// mail.Auth()

	// to_address := MergeSlice(message.to, message.cc)
	// to_address = MergeSlice(to_address, message.bcc)

	toAddress := MergeSlice(message.to, message.cc)
	toAddress = MergeSlice(toAddress, message.bcc)

	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	Header := make(map[string]string)
	Header["From"] = message.from
	Header["To"] = strings.Join(message.to, ";")
	Header["Cc"] = strings.Join(message.cc, ";")
	Header["Bcc"] = strings.Join(message.bcc, ";")
	Header["Subject"] = message.subject
	Header["Content-Type"] = "multipart/mixed;boundary=" + boundary
	Header["Mime-Version"] = "1.0"
	Header["Date"] = time.Now().String()
	mail.writeHeader(buffer, Header)
	// 寫入 HTML 內容
	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type:" + message.contentType + "\r\n"
	body += "\r\n" + message.body + "\r\n"
	buffer.WriteString(body)

	for _, name := range message.attachment.name {
		if message.attachment.withFile {
			attachment := "\r\n--" + boundary + "\r\n"
			attachment += "Content-Transfer-Encoding:base64\r\n"
			attachment += "Content-Disposition:attachment\r\n"
			attachment += "Content-Type:" + message.attachment.contentType + ";name=\"" + mime.BEncoding.Encode("UTF-8", name) + "\"\r\n"
			buffer.WriteString(attachment)
			defer func() {
				if err := recover(); err != nil {
					log1.Fatalln(err)
				}
			}()
			// mail.writeFile(buffer, message.attachment.name)
			mail.writeFile(buffer, global.EnvConfig.Files.ReportFile+"/"+name)
			// mail.writeFile(buffer, global.EnvConfig.Files.ReportFile+"/"+name+".pdf")
		}
	}

	// 決定發信方式
	addr := fmt.Sprintf("%s:%s", mail.host, mail.port)
	from := message.from
	msg := buffer.Bytes()

	// 分支邏輯來自原 SendEmailNotify
	switch {
	case global.EnvConfig.Email.Auth:
		var auth smtp.Auth
		if global.EnvConfig.Email.AuthType == "LoginAuth" {
			auth = LoginAuth(mail.user, mail.password)
		} else {
			auth = smtp.PlainAuth("", mail.user, mail.password, mail.host)
		}
		if err := smtp.SendMail(addr, auth, from, toAddress, msg); err != nil {
			return fmt.Errorf("SendMail with Auth failed: %w", err)
		}
		return nil

	case !global.EnvConfig.Email.DisableTLS:
		if err := smtp.SendMail(addr, nil, from, toAddress, msg); err != nil {
			return fmt.Errorf("SendMail with TLS but no Auth failed: %w", err)
		}
		return nil

	case global.EnvConfig.Email.DisableTLS:
		// 模擬 NoAuth + NoTLS
		c, err := smtp.Dial(addr)
		if err != nil {
			return fmt.Errorf("dial failed: %w", err)
		}
		defer c.Quit()

		if err = c.Mail(from); err != nil {
			return fmt.Errorf("MAIL FROM failed: %w", err)
		}
		for _, rcpt := range toAddress {
			if err = c.Rcpt(rcpt); err != nil {
				return fmt.Errorf("RCPT TO failed (%s): %w", rcpt, err)
			}
		}

		wc, err := c.Data()
		if err != nil {
			return fmt.Errorf("DATA failed: %w", err)
		}
		defer wc.Close()

		if _, err = wc.Write(msg); err != nil {
			return fmt.Errorf("write message failed: %w", err)
		}
		return nil

	default:
		return errors.New("no valid SMTP auth/TLS configuration found")
	}

	// to_address := MergeSlice(message.to, message.cc)
	// to_address = MergeSlice(to_address, message.bcc)

	// buffer.WriteString("\r\n--" + boundary + "--")
	// err := smtp.SendMail(mail.host+":"+mail.port, mail.auth, message.from, to_address, buffer.Bytes())
	// return err
}

func MergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}

func (mail SendMail) writeHeader(buffer *bytes.Buffer, Header map[string]string) string {
	header := ""
	for key, value := range Header {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)
	return header
}

// read and write the file to buffer
func (mail SendMail) writeFile(buffer *bytes.Buffer, fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(payload, file)
	buffer.WriteString("\r\n")
	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	// return "LOGIN", []byte{}, nil
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		}
	}
	return nil, nil
}
