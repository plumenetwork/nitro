package precompiles

// ArbHi provides a friendly greeting to anyone who calls it.
type ArbHi struct {
	Address   addr // 0x11a

	// Custom events
	Hi        func(ctx, mech, addr) error
	HiGasCost func(addr) (uint64, error)
}

func (con *ArbHi) SayHi(c ctx, evm mech) (string, error) {
	err := con.Hi(c, evm, c.caller)
	return "hi", err
}
