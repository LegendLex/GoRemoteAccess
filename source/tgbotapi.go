package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"tell-my-server-bot/misisapi"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	//Creating DB Table
	if os.Getenv("CREATE_TABLE") == "yes" && os.Getenv("DB_SWITCH") == "on" {
		if err := createTable(); err != nil {
			panic(err)
		}
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			commandEnd := strings.IndexByte(update.CallbackQuery.Data, ':')
			if commandEnd < 0 {
				continue
			}
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Принял. Обработал.")
			bot.Request(callback)
			switch update.CallbackQuery.Data[:commandEnd] {
			case "showScheduleForToday":
				today := time.Now()
				weekday := int(today.Weekday())
				if weekday == 0 {
					weekday = 7
				}
				schedule := misisapi.GetSchedule(update.CallbackQuery.Data[commandEnd+1:], today.Format("2006-01-02"))
				todayScheduleText := schedule.GetDay(weekday)
				msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, todayScheduleText)

				bot.Send(msg)
			case "showScheduleForTomorrow":
				tomorrow := time.Now().Add(24 * time.Hour)
				weekday := int(tomorrow.Weekday())
				if weekday == 0 {
					weekday = 7
				}
				schedule := misisapi.GetSchedule(update.CallbackQuery.Data[commandEnd+1:], tomorrow.Format("2006-01-02"))
				todayScheduleText := schedule.GetDay(weekday)
				msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, todayScheduleText)

				bot.Send(msg)
			case "saveAsDefault":
				if os.Getenv("DB_SWITCH") == "on" {
					err := saveGroupData(update.CallbackQuery.From.UserName, update.CallbackQuery.From.ID, update.CallbackQuery.Data[commandEnd+1:])
					if err != nil {
						msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "Database Error")
						bot.Send(msg)
					}

				} else {
					msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "Не могу сохранить группу, так как база данных не подключена.")

					bot.Send(msg)
				}
			default:

			}
		} else if update.Message != nil { // If we got a message
			switch update.Message.Text {
			// Welcome message
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					`Данный бот разработан в качестве персонального помощника.
					На данный момент поддерживаются команды:
					/schedule - вывод расписания для сохранённой группы.
					Для вывода расписания другой группы или смены сохранённой группы введите её название в формате "БПМ-20-1" без кавычек.`)

				bot.Send(msg)
				// Returns current time
			case "/time":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint(time.Now().Clock()))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
				// Returns schedule for today
			case "/schedule":
				msgText := ""
				groupId := ""
				if os.Getenv("DB_SWITCH") == "on" {
					userGroupId, err := getGroupData(update.Message.From.ID)
					if err != nil || userGroupId == "" {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "База данных недоступна или нет сохранённой группы")
						bot.Send(msg)
						break
					} else {
						msgText = "Варианты расписаний для вашей группы:"
						groupId = userGroupId
					}
				} else { // DB is off
					msgText = "База данных не подключена. Группа по умолчанию: БПМ-20-1"
					groupId = "6892"
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
				msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{
					tgbotapi.NewInlineKeyboardButtonData("Расписание на сегодня", fmt.Sprintf("showScheduleForToday:%v", groupId)),
					tgbotapi.NewInlineKeyboardButtonData("Расписание на завтра", fmt.Sprintf("showScheduleForTomorrow:%v", groupId)),
				})
				bot.Send(msg)

				// Work with zero-text messages (stickers, photos, docs)
			case "":
				//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				//msg.ReplyToMessageID = update.Message.MessageID
				//bot.Send(msg)

				// Default
			default:
				if groupIsConfirmed, groupID := misisapi.GetGroups(update.Message.Text); groupIsConfirmed {
					msgText := fmt.Sprintf("Группа \"%v\" найдена. Выберите действие ниже:", update.Message.Text)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
					buttons1 := []tgbotapi.InlineKeyboardButton{
						tgbotapi.NewInlineKeyboardButtonData("Расписание на сегодня", fmt.Sprintf("showScheduleForToday:%v", groupID)),
						tgbotapi.NewInlineKeyboardButtonData("Расписание на завтра", fmt.Sprintf("showScheduleForTomorrow:%v", groupID)),
					}
					buttons2 := []tgbotapi.InlineKeyboardButton{
						tgbotapi.NewInlineKeyboardButtonData("Сохранить группу по умолчанию", fmt.Sprintf("saveAsDefault:%v", groupID)),
					}
					keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons1, buttons2)
					msg.ReplyMarkup = keyboard
					bot.Send(msg)
				} else {
					msgText := "Введите корректное название группы в формате \"БПМ-20-1\""
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
					bot.Send(msg)
				}
			}
		}
	}
}
