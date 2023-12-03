package main

import (
	"botforlive/getdataex"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	//разные переменные
	screaming = true
	sheet     = "facel"

	// Текст заголовка
	firstMenu = "<b>Привет</b>\n\n Выбери объект."
	// првет тест для git
	//Да как так??
	// Текст кнопок
	uuug        = "Утилизация"
	facel       = "Факел"
	mckCeh6     = "МЦК Цех 6"
	mckCeh14    = "МЦК Цех 14"
	kondensatka = "Конденсатка"
	voadushka   = "Воздушечка"
	azotka      = "Азоточка"
	vodoblok    = "Водоблок"
	ten         = "десятка"
	nine        = "девятка"
	ppr         = "ППР"
	oneWeek     = "Первая неделя"
	twoWeek     = "Вторая неделя"
	threeWeek   = "третья неделя"
	fourWeek    = "четвёртая неделя"
	backButton  = "<= назад"

	// Store bot screaming status. статус бота наверное логи
	//screaming = false
	bot *tgbotapi.BotAPI

	// Keyboard layout for the first menu. One button, one row
	firstMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(uuug, uuug),
			tgbotapi.NewInlineKeyboardButtonData(mckCeh6, mckCeh6),
			tgbotapi.NewInlineKeyboardButtonData(kondensatka, kondensatka),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(facel, facel),
			tgbotapi.NewInlineKeyboardButtonData(mckCeh14, mckCeh14),
			tgbotapi.NewInlineKeyboardButtonData(voadushka, voadushka),
		),

		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(azotka, azotka),
			tgbotapi.NewInlineKeyboardButtonData(vodoblok, vodoblok),
			tgbotapi.NewInlineKeyboardButtonData(nine, nine),
			tgbotapi.NewInlineKeyboardButtonData(ten, ten),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(ppr, ppr),
		),
	)

	pprMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(oneWeek, oneWeek),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(twoWeek, twoWeek),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(threeWeek, threeWeek),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(fourWeek, fourWeek),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(backButton, backButton),
		),
	)
)

func main() {

	//Обработка ошибок
	var err error
	//токен с ключём
	bot, err = tgbotapi.NewBotAPI("6203992091:AAHKnR_ySFAvniox0UVb6yHlEI7WFZ22--c")
	if err != nil {
		// Abort if something is wrong
		log.Panic(err)
	}

	// Set this to true to log all interactions with telegram servers
	//Логи отключены
	bot.Debug = false

	//проверка соеденения
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Разобраться с контекстом
	// Create a new cancellable background context. Calling `cancel()` leads to the cancellation of the context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	//
	// `updates` is a golang channel which receives telegram updates
	updates := bot.GetUpdatesChan(u)

	// Запускаем функцию передаём апдейт и контекст
	// Pass cancellable context to goroutine
	go receiveUpdates(ctx, updates)

	// Tell the user the bot is online
	log.Println("Start listening for updates. Press enter to stop")

	// Wait for a newline symbol, then cancel handling updates
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()

}

// проверяет обновился чат или нет
func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		select {
		//останавливает цикл если пришел cancel
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		//передаём данные дальше если что-то есть запускаем функцию handleUpdate
		// receive update from channel and then handle it
		case update := <-updates:
			handleUpdate(update)
		}
	}
}

// проверяет кнопка или нет
// Если кнопка функция handleMessage
// если текст функция handleButton

func handleUpdate(update tgbotapi.Update) {
	switch {
	// Handle messages
	case update.Message != nil:
		handleMessage(update.Message)
	// Handle button clicks
	case update.CallbackQuery != nil:
		handleButton(update.CallbackQuery)
	}
}

// функция записывает в переменные uzr и текст выводит их в консоль и в чат
func handleMessage(message *tgbotapi.Message) {
	user := message.From
	text := message.Text

	if user == nil {
		return
	}

	//вывод в консоль
	// Print to console
	//
	log.Printf("%s wrote %s", user.FirstName, text)

	var err error
	//если команда запускаем фнкцию handleCommand
	if strings.HasPrefix(text, "/") {
		err = handleCommand(message.Chat.ID, text)
		//или если длинна текста больше 0 и screaming tru то выводим текст большими буквами
	} else if text == "Первая неделя" {
		fmt.Println("Привет")

	} else if screaming && len(text) > 0 {
		result := getdataex.Getdataex(text, sheet)
		fmt.Println(sheet)
		if len(result) == 0 {
			msg := tgbotapi.NewMessage(message.Chat.ID, "ничего не найденно попробуй ввести последние цифры позиции. В чек листах позиции указанны через тире пример 73-PDIRSA-180A")
			// To preserve markdown, we attach entities (bold, italic..)
			//msg.Entities = message.Entities
			//msg.ParseMode = tgbotapi.ModeHTML
			_, err = bot.Send(msg)
		}

		for _, res := range result {
			fmt.Println(res)

			msg := tgbotapi.NewMessage(message.Chat.ID, res)
			// To preserve markdown, we attach entities (bold, italic..)
			//msg.Entities = message.Entities
			msg.ParseMode = tgbotapi.ModeHTML
			_, err = bot.Send(msg)
		}

		err = handleCommand(message.Chat.ID, "/start")
		//иначе копируем всё, что летит обатно в чат в том виде в котором отправляли
	} else {
		// This is equivalent to forwarding, without the sender's name
		copyMsg := tgbotapi.NewCopyMessage(message.Chat.ID, message.Chat.ID, message.MessageID)
		_, err = bot.CopyMessage(copyMsg)
	}

	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}

}

// When we get a command, we react accordingly
// функция меняет статс большие или маленькие буквы запускает
// или запускает меню
func handleCommand(chatId int64, command string) error {
	var err error

	switch command {
	//Делает буквы большими

	case "/scream":
		screaming = true

	//Делает буквы маленькими
	case "/whisper":
		screaming = false

	// запускает меню
	case "/start":
		err = sendMenu(chatId)
	}

	return err
}

// функция которая отвечает за меню
func handleButton(query *tgbotapi.CallbackQuery) {
	var (
		text      string
		err       error
		inputppr  string
		inputppr2 string
	)
	//keyboard := tgbotapi.NewInlineKeyboardButtonData()
	markup := tgbotapi.NewInlineKeyboardMarkup()
	message := query.Message
	// переход по меню то что

	if query.Data == uuug {
		sheet = "uuug"
		fmt.Println(sheet)
		text = "<b>Поиск по объекту 930-20 цех 6. \nВведи позицию </b>"
		markup = firstMenuMarkup
	} else if query.Data == facel {
		sheet = "facel"
		fmt.Println(sheet)
		text = "<b>Поиск по объекту 930-09/10 цех 6. \nВведи позицию</b>"
		markup = firstMenuMarkup
	} else if query.Data == mckCeh6 {
		sheet = "ЦЕХ 6 (930-01)"
		fmt.Println(sheet)
		text = "<b>Поиск по объекту 930-01 цех 6. \nВведи позицию</b>"
		markup = firstMenuMarkup
	} else if query.Data == kondensatka {
		sheet = "88014"
		fmt.Println(sheet)
		text = "<b>Поиск по объекту 880-14. \nВведи позицию</b>"
		markup = firstMenuMarkup
	} else if query.Data == mckCeh14 {
		sheet = "ЦЕХ 14 (930-01)"
		fmt.Println(sheet)
		text = "<b>Поиск по объекту 930-01 цех 14. \nВведи позицию</b>"
		markup = firstMenuMarkup
	} else if query.Data == voadushka {
		sheet = "740-20"
		fmt.Println(sheet)
		text = "<b>Поиск по объекту 740-20. \nВведи позицию</b>"
		markup = firstMenuMarkup
	} else if query.Data == azotka {
		sheet = "730-30"
		fmt.Println(sheet)
		text = "<b>Поиск по объекту 730-30. \nВведи позицию</b>"
		markup = firstMenuMarkup
	} else if query.Data == vodoblok {
		sheet = "862-34"
		fmt.Println(sheet)
		text = "<b>Поиск по объекту 86-234. \nВведи позицию</b>"
		markup = firstMenuMarkup
	} else if query.Data == nine {
		sheet = "901-09"
		fmt.Println(sheet)
		text = "<b>Поиск по объекту 901-09. \nВведи позицию</b>"
		markup = firstMenuMarkup
	} else if query.Data == ten {
		sheet = "901-10"
		fmt.Println(sheet)
		text = "<b>Поиск по объекту 901-10. \nВведи позицию</b>"
		markup = firstMenuMarkup
	} else if query.Data == backButton {
		text = "Выбери объект"
		markup = firstMenuMarkup
	} else if query.Data == ppr {
		text = "<b> на дворе месяц " + time.Now().Month().String() + "</b>" + "\nВыбери неделю"
		markup = pprMenuMarkup
	} else if query.Data == oneWeek {
		markup = pprMenuMarkup
		inputppr, inputppr2 = getdataex.Getppr(3)
		fmt.Println(inputppr)
		text = oneWeek
	} else if query.Data == twoWeek {
		markup = pprMenuMarkup
		inputppr, inputppr2 = getdataex.Getppr(2)
		fmt.Println(inputppr)
		text = twoWeek
	} else if query.Data == threeWeek {
		markup = pprMenuMarkup
		inputppr, inputppr2 = getdataex.Getppr(1)
		fmt.Println(inputppr)
		text = threeWeek
	} else if query.Data == fourWeek {
		markup = pprMenuMarkup
		inputppr, inputppr2 = getdataex.Getppr(0)
		fmt.Println(inputppr)
		text = fourWeek

	}
	// сообщение из меню
	if inputppr != "" && inputppr2 != "" {

		pprmsg := tgbotapi.NewMessage(message.Chat.ID, inputppr)
		pprmsg1 := tgbotapi.NewMessage(message.Chat.ID, inputppr2)
		pprmsg.ParseMode = tgbotapi.ModeHTML
		pprmsg1.ParseMode = tgbotapi.ModeHTML
		_, err = bot.Send(pprmsg)
		_, err = bot.Send(pprmsg1)

		err = handleCommand(message.Chat.ID, "/start")
		if err != nil {
			log.Printf("An error occured: %s", err.Error())
		}
		inputppr = ""
		inputppr2 = ""
	}

	// Изменение меню
	callbackCfg := tgbotapi.NewCallback(query.ID, "")
	bot.Send(callbackCfg)

	// Replace menu text and keyboard
	msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, text, markup)
	msg.ParseMode = tgbotapi.ModeHTML
	bot.Send(msg)
}

// Отправка в меню
func sendMenu(chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, firstMenu)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = firstMenuMarkup
	_, err := bot.Send(msg)
	return err
}

func PPRMessage(message *tgbotapi.Message) {
	user := message.From
	text := message.Text

	if user == nil {
		return
	}

	//вывод в консоль
	// Print to console
	//
	log.Printf("%s wrote %s", user.FirstName, text)

	var err error

	msg := tgbotapi.NewMessage(message.Chat.ID, "ничего не найденно попробуй ввести последние цифры позиции. В чек листах позиции указанны через тире пример 73-PDIRSA-180A")
	// To preserve markdown, we attach entities (bold, italic..)
	//msg.Entities = message.Entities
	_, err = bot.Send(msg)

	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}

}
