@echo off
::start cmd /K
echo ######################Step1  start clone clash dashboard...
cd ./bin & if not exist clash-dashboard git clone https://github.com/Dreamacro/clash-dashboard.git -b gh-pages
echo.

echo ######################Step2  start update...
cd ./clash-dashboard & git fetch origin & git reset --hard origin/gh-pages
echo.

echo ######################Step3  view status
git status
echo.
pause
