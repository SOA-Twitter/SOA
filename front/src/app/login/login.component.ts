import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';
import { UserLogin } from '../model/userLogin';
import { DomSanitizer } from '@angular/platform-browser';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginForm!: FormGroup;
  userLogin!: UserLogin;

  constructor(
    private authService: AuthService,
    private router: Router,
    private fb: FormBuilder,
    private sanitizer: DomSanitizer
  ) {
    this.createForm()
  }

  createForm(){
    this.loginForm = this.fb.group({
      'email' : new FormControl(null, [Validators.required, Validators.pattern('^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$')]),
      'password': new FormControl('', [Validators.required, Validators.pattern('^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$')])
    });
  }

  ngOnInit(): void {}

  submit(){
    this.userLogin = new UserLogin(this.loginForm.value);
    this.authService.login(this.userLogin).subscribe((token: string) =>
    {
      this.router.navigateByUrl("/logged-home");
      token = token.replace("\"", "");
      token = token.replace("\"", "");
      localStorage.setItem('jwt token', token);
    },
    () => {window.alert('Invalid credentials!')});
  }


}
