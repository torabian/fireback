@echo off

if exist "%ProgramFiles(x86)%\WinRAR\WinRAR.exe" (
    echo "Using x86"
    
    "%ProgramFiles(x86)%\WinRAR\WinRAR.exe" a -IBCK -afzip -cfg- -ed -ep1 -k -m5 -r -tl ^
    "-sfx%ProgramFiles(x86)%\WinRAR\Zip.sfx" "-z.\winrar-config.txt" ^
    "..\..\artifacts\raspberry-pi-studio-desktop-windows\raspberry-pi-studio-desktop_latest_windows.exe" ^
    ".\build\bin\raspberry-pi-studio-desktop.exe" ^
    ".\raspberry-pi-studio-desktop-configuration.yml"
) else (
    if exist "%ProgramFiles%\WinRAR\WinRAR.exe" (
        echo "Using non x86"
        
        "%ProgramFiles%\WinRAR\WinRAR.exe" a -IBCK -afzip -cfg- -ed -ep1 -k -m5 -r -tl ^
        "-sfx%ProgramFiles%\WinRAR\Zip.sfx" "-z.\winrar-config.txt" ^
        ".\raspberry-pi-studio-desktop_latest_windows.exe" ^
        ".\build\bin\raspberry-pi-studio-desktop.exe" ^
        ".\raspberry-pi-studio-desktop-configuration.yml"
    ) else (
        echo "Found nothing"
    )
)
