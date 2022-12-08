import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';
import { recoveryMail } from '../model/recoveryMail';

@Component({
  selector: 'app-profile-recovery',
  templateUrl: './profile-recovery.component.html',
  styleUrls: ['./profile-recovery.component.css']
})
export class ProfileRecoveryComponent implements OnInit {

  form!: FormGroup;
  recoveryMail!: recoveryMail;

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private authService: AuthService
  ) { 
    this.createForm();
  }

  ngOnInit(): void {
  }

  createForm(){
    this.form = this.fb.group({
      'email' : new FormControl(null, [Validators.required, Validators.pattern('^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$')]),
    });
  }


  onSubmit(){
    this.recoveryMail = new recoveryMail(this.form.value);
    this.authService.recoveryMail(this.recoveryMail).subscribe( () => this.router.navigateByUrl("/login"))
  }
}
