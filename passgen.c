#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <openssl/sha.h>

char symbols[42] = {
	'(', 'E', 'H', '7', '4', 'A', 'l',
	'F', '0', 's', '3', 'k', '1', '#',
	'+', 't', 'x', '-', 'm', '8', '?',
	'I', 'y', 'n', '^', 'w', '%', 'D',
	'_', '5', '!', 'C', 'z', '$', 'J',
	'6', '*', 'B', '2', ')', 'G', '9'
};

void hash(unsigned char *p, unsigned char *buffer) {
	SHA256(p, strlen((char *) p), buffer);
}

void transform(unsigned char *buffer) {
	for (int i = 0; i < 32; i++) {
		buffer[i] = symbols[buffer[i] % 42];
	}
}

int main(int argc, char **argv) {

	int lines = argc > 1 ? atoi(argv[1]) : 10;

	if (lines < 1) lines = 10;

	unsigned char out[lines][33];

	memset(out, 0, sizeof out);

	int line_len = argc > 2 ? atoi(argv[2]) : 16;

	if (line_len < 1 || line_len > 32) line_len = 16;

	char *pass = getpass("password: ");

	hash((unsigned char *) pass, out[0]);

	for (int i = 1; i < lines; i++) {
		hash(out[i - 1], out[i]);
	}

	for (int i = 0; i < lines; i++) {
		transform(out[i]);
		out[i][line_len] = '\0';
		printf("#%d %s\n", i + 1, out[i]);
	}

	return 0;

}
