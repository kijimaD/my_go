#include <unistd.h>

int main(void) {
  pause();
  return 0;
}

/* 静的リンク */
// $ cc -static -o pause pause.c
// $ ls -l pause
// -rwxrwxr-x 1 orange orange 900240 Jan  5 22:51 pause
// $ ldd pause
// not a dynamic executable

/* 動的リンク */
/* 共有ライブラリlibc.soを使う */
// $ cc -o pause pause.c
// $ ls -l pause
// -rwxrwxr-x 1 orange orange 15960 Jan  5 22:53 pause
// $ ldd pause
// inux-vdso.so.1 (0x00007ffe5b9da000)
// ibc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fca488b9000)
// lib64/ld-linux-x86-64.so.2 (0x00007fca48afd000)

/* ELF */
/* $ readelf -h pause */
/* ELF Header: */
/*   Magic:   7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00  */
/*   Class:                             ELF64 */
/*   Data:                              2's complement, little endian */
/*   Version:                           1 (current) */
/*   OS/ABI:                            UNIX - System V */
/*   ABI Version:                       0 */
/*   Type:                              DYN (Position-Independent Executable file) */
/*   Machine:                           Advanced Micro Devices X86-64 */
/*   Version:                           0x1 */
/*   Entry point address:               0x1060 */
/*   Start of program headers:          64 (bytes into file) */
/*   Start of section headers:          13976 (bytes into file) */
/*   Flags:                             0x0 */
/*   Size of this header:               64 (bytes) */
/*   Size of program headers:           56 (bytes) */
/*   Number of program headers:         13 */
/*   Size of section headers:           64 (bytes) */
/*   Number of section headers:         31 */
/*   Section header string table index: 30 */
