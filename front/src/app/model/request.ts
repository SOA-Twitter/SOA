export class Request{
    username: string;

    constructor(obj?: any){
        this.username = obj && obj.username || null;
    }
}