package reversebits

func reverseBits(num uint32) uint32 {
	var out uint32
	for i := 0; i < 32; i++ {
		if num == 0 {
			out <<= 32 - i
			break
		}
		out = (out << 1) | (num & 1)
		num >>= 1
	}
	return out
}
