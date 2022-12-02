import { Component, OnInit, SecurityContext } from '@angular/core';
import { UntypedFormGroup, UntypedFormControl, Validators, UntypedFormBuilder } from '@angular/forms';
import { User } from 'src/app/model/user';
import { AuthService } from '../auth.service';
import { DomSanitizer } from '@angular/platform-browser';
import { Router } from '@angular/router';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {

  user!: User;
  register!: UntypedFormGroup;
  siteKey: string;

  constructor(private fb: UntypedFormBuilder, private authService: AuthService, private sanitizer: DomSanitizer, private router: Router){
    this.createForm();
    this.siteKey = '6Lcq4CYjAAAAAC28ZFxmcXD5w-D7UxBpQalorlnJ';
  }

  createForm(){
    this.register = this.fb.group({
      'email' : new UntypedFormControl(null, [Validators.required, Validators.pattern('^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$')]),
      'password' : new UntypedFormControl(null, [Validators.required, Validators.pattern('^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$')]),
      'username' : new UntypedFormControl(null, [Validators.required, Validators.pattern('(?=.{2,})[a-zA-Z0-9._]$')]),
      'firstName' : new UntypedFormControl(null, [Validators.required, Validators.pattern('(?=.{2,})[a-zA-Z]$')]),
      'lastName' : new UntypedFormControl(null, [Validators.required, Validators.pattern('(?=.{2,})[a-zA-Z]$')]),
      'gender' : new UntypedFormControl(null, [Validators.required, Validators.pattern('(?=.{4,})[a-zA-Z]$')]),
      'country' : new UntypedFormControl(null, [Validators.required]),
      'age' : new UntypedFormControl(null, [Validators.required, Validators.pattern('/^(?:1[8-9]|[2-9][0-9]|100)$$/')])
    });
  }

  ngOnInit() {}

  submit(){
    this.user = new User(this.register.value); 
    this.authService.signUp(this.user).subscribe(() => {this.router.navigateByUrl("/login");});
  }

}
