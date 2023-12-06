# path to executable
$programPath = "img.exe"

$paramsArray = @("-i test_images\cows.jpg -o test_images\cows.png",
"-i test_images\houses.jpg g_blur -o test_images\houses.png",
"-i test_images\view.jpg a_compositing -i2 test_images\houses.jpg -opacity 0.4 -o test_images\view.gif",
"-i test_images\view.jpg bresen_line -shape circle -pos 500x500 -size 300 -o test_images\view.png")

$timesToRun = 20

$totalTime = 0

# path to results
$resultsPath = "results.txt"

# delete last results
if (Test-Path $resultsPath) {
    Remove-Item $resultsPath
}


foreach ($params in $paramsArray) {
    for ($i = 0; $i -lt $timesToRun; $i++) {
        
        $startTime = Get-Date

        & $programPath $params

        $executionTime = ((Get-Date) - $startTime).TotalMilliseconds

        $totalTime += $executionTime
    }

    $meanTime = $totalTime / $timesToRun

    $totalTime = 0

    "Mean execution time for '$params': $meanTime ms" | Out-File -FilePath $resultsPath -Append
}
