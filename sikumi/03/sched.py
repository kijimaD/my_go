#!/usr/bin/python3
import sys
import time
import os
import plot_sched
def usage():
    print("""使い方: {} <プロセス数>
    * 論理CPU0上で<プロセス数>の数だけ同時に100ミリ秒程度CPUリソースを消費する負荷処理プロセスを起動した後に、すべてのプロセスの終了を待つ
    * "sched-<プロセス数>.jpg"というファイルに実行結果を示したグラフを書き出す
    * グラフのx軸は負荷処理プロセス開始からの経過時間[ミリ秒]、y軸は進捗[%]""".format(progname, file=sys.stderr))
    sys.exit(1)

# 実験に適した負荷を見積もるための前処理にかける負荷
