#!/usr/bin/python3

# プロセスを生成する様子を見る

import os, sys
ret = os.fork()
if ret == 0:
  print("子プロセス: pid={}, 親プロセスのpid={}".format(os.getpid(), os.getppid()))
  exit()
elif ret > 0:
  print("親プロセス: pid={}, 子プロセスのpid={}".format(os.getpid(), ret))
  exit()
sys.exit(1)

# 親プロセス: pid=2805200, 子プロセスのpid=2805201
# 子プロセス: pid=2805201, 親プロセスのpid=2805200
