import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';
import { recoverProfile } from '../model/recoverProfile';
import { recoverProfileWithUuid } from '../model/recoverProfileWithUuid';
import { uuid } from '../model/uuid';
import { ConfirmPasswordValidator } from './confirm-password.validator';

@Component({
  selector: 'app-recover-account',
  templateUrl: './recover-account.component.html',
  styleUrls: ['./recover-account.component.css']
})
export class RecoverAccountComponent implements OnInit {
  form!: FormGroup;
  // form1!: FormGroup;
  newPassword!: recoverProfileWithUuid;
  uuid!: uuid;

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private authService: AuthService
  ) { 
    this.createForm();
    // this.createForm1();
  }

  ngOnInit(): void {
  }

  createForm(){
    this.form = this.fb.group({
      'new_password' : new FormControl(null, [Validators.required, Validators.pattern('^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$')]),
      'repeated_password' : new FormControl(null, [Validators.required]),
      // TODO check if pattern works
      'recovery_uuid' : new FormControl(null, [Validators.required, Validators.pattern('^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$')])    
    },
    {
      validator: ConfirmPasswordValidator("new_password", "repeated_password")
    });
  }

  // createForm1(){
  //   this.form1 = this.fb.group({
  //     'recovery_uuid': new FormControl(null, [Validators.required, Validators.pattern('^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$')])
  //   })
  // }


  onSubmit(){
    this.newPassword = new recoverProfileWithUuid(this.form.value);
    this.authService.recoverAccount(this.newPassword).subscribe(() => {this.router.navigateByUrl("/login")})
  }

}
