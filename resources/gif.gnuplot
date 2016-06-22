set terminal png
set xrange[-10:130]
set yrange[-10:80]
#do for [i=500:500000:500]{set output sprintf("frame%d.png",i);plot "datos.dat" every ::1::i using 2:3 w l}
#plot "datos.dat" every ::1::100 using 2:3 w l
do for [i=50:1500:100]{set output sprintf("frame%d.png",i);plot "datos.dat" every ::1::i using 2:3 w l}
do for [i=1500:2000:200]{set output sprintf("frame%d.png",i);plot "datos.dat" every ::1::i using 2:3 w l}
do for [i=2000:80000:300]{set output sprintf("frame%d.png",i);plot "datos.dat" every ::1::i using 2:3 w l}
