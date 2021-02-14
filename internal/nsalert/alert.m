#import <stdlib.h>
#import <AppKit/NSAlert.h>

void info(char *title, char *message) {
	NSString* titleString = [NSString stringWithCString:title encoding:[NSString defaultCStringEncoding]];
	NSString* messageString = [NSString stringWithCString:message encoding:[NSString defaultCStringEncoding]];

	NSAlert *alert = [[NSAlert alloc] init];
	alert.messageText = titleString;
	alert.informativeText = messageString;
	alert.alertStyle = NSAlertStyleInformational;

	[alert runModal];
	[alert release];
}

void error(char *title, char *message) {
	NSString* titleString = [NSString stringWithCString:title encoding:[NSString defaultCStringEncoding]];
	NSString* messageString = [NSString stringWithCString:message encoding:[NSString defaultCStringEncoding]];

	NSAlert *alert = [[NSAlert alloc] init];
	alert.messageText = titleString;
	alert.informativeText = messageString;
	alert.alertStyle = NSAlertStyleCritical;

	[alert runModal];
	[alert release];
}

bool question(char *title, char *message, char *defaultButton, char*alternateButton) {
	NSString* titleString = [NSString stringWithCString:title encoding:[NSString defaultCStringEncoding]];
	NSString* messageString = [NSString stringWithCString:message encoding:[NSString defaultCStringEncoding]];
	NSString* defaultButtonString = [NSString stringWithCString:defaultButton encoding:[NSString defaultCStringEncoding]];
	NSString* alternateButtonString = [NSString stringWithCString:alternateButton encoding:[NSString defaultCStringEncoding]];

	NSAlert *alert = [[NSAlert alloc] init];
	alert.messageText = titleString;
	alert.informativeText = messageString;
	[alert addButtonWithTitle:defaultButtonString];
	[alert addButtonWithTitle:alternateButtonString];

	NSInteger button = [alert runModal];
	[alert release];
	if (button == NSAlertFirstButtonReturn) {
		return true;
	}
	return false;
}