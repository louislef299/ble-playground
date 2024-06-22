PI_BUILD_NAME="pi-peripheral"

pi:
	GOOS=linux GOARCH=arm GOARM=7 \
	go build -o $(PI_BUILD_NAME) ./peripheral
	scp $(PI_BUILD_NAME) pi5:/home/louis/

clean:
	rm $(PI_BUILD_NAME)
