export class ChangePassword{
    old_password: string | undefined;
    new_password: string | undefined;
    repeated_password: string | undefined;

    constructor(obj?: any){
        this.old_password && obj.old_password || null;
        this.new_password && obj.new_password || null;
        this.repeated_password && obj.repeated_password || null;
    }
}