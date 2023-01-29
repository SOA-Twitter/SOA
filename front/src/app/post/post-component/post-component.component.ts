import { HttpErrorResponse } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { AuthService } from 'src/app/auth.service';
import { Like } from 'src/app/model/like';
import { Post } from 'src/app/model/post';

declare var window:any;

@Component({
  selector: 'app-post-component',
  templateUrl: './post-component.component.html',
  styleUrls: ['./post-component.component.css']
})
export class PostComponentComponent implements OnInit {

  @Input()
  post!:Post;

  likes: Like[] = [];

  result: Like[] = [];

  formModal: any;

  constructor(private authService: AuthService) { }

  ngOnInit(): void {
    this.authService.getLikesByTweetId(this.post1.id).subscribe((likes)=>{this.likes = likes, console.log(JSON.stringify(this.likes))});
  }
  
  isLikedByMe():boolean{
    for (let like of this.likes){
      if(like.username === this.authService.getUsername() && like.liked){
        return true;
      }
    }
    return false;
  }

  get post1(){
    return this.post;
  }

  like(id: string){
    var checkBoxElem = document.getElementById(this.post1.id) as HTMLInputElement;
    // console.log(checkBoxElem.checked);
    checkBoxElem.disabled = true;

    this.authService.like(checkBoxElem.checked, id).subscribe(()=>{
      checkBoxElem.disabled = false;
    },
    (error: HttpErrorResponse) => {
      console.log(JSON.stringify(error));
      // checkBoxElem.checked = this.;
      checkBoxElem.disabled = false;
    })
  }

  showModal(){
    this.formModal = new window.bootstrap.Modal(
      document.getElementById(this.post1.id + 1)
    );

    this.result = this.likes.filter(like => like.liked === true);
    console.log(this.result);

    this.formModal.show();
  }
  
  closeModal(){
    this.formModal.hide();
  }
}
