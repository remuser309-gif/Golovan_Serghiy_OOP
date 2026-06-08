@echo off
cd /d D:\Golovan\boilerplate-go-back
set DB_NAME=iot_db
set DB_PASSWORD=postgres
"C:\Program Files\Go\bin\go.exe" run ./cmd/server/
pause
