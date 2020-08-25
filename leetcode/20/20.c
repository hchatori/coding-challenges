#include <stdio.h>
#include <string.h>

bool isValid(char * s) {
	for (int i = 0; i < strlen(s); i++) {
		printf("%s", i);
	}
	printf("\n");
	return false;
}

int main() {
	char s[] = "([)]"; // false
	char s[] = "{[]}"; // true
	printf("%d\n", isValid(s));
	return 0;
}