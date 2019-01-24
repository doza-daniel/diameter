#!/bin/bash

NAIVE="bin/naive"
SMART="bin/smart"

ITER="$1"
STEP="$2"

function run {
    { time $1 "$2" > /dev/null; } |& grep real | sed -E 's/real\s*/scale=4\n/; s/m/*60+/; s/s$//' \
        | bc | sed 's/^\./0./'
}

function filename {
    printf "%s-%dstep" $(echo $ITER | sed 's/000$/K/') $STEP
}

[ $# -ne 2 ] && exit 1

[ ! "$(command -v bc)" ] && exit 1
[ ! "$(command -v gnuplot)" ] && exit 1

{ [ ! -f $NAIVE ] || [ ! -f $SMART  ]; } && exit 1

for i in $(seq $STEP $STEP $ITER); do
    /tmp/pointgen -n $i -u 1000 -l -1000 > "/tmp/test.json";
    naivetime=$(run "$NAIVE" "/tmp/test.json")
    smarttime=$(run "$SMART" "/tmp/test.json")

    printf "%d %f %f\n" $i $naivetime $smarttime
done > $(filename).dat

gnuplot -e "filename='$(filename).dat'" plot.gp > "$(filename).png"
