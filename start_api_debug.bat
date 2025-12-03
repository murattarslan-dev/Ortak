@echo off
title Ortak API Server (Debug Mode)
color 0C
echo.
echo ========================================
echo    ORTAK - Debug Mode
echo ========================================
echo.
echo Checking port 8080...
for /f "tokens=5" %%a in ('netstat -aon ^| findstr :8080 2^>nul') do (
    echo Killing existing process on port 8080...
    taskkill /f /pid %%a >nul 2>&1
)
echo.
echo Starting API server in DEBUG mode...
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
set "GIN_MODE=debug"
go run cmd/api/main.go

echo.
echo API server stopped.
pause