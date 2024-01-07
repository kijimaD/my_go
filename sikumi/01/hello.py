#!/usr/bin/python3
print("hello world")

# strace ./hello.py

# ================

# execve("./hello.py", ["./hello.py"], 0x7ffdd04a2fd0 /* 91 vars */) = 0
# brk(NULL)                               = 0x557705459000
# arch_prctl(0x3001 /* ARCH_??? */, 0x7fff6aecf150) = -1 EINVAL (Invalid argument)
# mmap(NULL, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0f345a7000
# access("/etc/ld.so.preload", R_OK)      = -1 ENOENT (No such file or directory)
# openat(AT_FDCWD, "/etc/ld.so.cache", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=85349, ...}, AT_EMPTY_PATH) = 0
# mmap(NULL, 85349, PROT_READ, MAP_PRIVATE, 3, 0) = 0x7f0f34592000
# close(3)                                = 0
# openat(AT_FDCWD, "/lib/x86_64-linux-gnu/libm.so.6", O_RDONLY|O_CLOEXEC) = 3
# read(3, "\177ELF\2\1\1\3\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0\0\0\0\0\0\0\0\0"..., 832) = 832
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=940560, ...}, AT_EMPTY_PATH) = 0
# mmap(NULL, 942344, PROT_READ, MAP_PRIVATE|MAP_DENYWRITE, 3, 0) = 0x7f0f344ab000
# mmap(0x7f0f344b9000, 507904, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0xe000) = 0x7f0f344b9000
# mmap(0x7f0f34535000, 372736, PROT_READ, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x8a000) = 0x7f0f34535000
# mmap(0x7f0f34590000, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0xe4000) = 0x7f0f34590000
# close(3)                                = 0
# openat(AT_FDCWD, "/lib/x86_64-linux-gnu/libexpat.so.1", O_RDONLY|O_CLOEXEC) = 3
# read(3, "\177ELF\2\1\1\0\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0\0\0\0\0\0\0\0\0"..., 832) = 832
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=194872, ...}, AT_EMPTY_PATH) = 0
# mmap(NULL, 196792, PROT_READ, MAP_PRIVATE|MAP_DENYWRITE, 3, 0) = 0x7f0f3447a000
# mprotect(0x7f0f3447e000, 172032, PROT_NONE) = 0
# mmap(0x7f0f3447e000, 126976, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x4000) = 0x7f0f3447e000
# mmap(0x7f0f3449d000, 40960, PROT_READ, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x23000) = 0x7f0f3449d000
# mmap(0x7f0f344a8000, 12288, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x2d000) = 0x7f0f344a8000
# close(3)                                = 0
# openat(AT_FDCWD, "/lib/x86_64-linux-gnu/libz.so.1", O_RDONLY|O_CLOEXEC) = 3
# read(3, "\177ELF\2\1\1\3\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0\0\0\0\0\0\0\0\0"..., 832) = 832
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=108936, ...}, AT_EMPTY_PATH) = 0
# mmap(NULL, 110776, PROT_READ, MAP_PRIVATE|MAP_DENYWRITE, 3, 0) = 0x7f0f3445e000
# mprotect(0x7f0f34460000, 98304, PROT_NONE) = 0
# mmap(0x7f0f34460000, 69632, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x2000) = 0x7f0f34460000
# mmap(0x7f0f34471000, 24576, PROT_READ, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x13000) = 0x7f0f34471000
# mmap(0x7f0f34478000, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x19000) = 0x7f0f34478000
# close(3)                                = 0
# openat(AT_FDCWD, "/lib/x86_64-linux-gnu/libc.so.6", O_RDONLY|O_CLOEXEC) = 3
# read(3, "\177ELF\2\1\1\3\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0P\237\2\0\0\0\0\0"..., 832) = 832
# pread64(3, "\6\0\0\0\4\0\0\0@\0\0\0\0\0\0\0@\0\0\0\0\0\0\0@\0\0\0\0\0\0\0"..., 784, 64) = 784
# pread64(3, "\4\0\0\0 \0\0\0\5\0\0\0GNU\0\2\0\0\300\4\0\0\0\3\0\0\0\0\0\0\0"..., 48, 848) = 48
# pread64(3, "\4\0\0\0\24\0\0\0\3\0\0\0GNU\0i8\235HZ\227\223\333\350s\360\352,\223\340."..., 68, 896) = 68
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=2216304, ...}, AT_EMPTY_PATH) = 0
# pread64(3, "\6\0\0\0\4\0\0\0@\0\0\0\0\0\0\0@\0\0\0\0\0\0\0@\0\0\0\0\0\0\0"..., 784, 64) = 784
# mmap(NULL, 2260560, PROT_READ, MAP_PRIVATE|MAP_DENYWRITE, 3, 0) = 0x7f0f34236000
# mmap(0x7f0f3425e000, 1658880, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x28000) = 0x7f0f3425e000
# mmap(0x7f0f343f3000, 360448, PROT_READ, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x1bd000) = 0x7f0f343f3000
# mmap(0x7f0f3444b000, 24576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x214000) = 0x7f0f3444b000
# mmap(0x7f0f34451000, 52816, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7f0f34451000
# close(3)                                = 0
# mmap(NULL, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0f34234000
# arch_prctl(ARCH_SET_FS, 0x7f0f34235000) = 0
# set_tid_address(0x7f0f342352d0)         = 2799555
# set_robust_list(0x7f0f342352e0, 24)     = 0
# rseq(0x7f0f342359a0, 0x20, 0, 0x53053053) = 0
# mprotect(0x7f0f3444b000, 16384, PROT_READ) = 0
# mprotect(0x7f0f34478000, 4096, PROT_READ) = 0
# mprotect(0x7f0f344a8000, 8192, PROT_READ) = 0
# mprotect(0x7f0f34590000, 4096, PROT_READ) = 0
# mprotect(0x557703aea000, 28672, PROT_READ) = 0
# mprotect(0x7f0f345e1000, 8192, PROT_READ) = 0
# prlimit64(0, RLIMIT_STACK, NULL, {rlim_cur=9788*1024, rlim_max=RLIM64_INFINITY}) = 0
# munmap(0x7f0f34592000, 85349)           = 0
# getrandom("\x13\xb2\xd8\x5f\xc4\xf1\x4c\x96", 8, GRND_NONBLOCK) = 8
# brk(NULL)                               = 0x557705459000
# brk(0x55770547a000)                     = 0x55770547a000
# openat(AT_FDCWD, "/usr/lib/locale/locale-archive", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=6784496, ...}, AT_EMPTY_PATH) = 0
# mmap(NULL, 6784496, PROT_READ, MAP_PRIVATE, 3, 0) = 0x7f0f33bbb000
# close(3)                                = 0
# openat(AT_FDCWD, "/usr/lib/x86_64-linux-gnu/gconv/gconv-modules.cache", O_RDONLY) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=27002, ...}, AT_EMPTY_PATH) = 0
# mmap(NULL, 27002, PROT_READ, MAP_SHARED, 3, 0) = 0x7f0f345a0000
# close(3)                                = 0
# futex(0x7f0f34450a6c, FUTEX_WAKE_PRIVATE, 2147483647) = 0
# getcwd("/home/orange/Project/my_go/sikumi/01", 4096) = 37
# getrandom("\x12\xf1\x71\xe8\xb8\x86\x44\x61\x90\xed\xef\xc7\x31\x13\x8b\x16\x2a\xbf\x76\x87\x2a\x51\x48\xae", 24, GRND_NONBLOCK) = 24
# mmap(NULL, 1048576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0f33abb000
# mmap(NULL, 266240, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0f33a7a000
# mmap(NULL, 135168, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0f33a59000
# brk(0x55770549b000)                     = 0x55770549b000
# brk(0x5577054be000)                     = 0x5577054be000
# readlink("/usr/bin/python3", "python3.10", 4096) = 10
# readlink("/usr/bin/python3.10", 0x7fff6aec9ab0, 4096) = -1 EINVAL (Invalid argument)
# openat(AT_FDCWD, "/usr/bin/pyvenv.cfg", O_RDONLY) = -1 ENOENT (No such file or directory)
# openat(AT_FDCWD, "/usr/pyvenv.cfg", O_RDONLY) = -1 ENOENT (No such file or directory)
# newfstatat(AT_FDCWD, "/usr/bin/Modules/Setup.local", 0x7fff6aecaa80, 0) = -1 ENOENT (No such file or directory)
# newfstatat(AT_FDCWD, "/usr/bin/lib/python3.10/os.py", 0x7fff6aeca980, 0) = -1 ENOENT (No such file or directory)
# newfstatat(AT_FDCWD, "/usr/bin/lib/python3.10/os.pyc", 0x7fff6aeca980, 0) = -1 ENOENT (No such file or directory)
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/os.py", {st_mode=S_IFREG|0644, st_size=39557, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/bin/pybuilddir.txt", O_RDONLY) = -1 ENOENT (No such file or directory)
# newfstatat(AT_FDCWD, "/usr/bin/lib/python3.10/lib-dynload", 0x7fff6aec9b00, 0) = -1 ENOENT (No such file or directory)
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/lib-dynload", {st_mode=S_IFDIR|0755, st_size=12288, ...}, 0) = 0
# mmap(NULL, 1048576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0f33959000
# sysinfo({uptime=3006421, loads=[165664, 148736, 158976], totalram=33358684160, freeram=2632384512, sharedram=1545232384, bufferram=2868572160, totalswap=2147479552, freeswap=1544212480, procs=1897, totalhigh=0, freehigh=0, mem_unit=1}) = 0
# openat(AT_FDCWD, "/etc/localtime", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=309, ...}, AT_EMPTY_PATH) = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=309, ...}, AT_EMPTY_PATH) = 0
# brk(0x5577054df000)                     = 0x5577054df000
# read(3, "TZif2\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\4\0\0\0\4\0\0\0\0"..., 4096) = 309
# lseek(3, -176, SEEK_CUR)                = 133
# read(3, "TZif2\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\4\0\0\0\4\0\0\0\0"..., 4096) = 176
# close(3)                                = 0
# brk(0x5577054de000)                     = 0x5577054de000
# newfstatat(AT_FDCWD, "/usr/lib/python310.zip", 0x7fff6aecd340, 0) = -1 ENOENT (No such file or directory)
# newfstatat(AT_FDCWD, "/usr/lib", {st_mode=S_IFDIR|0755, st_size=4096, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python310.zip", 0x7fff6aecd0b0, 0) = -1 ENOENT (No such file or directory)
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10", O_RDONLY|O_NONBLOCK|O_CLOEXEC|O_DIRECTORY) = 3
# newfstatat(3, "", {st_mode=S_IFDIR|0755, st_size=20480, ...}, AT_EMPTY_PATH) = 0
# getdents64(3, 0x5577054bec70 /* 206 entries */, 32768) = 6896
# getdents64(3, 0x5577054bec70 /* 0 entries */, 32768) = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/encodings/__init__.cpython-310-x86_64-linux-gnu.so", 0x7fff6aecd4d0, 0) = -1 ENOENT (No such file or directory)
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/encodings/__init__.abi3.so", 0x7fff6aecd4d0, 0) = -1 ENOENT (No such file or directory)
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/encodings/__init__.so", 0x7fff6aecd4d0, 0) = -1 ENOENT (No such file or directory)
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/encodings/__init__.py", {st_mode=S_IFREG|0644, st_size=5620, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/encodings/__init__.py", {st_mode=S_IFREG|0644, st_size=5620, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/encodings/__pycache__/__init__.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# fcntl(3, F_GETFD)                       = 0x1 (flags FD_CLOEXEC)
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=3875, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecd7c0)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=3875, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\4[\205d\364\25\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 3876) = 3875
# read(3, "", 1)                          = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/codecs.py", {st_mode=S_IFREG|0644, st_size=36714, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/codecs.py", {st_mode=S_IFREG|0644, st_size=36714, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/__pycache__/codecs.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=33219, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecc930)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=33219, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\4[\205dj\217\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 33220) = 33219
# read(3, "", 1)                          = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/encodings", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/encodings", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/encodings", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/encodings", O_RDONLY|O_NONBLOCK|O_CLOEXEC|O_DIRECTORY) = 3
# newfstatat(3, "", {st_mode=S_IFDIR|0755, st_size=20480, ...}, AT_EMPTY_PATH) = 0
# getdents64(3, 0x5577054cd430 /* 125 entries */, 32768) = 4224
# getdents64(3, 0x5577054cd430 /* 0 entries */, 32768) = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/encodings/aliases.py", {st_mode=S_IFREG|0644, st_size=15677, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/encodings/aliases.py", {st_mode=S_IFREG|0644, st_size=15677, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/encodings/__pycache__/aliases.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=10921, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecc250)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=10921, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\4[\205d==\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 10922) = 10921
# read(3, "", 1)                          = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/encodings", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/encodings/utf_8.py", {st_mode=S_IFREG|0644, st_size=1005, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/encodings/utf_8.py", {st_mode=S_IFREG|0644, st_size=1005, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/encodings/__pycache__/utf_8.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=1597, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecd7f0)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=1597, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\4[\205d\355\3\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 1598) = 1597
# read(3, "", 1)                          = 0
# close(3)                                = 0
# rt_sigaction(SIGPIPE, {sa_handler=SIG_IGN, sa_mask=[], sa_flags=SA_RESTORER|SA_ONSTACK, sa_restorer=0x7f0f34278520}, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGXFSZ, {sa_handler=SIG_IGN, sa_mask=[], sa_flags=SA_RESTORER|SA_ONSTACK, sa_restorer=0x7f0f34278520}, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGHUP, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGINT, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGQUIT, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGILL, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGTRAP, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGABRT, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGBUS, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGFPE, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGKILL, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGUSR1, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGSEGV, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGUSR2, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGPIPE, NULL, {sa_handler=SIG_IGN, sa_mask=[], sa_flags=SA_RESTORER|SA_ONSTACK, sa_restorer=0x7f0f34278520}, 8) = 0
# rt_sigaction(SIGALRM, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGTERM, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGSTKFLT, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGCHLD, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGCONT, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGSTOP, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGTSTP, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGTTIN, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGTTOU, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGURG, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGXCPU, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGXFSZ, NULL, {sa_handler=SIG_IGN, sa_mask=[], sa_flags=SA_RESTORER|SA_ONSTACK, sa_restorer=0x7f0f34278520}, 8) = 0
# rt_sigaction(SIGVTALRM, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGPROF, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGWINCH, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGIO, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGPWR, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGSYS, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_2, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_3, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_4, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_5, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_6, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_7, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_8, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_9, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_10, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_11, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_12, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_13, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_14, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_15, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_16, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_17, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_18, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_19, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_20, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_21, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_22, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_23, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_24, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_25, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_26, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_27, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_28, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_29, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_30, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_31, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGRT_32, NULL, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# rt_sigaction(SIGINT, {sa_handler=0x55770381a410, sa_mask=[], sa_flags=SA_RESTORER|SA_ONSTACK, sa_restorer=0x7f0f34278520}, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=0}, 8) = 0
# newfstatat(0, "", {st_mode=S_IFCHR|0620, st_rdev=makedev(0x88, 0xf), ...}, AT_EMPTY_PATH) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/io.py", {st_mode=S_IFREG|0644, st_size=4196, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/io.py", {st_mode=S_IFREG|0644, st_size=4196, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/__pycache__/io.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=3663, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecd8b0)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=3663, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\4[\205dd\20\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 3664) = 3663
# read(3, "", 1)                          = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/abc.py", {st_mode=S_IFREG|0644, st_size=6522, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/abc.py", {st_mode=S_IFREG|0644, st_size=6522, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/__pycache__/abc.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=6751, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecca20)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=6751, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\4[\205dz\31\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 6752) = 6751
# read(3, "", 1)                          = 0
# close(3)                                = 0
# dup(0)                                  = 3
# close(3)                                = 0
# newfstatat(0, "", {st_mode=S_IFCHR|0620, st_rdev=makedev(0x88, 0xf), ...}, AT_EMPTY_PATH) = 0
# ioctl(0, TCGETS, {B38400 opost isig icanon echo ...}) = 0
# lseek(0, 0, SEEK_CUR)                   = -1 ESPIPE (Illegal seek)
# ioctl(0, TCGETS, {B38400 opost isig icanon echo ...}) = 0
# dup(1)                                  = 3
# close(3)                                = 0
# newfstatat(1, "", {st_mode=S_IFCHR|0620, st_rdev=makedev(0x88, 0xf), ...}, AT_EMPTY_PATH) = 0
# ioctl(1, TCGETS, {B38400 opost isig icanon echo ...}) = 0
# lseek(1, 0, SEEK_CUR)                   = -1 ESPIPE (Illegal seek)
# ioctl(1, TCGETS, {B38400 opost isig icanon echo ...}) = 0
# dup(2)                                  = 3
# close(3)                                = 0
# newfstatat(2, "", {st_mode=S_IFCHR|0620, st_rdev=makedev(0x88, 0xf), ...}, AT_EMPTY_PATH) = 0
# ioctl(2, TCGETS, {B38400 opost isig icanon echo ...}) = 0
# lseek(2, 0, SEEK_CUR)                   = -1 ESPIPE (Illegal seek)
# ioctl(2, TCGETS, {B38400 opost isig icanon echo ...}) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/site.py", {st_mode=S_IFREG|0644, st_size=23667, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/site.py", {st_mode=S_IFREG|0644, st_size=23667, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/__pycache__/site.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=17922, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecd8b0)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=17922, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\4[\205ds\\\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 17923) = 17922
# read(3, "", 1)                          = 0
# close(3)                                = 0
# brk(0x5577054ff000)                     = 0x5577054ff000
# brk(0x5577054fe000)                     = 0x5577054fe000
# brk(0x5577054fa000)                     = 0x5577054fa000
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/os.py", {st_mode=S_IFREG|0644, st_size=39557, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/os.py", {st_mode=S_IFREG|0644, st_size=39557, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/__pycache__/os.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=31599, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecca20)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=31599, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\4[\205d\205\232\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 31600) = 31599
# read(3, "", 1)                          = 0
# close(3)                                = 0
# mmap(NULL, 151552, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f0f33934000
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/stat.py", {st_mode=S_IFREG|0644, st_size=5485, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/stat.py", {st_mode=S_IFREG|0644, st_size=5485, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/__pycache__/stat.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=4273, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecbb90)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=4273, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\4[\205dm\25\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 4274) = 4273
# read(3, "", 1)                          = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/_collections_abc.py", {st_mode=S_IFREG|0644, st_size=32284, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/_collections_abc.py", {st_mode=S_IFREG|0644, st_size=32284, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/__pycache__/_collections_abc.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=32925, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecbb90)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=32925, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\4[\205d\34~\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 32926) = 32925
# read(3, "", 1)                          = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/posixpath.py", {st_mode=S_IFREG|0644, st_size=16250, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/posixpath.py", {st_mode=S_IFREG|0644, st_size=16250, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/__pycache__/posixpath.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=10530, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecbb90)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=10530, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\4[\205dz?\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 10531) = 10530
# read(3, "", 1)                          = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/genericpath.py", {st_mode=S_IFREG|0644, st_size=4975, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/genericpath.py", {st_mode=S_IFREG|0644, st_size=4975, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/__pycache__/genericpath.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=3907, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecad00)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=3907, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\4[\205do\23\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 3908) = 3907
# read(3, "", 1)                          = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/_sitebuiltins.py", {st_mode=S_IFREG|0644, st_size=3128, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/_sitebuiltins.py", {st_mode=S_IFREG|0644, st_size=3128, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/__pycache__/_sitebuiltins.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=3547, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecca20)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=3547, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\4[\205d8\f\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 3548) = 3547
# read(3, "", 1)                          = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/bin/pyvenv.cfg", 0x7fff6aecd3c0, 0) = -1 ENOENT (No such file or directory)
# newfstatat(AT_FDCWD, "/usr/pyvenv.cfg", 0x7fff6aecd3c0, 0) = -1 ENOENT (No such file or directory)
# geteuid()                               = 1000
# getuid()                                = 1000
# getegid()                               = 1000
# getgid()                                = 1000
# newfstatat(AT_FDCWD, "/home/orange/.local/lib/python3.10/site-packages", {st_mode=S_IFDIR|0775, st_size=4096, ...}, 0) = 0
# openat(AT_FDCWD, "/home/orange/.local/lib/python3.10/site-packages", O_RDONLY|O_NONBLOCK|O_CLOEXEC|O_DIRECTORY) = 3
# newfstatat(3, "", {st_mode=S_IFDIR|0775, st_size=4096, ...}, AT_EMPTY_PATH) = 0
# getdents64(3, 0x5577054e4ba0 /* 10 entries */, 32768) = 352
# getdents64(3, 0x5577054e4ba0 /* 0 entries */, 32768) = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/local/lib/python3.10/dist-packages", {st_mode=S_IFDIR|0755, st_size=4096, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/local/lib/python3.10/dist-packages", O_RDONLY|O_NONBLOCK|O_CLOEXEC|O_DIRECTORY) = 3
# newfstatat(3, "", {st_mode=S_IFDIR|0755, st_size=4096, ...}, AT_EMPTY_PATH) = 0
# getdents64(3, 0x5577054e4ba0 /* 2 entries */, 32768) = 48
# getdents64(3, 0x5577054e4ba0 /* 0 entries */, 32768) = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3/dist-packages", {st_mode=S_IFDIR|0755, st_size=12288, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3/dist-packages", O_RDONLY|O_NONBLOCK|O_CLOEXEC|O_DIRECTORY) = 3
# newfstatat(3, "", {st_mode=S_IFDIR|0755, st_size=12288, ...}, AT_EMPTY_PATH) = 0
# getdents64(3, 0x5577054e4ba0 /* 245 entries */, 32768) = 9912
# getdents64(3, 0x5577054e4ba0 /* 0 entries */, 32768) = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/dist-packages", {st_mode=S_IFDIR|0755, st_size=4096, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/dist-packages", O_RDONLY|O_NONBLOCK|O_CLOEXEC|O_DIRECTORY) = 3
# newfstatat(3, "", {st_mode=S_IFDIR|0755, st_size=4096, ...}, AT_EMPTY_PATH) = 0
# getdents64(3, 0x5577054e4ba0 /* 3 entries */, 32768) = 96
# getdents64(3, 0x5577054e4ba0 /* 0 entries */, 32768) = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/sitecustomize.py", {st_mode=S_IFREG|0644, st_size=155, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/sitecustomize.py", {st_mode=S_IFREG|0644, st_size=155, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/__pycache__/sitecustomize.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=225, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecc660)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=225, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\223\21Hb\233\0\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 226) = 225
# read(3, "", 1)                          = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/lib-dynload", {st_mode=S_IFDIR|0755, st_size=12288, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/lib-dynload", {st_mode=S_IFDIR|0755, st_size=12288, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/lib-dynload", {st_mode=S_IFDIR|0755, st_size=12288, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/lib-dynload", O_RDONLY|O_NONBLOCK|O_CLOEXEC|O_DIRECTORY) = 3
# newfstatat(3, "", {st_mode=S_IFDIR|0755, st_size=12288, ...}, AT_EMPTY_PATH) = 0
# getdents64(3, 0x5577054e4ba0 /* 49 entries */, 32768) = 3080
# getdents64(3, 0x5577054e4ba0 /* 0 entries */, 32768) = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/home/orange/.local/lib/python3.10/site-packages", {st_mode=S_IFDIR|0775, st_size=4096, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/home/orange/.local/lib/python3.10/site-packages", {st_mode=S_IFDIR|0775, st_size=4096, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/home/orange/.local/lib/python3.10/site-packages", {st_mode=S_IFDIR|0775, st_size=4096, ...}, 0) = 0
# openat(AT_FDCWD, "/home/orange/.local/lib/python3.10/site-packages", O_RDONLY|O_NONBLOCK|O_CLOEXEC|O_DIRECTORY) = 3
# newfstatat(3, "", {st_mode=S_IFDIR|0775, st_size=4096, ...}, AT_EMPTY_PATH) = 0
# getdents64(3, 0x5577054e4ba0 /* 10 entries */, 32768) = 352
# getdents64(3, 0x5577054e4ba0 /* 0 entries */, 32768) = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/local/lib/python3.10/dist-packages", {st_mode=S_IFDIR|0755, st_size=4096, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/local/lib/python3.10/dist-packages", {st_mode=S_IFDIR|0755, st_size=4096, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/local/lib/python3.10/dist-packages", {st_mode=S_IFDIR|0755, st_size=4096, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/local/lib/python3.10/dist-packages", O_RDONLY|O_NONBLOCK|O_CLOEXEC|O_DIRECTORY) = 3
# newfstatat(3, "", {st_mode=S_IFDIR|0755, st_size=4096, ...}, AT_EMPTY_PATH) = 0
# getdents64(3, 0x5577054e4ba0 /* 2 entries */, 32768) = 48
# getdents64(3, 0x5577054e4ba0 /* 0 entries */, 32768) = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3/dist-packages", {st_mode=S_IFDIR|0755, st_size=12288, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3/dist-packages", {st_mode=S_IFDIR|0755, st_size=12288, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3/dist-packages", {st_mode=S_IFDIR|0755, st_size=12288, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3/dist-packages", O_RDONLY|O_NONBLOCK|O_CLOEXEC|O_DIRECTORY) = 3
# newfstatat(3, "", {st_mode=S_IFDIR|0755, st_size=12288, ...}, AT_EMPTY_PATH) = 0
# getdents64(3, 0x5577054e4ba0 /* 245 entries */, 32768) = 9912
# getdents64(3, 0x5577054e4ba0 /* 0 entries */, 32768) = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3/dist-packages/apport_python_hook.py", {st_mode=S_IFREG|0644, st_size=8063, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3/dist-packages/apport_python_hook.py", {st_mode=S_IFREG|0644, st_size=8063, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3/dist-packages/__pycache__/apport_python_hook.cpython-310.pyc", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=4661, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecb7d0)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=4661, ...}, AT_EMPTY_PATH) = 0
# read(3, "o\r\r\n\0\0\0\0\324Axa\177\37\0\0\343\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"..., 4662) = 4661
# read(3, "", 1)                          = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10", {st_mode=S_IFDIR|0755, st_size=20480, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/lib-dynload", {st_mode=S_IFDIR|0755, st_size=12288, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/home/orange/.local/lib/python3.10/site-packages", {st_mode=S_IFDIR|0775, st_size=4096, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/local/lib/python3.10/dist-packages", {st_mode=S_IFDIR|0755, st_size=4096, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3/dist-packages", {st_mode=S_IFDIR|0755, st_size=12288, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/dist-packages", {st_mode=S_IFDIR|0755, st_size=4096, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/dist-packages", {st_mode=S_IFDIR|0755, st_size=4096, ...}, 0) = 0
# newfstatat(AT_FDCWD, "/usr/lib/python3.10/dist-packages", {st_mode=S_IFDIR|0755, st_size=4096, ...}, 0) = 0
# openat(AT_FDCWD, "/usr/lib/python3.10/dist-packages", O_RDONLY|O_NONBLOCK|O_CLOEXEC|O_DIRECTORY) = 3
# newfstatat(3, "", {st_mode=S_IFDIR|0755, st_size=4096, ...}, AT_EMPTY_PATH) = 0
# getdents64(3, 0x5577054e4ba0 /* 3 entries */, 32768) = 96
# getdents64(3, 0x5577054e4ba0 /* 0 entries */, 32768) = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/home/orange/Project/my_go/sikumi/01/./hello.py", {st_mode=S_IFREG|0775, st_size=61, ...}, 0) = 0
# openat(AT_FDCWD, "/home/orange/Project/my_go/sikumi/01/./hello.py", O_RDONLY|O_CLOEXEC) = 3
# newfstatat(3, "", {st_mode=S_IFREG|0775, st_size=61, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aece700)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# lseek(3, -22, SEEK_END)                 = 39
# lseek(3, 0, SEEK_CUR)                   = 39
# read(3, "\n\n# strace ./hello.py\n", 4096) = 22
# lseek(3, 0, SEEK_END)                   = 61
# lseek(3, 0, SEEK_CUR)                   = 61
# lseek(3, 0, SEEK_SET)                   = 0
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0775, st_size=61, ...}, AT_EMPTY_PATH) = 0
# read(3, "#!/usr/bin/python3\nprint(\"hello "..., 62) = 61
# read(3, "", 1)                          = 0
# lseek(3, 0, SEEK_SET)                   = 0
# close(3)                                = 0
# newfstatat(AT_FDCWD, "/home/orange/Project/my_go/sikumi/01/./hello.py", {st_mode=S_IFREG|0775, st_size=61, ...}, 0) = 0
# readlink("./hello.py", 0x7fff6aebe060, 4096) = -1 EINVAL (Invalid argument)
# getcwd("/home/orange/Project/my_go/sikumi/01", 1024) = 37
# readlink("/home/orange/Project/my_go/sikumi/01/hello.py", 0x7fff6aebd7e0, 1023) = -1 EINVAL (Invalid argument)
# openat(AT_FDCWD, "/home/orange/Project/my_go/sikumi/01/./hello.py", O_RDONLY) = 3
# ioctl(3, FIOCLEX)                       = 0
# newfstatat(3, "", {st_mode=S_IFREG|0775, st_size=61, ...}, AT_EMPTY_PATH) = 0
# ioctl(3, TCGETS, 0x7fff6aecf010)        = -1 ENOTTY (Inappropriate ioctl for device)
# lseek(3, 0, SEEK_CUR)                   = 0
# newfstatat(3, "", {st_mode=S_IFREG|0775, st_size=61, ...}, AT_EMPTY_PATH) = 0
# read(3, "#!/usr/bin/python3\nprint(\"hello "..., 4096) = 61
# lseek(3, 0, SEEK_SET)                   = 0
# read(3, "#!/usr/bin/python3\nprint(\"hello "..., 4096) = 61
# read(3, "", 4096)                       = 0
# close(3)                                = 0
# write(1, "hello world\n", 12hello world
# )           = 12
# rt_sigaction(SIGINT, {sa_handler=SIG_DFL, sa_mask=[], sa_flags=SA_RESTORER|SA_ONSTACK, sa_restorer=0x7f0f34278520}, {sa_handler=0x55770381a410, sa_mask=[], sa_flags=SA_RESTORER|SA_ONSTACK, sa_restorer=0x7f0f34278520}, 8) = 0
# munmap(0x7f0f33934000, 151552)          = 0
# exit_group(0)                           = ?
# +++ exited with 0 +++
