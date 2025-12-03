@echo off
title Ortak API Builder
color 0E
echo.
echo ========================================
echo    ORTAK - Build and Run
echo ========================================
echo.

cd /d "%~dp0"

echo Checking port 8080...
for /f "tokens=5" %%a in ('netstat -aon ^| findstr :8080 2^>nul') do (
    echo Killing existing process on port 8080 (PID: %%a)
    taskkill /f /pid %%a >nul 2>&1
)

echo.
echo Building Ortak API...
go build -o ortak-api.exe cmd/api/main.go

if errorlevel 1 (
    echo Build failed!
    pause
    exit /b 1
)

echo Build successful!
echo.
echo Starting API server in new terminal...

start "Ortak API Server" cmd /k "title Ortak API Server && color 0A && echo API Server Running on http://localhost:8080 && echo Press Ctrl+C to stop && set "GIN_MODE=release" && ortak-api.exe"

echo.
echo API server started in new terminal window!
echo Check the new window for server status.
echo.
pause