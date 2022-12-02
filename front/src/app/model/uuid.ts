export class uuid{
    recovery_uuid:string;
    
    constructor(obj?: any){
        this.recovery_uuid = obj && obj.recovery_uuid || null;
    }
}