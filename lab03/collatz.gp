set term png
set output "collatz.png"
set xlabel "Liczba"
set ylabel "Długość ciągu Collatza"
plot "collatz.dat" using 1:2 with points pt 7 ps 0.3 lc rgb "black"
