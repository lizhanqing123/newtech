#include <stdio.h>
#include <string.h>




void del_space( char *str )
{
	int i = 0;
	char *p = str;
	while(*p)
	{
		if ( *p !=  ' ' ) str[i++] = *p;

		p++;
	}
	str[i] = '\0';
}


int main(int argc, int argv[] )
{
	int a = 10;
	int const *p = &a;
	p++;
	return 0;
}

