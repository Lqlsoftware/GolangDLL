#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// import libinterop
#include "libinterop.h"

//  convert str to GoString
void str2GoString(char *str, GoString *gostr)
{
    gostr->p         = str;
    gostr->n         = strlen(str);
}

//  convert GoString to str
void GoString2str(GoString *gostr, char* cstr)
{
    memcpy(cstr, gostr->p, gostr->n);
    cstr[gostr->n]   = '\0';
}

int main()
{
    // Print massage
    printf("#########################################\n");
    printf("### C/C++ Calling Golang Shared-C lib ###\n");
    printf("#########################################\n");

    // Call the external functions
    //   Function "Init" to create named pipe
    Init();
    //   Function "Send" to Send string.
    char str[1000];
    while (1) {
        // read from input
        printf("Input message: ");
        gets(str);
        // Send message
        Send(str);
    }

    return 0;
}
