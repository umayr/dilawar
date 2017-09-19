@echo off

SET APP_NAME=dilawar
SET DIR_PATH=%PROGRAMDATA%\dilawar
SET DOWNLOADS_PATH=C:%HOMEPATH%\Downloads
SET RELEASE_PATH=https://github.com/umayr/dilawar/releases/download/
SET VERSION=0.1.0

echo **********************************************************
echo *********WELCOME TO %APP_NAME% APP INSTALLER~WINDOWS*********
echo **********************************************************
echo. 
echo.
echo BELOW IS THE LIST OF SUPPORTED ARCHITECTURE:
echo 1. x86 BIT ARCHITECTURE
echo 2. x64 BIT ARCHITECTURE
echo.
SET /p option="PLEASE SELECT YOUR ARCHITECTURE FROM THE OPTIONS GIVEN ABOVE: "

if %option% == 1 goto x86
if %option% == 2 goto x64
goto eof

:x86 
SET EXE_NAME=dilawar_windows_386.exe
goto install

:x64
SET EXE_NAME=dilawar_windows_amd64.exe
goto install

:install
SET EXE_PATH=%RELEASE_PATH%/%VERSION%/%EXE_NAME%
ECHO DOWNLOADING THE RELEVANT PRECOMPILED EXECUTABLE FOR YOUR ARCHITECTURE . . . . 
powershell -command "& { (New-Object Net.WebClient).DownloadFile('%EXE_PATH%', '%DOWNLOADS_PATH%\dilawar.exe') }"

MKDIR %DIR_PATH%
MOVE %DOWNLOADS_PATH%\dilawar.exe %DIR_PATH%\
SETX PATH "%DIR_PATH%;%PATH%"

ECHO %APP_NAME% HAS BEEN SUCCESSFULLY INSTALLED IN YOUR PC.
SET /p exit="PRESS ANY KEY TO CONTINUE. . . "

:eof