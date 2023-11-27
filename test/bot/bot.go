package main

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// 替换为您的机器人API令牌
	bot, err := tgbotapi.NewBotAPI("6370023079:AAHnSo3TzXxDbvVTlU5sFkkv4ayebkpGZO4")
	if err != nil {
		log.Panic(err)
	}

	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60

	// updates, err := bot.GetUpdatesChan(u)

	// for update := range updates {
	// 	if update.Message != nil {
	// 		// 获取消息的聊天ID
	// 		chatID := update.Message.Chat.ID

	// 		// 发送消息给用户
	// 		msg := tgbotapi.NewMessage(chatID, "这是机器人主动发送的消息！")
	// 		bot.Send(msg)
	// 	}

	// 	// 等待一段时间后进行下一次轮询
	// 	time.Sleep(5 * time.Second)
	// }

	// 存储用户聊天 ID 的映射
	userChatIDs := map[int64]struct{}{}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)


	msg:=tgbotapi.NewMessage(6002330934, "这是机器人主动发送的消息！")
	bot.Send(msg);



	for update := range updates {
		if update.Message != nil {
			// 获取消息的聊天ID
			chatID := update.Message.Chat.ID

			// 将聊天ID存储到映射中，确保唯一性
			userChatIDs[chatID] = struct{}{}

			// 处理用户发来的消息，可以根据需要进行其他逻辑处理
			handleIncomingMessage(update.Message, bot)
		}
	}

	// 定期向所有用户发送消息
	for chatID := range userChatIDs {
		// 创建要发送的消息
		msg := tgbotapi.NewMessage(chatID, "这是定期发送的消息！")

		// 发送消息
		bot.Send(msg)

		// 等待一段时间后进行下一次发送
		time.Sleep(10 * time.Second)
	}

	
}

func handleIncomingMessage(msg *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	// 在这里处理用户发来的消息，根据需要进行其他逻辑
	log.Printf("[%s] %s %s", msg.From.UserName, msg.Text, msg.Chat.ID)

	// 例如，可以回复用户
	reply := tgbotapi.NewMessage(msg.Chat.ID, "机器人收到您的消息了！")
	bot.Send(reply)
}
