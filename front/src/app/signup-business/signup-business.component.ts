import { Component, OnInit } from '@angular/core';
import { UntypedFormBuilder, UntypedFormControl, UntypedFormGroup, Validators } from '@angular/forms';
import { DomSanitizer } from '@angular/platform-browser';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';
import { BuisnessUser } from '../model/businessUser';

@Component({
  selector: 'app-signup-business',
  templateUrl: './signup-business.component.html',
  styleUrls: ['./signup-business.component.css']
})
export class SignupBusinessComponent implements OnInit {

  user!: BuisnessUser;
  register!: UntypedFormGroup;
  siteKey: string;

  constructor(private fb: UntypedFormBuilder, private authService: AuthService, private sanitizer: DomSanitizer, private router: Router){
    this.createForm();
    this.siteKey = '6Lcq4CYjAAAAAC28ZFxmcXD5w-D7UxBpQalorlnJ';
  }

  createForm(){
    this.register = this.fb.group({
      'company_name' : new UntypedFormControl(null, [Validators.required, Validators.pattern('[a-zA-Z-\']+')]),
      'company_website' : new UntypedFormControl(null, [Validators.required, Validators.pattern('[a-zA-Z-\'_./]+')]),
      'email' : new UntypedFormControl(null, [Validators.required, Validators.pattern('^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$')]),
      'password' : new UntypedFormControl(null, [Validators.required, Validators.pattern('^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$')]),
      'username' : new UntypedFormControl(null, [Validators.required, Validators.pattern('[a-zA-Z0-9._]{2,}$')])
    });
  }

  ngOnInit(): void {}

  submit(){
    this.user = new BuisnessUser(this.register.value);
    this.authService.signUpBusiness(this.user).subscribe(() => {this.router.navigateByUrl("/login");});
  }
}
