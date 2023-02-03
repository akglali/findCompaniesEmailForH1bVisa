package sendEmailToFoundCompanies

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"
)

func SendEmail() {
	// to be able to read .env file we should load first
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Something went wrong with env file Env file ")
		return
	}

	f, err := excelize.OpenFile("employerH1B.xlsx")

	if err != nil {
		log.Fatal(err)
	}
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		//companyName := row[1]
		companyEmails := row[4]

		// Remove square brackets from string
		array := companyEmails[1 : len(companyEmails)-1]

		// Split the string into a slice of strings using the space character as a delimiter
		emailArray := strings.Split(array, " ")
		for i, email := range emailArray {
			fmt.Println(i, email)
		}

	}
}

//by default, it is .env, so we don't have to write

func EmailTemplate(companyName, email string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Something went wrong with env file Env file ")
		return
	}
	from := os.Getenv("MAIL")
	password := os.Getenv("PASSWD")
	host := "smtp.gmail.com"
	port := "587"
	name := "Ali Akgol"
	body := "To Whom It May Concern,\n\n" +
		"I hope this email finds you well. I am writing to express my interest in a software engineering role at " + companyName +
		" As a recently graduated computer engineer with a 3.35 GPA, I am eager to put my skills and knowledge to use in a challenging and dynamic work environment." +
		"\n\nOne of my main objectives is to pursue a career in the United States, and I believe that " + companyName + " would provide me with the opportunities I am seeking." +
		" I am highly motivated, ambitious and eager to take on new challenges, and I believe that I would be an asset to your team.\n\n" +
		"I have a strong foundation in various programming languages such as Go, Java, Python, TypeScript, JavaScript, SQL, and Dart." +
		" I am proficient in front-end frameworks like React, Angular, NextJS, and Flutter, and I have experience in creating robust and scalable" +
		" web applications using backend frameworks like ExpressJS, NestJS, and Gin.\n\nI am confident that my technical skills, combined with my passion" +
		" for software engineering and my commitment to continuous learning and growth, make me a strong candidate for this position. I would be honored " +
		"to have the opportunity to bring my skills and experience to your team and to be a part of a company that is making a difference in the industry.\n\n" +
		"I would appreciate the opportunity to discuss my qualifications and how I can contribute to the success of your company. I am available for an interview at your convenience. " +
		"I have attached my resume and portfolio for your review.\n\n" +
		"I was searching for a company that offers H1B visa sponsorship, and I came across the website h1bgrader," +
		" which lists all the companies that have sponsored visas in the past. However, the website did not provide any " +
		"email addresses to contact these companies. So, I teamed up with my friend Sultan and wrote a code to find the websites" +
		" of these companies using their names. We then searched for email addresses on these websites by writing a function that" +
		" searched for them in the text. To send the emails, we created a template email and wrote another function to send the emails. " +
		"Although writing the code wasn't particularly difficult, I believe it demonstrates my dedication and determination to make my dream of" +
		" starting my career in the US a reality. The code is available on GitHub! \n\n" +
		"\n\nThank you for considering my application. I look forward to hearing from you soon." +
		"\n\nBest regards," +
		"\n\n" + name + "\r\n"
	// Connect to the SMTP server
	auth := smtp.PlainAuth("", from, password, host)
	to := email
	// Read the attachment file
	file, err := ioutil.ReadFile("resume.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Encode the attachment to base64
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(encoded, file)
	// Create the message
	var message bytes.Buffer
	message.WriteString("From: " + from + "\r\n")
	message.WriteString("To: " + to + "\r\n")
	message.WriteString("Subject: H1B Visa Sponsorship\r\n")
	message.WriteString("MIME-Version: 1.0\r\n")
	message.WriteString("Content-Type: multipart/mixed; boundary=frontier\r\n")
	message.WriteString("\r\n")
	message.WriteString("--frontier\r\n")
	message.WriteString("Content-Type: text/plain\r\n")
	message.WriteString("\r\n")
	message.WriteString(body)
	message.WriteString("\r\n")
	message.WriteString("--frontier\r\n")
	message.WriteString("Content-Type: application/pdf\r\n")
	message.WriteString("Content-Transfer-Encoding: base64\r\n")
	message.WriteString("Content-Disposition: attachment; filename=resume.pdf\r\n")
	message.WriteString("\r\n")
	message.Write(encoded)
	message.WriteString("\r\n")
	message.WriteString("--frontier--\r\n")
	// Send the email
	err = smtp.SendMail(host+":"+port, auth, from, []string{to}, message.Bytes())
	if err != nil {
		fmt.Println("Email Couldn't send" + companyName)
		return
	}

	fmt.Println("Email sent successfully!")
}