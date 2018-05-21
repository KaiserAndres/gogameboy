package gogameboy

type Registers struct {
	flags, flagsP, A, B, C, D, E, F, Ap, Bp, Cp, Dp, Ep, Fp, H, L, Hp, Lp byte
	IX, IY, SP, PC, Interrupt, I uint16
}
