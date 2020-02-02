GOOS=linux GOARM=5 GOARCH=arm go build -o gpioutils $1 &&
arm-none-eabi-strip ./gpioutils &&
upx ./gpioutils