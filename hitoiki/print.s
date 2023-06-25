TEXT     ·myPrint(SB),0,$0
    MOVQ a1+8(FP), DI
    MOVQ a2+16(FP), SI
    MOVQ a3+24(FP), DX
    MOVQ trap+0(FP), AX
    SYSCALL
    RET
