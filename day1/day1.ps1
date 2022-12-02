[string[]]$filecontent = Get-Content -Path "C:\Users\1294643814A\OneDrive - United States Air Force\advent_of_code\day_1\day_1_input.txt"

$elf = 0
[int]$count = 0
[int]$newcount = 0

foreach ($a in $filecontent) {

    if ($a -eq "") {
        if ($newcount -gt $count) {

            
            write-host $elf': ' $newcount
            $count = $newcount
        }
  
        $elf++
        $newcount = 0
    }
    else {
        [int]$intme = [int]$a
        $newcount = $intme + $newcount
    }

}



#162: 69177 is the right answer
