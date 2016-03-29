CC=gcc
CFLAGS=-g -Wall -O2
EXT=c
TARGET=binary-tree-level-order-traversal

.PHONY: c cpp clean

c: CC=gcc
c: EXT=c
c: $(TARGET)

cpp: CC=g++
cpp: EXT=cpp
cpp: CFLAGS+=-std=c++11
cpp: $(TARGET)

$(TARGET):
	$(CC) $(CFLAGS) -o $@ $@.$(EXT)

clean:
	rm -f $(TARGET)
