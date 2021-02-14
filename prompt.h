#ifndef PROMPT
#define PROMPT

#import <stdlib.h>
#import <stdbool.h>

void message(char *title, char *message);

void error(char *title, char *message);

bool question(char *title, char *message, char *defaultButton, char *alternateButton);

#endif /* !PROMPT */