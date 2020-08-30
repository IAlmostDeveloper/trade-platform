package configs

import "os"

var SmtpClientHost = "smtp.yandex.ru"
var SmtpClientAddress = "smtp.yandex.ru:25"
var SmtpClientEmail = "xsollatradeplatform2@yandex.ru"
var SmtpClientPassword =  os.Getenv("TRADE_PLATFORM_SMTP_PASSWORD")

