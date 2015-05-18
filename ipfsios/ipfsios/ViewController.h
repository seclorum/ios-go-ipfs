//
//  ViewController.h
//  ipfsios
//
//  Created by Jay Vaughan on 4/22/15.
//  Copyright (c) 2015 Jay Vaughan. All rights reserved.
//

#import <UIKit/UIKit.h>

@interface ViewController : UIViewController
@property (strong, nonatomic) IBOutlet UITextField *valB;
@property (strong, nonatomic) IBOutlet UITextField *valA;
@property (strong, nonatomic) IBOutlet UILabel *resultLabel;
- (IBAction)addGoButtonPressed:(id)sender;

@end

