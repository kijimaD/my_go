https://qiita.com/Akatsuki_py/items/8041fba499d54d59e0dd

```
go build -work -a -p 1 -x main.go
```

```
WORK=/tmp/go-build3679823934
mkdir -p $WORK/b004/
cat >/tmp/go-build3679823934/b004/importcfg << 'EOF' # internal
# import config
EOF
cd /home/silver/Project/my_go/build
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b004/_pkg_.a -trimpath "$WORK/b004=>" -p internal/goarch -std -+ -complete -buildid KbGxOhsYdL57je7iUeBM/KbGxOhsYdL57je7iUeBM -goversion go1.20.1 -c=12 -nolocalimports -importcfg $WORK/b004/importcfg -pack /usr/local/go/src/internal/goarch/goarch.go /usr/local/go/src/internal/goarch/goarch_amd64.go /usr/local/go/src/internal/goarch/zgoarch_amd64.go
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b004/_pkg_.a # internal
cp $WORK/b004/_pkg_.a /home/silver/.cache/go-build/98/98139ed2e1f6bcea287bcda02735e8168014f8ebd5644b652217998b47389f4f-d # internal
mkdir -p $WORK/b003/
cat >/tmp/go-build3679823934/b003/go_asm.h << 'EOF' # internal
EOF
cd /usr/local/go/src/internal/abi
/usr/local/go/pkg/tool/linux_amd64/asm -p internal/abi -trimpath "$WORK/b003=>" -I $WORK/b003/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -gensymabis -o $WORK/b003/symabis ./abi_test.s
cat >/tmp/go-build3679823934/b003/importcfg << 'EOF' # internal
# import config
packagefile internal/goarch=/tmp/go-build3679823934/b004/_pkg_.a
EOF
cd /home/silver/Project/my_go/build
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b003/_pkg_.a -trimpath "$WORK/b003=>" -p internal/abi -std -+ -buildid zmTnvzpSdT2x1KJhh_88/zmTnvzpSdT2x1KJhh_88 -goversion go1.20.1 -symabis $WORK/b003/symabis -c=12 -nolocalimports -importcfg $WORK/b003/importcfg -pack -asmhdr $WORK/b003/go_asm.h /usr/local/go/src/internal/abi/abi.go /usr/local/go/src/internal/abi/abi_amd64.go
cd /usr/local/go/src/internal/abi
/usr/local/go/pkg/tool/linux_amd64/asm -p internal/abi -trimpath "$WORK/b003=>" -I $WORK/b003/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b003/abi_test.o ./abi_test.s
/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b003/_pkg_.a $WORK/b003/abi_test.o # internal
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b003/_pkg_.a # internal
cp $WORK/b003/_pkg_.a /home/silver/.cache/go-build/8e/8ec4968fc2f6138293acbc01913ec2aac072b016dfce5d21e2e0335321fe578d-d # internal
mkdir -p $WORK/b007/
cat >/tmp/go-build3679823934/b007/go_asm.h << 'EOF' # internal
EOF
cd /usr/local/go/src/internal/cpu
/usr/local/go/pkg/tool/linux_amd64/asm -p internal/cpu -trimpath "$WORK/b007=>" -I $WORK/b007/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -gensymabis -o $WORK/b007/symabis ./cpu.s ./cpu_x86.s
cat >/tmp/go-build3679823934/b007/importcfg << 'EOF' # internal
# import config
EOF
cd /home/silver/Project/my_go/build
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b007/_pkg_.a -trimpath "$WORK/b007=>" -p internal/cpu -std -+ -buildid hBUOX2BEel8L8T3qqBLS/hBUOX2BEel8L8T3qqBLS -goversion go1.20.1 -symabis $WORK/b007/symabis -c=12 -nolocalimports -importcfg $WORK/b007/importcfg -pack -asmhdr $WORK/b007/go_asm.h /usr/local/go/src/internal/cpu/cpu.go /usr/local/go/src/internal/cpu/cpu_x86.go
cd /usr/local/go/src/internal/cpu
/usr/local/go/pkg/tool/linux_amd64/asm -p internal/cpu -trimpath "$WORK/b007=>" -I $WORK/b007/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b007/cpu.o ./cpu.s
/usr/local/go/pkg/tool/linux_amd64/asm -p internal/cpu -trimpath "$WORK/b007=>" -I $WORK/b007/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b007/cpu_x86.o ./cpu_x86.s
/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b007/_pkg_.a $WORK/b007/cpu.o $WORK/b007/cpu_x86.o # internal
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b007/_pkg_.a # internal
cp $WORK/b007/_pkg_.a /home/silver/.cache/go-build/8c/8c6dce5c0d4aaa979a98f0f59bc29c923d0dde2d28238b8170141278cf55a420-d # internal
mkdir -p $WORK/b006/
cat >/tmp/go-build3679823934/b006/go_asm.h << 'EOF' # internal
EOF
cd /usr/local/go/src/internal/bytealg
/usr/local/go/pkg/tool/linux_amd64/asm -p internal/bytealg -trimpath "$WORK/b006=>" -I $WORK/b006/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -gensymabis -o $WORK/b006/symabis ./compare_amd64.s ./count_amd64.s ./equal_amd64.s ./index_amd64.s ./indexbyte_amd64.s
cat >/tmp/go-build3679823934/b006/importcfg << 'EOF' # internal
# import config
packagefile internal/cpu=/tmp/go-build3679823934/b007/_pkg_.a
EOF
cd /home/silver/Project/my_go/build
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b006/_pkg_.a -trimpath "$WORK/b006=>" -p internal/bytealg -std -+ -buildid dJ1BzWuR9zwp6uFzLgw0/dJ1BzWuR9zwp6uFzLgw0 -goversion go1.20.1 -symabis $WORK/b006/symabis -c=12 -nolocalimports -importcfg $WORK/b006/importcfg -pack -asmhdr $WORK/b006/go_asm.h /usr/local/go/src/internal/bytealg/bytealg.go /usr/local/go/src/internal/bytealg/compare_native.go /usr/local/go/src/internal/bytealg/count_native.go /usr/local/go/src/internal/bytealg/equal_generic.go /usr/local/go/src/internal/bytealg/equal_native.go /usr/local/go/src/internal/bytealg/index_amd64.go /usr/local/go/src/internal/bytealg/index_native.go /usr/local/go/src/internal/bytealg/indexbyte_native.go
cd /usr/local/go/src/internal/bytealg
/usr/local/go/pkg/tool/linux_amd64/asm -p internal/bytealg -trimpath "$WORK/b006=>" -I $WORK/b006/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b006/compare_amd64.o ./compare_amd64.s
/usr/local/go/pkg/tool/linux_amd64/asm -p internal/bytealg -trimpath "$WORK/b006=>" -I $WORK/b006/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b006/count_amd64.o ./count_amd64.s
/usr/local/go/pkg/tool/linux_amd64/asm -p internal/bytealg -trimpath "$WORK/b006=>" -I $WORK/b006/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b006/equal_amd64.o ./equal_amd64.s
/usr/local/go/pkg/tool/linux_amd64/asm -p internal/bytealg -trimpath "$WORK/b006=>" -I $WORK/b006/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b006/index_amd64.o ./index_amd64.s
/usr/local/go/pkg/tool/linux_amd64/asm -p internal/bytealg -trimpath "$WORK/b006=>" -I $WORK/b006/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b006/indexbyte_amd64.o ./indexbyte_amd64.s
/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b006/_pkg_.a $WORK/b006/compare_amd64.o $WORK/b006/count_amd64.o $WORK/b006/equal_amd64.o $WORK/b006/index_amd64.o $WORK/b006/indexbyte_amd64.o # internal
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b006/_pkg_.a # internal
cp $WORK/b006/_pkg_.a /home/silver/.cache/go-build/c8/c8d52043c509f288e4b391dde6b152cc6b51b467463968c542fc4b909dcce9ce-d # internal
mkdir -p $WORK/b008/
cat >/tmp/go-build3679823934/b008/importcfg << 'EOF' # internal
# import config
EOF
cd /home/silver/Project/my_go/build
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b008/_pkg_.a -trimpath "$WORK/b008=>" -p internal/coverage/rtcov -std -+ -complete -buildid ogwZcPSqw6_ID4BJ32su/ogwZcPSqw6_ID4BJ32su -goversion go1.20.1 -c=12 -nolocalimports -importcfg $WORK/b008/importcfg -pack /usr/local/go/src/internal/coverage/rtcov/rtcov.go
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b008/_pkg_.a # internal
cp $WORK/b008/_pkg_.a /home/silver/.cache/go-build/4a/4aff17cb156129b8be1ad5a6261b6eed4f8164ba5dfdaccd459213dc4d0945cb-d # internal
mkdir -p $WORK/b009/
cat >/tmp/go-build3679823934/b009/importcfg << 'EOF' # internal
# import config
EOF
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b009/_pkg_.a -trimpath "$WORK/b009=>" -p internal/goexperiment -std -complete -buildid dLHhzpJZAPHIVXi7uKfP/dLHhzpJZAPHIVXi7uKfP -goversion go1.20.1 -c=12 -nolocalimports -importcfg $WORK/b009/importcfg -pack /usr/local/go/src/internal/goexperiment/exp_arenas_off.go /usr/local/go/src/internal/goexperiment/exp_boringcrypto_off.go /usr/local/go/src/internal/goexperiment/exp_coverageredesign_on.go /usr/local/go/src/internal/goexperiment/exp_fieldtrack_off.go /usr/local/go/src/internal/goexperiment/exp_heapminimum512kib_off.go /usr/local/go/src/internal/goexperiment/exp_pagetrace_off.go /usr/local/go/src/internal/goexperiment/exp_preemptibleloops_off.go /usr/local/go/src/internal/goexperiment/exp_regabiargs_on.go /usr/local/go/src/internal/goexperiment/exp_regabiwrappers_on.go /usr/local/go/src/internal/goexperiment/exp_staticlockranking_off.go /usr/local/go/src/internal/goexperiment/exp_unified_on.go /usr/local/go/src/internal/goexperiment/flags.go
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b009/_pkg_.a # internal
cp $WORK/b009/_pkg_.a /home/silver/.cache/go-build/8f/8fcad58e667cc287c2d2dd3a8242ee1b11a7dc6c87f264a5c1c6887327799bc7-d # internal
mkdir -p $WORK/b010/
cat >/tmp/go-build3679823934/b010/importcfg << 'EOF' # internal
# import config
EOF
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b010/_pkg_.a -trimpath "$WORK/b010=>" -p internal/goos -std -+ -complete -buildid nzP2kI4sAeAyR1Mx9Eop/nzP2kI4sAeAyR1Mx9Eop -goversion go1.20.1 -c=12 -nolocalimports -importcfg $WORK/b010/importcfg -pack /usr/local/go/src/internal/goos/goos.go /usr/local/go/src/internal/goos/unix.go /usr/local/go/src/internal/goos/zgoos_linux.go
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b010/_pkg_.a # internal
cp $WORK/b010/_pkg_.a /home/silver/.cache/go-build/88/8896592ebc9bac9bdf84f1d745e10c35629fe48b5ede229aaa3730d192d00e71-d # internal
mkdir -p $WORK/b011/
cat >/tmp/go-build3679823934/b011/go_asm.h << 'EOF' # internal
EOF
cd /usr/local/go/src/runtime/internal/atomic
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime/internal/atomic -trimpath "$WORK/b011=>" -I $WORK/b011/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -gensymabis -o $WORK/b011/symabis ./atomic_amd64.s
cat >/tmp/go-build3679823934/b011/importcfg << 'EOF' # internal
# import config
EOF
cd /home/silver/Project/my_go/build
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b011/_pkg_.a -trimpath "$WORK/b011=>" -p runtime/internal/atomic -std -+ -buildid rtWiyacP3g_amYBSwg-k/rtWiyacP3g_amYBSwg-k -goversion go1.20.1 -symabis $WORK/b011/symabis -c=12 -nolocalimports -importcfg $WORK/b011/importcfg -pack -asmhdr $WORK/b011/go_asm.h /usr/local/go/src/runtime/internal/atomic/atomic_amd64.go /usr/local/go/src/runtime/internal/atomic/doc.go /usr/local/go/src/runtime/internal/atomic/stubs.go /usr/local/go/src/runtime/internal/atomic/types.go /usr/local/go/src/runtime/internal/atomic/types_64bit.go /usr/local/go/src/runtime/internal/atomic/unaligned.go
cd /usr/local/go/src/runtime/internal/atomic
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime/internal/atomic -trimpath "$WORK/b011=>" -I $WORK/b011/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b011/atomic_amd64.o ./atomic_amd64.s
/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b011/_pkg_.a $WORK/b011/atomic_amd64.o # internal
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b011/_pkg_.a # internal
cp $WORK/b011/_pkg_.a /home/silver/.cache/go-build/6c/6cffd83bf996524770297179443e2f354e6ebaf35ae6cb7676380666f90cfd50-d # internal
mkdir -p $WORK/b012/
cat >/tmp/go-build3679823934/b012/importcfg << 'EOF' # internal
# import config
packagefile internal/goarch=/tmp/go-build3679823934/b004/_pkg_.a
EOF
cd /home/silver/Project/my_go/build
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b012/_pkg_.a -trimpath "$WORK/b012=>" -p runtime/internal/math -std -+ -complete -buildid hLPNcq_T1hSo11JqagDZ/hLPNcq_T1hSo11JqagDZ -goversion go1.20.1 -c=12 -nolocalimports -importcfg $WORK/b012/importcfg -pack /usr/local/go/src/runtime/internal/math/math.go
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b012/_pkg_.a # internal
cp $WORK/b012/_pkg_.a /home/silver/.cache/go-build/29/2972918356ed563072b08044c6f6886edbc9498ce1a68990112131af2d26d61f-d # internal
mkdir -p $WORK/b013/
cat >/tmp/go-build3679823934/b013/importcfg << 'EOF' # internal
# import config
packagefile internal/goarch=/tmp/go-build3679823934/b004/_pkg_.a
packagefile internal/goos=/tmp/go-build3679823934/b010/_pkg_.a
EOF
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b013/_pkg_.a -trimpath "$WORK/b013=>" -p runtime/internal/sys -std -+ -complete -buildid PuEU1FadyrYy_YAyjLMp/PuEU1FadyrYy_YAyjLMp -goversion go1.20.1 -c=12 -nolocalimports -importcfg $WORK/b013/importcfg -pack /usr/local/go/src/runtime/internal/sys/consts.go /usr/local/go/src/runtime/internal/sys/consts_norace.go /usr/local/go/src/runtime/internal/sys/intrinsics.go /usr/local/go/src/runtime/internal/sys/intrinsics_common.go /usr/local/go/src/runtime/internal/sys/nih.go /usr/local/go/src/runtime/internal/sys/sys.go /usr/local/go/src/runtime/internal/sys/zversion.go
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b013/_pkg_.a # internal
cp $WORK/b013/_pkg_.a /home/silver/.cache/go-build/2c/2c4c5958e67297b81dc412ac01e9004b24a97319a009ae588c3f776f9aeae230-d # internal
mkdir -p $WORK/b014/
cat >/tmp/go-build3679823934/b014/go_asm.h << 'EOF' # internal
EOF
cd /usr/local/go/src/runtime/internal/syscall
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime/internal/syscall -trimpath "$WORK/b014=>" -I $WORK/b014/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -gensymabis -o $WORK/b014/symabis ./asm_linux_amd64.s
cat >/tmp/go-build3679823934/b014/importcfg << 'EOF' # internal
# import config
EOF
cd /home/silver/Project/my_go/build
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b014/_pkg_.a -trimpath "$WORK/b014=>" -p runtime/internal/syscall -std -+ -buildid 96mbWMY-iY8_cYYYZ32n/96mbWMY-iY8_cYYYZ32n -goversion go1.20.1 -symabis $WORK/b014/symabis -c=12 -nolocalimports -importcfg $WORK/b014/importcfg -pack -asmhdr $WORK/b014/go_asm.h /usr/local/go/src/runtime/internal/syscall/defs_linux.go /usr/local/go/src/runtime/internal/syscall/defs_linux_amd64.go /usr/local/go/src/runtime/internal/syscall/syscall_linux.go
cd /usr/local/go/src/runtime/internal/syscall
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime/internal/syscall -trimpath "$WORK/b014=>" -I $WORK/b014/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b014/asm_linux_amd64.o ./asm_linux_amd64.s
/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b014/_pkg_.a $WORK/b014/asm_linux_amd64.o # internal
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b014/_pkg_.a # internal
cp $WORK/b014/_pkg_.a /home/silver/.cache/go-build/ed/ed791b27aa25ed20850a4ea7f0ea9bc3a67ebaa14b37182d8b8b0520172d61fd-d # internal
mkdir -p $WORK/b002/
cat >/tmp/go-build3679823934/b002/go_asm.h << 'EOF' # internal
EOF
cd /usr/local/go/src/runtime
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b002=>" -I $WORK/b002/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -gensymabis -o $WORK/b002/symabis ./asm.s ./asm_amd64.s ./duff_amd64.s ./memclr_amd64.s ./memmove_amd64.s ./preempt_amd64.s ./rt0_linux_amd64.s ./sys_linux_amd64.s ./time_linux_amd64.s
cat >/tmp/go-build3679823934/b002/importcfg << 'EOF' # internal
# import config
packagefile internal/abi=/tmp/go-build3679823934/b003/_pkg_.a
packagefile internal/bytealg=/tmp/go-build3679823934/b006/_pkg_.a
packagefile internal/coverage/rtcov=/tmp/go-build3679823934/b008/_pkg_.a
packagefile internal/cpu=/tmp/go-build3679823934/b007/_pkg_.a
packagefile internal/goarch=/tmp/go-build3679823934/b004/_pkg_.a
packagefile internal/goexperiment=/tmp/go-build3679823934/b009/_pkg_.a
packagefile internal/goos=/tmp/go-build3679823934/b010/_pkg_.a
packagefile runtime/internal/atomic=/tmp/go-build3679823934/b011/_pkg_.a
packagefile runtime/internal/math=/tmp/go-build3679823934/b012/_pkg_.a
packagefile runtime/internal/sys=/tmp/go-build3679823934/b013/_pkg_.a
packagefile runtime/internal/syscall=/tmp/go-build3679823934/b014/_pkg_.a
EOF
cd /home/silver/Project/my_go/build
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b002/_pkg_.a -trimpath "$WORK/b002=>" -p runtime -std -+ -buildid dpLjwSjPgAAnEG0krm0v/dpLjwSjPgAAnEG0krm0v -goversion go1.20.1 -symabis $WORK/b002/symabis -c=12 -nolocalimports -importcfg $WORK/b002/importcfg -pack -asmhdr $WORK/b002/go_asm.h /usr/local/go/src/runtime/alg.go /usr/local/go/src/runtime/arena.go /usr/local/go/src/runtime/asan0.go /usr/local/go/src/runtime/atomic_pointer.go /usr/local/go/src/runtime/cgo.go /usr/local/go/src/runtime/cgo_mmap.go /usr/local/go/src/runtime/cgo_sigaction.go /usr/local/go/src/runtime/cgocall.go /usr/local/go/src/runtime/cgocallback.go /usr/local/go/src/runtime/cgocheck.go /usr/local/go/src/runtime/chan.go /usr/local/go/src/runtime/checkptr.go /usr/local/go/src/runtime/compiler.go /usr/local/go/src/runtime/complex.go /usr/local/go/src/runtime/covercounter.go /usr/local/go/src/runtime/covermeta.go /usr/local/go/src/runtime/cpuflags.go /usr/local/go/src/runtime/cpuflags_amd64.go /usr/local/go/src/runtime/cpuprof.go /usr/local/go/src/runtime/cputicks.go /usr/local/go/src/runtime/create_file_unix.go /usr/local/go/src/runtime/debug.go /usr/local/go/src/runtime/debugcall.go /usr/local/go/src/runtime/debuglog.go /usr/local/go/src/runtime/debuglog_off.go /usr/local/go/src/runtime/defs_linux_amd64.go /usr/local/go/src/runtime/env_posix.go /usr/local/go/src/runtime/error.go /usr/local/go/src/runtime/exithook.go /usr/local/go/src/runtime/extern.go /usr/local/go/src/runtime/fastlog2.go /usr/local/go/src/runtime/fastlog2table.go /usr/local/go/src/runtime/float.go /usr/local/go/src/runtime/hash64.go /usr/local/go/src/runtime/heapdump.go /usr/local/go/src/runtime/histogram.go /usr/local/go/src/runtime/iface.go /usr/local/go/src/runtime/lfstack.go /usr/local/go/src/runtime/lfstack_64bit.go /usr/local/go/src/runtime/lock_futex.go /usr/local/go/src/runtime/lockrank.go /usr/local/go/src/runtime/lockrank_off.go /usr/local/go/src/runtime/malloc.go /usr/local/go/src/runtime/map.go /usr/local/go/src/runtime/map_fast32.go /usr/local/go/src/runtime/map_fast64.go /usr/local/go/src/runtime/map_faststr.go /usr/local/go/src/runtime/mbarrier.go /usr/local/go/src/runtime/mbitmap.go /usr/local/go/src/runtime/mcache.go /usr/local/go/src/runtime/mcentral.go /usr/local/go/src/runtime/mcheckmark.go /usr/local/go/src/runtime/mem.go /usr/local/go/src/runtime/mem_linux.go /usr/local/go/src/runtime/metrics.go /usr/local/go/src/runtime/mfinal.go /usr/local/go/src/runtime/mfixalloc.go /usr/local/go/src/runtime/mgc.go /usr/local/go/src/runtime/mgclimit.go /usr/local/go/src/runtime/mgcmark.go /usr/local/go/src/runtime/mgcpacer.go /usr/local/go/src/runtime/mgcscavenge.go /usr/local/go/src/runtime/mgcstack.go /usr/local/go/src/runtime/mgcsweep.go /usr/local/go/src/runtime/mgcwork.go /usr/local/go/src/runtime/mheap.go /usr/local/go/src/runtime/mpagealloc.go /usr/local/go/src/runtime/mpagealloc_64bit.go /usr/local/go/src/runtime/mpagecache.go /usr/local/go/src/runtime/mpallocbits.go /usr/local/go/src/runtime/mprof.go /usr/local/go/src/runtime/mranges.go /usr/local/go/src/runtime/msan0.go /usr/local/go/src/runtime/msize.go /usr/local/go/src/runtime/mspanset.go /usr/local/go/src/runtime/mstats.go /usr/local/go/src/runtime/mwbbuf.go /usr/local/go/src/runtime/nbpipe_pipe2.go /usr/local/go/src/runtime/netpoll.go /usr/local/go/src/runtime/netpoll_epoll.go /usr/local/go/src/runtime/os_linux.go /usr/local/go/src/runtime/os_linux_generic.go /usr/local/go/src/runtime/os_linux_noauxv.go /usr/local/go/src/runtime/os_linux_x86.go /usr/local/go/src/runtime/os_nonopenbsd.go /usr/local/go/src/runtime/pagetrace_off.go /usr/local/go/src/runtime/panic.go /usr/local/go/src/runtime/plugin.go /usr/local/go/src/runtime/preempt.go /usr/local/go/src/runtime/preempt_nonwindows.go /usr/local/go/src/runtime/print.go /usr/local/go/src/runtime/proc.go /usr/local/go/src/runtime/profbuf.go /usr/local/go/src/runtime/proflabel.go /usr/local/go/src/runtime/race0.go /usr/local/go/src/runtime/rdebug.go /usr/local/go/src/runtime/relax_stub.go /usr/local/go/src/runtime/retry.go /usr/local/go/src/runtime/runtime.go /usr/local/go/src/runtime/runtime1.go /usr/local/go/src/runtime/runtime2.go /usr/local/go/src/runtime/runtime_boring.go /usr/local/go/src/runtime/rwmutex.go /usr/local/go/src/runtime/select.go /usr/local/go/src/runtime/sema.go /usr/local/go/src/runtime/signal_amd64.go /usr/local/go/src/runtime/signal_linux_amd64.go /usr/local/go/src/runtime/signal_unix.go /usr/local/go/src/runtime/sigqueue.go /usr/local/go/src/runtime/sigqueue_note.go /usr/local/go/src/runtime/sigtab_linux_generic.go /usr/local/go/src/runtime/sizeclasses.go /usr/local/go/src/runtime/slice.go /usr/local/go/src/runtime/softfloat64.go /usr/local/go/src/runtime/stack.go /usr/local/go/src/runtime/stkframe.go /usr/local/go/src/runtime/string.go /usr/local/go/src/runtime/stubs.go /usr/local/go/src/runtime/stubs2.go /usr/local/go/src/runtime/stubs3.go /usr/local/go/src/runtime/stubs_amd64.go /usr/local/go/src/runtime/stubs_linux.go /usr/local/go/src/runtime/symtab.go /usr/local/go/src/runtime/sys_nonppc64x.go /usr/local/go/src/runtime/sys_x86.go /usr/local/go/src/runtime/time.go /usr/local/go/src/runtime/time_nofake.go /usr/local/go/src/runtime/timeasm.go /usr/local/go/src/runtime/tls_stub.go /usr/local/go/src/runtime/trace.go /usr/local/go/src/runtime/traceback.go /usr/local/go/src/runtime/type.go /usr/local/go/src/runtime/typekind.go /usr/local/go/src/runtime/unsafe.go /usr/local/go/src/runtime/utf8.go /usr/local/go/src/runtime/vdso_elf64.go /usr/local/go/src/runtime/vdso_linux.go /usr/local/go/src/runtime/vdso_linux_amd64.go /usr/local/go/src/runtime/write_err.go
cp /usr/local/go/src/runtime/asm_amd64.h $WORK/b002/asm_GOARCH.h
cd /usr/local/go/src/runtime
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b002=>" -I $WORK/b002/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b002/asm.o ./asm.s
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b002=>" -I $WORK/b002/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b002/asm_amd64.o ./asm_amd64.s
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b002=>" -I $WORK/b002/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b002/duff_amd64.o ./duff_amd64.s
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b002=>" -I $WORK/b002/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b002/memclr_amd64.o ./memclr_amd64.s
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b002=>" -I $WORK/b002/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b002/memmove_amd64.o ./memmove_amd64.s
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b002=>" -I $WORK/b002/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b002/preempt_amd64.o ./preempt_amd64.s
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b002=>" -I $WORK/b002/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b002/rt0_linux_amd64.o ./rt0_linux_amd64.s
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b002=>" -I $WORK/b002/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b002/sys_linux_amd64.o ./sys_linux_amd64.s
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b002=>" -I $WORK/b002/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -compiling-runtime -D GOAMD64_v1 -o $WORK/b002/time_linux_amd64.o ./time_linux_amd64.s
/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b002/_pkg_.a $WORK/b002/asm.o $WORK/b002/asm_amd64.o $WORK/b002/duff_amd64.o $WORK/b002/memclr_amd64.o $WORK/b002/memmove_amd64.o $WORK/b002/preempt_amd64.o $WORK/b002/rt0_linux_amd64.o $WORK/b002/sys_linux_amd64.o $WORK/b002/time_linux_amd64.o # internal
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b002/_pkg_.a # internal
cp $WORK/b002/_pkg_.a /home/silver/.cache/go-build/aa/aa2fe26fb6d31e6b4a758a274c4f0743fac08b996fb5b8d1204d6b1572bcd8f5-d # internal
mkdir -p $WORK/b001/
cat >/tmp/go-build3679823934/b001/importcfg << 'EOF' # internal
# import config
packagefile runtime=/tmp/go-build3679823934/b002/_pkg_.a
EOF
cd /home/silver/Project/my_go/build
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b001/_pkg_.a -trimpath "$WORK/b001=>" -p main -complete -buildid CagZ0ozfLQsQK3rPCcHh/CagZ0ozfLQsQK3rPCcHh -goversion go1.20.1 -c=12 -nolocalimports -importcfg $WORK/b001/importcfg -pack ./main.go
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b001/_pkg_.a # internal
cp $WORK/b001/_pkg_.a /home/silver/.cache/go-build/1f/1f8f5ce96502267a5be7c78452b09f32e53d243993d30af0284049c5eafb3109-d # internal
cat >/tmp/go-build3679823934/b001/importcfg.link << 'EOF' # internal
packagefile command-line-arguments=/tmp/go-build3679823934/b001/_pkg_.a
packagefile runtime=/tmp/go-build3679823934/b002/_pkg_.a
packagefile internal/abi=/tmp/go-build3679823934/b003/_pkg_.a
packagefile internal/bytealg=/tmp/go-build3679823934/b006/_pkg_.a
packagefile internal/coverage/rtcov=/tmp/go-build3679823934/b008/_pkg_.a
packagefile internal/cpu=/tmp/go-build3679823934/b007/_pkg_.a
packagefile internal/goarch=/tmp/go-build3679823934/b004/_pkg_.a
packagefile internal/goexperiment=/tmp/go-build3679823934/b009/_pkg_.a
packagefile internal/goos=/tmp/go-build3679823934/b010/_pkg_.a
packagefile runtime/internal/atomic=/tmp/go-build3679823934/b011/_pkg_.a
packagefile runtime/internal/math=/tmp/go-build3679823934/b012/_pkg_.a
packagefile runtime/internal/sys=/tmp/go-build3679823934/b013/_pkg_.a
packagefile runtime/internal/syscall=/tmp/go-build3679823934/b014/_pkg_.a
modinfo "0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nbuild\t-buildmode=exe\nbuild\t-compiler=gc\nbuild\tCGO_ENABLED=1\nbuild\tCGO_CFLAGS=\nbuild\tCGO_CPPFLAGS=\nbuild\tCGO_CXXFLAGS=\nbuild\tCGO_LDFLAGS=\nbuild\tGOARCH=amd64\nbuild\tGOOS=linux\nbuild\tGOAMD64=v1\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"
EOF
mkdir -p $WORK/b001/exe/
cd .
/usr/local/go/pkg/tool/linux_amd64/link -o $WORK/b001/exe/a.out -importcfg $WORK/b001/importcfg.link -buildmode=exe -buildid=BEgVX_wAkfDDbvbYov3s/CagZ0ozfLQsQK3rPCcHh/4d1loMPwJkrjWU4Ryh2R/BEgVX_wAkfDDbvbYov3s -extld=gcc $WORK/b001/_pkg_.a
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b001/exe/a.out # internal
cp $WORK/b001/exe/a.out main
```
