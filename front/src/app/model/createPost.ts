export class createPost {
    text:string;

    constructor(obj?: any){
        this.text = obj && obj.text || null;
    }
}
