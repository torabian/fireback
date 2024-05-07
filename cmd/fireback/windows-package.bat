@echo off

if exist "%ProgramFiles(x86)%\WinRAR\WinRAR.exe" (
    echo "Using x86"
    
    "%ProgramFiles(x86)%\WinRAR\WinRAR.exe" a -IBCK -afzip -cfg- -ed -ep1 -k -m5 -r -tl ^
    "-sfx%ProgramFiles(x86)%\WinRAR\Zip.sfx" "-z.\backend\cmd\fireback\winrar-config.txt" ^
    ".\fireback_latest_windows.exe" ^
    ".\backend\artifacts\fireback\release-win32\fireback.exe"
) else (
    if exist "%ProgramFiles%\WinRAR\WinRAR.exe" (
        echo "Using non x86"
        
        "%ProgramFiles%\WinRAR\WinRAR.exe" a -IBCK -afzip -cfg- -ed -ep1 -k -m5 -r -tl ^
        "-sfx%ProgramFiles%\WinRAR\Zip.sfx" "-z.\backend\cmd\fireback\winrar-config.txt" ^
        ".\fireback_latest_windows.exe" ^
        ".\backend\artifacts\fireback\release-win32\fireback.exe" ^
    ) else (
        echo "Found nothing"
    )
)
