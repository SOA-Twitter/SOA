export class UserIsFollowed{
    is_followed: boolean;

    constructor(obj?: any){
        this.is_followed = obj && obj.is_followed || null;
    }
}