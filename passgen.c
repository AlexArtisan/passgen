#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <openssl/sha.h>

#define MAX(X, Y) (((X) > (Y)) ? (X) : (Y))

char symbols[] = {
	'(', 'E', 'H', '7', '4', 'A', 'l',
	'F', '0', 's', '3', 'k', '1', '#',
	'+', 't', 'x', '-', 'm', '8', '?',
	'I', 'y', 'n', '^', 'w', '%', 'D',
	'_', '5', '!', 'C', 'z', '$', 'J',
	'6', '*', 'B', '2', ')', 'G', '9'
};

int parser(char *arg) {
	return abs(atoi(arg));
}

void worker(int line, unsigned char *password, unsigned char *string, unsigned char buffers[2][33], int width, int last) {

	unsigned char *combined = malloc(strlen((char *) password) + strlen((char *) string) + 1);

	strcpy((char *) combined, (const char *) password);
	strcat((char *) combined, (const char *) string);

	SHA256(combined, strlen((char *) combined), buffers[1]);

	free(combined);

	for (int i = 0; i < 32; i++) {
	  buffers[1][i] = symbols[buffers[1][i] % 42];
	}

	strcpy((char *) buffers[0], (const char *) buffers[1]);

	buffers[1][width] = '\0';

	if (last == 0 || last == line) {
		printf("#%d %s\n", line, buffers[1]);
	}

}

int main(int argc, char **argv) {

	int lines = 10, width = 32, last = 0;

	if (argc > 1) {
		lines = MAX(parser(argv[1]), 1);
	}

	if (argc > 2) {
		width = parser(argv[2]) % 33;
	}

	if (argc > 3 && strcmp(argv[3], "-l") == 0) {
		last = lines;
	}

	unsigned char out[2][33] = {};

	char *prompt = getpass("password: ");

	unsigned char password[strlen(prompt)];

	strcpy((char *) password, prompt);

	worker(1, password, password, out, 0, -1);

	printf("check -> %d\n", (int)out[0][0]);

	worker(1, password, password, out, width, last);

	for (int i = 2; i <= lines; i++) {
		worker(i, password, out[0], out, width, last);
	}

	return 0;

}
