import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

// import { NgxWebstorageModule } from 'ngx-webstorage';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './navbar/navbar.component';
import { UnloggedHomeComponent } from './unlogged-home/unlogged-home.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { PostComponentComponent } from './post/post-component/post-component.component';
import { LoggedHomeComponent } from './logged-home/logged-home.component';
import { LoggedNavbarComponent } from './logged-navbar/logged-navbar.component';
import { CreatePostComponent } from './create-post/create-post.component';
import { NgxCaptchaModule } from 'ngx-captcha';
import { MyProfileComponent } from './my-profile/my-profile.component';
import { ChangePasswordComponent } from './change-password/change-password.component';
import { SignupBusinessComponent } from './signup-business/signup-business.component';
import { ProfileRecoveryComponent } from './profile-recovery/profile-recovery.component';
import { RecoverAccountComponent } from './recover-account/recover-account.component';
import { Guard } from './guard/guard';
import { httpInterceptorProviders } from 'src';
import { LikesComponent } from "./post/likes/likes.component";
import { RequestsComponent } from "./my-profile/requests/requests.component";

@NgModule({
    declarations: [
        AppComponent,
        NavbarComponent,
        UnloggedHomeComponent,
        LoginComponent,
        SignupComponent,
        PostComponentComponent,
        LoggedHomeComponent,
        LoggedNavbarComponent,
        CreatePostComponent,
        MyProfileComponent,
        ChangePasswordComponent,
        SignupBusinessComponent,
        ProfileRecoveryComponent,
        RecoverAccountComponent,
        LikesComponent,
        RequestsComponent
    ],
    providers: [httpInterceptorProviders, Guard],
    bootstrap: [AppComponent],
    imports: [
        BrowserModule,
        AppRoutingModule,
        ReactiveFormsModule,
        HttpClientModule,
        NgxCaptchaModule,
    ]
})
export class AppModule { }
