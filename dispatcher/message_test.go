package dispatcher

type TestMessage struct {
	id   int
	data string
}

func (t *TestMessage) Route() (open bool, key string) {
	if t.id > 10 {
		return true, "{\"id\":123,\"name\": \"Ailey\"}"
	}
	return false, "{\"id\":123,\"name\": \"Terry\"}"
}
