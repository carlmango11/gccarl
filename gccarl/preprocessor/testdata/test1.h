#define FOUR 4
FOUR

#ifndef NO
ifndef
#endif

#ifdef NO
    wrong
    #ifndef FOUR
        nested wrong
    #endif
#endif

#undef FOUR
FOUR

#include <lib1.h>
FOUR