BINARY_DIR :=./tmp
PLIST :=com.github.walkersumida.batteryalarm.plist

install:
	go build -o $(BINARY_DIR)/batteryalarm main.go
	sudo mv $(BINARY_DIR)/batteryalarm /usr/local/bin
	cp ./$(PLIST) ~/Library/LaunchAgents
	launchctl load ~/Library/LaunchAgents/$(PLIST)

update:
	go build -o $(BINARY_DIR)/batteryalarm main.go
	sudo mv $(BINARY_DIR)/batteryalarm /usr/local/bin
	cp ./$(PLIST) ~/Library/LaunchAgents
	launchctl unload ~/Library/LaunchAgents/$(PLIST)
	launchctl load ~/Library/LaunchAgents/$(PLIST)

uninstall:
	sudo rm /usr/local/bin/batteryalarm
	rm ~/Library/LaunchAgents/$(PLIST)
