#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <stdio.h> 

/* 

	Step 1: Have two pointers, beg and end, one for "(" and one for ")". 
	Iterating through expression from beginning to end, find the first 
	instance of "(" and set beg to that index. Iterating through expression
	from end to beginning, find the first instance of ")" and set end to that
	index. 
	
	Step 2: Recursively call the function on the string found between beg and
	end (currString). 
	
	Step 3: When there are no more parentheticals, take the index 
	value of (beg - 1) which should give me the special character, and 
	evaluate currString given the logical operator. (Need to write if 
	statements for each logical operator.) 
		Note: 
		"&(expr1, expr2, ...)" only evaluates to true if all true.
		"|(expr1, expr2)" only evaluates to false if all false.
	
	Step 4: Store the evaluated boolean in currBool. 
	
	Step 5: Pop back up and evaluate currString given currBool. Continue popping
	back up until expression is fully parsed, then return currBool. 
	
	*/ 

bool parseBoolExpr(char* expression){
	int beg = 0;
	int end = 0;
	// for (int i = 0; i < strlen(expression); i++) {
	// 	printf("%s", expression[i]);
	// }
	return false;
}

int main() {
	int expLen = 16;
	// char expression[] = "|(&(t,f,t),!(t))"; // return false
	// char expression[] = "|(f,t)"; // return true
	// parseBoolExpr(expression);
	char* expression = (char*)malloc(expLen * sizeof(char));
	return 0;
}