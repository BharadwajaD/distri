package system

const PROCESS_COUNT = 4


type System struct{
    Processes [PROCESS_COUNT]Process
}

func NewSystems() System {

    var processes [PROCESS_COUNT]Process;

    for i := 0; i < PROCESS_COUNT; i++ {
        processes[i] = NewProcess(i);
    }

    return System{
        Processes: processes,
    }
}
