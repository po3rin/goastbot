package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/pkg/errors"
	"github.com/po3rin/godocbot/gendoc"
	"github.com/po3rin/godocbot/logger"
)

var msg = `aaaaaaa
fffff`

// Handler handle request.
func Handler(w http.ResponseWriter, r *http.Request) {
	channelSecret := os.Getenv("CHANNEL_SECRET")
	channelAccessToken := os.Getenv("CHANNEL_ACCESS_TOKEN")

	client := &http.Client{
		Timeout: time.Duration(15 * time.Second),
	}
	bot, err := linebot.New(channelSecret, channelAccessToken, linebot.WithHTTPClient(client))
	if err != nil {
		err = errors.Wrapf(err, "Failed to init Client: %+v", r)
		logger.Error(err)
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		fmt.Fprintf(w, "%v\n", err)
		return
	}
	received, err := bot.ParseRequest(r)
	if err != nil {
		err = errors.Wrapf(err, "Failed to parse request: %+v", r)
		logger.Error(err)
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		fmt.Fprintf(w, "%v\n", err)
		return
	}

	if len(received) == 0 {
		logger.Warnf("Warn to recieve no events")
	}

	for _, event := range received {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				source := event.Source
				if source.Type == linebot.EventSourceTypeUser {
					doc, err := gendoc.GenDoc(message.Text)
					if err != nil {
						logger.Error(err)
						postMessage := linebot.NewTextMessage("not found ...")
						if _, err = bot.ReplyMessage(event.ReplyToken, postMessage).Do(); err != nil {
							err = errors.Wrapf(err, "Failed to reply message: %+v", r)
							logger.Error(err)
							w.WriteHeader(500)
							fmt.Fprintf(w, "Done: %v", err)
							return
						}
						w.WriteHeader(404)
						fmt.Fprintf(w, "Done: %v", err)
						return
					}
					postMessage := linebot.NewTextMessage(doc)
					if _, err = bot.ReplyMessage(event.ReplyToken, postMessage).Do(); err != nil {
						err = errors.Wrapf(err, "Failed to reply message: %+v", r)
						logger.Error(err)
						w.WriteHeader(500)
						fmt.Fprintf(w, "Done: %v", err)
						return
					}
				} else {
					logger.Warnf("Warn to recieve unsupported source type: %+v\n", source.Type)
				}
			default:
				logger.Warnf("Warn to recieve unsupported message: %+v\n", message)
			}
		}
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "Done")
}
