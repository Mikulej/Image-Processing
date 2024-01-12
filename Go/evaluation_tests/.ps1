# path to executable
$programPath = "img.exe"

$paramsArray = @("-i test_images\1k.jpg -o test_images\1k2.png",
"-i test_images\2k.jpg -o test_images\2k2.png",
"-i test_images\4k.jpg -o test_images\4k2.png",
"-i test_images\1k.jpg g_blur -ksize 11 -sigma 1 -o test_images\1k.png",
"-i test_images\2k.jpg g_blur -ksize 11 -sigma 1 -o test_images\2k.png",
"-i test_images\4k.jpg g_blur -ksize 11 -sigma 1 -o test_images\4k.png",
"-i test_images\1k.jpg a_compositing -i2 test_images\1k_a.jpg -opacity 0.4 -o test_images\1k.gif",
"-i test_images\1k.jpg bresen_line -shape circle -pos 500x500 -size 300 -o test_images\1k1.png",
"-i test_images\2k.jpg a_compositing -i2 test_images\2k_a.jpg -opacity 0.4 -o test_images\2k.gif",
"-i test_images\2k.jpg bresen_line -shape circle -pos 500x500 -size 300 -o test_images\2k1.png",
"-i test_images\4k.jpg a_compositing -i2 test_images\4k_a.jpg -opacity 0.4 -o test_images\4k.gif",
"-i test_images\4k.jpg bresen_line -shape circle -pos 500x500 -size 300 -o test_images\4k1.png")

$timesToRun = 10

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

        Start-Process -FilePath $programPath -ArgumentList $params -NoNewWindow -Wait

        $executionTime = ((Get-Date) - $startTime).TotalMilliseconds

        $totalTime += $executionTime
    }

    $meanTime = $totalTime / $timesToRun

    $totalTime = 0

    "Mean execution time for '$params': $meanTime ms" | Out-File -FilePath $resultsPath -Append
}
