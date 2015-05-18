//
//  main.m
//  ipfsios
//
//  Created by Jay Vaughan on 4/22/15.
//  Copyright (c) 2015 Jay Vaughan. All rights reserved.
//

#import <UIKit/UIKit.h>
#import "AppDelegate.h"

int iosmain(int argc, char * argv[]) {
    @autoreleasepool {
        NSLog(@"Startup of obj-c ::");
        return UIApplicationMain(argc, argv, nil, NSStringFromClass([AppDelegate class]));
    }
}
