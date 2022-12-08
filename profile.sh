#!/bin/zsh

rm times.txt
touch times.txt

cd cmd

for i in {1..8}
do
  pwd
  cd "day$i"
  go build -o "build"
  echo "day$i :" >> ../../times.txt
  \time -al -o ../../times.txt ./build
  cd ..
done
