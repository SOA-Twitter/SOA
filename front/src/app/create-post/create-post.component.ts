import { Component, OnInit } from '@angular/core';
import { UntypedFormControl, UntypedFormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-create-post',
  templateUrl: './create-post.component.html',
  styleUrls: ['./create-post.component.css']
})
export class CreatePostComponent implements OnInit {
  postForm: UntypedFormGroup = new UntypedFormGroup({
    username: new UntypedFormControl('', Validators.required),
    communityId: new UntypedFormControl('', Validators.required),
    title: new UntypedFormControl('', Validators.required),
    text: new UntypedFormControl('', Validators.required)
  });

  constructor() { }

  ngOnInit(): void {
  }

}
