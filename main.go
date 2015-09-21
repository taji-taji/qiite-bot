package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func main() {
	const (
		tCkey = "MT7FaKikyXrcsJs1Zeb63wp45"
		tCsec = "gdImSZEePjyljGOnfrBYS2qRJXzEUIvnJYKso8OOPGGEhbSsgD"
		tAtok = "3624925453-Z4Zb6rNeJUwfHtFK3oEvsHnUfn5D3p7PsqhzFjU"
		tAsec = "AzURW2USvEW3rSjKRHB6drHQUTu26t2J5N4IhSzlUqCl6"
		qUrl  = "https://qiita.com/api/v2/"
	)
	anaconda.SetConsumerKey(tCkey)
	anaconda.SetConsumerSecret(tCsec)
	api := anaconda.NewTwitterApi(tAtok, tAsec)

	v := url.Values{}
	stream := api.UserStream(v)

	for {
		// 受信待ち
		select {
		case item := <-stream.C:
			switch status := item.(type) {
			case anaconda.Tweet:
				// Tweet を受信
				// リプライテキスト
				replyText := status.Text
				replyToScreenName := status.InReplyToScreenName
				if replyToScreenName == "qiite_bot" {
					re, _ := regexp.Compile("^@" + replyToScreenName)
					replyText = re.ReplaceAllString(replyText, "")
					replyText = strings.TrimSpace(replyText)
					requestUrl := qUrl + "items?page=1&per_page=100&query=" + replyText
					fmt.Printf("%v", requestUrl)
					res, err := http.Get(requestUrl)
					if err != nil {
						fmt.Printf("%#v\n", err)
					} else {
						fmt.Printf("%#v\n", res)
					}
				}
				// ツイートユーザー
				senduser := status.User
				fmt.Printf("%#v\n", replyToScreenName)
				fmt.Printf("%s: %s\n", senduser.ScreenName, replyText)
			default:
			}
		}
	}
}
