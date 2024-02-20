package system

type System struct {
	Processes []Process
}

func NewSystems(pcount int) System {

	processes := make([]Process, pcount)

	for i := 0; i < pcount; i++ {
		processes[i] = NewProcess(i, pcount)
	}

	return System{
		Processes: processes,
	}
}
