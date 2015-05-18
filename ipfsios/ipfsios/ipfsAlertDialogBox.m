//
//  ipfsAlertDialogBox.m
//  ipfsios
//
//  Created by Jay Vaughan on 4/23/15.
//  Copyright (c) 2015 Jay Vaughan. All rights reserved.
//

#import <UIKit/UIKit.h>

@interface ipfsAlertDelegate : NSObject<UIAlertViewDelegate>
@property int buttonTapped;
@property NSCondition* condition;
-(int) waitForButtonTap;
@end

@implementation ipfsAlertDelegate
- (id) init {
    self = [super init];
    if (self) {
        self.buttonTapped = -1;
        self.condition = [[NSCondition alloc] init];
    }
    return self;
}


-(void)alertView:(UIAlertView *)alertView didDismissWithButtonIndex:(NSInteger)buttonIndex {
    self.buttonTapped = buttonIndex;
    
    [self.condition lock];
    [self.condition broadcast];
    [self.condition unlock];
}

-(int)waitForButtonTap {
    [self.condition lock];
    while (self.buttonTapped < 0)
        [self.condition wait];
    [self.condition unlock];
    return self.buttonTapped;
}


@end


char* PopUpDialogBox(char* msg){
    BOOL userPressedYes = NO;
    @autoreleasepool {
        ipfsAlertDelegate* myDelegate = [[ipfsAlertDelegate alloc] init];
        UIAlertView* alert = [[UIAlertView alloc] initWithTitle:@"web debug"
                                                        message:[NSString stringWithFormat:@"Response to request:\n%s", msg]
                                                       delegate:myDelegate cancelButtonTitle:@"No" otherButtonTitles:@"Yes", nil];
        [alert performSelectorOnMainThread:@selector(show) withObject:nil waitUntilDone: NO];
        userPressedYes = [myDelegate waitForButtonTap] == 1;
    }
    return userPressedYes ? "YES" : "NO";
}