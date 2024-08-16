CC = go
SRC = main.go
TARGET = dump-remover
PREFIX = /usr
INSTALL_DIR = /opt/dump-remover

all: $(TARGET)

$(TARGET): $(SRC)
	$(CC) build -o $(TARGET) $(SRC)

run: $(SRC)
	$(CC) run $(SRC)

install:
	install -d 	   $(INSTALL_DIR)
	install -m 644 $(TARGET) 				$(INSTALL_DIR)
	ln -s 		   $(INSTALL_DIR)/$(TARGET) $(PREFIX)/bin/$(TARGET)

uninstall:
	@rm -f 	$(PREFIX)/bin/$(TARGET)
	@rm -rf $(INSTALL_DIR)

clean:
	@rm -f $(TARGET)

.PHONY: all, clean
