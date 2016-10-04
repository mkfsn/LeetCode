CC=gcc
CFLAGS=-g -Wall -O2
EXT=c
TARGET=$(shell basename $(CURDIR))

.PHONY: c cpp clean run

c: CC=gcc
c: EXT=c
c: $(TARGET)

cpp: CC=g++
cpp: EXT=cpp
cpp: CFLAGS+=-std=c++11
cpp: $(TARGET)

$(TARGET):
ifeq ($(DEBUG), TRUE)
	$(CC) $(CFLAGS) -DDEBUG -o $@ $@.$(EXT)
else
	$(CC) $(CFLAGS) -o $@ $@.$(EXT)
endif

clean:
	rm -rf $(TARGET) $(TARGET).dSYM

run:
	./$(TARGET)
