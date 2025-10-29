package day23

func Solve(file string) int {
	program := ParseInput(file)
	network := InitNetwork(program, 50)

	for {
		for id := range network {
			computer := network[id]
			computer.ProcessPacket()
			exit, packet := computer.SendPackets(network)

			if exit {
				return packet[1]
			}
		}
	}

	return 0
}
