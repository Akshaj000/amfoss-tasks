## Commands used in ubuntu terminal.
also vi file.txt can be used to edit files.

/c$ mkdir Coordinate-Location

/c$ cd Coordinate-Location

/c/Coordinate-Location$ mkdir North

/c/Coordinate-Location$ cd North

/c/Coordinate-Location/North$ touch NDegree.txt

/c/Coordinate-Location/North$ cat > NDegree.txt

9°

/c/Coordinate-Location/North$ touch NMinutes.txt

/c/Coordinate-Location/North$ cat > NMinutes.txt

5'

/c/Coordinate-Location/North$ touch NSeconds.txt

/c/Coordinate-Location/North$ cat > NSeconds.txt

38.8"

/c/Coordinate-Location/North$ touch NorthCoordinate.txt

/c/Coordinate-Location/North$ cat NDegree.txt NMinutes.txt NSeconds.txt > NorthCoordinate.txt

/c/Coordinate-Location/North$ sudo mv NorthCoordinate.txt /mnt/c/Coordinate-Location

/c/Coordinate-Location/North$cd ..

/c/Coordinate-Location$sudo mv NorthCoordinate.txt North.txt

/c/Coordinate-Location$ mkdir East

/c/Coordinate-Location$ cd East

/c/Coordinate-Location/East$ touch EDegree.txt

/c/Coordinate-Location/East$ cat > EDegree.txt

76°

/c/Coordinate-Location/East$ touch EMinutes.txt

/c/Coordinate-Location/East$ cat > EMinutes.txt

29'

/c/Coordinate-Location/East$ touch ESeconds.txt

/c/Coordinate-Location/East$ cat > ESeconds.txt

30.8"

/c/Coordinate-Location/East$ touch EastCoordinate.txt

/c/Coordinate-Location/East$ cat EDegree.txt EMinutes.txt ESeconds.txt > EastCoordinate.txt

/c/Coordinate-Location/East$ sudo mv EastCoordinate.txt /mnt/c/Coordinate-Location

/c/Coordinate-Location/East$ cd ..

/c/Coordinate-Location$ sudo mv EastCoordinate.txt East.txt

/c/Coordinate-Location$ cat North.txt East.txt > Location-Coordinate.txt

/c/Coordinate-Location$ touch solution.md

/c/Coordinate-Location$ cd ..

/c$ sudo mv Coordinate-Location /mnt/c/amfoss-tasks/task-01

/c$ cd amfoss-tasks

/c$ git init

/c$ git add .

/c$ git status

/c$ git commit -m "Commited"

/c$ git origin "https://github.com/Akshaj000/amfoss-tasks"

/c$ git branch master

/c$ git checkout master

/c$ git push origin master







