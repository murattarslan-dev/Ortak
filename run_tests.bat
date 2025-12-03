@echo off
title Ortak API Tests
color 0B
echo.
echo ========================================
echo    ORTAK - Unit Test Runner
echo ========================================
echo.

cd /d "%~dp0"
call test_runner.bat

echo.
echo ========================================
echo All tests completed!
echo ========================================
pause