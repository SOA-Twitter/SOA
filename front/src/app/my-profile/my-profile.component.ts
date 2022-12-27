import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router';
import { AuthService } from '../auth.service';
import { Post } from '../model/post';
import { Request } from '../model/request';

declare var window:any;

@Component({
  selector: 'app-my-profile',
  templateUrl: './my-profile.component.html',
  styleUrls: ['./my-profile.component.css']
})
export class MyProfileComponent implements OnInit {

  posts: Post[] = [];
  usernames: Request[] = [];
  isFollowedByLoggedUser = false;

  formModal: any;

  currentUser = {
    "username": "",
    "first_name": "",
    "last_name": "",
    "email": "",
    "gender": "",
    "country": "",
    "age": 0,
    "company_name": "",
    "company_website": "",
    "private": false,
    "role": ""
  };

  constructor(
    private service: AuthService,
    private route: ActivatedRoute,
  ) { }

  ngOnInit(): void {

    this.route.paramMap.subscribe((params: ParamMap) => {
      let username = params.get('username') || ""

      this.service.getProfileDetails((username)).subscribe((userDetails) => {
        this.currentUser = {
          "username": userDetails.username,
          "first_name": userDetails.first_name,
          "last_name": userDetails.last_name,
          "email": userDetails.email,
          "gender": userDetails.gender,
          "country": userDetails.country,
          "age": userDetails.age,
          "company_name": userDetails.company_name,
          "company_website": userDetails.company_website,
          "private": userDetails.private,
          "role": userDetails.role
        };

      },
        (error) => { window.alert('Error: ' + error) });

    })

    this.route.paramMap.subscribe((params: ParamMap) => {
      let username = params.get('username') || ""
      this.service.getPostByLoggedUser((username)).subscribe(posts => { this.posts = posts })
    });

    this.isFollowedByLoggedUser = this.isFollowed();
  }

  protected canEditPrivacy(): boolean {
    return this.currentUser.username === this.service.getUsername();
  }

  protected changePrivacy() {
    let checkBoxElem = document.getElementById('privacy-checkbox') as HTMLInputElement;
    checkBoxElem.disabled = true;

    this.service.changeProfilePrivacy(checkBoxElem.checked).subscribe(
      (privacy) => {
        console.log(JSON.stringify(privacy));
        checkBoxElem.disabled = false;
      },
      (error: HttpErrorResponse) => {
        console.log(JSON.stringify(error));
        alert("An error occured during privacy change request.");
        // Revert checkbox value to previous state, before the request
        checkBoxElem.checked = this.currentUser.private;
        checkBoxElem.disabled = false;
      }
    );
  }

  showModal(){
    this.formModal = new window.bootstrap.Modal(
      document.getElementById("req")
    );

    this.service.getRequests().subscribe((usernames) => {this.usernames = usernames});

    this.formModal.show();
  }

  closeModal(){
    this.formModal.hide();
  }

  followUser(){
    let followbtn = document.getElementById("followBtn") as HTMLInputElement
      let usernameFromUrl = this.route.snapshot.params['username'];
      this.service.followUser(usernameFromUrl).subscribe(() => {
        followbtn.disabled = true
        followbtn.value = "Requested"
      })
  }

  protected removeHandledFollowRequest(requester_username: string) {

    this.usernames = this.usernames.filter( (request_iter)=> {
      request_iter.username != requester_username;
    })
  }

  private isFollowed() : boolean{

    let isFollowedInfo: boolean;
    let followbtn = document.getElementById("followBtn") as HTMLInputElement;

    this.service.isFollowed(this.route.snapshot.params['username']).subscribe(
      (isFollowedObj) => {
        isFollowedInfo = isFollowedObj.is_followed;
        if (isFollowedInfo) {
          followbtn.value = 'Following';
          followbtn.disabled = true;
          return true;
        }
        else{
          return false;
        }
      },
      (error: HttpErrorResponse) => {
        console.log(JSON.stringify(error));
      }
    );
    return false;
  }
}
