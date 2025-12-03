@echo off
title Ortak API Server
color 0A
echo.
echo ========================================
echo    ORTAK - Takim ve Gorev Yonetimi
echo ========================================
echo.
echo Checking port 8080...
for /f "tokens=5" %%a in ('netstat -aon ^| findstr :8080') do (
    echo Killing existing process on port 8080...
    taskkill /f /pid %%a >nul 2>&1
)
echo.
echo Starting API server...
echo Server will be available at: http://localhost:8080
echo.
echo Available endpoints:
echo   POST /api/v1/register
echo   POST /api/v1/login
echo   GET  /api/v1/users
echo   POST /api/v1/users
echo   GET  /api/v1/teams
echo   POST /api/v1/teams
echo   GET  /api/v1/tasks
echo   POST /api/v1/tasks
echo   GET  /api/v1/health
echo.
echo ========================================
echo.

cd /d "%~dp0"
set "GIN_MODE=release"
go run cmd/api/main.go

echo.
echo API server stopped.
pause