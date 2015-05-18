//
//  ViewController.m
//  ipfsios
//
//  Created by Jay Vaughan on 4/22/15.
//  Copyright (c) 2015 Jay Vaughan. All rights reserved.
//

#import "ViewController.h"
#include "_cgo_export.h"

@interface ViewController ()


@end

@implementation ViewController

- (void) viewDidLoad {
    [super viewDidLoad];
}

- (void) didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
}


- (IBAction)addGoButtonPressed:(id)sender {
    NSLog(@"val a: %d, val b: %d", self.valA.text.intValue, self.valB.text.intValue);
    int result = AddUsingGo(self.valA.text.intValue,
                            self.valB.text.intValue);
    self.resultLabel.text = @(result).stringValue;
    NSLog(@"Result: %@", @(result).stringValue);
}
@end
