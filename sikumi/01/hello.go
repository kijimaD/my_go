package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

// go build hello.go
// strace -o hello.log ./hello

// ================

// execve("./hello", ["./hello"], 0x7fffa04f7220 /* 91 vars */) = 0
// arch_prctl(ARCH_SET_FS, 0x52eb50)       = 0
// sched_getaffinity(0, 8192, [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15]) = 8
// openat(AT_FDCWD, "/sys/kernel/mm/transparent_hugepage/hpage_pmd_size", O_RDONLY) = 3
// read(3, "2097152\n", 20)                = 8
// close(3)                                = 0
// mmap(NULL, 262144, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7fed06737000
// mmap(NULL, 131072, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7fed06717000
// mmap(NULL, 1048576, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7fed06617000
// mmap(NULL, 8388608, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7fed05e17000
// mmap(NULL, 67108864, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7fed01e17000
// mmap(NULL, 536870912, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7fece1e17000
// mmap(NULL, 8388608, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7fece1617000
// mmap(0xc000000000, 67108864, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0xc000000000
// mmap(NULL, 33554432, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7fecdf617000
// mmap(NULL, 1133584, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7fecdf502000
// mmap(0xc000000000, 4194304, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0xc000000000
// mmap(0x7fed06717000, 131072, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7fed06717000
// mmap(0x7fed06697000, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7fed06697000
// mmap(0x7fed0621d000, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7fed0621d000
// mmap(0x7fed03e47000, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7fed03e47000
// mmap(0x7fecf1f97000, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7fecf1f97000
// mmap(0x7fece1617000, 4222976, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7fece1617000
// madvise(0x7fece1800000, 2097152, MADV_HUGEPAGE) = 0
// mmap(NULL, 1048576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7fecdf402000
// mmap(NULL, 65536, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7fecdf3f2000
// mmap(NULL, 65536, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7fecdf3e2000
// rt_sigprocmask(SIG_SETMASK, NULL, [], 8) = 0
// sigaltstack(NULL, {ss_sp=NULL, ss_flags=SS_DISABLE, ss_size=0}) = 0
// sigaltstack({ss_sp=0xc000008000, ss_flags=0, ss_size=32768}, NULL) = 0
// rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
// gettid()                                = 2799094
// rt_sigaction(SIGHUP, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGHUP, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGINT, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGINT, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGQUIT, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGQUIT, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGILL, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGILL, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGTRAP, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGTRAP, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGABRT, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGABRT, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGBUS, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGBUS, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGFPE, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGFPE, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGUSR1, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGUSR1, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGSEGV, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGSEGV, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGUSR2, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGUSR2, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGPIPE, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGPIPE, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGALRM, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGALRM, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGTERM, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGTERM, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGSTKFLT, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGSTKFLT, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGCHLD, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGCHLD, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGURG, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGURG, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGXCPU, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGXCPU, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGXFSZ, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGXFSZ, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGVTALRM, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGVTALRM, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGPROF, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGPROF, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGWINCH, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGWINCH, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGIO, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGIO, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGPWR, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGPWR, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGSYS, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGSYS, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRTMIN, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_1, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_1, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_2, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_3, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_3, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_4, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_4, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_5, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_5, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_6, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_6, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_7, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_7, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_8, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_8, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_9, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_9, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_10, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_10, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_11, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_11, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_12, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_12, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_13, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_13, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_14, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_14, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_15, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_15, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_16, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_16, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_17, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_17, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_18, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_18, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_19, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_19, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_20, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_20, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_21, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_21, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_22, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_22, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_23, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_23, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_24, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_24, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_25, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_25, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_26, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_26, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_27, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_27, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_28, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_28, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_29, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_29, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_30, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_30, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_31, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_31, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigaction(SIGRT_32, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
// rt_sigaction(SIGRT_32, {sa_handler=0x45f260, sa_mask=~[], sa_flags=SA_RESTORER|SA_ONSTACK|SA_RESTART|SA_SIGINFO, sa_restorer=0x45f3a0}, NULL, 8) = 0
// rt_sigprocmask(SIG_SETMASK, ~[], [], 8) = 0
// clone(child_stack=0xc000078000, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS, tls=0xc000068090) = 2799095
// rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
// rt_sigprocmask(SIG_SETMASK, ~[], [], 8) = 0
// clone(child_stack=0xc00007a000, flags=CLONE_VM|CLONE_FS|CLONE_FILES|CLONE_SIGHAND|CLONE_THREAD|CLONE_SYSVSEM|CLONE_SETTLS, tls=0xc000068490) = 2799096
// rt_sigprocmask(SIG_SETMASK, [], NULL, 8) = 0
// --- SIGURG {si_signo=SIGURG, si_code=SI_TKILL, si_pid=2799094, si_uid=1000} ---
// rt_sigreturn({mask=[]})                 = 0
// --- SIGURG {si_signo=SIGURG, si_code=SI_TKILL, si_pid=2799094, si_uid=1000} ---
// rt_sigreturn({mask=[]})                 = 0
// --- SIGURG {si_signo=SIGURG, si_code=SI_TKILL, si_pid=2799094, si_uid=1000} ---
// rt_sigreturn({mask=[]})                 = 0
// --- SIGURG {si_signo=SIGURG, si_code=SI_TKILL, si_pid=2799094, si_uid=1000} ---
// rt_sigreturn({mask=[]})                 = 0
// futex(0xc000088148, FUTEX_WAKE_PRIVATE, 1) = 1
// futex(0xc000088148, FUTEX_WAKE_PRIVATE, 1) = 1
// futex(0xc000068548, FUTEX_WAKE_PRIVATE, 1) = 1
// futex(0x52ec08, FUTEX_WAIT_PRIVATE, 0, NULL) = -1 EAGAIN (Resource temporarily unavailable)
// fcntl(0, F_GETFL)                       = 0x8002 (flags O_RDWR|O_LARGEFILE)
// futex(0xc000068548, FUTEX_WAKE_PRIVATE, 1) = 1
// fcntl(1, F_GETFL)                       = 0x8002 (flags O_RDWR|O_LARGEFILE)
// fcntl(2, F_GETFL)                       = 0x8002 (flags O_RDWR|O_LARGEFILE)
// getrlimit(RLIMIT_NOFILE, {rlim_cur=1024, rlim_max=1024*1024}) = 0
// setrlimit(RLIMIT_NOFILE, {rlim_cur=1024*1024, rlim_max=1024*1024}) = 0
// write(1, "hello world\n", 12hello world
// )           = 12
// exit_group(0)                           = ?
// +++ exited with 0 +++
