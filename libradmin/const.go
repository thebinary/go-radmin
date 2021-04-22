package libradmin

const FR_CHANNEL_MAGIC = 0xf7eead16
const (
	FR_CHANNEL_STDIN = iota
	FR_CHANNEL_STDOUT
	FR_CHANNEL_STDERR
	FR_CHANNEL_CMD_STATUS
	FR_CHANNEL_INIT_ACK
	FR_CHANNEL_AUTH_CHALLENGE
	FR_CHANNEL_AUTH_RESPONSE
	FR_CHANNEL_WANT_MORE
)
const FR_SEPRATOR = 0x9
