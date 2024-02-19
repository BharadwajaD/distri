package system

type Message interface {}

type Process struct {
    id int
    channels [PROCESS_COUNT]chan Message
}

func NewProcess(id int) Process{

    var channels [PROCESS_COUNT]chan Message

    for i := 0; i < PROCESS_COUNT; i++ {
        channels[i] = make(chan Message)
    }

    return Process{
        id: id,
        channels: channels,
    }
}

func (p* Process) send(id int, msg *Message) {
    p.channels[id] <- msg
}

func (p* Process) recv(id int) Message {
    return  <-p.channels[id]
}
