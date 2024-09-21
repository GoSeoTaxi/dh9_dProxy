package model

type NewProxy struct {
	Type  string // Тип прокси (например, socks5)
	Login string // Логин пользователя
	Pass  string // Пароль пользователя
	Addr  string // Адрес
	Port  string // Порт
}
