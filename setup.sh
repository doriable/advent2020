#!/bin/bash

DAY=$1
YEAR=$2

if [ -z $YEAR ]
then
    YEAR="2020"
fi

mkdir -p day${DAY}
curl -s --cookie "session=$AOC_SESSION" https://adventofcode.com/${YEAR}/day/${DAY}/input > day${DAY}/input
cp template.go day${DAY}/solution.go
