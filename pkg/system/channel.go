package system

type Message interface{}

type Channel struct {
	comm chan Message
}

func NewChannel() Channel {
	return Channel{
		comm: make(chan Message),
	}
}

func (c *Channel) send(msg *Message) {
	c.comm <- msg
}

func (c *Channel) recv() Message {
	return <-c.comm
}
