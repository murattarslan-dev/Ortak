@echo off
title Ortak API Tests
color 0B
echo.
echo ========================================
echo    ORTAK - Unit Test Runner
echo ========================================
echo.

cd /d "%~dp0"

set modules=./internal/auth/service ./internal/auth/handler ./internal/user/service ./internal/user/handler ./internal/team/service ./internal/team/handler ./internal/task/service ./internal/task/handler

for %%m in (%modules%) do (
    echo Running tests for %%m...
    go test -v %%m
    if errorlevel 1 (
        echo Tests failed for %%m
    ) else (
        echo Tests passed for %%m
    )
    echo ---
)

echo.
echo ========================================
echo Tests completed!
echo ========================================
pause