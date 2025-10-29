package day23

func Solve(file string, part int) int {
	program := ParseInput(file)
	network := InitNetwork(program, 50)

	for {
		for _, computer := range network.computers {
			computer.MergeIncoming()
			computer.ProcessPacket()
			is255, packet := computer.SendPackets(network)

			if is255 {
				if part == 1 {
					return packet[1]
				}

				if part == 2 {
					network.NAT = packet
				}

			}
		}

		if network.IsIdle() {
			panic("IDLE")
			if exit, y := network.Reset(); exit {
				return y
			}
		}
	}

	return 0
}
