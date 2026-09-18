package messageSender

import (
	_ "github.com/icelee123/komari125/utils/messageSender/bark"
	_ "github.com/icelee123/komari125/utils/messageSender/email"
	_ "github.com/icelee123/komari125/utils/messageSender/empty"
	_ "github.com/icelee123/komari125/utils/messageSender/javascript"
	_ "github.com/icelee123/komari125/utils/messageSender/serverchan3"
	_ "github.com/icelee123/komari125/utils/messageSender/serverchanturbo"
	_ "github.com/icelee123/komari125/utils/messageSender/telegram"
	_ "github.com/icelee123/komari125/utils/messageSender/webhook"
)

func All() {
}
