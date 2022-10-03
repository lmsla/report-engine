package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/smtp"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"strings"
	"time"
)

// define email interface, and implemented auth and send method
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
    name        string
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

// 用 group name 取出群組中的 member
func GetMemberList(groupName string) ([]string) {
	inventory,err := GetMemberByGroupName(groupName)
	if err != nil {
		fmt.Println(err)
	}
	var memberList []string
	for _,data := range inventory {
		memberList = append(memberList, data.Email)
	}
	return memberList

}


// 取出filename by schedule ID 
func GetFilenameByscheduleID(scheduleID int) (entities.FileHistory, error){
    var instance entities.FileHistory
	instance.ScheduleID = scheduleID
	err := global.Mysql.Where("schedule_id= ? AND filetype = ?",scheduleID,"pdf" ).First(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil
}



func Sendmail(scheduleID int ,group []entities.Group) {
    // user := "rabot6201@gmail.com"
    // password := "mohptlqcqeiisshx"
    // host := "smtp.gmail.com"
    // port := "587"
    for _,group := range group {
        fmt.Println(group.Name)
    
    memberlist := GetMemberList(group.Name)
    fmt.Println(memberlist)
	user := global.EnvConfig.Email.User
    password := global.EnvConfig.Email.Password
    host := global.EnvConfig.Email.Host
    port := global.EnvConfig.Email.Port

    // fmt.Println(user)
    // fmt.Println(password)
    // fmt.Println(host)
    // fmt.Println(port)

    // //用 ScheduleID 取出 Schedule 的相關資料
    // inventory,err := GetScheduleBysSheduleID(scheduleID)
	// if err != nil {
	// 	fmt.Println(err)
	// }
    // //用 ReportName 取出 ReportID
    // inventory1,err := GetReportByReportName(inventory.Report)
    // 	if err != nil {
	// 	fmt.Println(err)
	// }
    // //用 ReportID 取出 Report 的相關資料 
	// inventory2, err := GetReportByReportID(inventory1.ReportID)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(inventory.Name)

	// fromtimeconverted := Timeconverter(inventory2.From+"+8h")
	// totimeconverted := Timeconverter(inventory2.To+"+8h")
	// str1 := fromtimeconverted[0:16] 
	// str2 := totimeconverted[0:16]
    // pdfPath := fmt.Sprintf("%s/%s %s~%s.pdf",global.EnvConfig.Reportengine.PdfPath,inventory1.Name,str1,str2)
    pdfName,err := GetFilenameByscheduleID(scheduleID)
    if err != nil {
		fmt.Println(err)
	}
    pdfPath := fmt.Sprintf("%s/%s.pdf", global.EnvConfig.Reportengine.PdfPath, pdfName.PdfName)
    fmt.Println(scheduleID)
    fmt.Println(pdfName.PdfName)

    var mail Mail
    mail = &SendMail{user: user, password: password, host: host, port: port}
    message := Message{from: user,
        // to:          []string{"russell.chen@bimap.co"},
        to:          memberlist,
        // cc:          []string{},
        // bcc:         []string{},
        subject:     "HELLO WORLD",
        body:        "report",
        contentType: "text/plain;charset=utf-8",
        // attachment: Attachment{
        //     name:        "test.jpg",
        //     contentType: "image/jpg",
        //     withFile:    true,
        // },
        attachment: Attachment{
            // name:        "/Users/chen/Documents/gitlab/git-out/product/report-backend/pdf/pdf_file/report01 2022-09-07 08:00:00~2022-09-07 16:24:25.pdf",
            name: pdfPath,
            contentType: "application/octet-stream",
            withFile:    true,
        },
    }
    err = newFunction(mail, message)
    if err != nil {
        fmt.Println("Send mail error!")
        fmt.Println(err)
    } else {
        fmt.Println("Send mail success!")
    }
}
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
    mail.Auth()
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

    body := "\r\n--" + boundary + "\r\n"
    body += "Content-Type:" + message.contentType + "\r\n"
    body += "\r\n" + message.body + "\r\n"
    buffer.WriteString(body)

    if message.attachment.withFile {
        attachment := "\r\n--" + boundary + "\r\n"
        attachment += "Content-Transfer-Encoding:base64\r\n"
        attachment += "Content-Disposition:attachment\r\n"
        attachment += "Content-Type:" + message.attachment.contentType + ";name=\"" + mime.BEncoding.Encode("UTF-8", message.attachment.name) + "\"\r\n"
        buffer.WriteString(attachment)
        defer func() {
            if err := recover(); err != nil {
                log.Fatalln(err)
            }
        }()
        mail.writeFile(buffer, message.attachment.name)
    }

    to_address := MergeSlice(message.to, message.cc)
    to_address = MergeSlice(to_address, message.bcc)

    buffer.WriteString("\r\n--" + boundary + "--")
    err := smtp.SendMail(mail.host+":"+mail.port, mail.auth, message.from, to_address, buffer.Bytes())
    return err
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
    file, err := ioutil.ReadFile(fileName)
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