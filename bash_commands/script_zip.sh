#!/bin/bash

today=$(date +%Y_%m_%d__%H_%M_%S);
for i in */; 
do zip -r "${i%/}.zip" "$i"
mkdir -p ./backup/data/$today/ 
mv ${i%/}.zip backup/data/$today/; 
done 
