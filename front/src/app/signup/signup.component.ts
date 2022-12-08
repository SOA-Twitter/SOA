import { Component, OnInit, SecurityContext } from '@angular/core';
import {  Validators, FormGroup, FormBuilder, FormControl } from '@angular/forms';
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
  register!: FormGroup;
  siteKey: string;

  constructor(private fb: FormBuilder, private authService: AuthService, private sanitizer: DomSanitizer, private router: Router){
    this.createForm();
    this.siteKey = '6Lcq4CYjAAAAAC28ZFxmcXD5w-D7UxBpQalorlnJ';
  }

  createForm(){
    this.register = this.fb.group({
      'email' : new FormControl(null, [Validators.required, Validators.pattern('^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$')]),
      'password' : new FormControl(null, [Validators.required, Validators.pattern('^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$')]),
      'username' : new FormControl(null, [Validators.required, Validators.pattern('[a-zA-Z0-9._]{2,}$')]),
      'first_name' : new FormControl(null, [Validators.required, Validators.pattern('[a-zA-Z-\']{2,}')]),
      'last_name' : new FormControl(null, [Validators.required, Validators.pattern('[a-zA-Z-\']{2,}')]),
      'gender' : new FormControl(null, [Validators.required]),
      'country' : new FormControl(null, [Validators.required, Validators.pattern('[a-zA-Z-]{4,}')]),
      'age' : new FormControl(null, [Validators.required, Validators.min(18), Validators.max(100)])
    });
  }

  ngOnInit() {}

  submit(){
    this.user = new User(this.register.value); 
    this.authService.signUp(this.user).subscribe((res) => 
    {
      if (res.toString() == 'Created' || res.toString() == '201' || res.toString() == 'Ok' || res.toString() == '200') {
        this.router.navigateByUrl("/login");
      }else{
        alert("An error occurred while registrating. Please try again later!");
      }
      
    
    });
  }

}
