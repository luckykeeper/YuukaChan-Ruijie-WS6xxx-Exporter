@REM darwin386
set GOOS=darwin
set GOARCH=386
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM darwinamd64
set GOOS=darwin
set GOARCH=amd64
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM darwinarm
set GOOS=darwin
set GOARCH=arm
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM darwinarm64
set GOOS=darwin
set GOARCH=arm64
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM freebsdamd64
set GOOS=freebsd
set GOARCH=amd64
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM linux
set GOOS=linux
set GOARCH=amd64
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM linux386
set GOOS=linux
set GOARCH=386
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM linuxarm64
set GOOS=linux
set GOARCH=arm64
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM linuxppc64le
set GOOS=linux
set GOARCH=ppc64le
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM linuxppc64
set GOOS=linux
set GOARCH=ppc64
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM linuxmips
set GOOS=linux
set GOARCH=mips
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM linuxmipsle
set GOOS=linux
set GOARCH=mipsle
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM linuxmips64
set GOOS=linux
set GOARCH=mips64
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM linuxmips64le
set GOOS=linux
set GOARCH=mips64le
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM openbsd
set GOOS=openbsd
set GOARCH=amd64
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM openbsd
set GOOS=openbsd
set GOARCH=arm
go build -o build/yuukaExporter_%GOOS%_%GOARCH% yuukaExporter.go
@REM Windows386
set GOOS=windows
set GOARCH=386
go build -o build/yuukaExporter_%GOOS%_%GOARCH%.exe yuukaExporter.go
@REM Windowsamd64
set GOOS=windows
set GOARCH=amd64
go build -o build/yuukaExporter_%GOOS%_%GOARCH%.exe yuukaExporter.go