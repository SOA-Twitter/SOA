import { HttpClient, HttpStatusCode } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { Observable } from "rxjs";
import { BuisnessUser } from "./model/businessUser";
import { ChangePassword } from "./model/changePassword";
import { recoverProfileWithUuid } from "./model/recoverProfileWithUuid";
import { recoveryMail } from "./model/recoveryMail";
import { User } from "./model/user";
import { UserLogin } from "./model/userLogin";
import { createPost } from "./model/createPost";
import { Post } from "./model/post";
import { Like } from "./model/like";

@Injectable({
    providedIn: 'root'
})

export class AuthService {

    constructor (
        private httpClient: HttpClient,
        private router: Router
        ){}

    private readonly path:string = "https://localhost:8081/auth/";

    signUp(user: User): Observable<HttpStatusCode>{
        return this.httpClient.post<HttpStatusCode>(this.path + 'register', user);
    }

    signUpBusiness(user: BuisnessUser): Observable<BuisnessUser>{
        return this.httpClient.post<BuisnessUser>(this.path + 'register', user);
    }

    login(user: UserLogin): Observable<any>{
        return this.httpClient.post(this.path + 'login', user, {responseType:'text'});
    }

    recoveryMail(recoveryMail: recoveryMail): Observable<recoveryMail>{
        return this.httpClient.post<recoveryMail>(this.path + 'recoverEmail', recoveryMail)
    }

    recoverAccount(newPassword: recoverProfileWithUuid): Observable<any>{
        return this.httpClient.post(this.path + 'recover', newPassword)
    }

    changePassword(changePassword: ChangePassword): Observable<any>{
      return this.httpClient.post<ChangePassword>('https://localhost:8081/profile/changePassword', changePassword)
    }

    createPost(post: createPost): Observable<createPost>{
      return this.httpClient.post<createPost>("https://localhost:8081/tweet/postTweets", post)
    }

    getProfileDetails(username: string): Observable<any>{
      return this.httpClient.get('https://localhost:8081/profile/' + username)
    }

    getPostByLoggedUser(username: string): Observable<Post[]>{
      return this.httpClient.get<Post[]>(`https://localhost:8081/tweet/getTweets/${username}`)
    }

    changeProfilePrivacy(privacy: boolean): Observable<any>{
      return this.httpClient.put('https://localhost:8081/profile/privacy', { "private" : privacy })
    }

    getLikesByTweetId(id: string): Observable<Like[]>{
      return this.httpClient.get<Like[]>(`https://localhost:8081/tweet/getTweetLikes/${id}`);
    }

    like(liked: boolean, id: string): Observable<any>{
      return this.httpClient.put('https://localhost:8081/tweet/like/' + id, { "liked" : liked })
    }

    followUser(username: string): Observable<any>{
      return this.httpClient.post('https://localhost:8081/social/follow', {"username" : username})
    }

    getRequests(): Observable<Request[]>{
      return this.httpClient.get<Request[]>('https://localhost:8081/social/pending')
    }

    acceptReq(username: string): Observable<any>{
      return this.httpClient.put('https://localhost:8081/social/accept', {"username" : username})
    }

    decline(username: string): Observable<any>{
      return this.httpClient.put('https://localhost:8081/social/decline', {"username" : username})
    }

    isFollowed(username: string): Observable<any>{
      return this.httpClient.get('https://localhost:8081/social/isFollowed/'+ username)
    }

    getHomePageForLoggedUser(): Observable<Post[]>{
      return this.httpClient.get<Post[]>("https://localhost:8081/social/homeFeed")
    }

    get isAuthenticated() {
        return this.getToken();
    }

    logout() {
      localStorage.removeItem('jwt token')
      this.router.navigateByUrl('/login');
    }

    getUsername(): string {
      let token = this.parseToken();
      if (token) {
        return this.parseToken()['Username']
      }
      return "";
    }

    getEmail(): string {
      let token = this.parseToken();
      if (token) {
        return this.parseToken()['Email']
      }
      return "";
    }

    getRoles() {
      let token = this.parseToken();
      if (token) {
        return this.parseToken()['Role'];
      }
      return []
    }

    private parseToken() {
      let jwt = localStorage.getItem('jwt token');
      if (jwt !== null) {
        let jwtData = jwt.split('.')[1];
        let decodedJwtJsonData = atob(jwtData);
        let decodedJwtData = JSON.parse(decodedJwtJsonData);
        return decodedJwtData;
      }
    }

    tokenIsPresent(): boolean {
      let token = this.getToken()
      return token != null;
    }

    getToken() {
      let token = localStorage.getItem('jwt token');
      return token
    }

}
