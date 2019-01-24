#!/bin/bash

NAIVE="/tmp/naive"
SMART="/tmp/smart"

ITER="$1"
STEP="$2"

function run {
    { time $1 "$2" > /dev/null; } |& grep real | awk '{print $2}' | sed 's/0m//; s/s//'
}

function filename {
    printf "%s-%dstep" $(echo $ITER | sed 's/000$/K/') $STEP
}

[ $# -ne 2 ] && exit 1

for i in $(seq $STEP $STEP $ITER); do
    /tmp/pointgen -n $i -u 1000 -l -1000 > "/tmp/test.json";
    naivetime=$(run "$NAIVE" "/tmp/test.json")
    smarttime=$(run "$SMART" "/tmp/test.json")

    printf "%d %f %f\n" $i $naivetime $smarttime
done > $(filename).dat

gnuplot -e "filename='$(filename).dat'" plot.plg > "$(filename).png"
