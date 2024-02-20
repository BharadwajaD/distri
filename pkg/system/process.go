package system

import "fmt"


type Process struct {
	id       int
	channels []Channel
}

func NewProcess(id, pcount int) Process {

	channels := make([]Channel, pcount)

	for i := 0; i < pcount; i++ {
		if id != i {
			channels[i] = NewChannel()
		}
	}
	return Process{
		id: id,
	}
}

func (p *Process) Exec(fn func()) {

    fmt.Printf("From pid: %d\n", p.id);
	fn()

}

func (p *Process) Send(to int, msg *Message) {
	p.channels[to].send(msg)
}

func (p *Process) Recv(from int) Message {
	return p.channels[from].recv()
}
