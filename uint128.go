package Uint128

type Uint128 struct {
	Hi uint64
	Lo uint64
}

func (N *Uint128) Increment() {
	var T uint64 = N.Lo + 1
	N.Hi -= ((T ^ N.Lo) & T) >> 63
	N.Lo = T
}

func (N *Uint128) Decrement() {
	var T uint64 = N.Lo - 1
	N.Hi -= ((T ^ N.Lo) & T) >> 63
	N.Lo = T
}

func AddUint128(N *Uint128, M *Uint128) Uint128 {
	var ans = Uint128{
		Hi: 0,
		Lo: 0,
	}
	var C uint64 = (((N.Lo & M.Lo) & 1) + (N.Lo >> 1) + (M.Lo >> 1)) >> 63
	ans.Hi = N.Hi + M.Hi + C
	ans.Lo = N.Lo + M.Lo
	return ans
}

func SubtractUint128(N *Uint128, M *Uint128) Uint128 {
	var ans = Uint128{
		Hi: 0,
		Lo: 0,
	}
	ans.Lo = N.Lo - M.Lo
	var C uint64 = (((ans.Lo & M.Lo) & 1) + (M.Lo >> 1) + (ans.Lo >> 1)) >> 63
	ans.Hi = N.Hi - (M.Hi + C)
	return ans
}
