package config

var (
	Key = "ab9312a52781f4b7c7edf4341ef940daff94c567ffa503c3db8125fec68c4225"
)

type BodyStructure map[string]string

const (
	ERRORS_NULL = iota
)

const PROXY_PORT = "8088"

// database stuff

// users (id, email, password, phone, address)

// shops (id, user_id, address)

// computers (id, user_id, shop_id, computer_name, hashed_id, unlock_code, activation_date, expiration_date, duration)

// payments (id, user_id, shops, amount, gateway, date, status)

// blocklist (id, pattern)

// user_blocklist (id, pattern)

// shop_blocklist (id, pattern)
