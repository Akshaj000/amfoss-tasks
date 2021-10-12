/mnt/c$ mkdir Coordinate-Location
/mnt/c$ cd Coordinate-Location
/mnt/c/Coordinate-Location$ mkdir North
/mnt/c/Coordinate-Location$ cd North
/mnt/c/Coordinate-Location/North$ touch NDegree.txt
/mnt/c/Coordinate-Location/North$ cat > NDegree.txt
9°
/mnt/c/Coordinate-Location/North$ touch NMinutes.txt
/mnt/c/Coordinate-Location/North$ cat > NMinutes.txt
5'
/mnt/c/Coordinate-Location/North$ touch NSeconds.txt
/mnt/c/Coordinate-Location/North$ cat > NSeconds.txt
38.8"
/mnt/c/Coordinate-Location/North$ touch NorthCoordinate.txt
/mnt/c/Coordinate-Location/North$ cat NDegree.txt NMinutes.txt NSeconds.txt > NorthCoordinate.txt
/mnt/c/Coordinate-Location/North$ sudo mv NorthCoordinate.txt /mnt/c/Coordinate-Location
/mnt/c/Coordinate-Location/North$cd ..
/mnt/c/Coordinate-Location$sudo mv NorthCoordinate.txt North.txt
/mnt/c/Coordinate-Location$ mkdir East
/mnt/c/Coordinate-Location$ cd East
/mnt/c/Coordinate-Location/East$ touch EDegree.txt
/mnt/c/Coordinate-Location/East$ cat > EDegree.txt
76°
/mnt/c/Coordinate-Location/East$ touch EMinutes.txt
/mnt/c/Coordinate-Location/East$ cat > EMinutes.txt
29'
/mnt/c/Coordinate-Location/East$ touch ESeconds.txt
/mnt/c/Coordinate-Location/East$ cat > ESeconds.txt
30.8"
/mnt/c/Coordinate-Location/East$ touch EastCoordinate.txt
/mnt/c/Coordinate-Location/East$ cat EDegree.txt EMinutes.txt ESeconds.txt > EastCoordinate.txt
/mnt/c/Coordinate-Location/East$ sudo mv EastCoordinate.txt /mnt/c/Coordinate-Location
/mnt/c/Coordinate-Location/East$ cd ..
/mnt/c/Coordinate-Location$ sudo mv EastCoordinate.txt East.txt
/mnt/c/Coordinate-Location$ cat North.txt East.txt > Location-Coordinate.txt
/mnt/c/Coordinate-Location$ touch solution.md
/mnt/c/Coordinate-Location$







