import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';
import { ChangePassword } from '../model/changePassword';
import { ConfirmPasswordValidator } from '../recover-account/confirm-password.validator';

@Component({
  selector: 'app-change-password',
  templateUrl: './change-password.component.html',
  styleUrls: ['./change-password.component.css']
})
export class ChangePasswordComponent implements OnInit {

  form!: FormGroup;
  newPassword!: ChangePassword;

  constructor(private fb: FormBuilder, private authService: AuthService, private router: Router,) {
    this.createForm();
  }

  ngOnInit(): void {
  }

  createForm(){
    this.form = this.fb.group({
      'old_password': new FormControl('', [Validators.required, Validators.pattern('^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$')]),
      'new_password': new FormControl('', [Validators.required, Validators.pattern('^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$')]),
      'repeated_password': new FormControl('', [Validators.required, Validators.pattern('^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$')])
    },
    {
      validator: ConfirmPasswordValidator("new_password", "repeated_password")
    });
  }

  onSubmit(){
    this.newPassword = new ChangePassword(this.form.value);
    console.log(this.form.value);
    this.authService.changePassword(this.newPassword).subscribe(() =>
    {
      this.authService.logout();
      this.router.navigateByUrl("/login");
    })
  }

}
