class GeneralError extends Error{
        constructor(message){
                super(message);
                this.name = this.constructor.name;
                this.message = message;
                if(typeof Error.captureStackTrace === 'function'){
                        Error.captureStackTrace(this, this.constructor)
                }
                else{
                        this.stack = (new Error(message)).stack;
                }
        }
}

export default class InvalidLinkError extends GeneralError{
        constructor(message){
                super(message);
        }
}