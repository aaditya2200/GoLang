package temp

type in_addr struct {
	s_addr uint32
}

type sockaddr_in struct {
	s_len uint8
	sin_port uint16
	sin_zero string
	sin_addr in_addr
	sin_family uint16
}

