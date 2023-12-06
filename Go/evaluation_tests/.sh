#!/bin/bash

# path to executable
programPath="img.exe"

paramsArray=("-i test_images\cows.jpg -o test_images\cows.png"
"-i test_images\houses.jpg g_blur -o test_images\houses.png"
"-i test_images\view.jpg a_compositing -i2 test_images\houses.jpg -opacity 0.4 -o test_images\view.gif"
"-i test_images\view.jpg bresen_line -shape circle -pos 500x500 -size 300 -o test_images\view.png")

timesToRun=20

# path to results
resultsPath="results.txt"

# delete last results
if [ -f "$resultsPath" ]; then
    rm $resultsPath
fi

for params in "${paramsArray[@]}"; do
    totalTime=0

    for (( i=0; i<$timesToRun; i++ )); do
        startTime=$(date +%s%3N)

        $programPath $params

        endTime=$(date +%s%3N)
        executionTime=$((endTime-startTime))

        totalTime=$(($totalTime + $executionTime))
    done

    meanTime=$(($totalTime / $timesToRun))

    echo "Mean execution time for parameters '$params': $meanTime ms" >> $resultsPath
done
