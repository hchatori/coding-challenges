#include <iostream>
#define NIL -1
#define MAX 100

int memo[MAX];

int fib(int n)
{
	if (memo[n] != NIL) 
	{
		return memo[n];
	}

	if (n <= 1)
	{
		memo[n] = n;
	} else 
	{
		memo[n] = fib(n-1) + fib (n-2);
	}
	
	return memo[n];
}

void initialize()
{
	for (int i = 0; i <= MAX; i++) 
	{
		memo[i] = NIL;
	}

}

int main() 
{
	int n = 40;
	initialize();
	std::cout << fib(n) << "\n";
	return 0;
}