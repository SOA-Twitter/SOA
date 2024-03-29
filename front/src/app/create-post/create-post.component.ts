import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../auth.service';
import { createPost } from '../model/createPost';

@Component({
  selector: 'app-create-post',
  templateUrl: './create-post.component.html',
  styleUrls: ['./create-post.component.css']
})
export class CreatePostComponent implements OnInit {

  postForm!: FormGroup;
  post!:createPost;

  constructor(private fb: FormBuilder, private authService: AuthService, private router: Router) {
    this.createForm();
  }

  ngOnInit(): void {}

  createForm(){
    this.postForm = this.fb.group({
      'text' : new FormControl(null, [Validators.required, Validators.pattern('[a-zA-Z-\'_./0-9]+')]),
    })
  }

  createPost(){
    this.post = new createPost(this.postForm.value)
    this.authService.createPost(this.post).subscribe(()=>{this.router.navigateByUrl("/logged-home");})
  }

}
