// Copyright © 2018 Bassirou Sarr
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send an email with sendgrid",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		message := constructMessageFromCLI(cmd)
		_, err := client.Send(message)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n%v\n",
			"Message was sent successfully",
			message)
		return err
	},
}

func constructMessageFromCLI(cmd *cobra.Command) *mail.SGMailV3 {
	from := mail.NewEmail(cmd.Flag("sender-name").Value.String(),
		cmd.Flag("sender-addr").Value.String())
	subject := cmd.Flag("subject").Value.String()
	to := mail.NewEmail(cmd.Flag("receiver-name").Value.String(),
		cmd.Flag("receiver-addr").Value.String())
	plainTextContent := cmd.Flag("text").Value.String()
	htmlContent := cmd.Flag("html").Value.String()
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	return message
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringP("sender-name", "", "example", "sender's name")
	sendCmd.Flags().StringP("sender-addr", "", "example@example.com", "sender's email address")
	sendCmd.Flags().StringP("receiver-name", "", "example", "receiver's name")
	sendCmd.Flags().StringP("receiver-addr", "", "example@example.com", "receiver's email address")
	sendCmd.Flags().StringP("subject", "", "Sending with SendGrid is Fun", "email's subject")
	sendCmd.Flags().StringP("html", "", "<strong>and easy to do anywhere, even with Go</strong>", "email's plain text content")
	sendCmd.Flags().StringP("text", "", "and easy to do anywhere, even with Go", "email's plain text content")

}
