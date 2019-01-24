set term png

set xlabel "Input Size"
set ylabel "Time (seconds)"

# set xtics 5000

set style line 1 lc rgb 'blue' pt 7
set style line 2 lc rgb 'green' pt 7

if (!exists("filename")) filename='bench.dat'
plot filename using 1:2 t 'naive' w p ls 1, filename using 1:3 t 'smart' w p ls 2

