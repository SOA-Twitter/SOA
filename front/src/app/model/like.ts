export class Like {
    tweetId:string;
    username:string;
    liked:boolean;

    constructor(obj?: any){
        this.tweetId = obj && obj.tweetId || null;
        this.username = obj && obj.username || null;
        this.liked = obj && obj.liked || null;
    }
}
