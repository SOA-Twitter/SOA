export class Post{
    id:string;
    username:string;
    text:string;
    creationDate:Date;

    constructor(obj?: any){
        this.id = obj && obj.id || null;
        this.username = obj && obj.username || null;
        this.text = obj && obj.text || null;
        this.creationDate = obj && obj.creationDate || null;
    }
}