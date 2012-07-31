TARG=test

.PHONY: all clean

all: $(TARG)

$(TARG):
	gcc -o $@ $(wildcard *.c) `sdl2-config --cflags --libs`

clean:
	rm -f $(TARG)
