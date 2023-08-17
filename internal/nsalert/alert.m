#import <stdlib.h>
#import <AppKit/NSAlert.h>

void alert(char *title, char *message, NSAlertStyle style) {
	NSString* titleString = [NSString stringWithCString:title encoding:NSUTF8StringEncoding];
	NSString* messageString = [NSString stringWithCString:message encoding:NSUTF8StringEncoding];
	if([NSThread isMainThread]) {
		NSAlert *alert = [[NSAlert alloc] init];
		alert.messageText = titleString;
		alert.informativeText = messageString;
		alert.alertStyle = NSAlertStyleInformational;

		[alert runModal];
		[alert release];
		return;
	} else {
		dispatch_group_t wg = dispatch_group_create();
		dispatch_group_enter(wg);
		dispatch_async(dispatch_get_main_queue(), ^(void){
			NSAlert *alert = [[NSAlert alloc] init];
			alert.messageText = titleString;
			alert.informativeText = messageString;
			alert.alertStyle = NSAlertStyleInformational;

			[alert runModal];
			[alert release];
			dispatch_group_leave(wg);
			return;
		});
		dispatch_group_wait(wg, DISPATCH_TIME_FOREVER);
	}
}

void info(char *title, char *message) {
	alert(title, message, NSAlertStyleInformational);
}

void error(char *title, char *message) {
	alert(title, message, NSAlertStyleCritical);
}

bool question(char *title, char *message, char *defaultButton, char*alternateButton) {
	NSString* titleString = [NSString stringWithCString:title encoding:NSUTF8StringEncoding];
	NSString* messageString = [NSString stringWithCString:message encoding:NSUTF8StringEncoding];
	NSString* defaultButtonString = [NSString stringWithCString:defaultButton encoding:NSUTF8StringEncoding];
	NSString* alternateButtonString = [NSString stringWithCString:alternateButton encoding:NSUTF8StringEncoding];

	__block NSInteger button;
	if([NSThread isMainThread]) {
		NSAlert *alert = [[NSAlert alloc] init];
		alert.messageText = titleString;
		alert.informativeText = messageString;
		[alert addButtonWithTitle:defaultButtonString];
		[alert addButtonWithTitle:alternateButtonString];

		button = [alert runModal];
		[alert release];
	} else {
		dispatch_group_t wg = dispatch_group_create();
		dispatch_group_enter(wg);
		dispatch_async(dispatch_get_main_queue(), ^(void){

			NSAlert *alert = [[NSAlert alloc] init];
			alert.messageText = titleString;
			alert.informativeText = messageString;
			[alert addButtonWithTitle:defaultButtonString];
			[alert addButtonWithTitle:alternateButtonString];

			button = [alert runModal];
			[alert release];

			dispatch_group_leave(wg);
			return;
		});
		dispatch_group_wait(wg, DISPATCH_TIME_FOREVER);
	}

	if (button == NSAlertFirstButtonReturn) {
		return true;
	}
	return false;
}