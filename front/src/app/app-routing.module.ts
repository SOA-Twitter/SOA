import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UnloggedHomeComponent } from './unlogged-home/unlogged-home.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { LoggedHomeComponent } from './logged-home/logged-home.component';
import { CreatePostComponent } from './create-post/create-post.component';
import { MyProfileComponent } from './my-profile/my-profile.component';
import { ChangePasswordComponent } from './change-password/change-password.component';
import { SignupBusinessComponent } from './signup-business/signup-business.component';
import { MyProfileBusinessComponent } from './my-profile-business/my-profile-business.component';
import { ProfileRecoveryComponent } from './profile-recovery/profile-recovery.component';

const routes: Routes = [
    {path:'unlogged-home', component:UnloggedHomeComponent},
    {path:'login', component:LoginComponent},
    {path:'signup', component:SignupComponent},
    {path:'logged-home', component:LoggedHomeComponent},
    {path:'create-post', component:CreatePostComponent},
    {path:'my-profile', component:MyProfileComponent},
    {path:'change-password', component:ChangePasswordComponent},
    {path:'signup-business', component:SignupBusinessComponent},
    {path:'my-profile-business', component:MyProfileBusinessComponent},
    {path:'profile-recovery', component:ProfileRecoveryComponent}
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
  })
  export class AppRoutingModule { }