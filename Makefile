all:
	gcc passgen.c -lssl -lcrypto -o passgen
clean:
	rm passgen
