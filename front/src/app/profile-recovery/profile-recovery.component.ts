import { Component, OnInit } from '@angular/core';
import { UntypedFormControl, UntypedFormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-profile-recovery',
  templateUrl: './profile-recovery.component.html',
  styleUrls: ['./profile-recovery.component.css']
})
export class ProfileRecoveryComponent implements OnInit {

  form!: UntypedFormGroup;

  constructor() { 
    this.createForm();
  }

  ngOnInit(): void {
  }

  createForm(){
    this.form = new UntypedFormGroup({
      'email' : new UntypedFormControl(null, [Validators.required, Validators.pattern('^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$')]),
    });
  }

}
