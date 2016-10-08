export class GeneralError extends Error{
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

export class InvalidLinkError extends GeneralError{
        constructor(message){
                super(message);
        }
}

export class NoticeUserError{
        constructor(message, store){
                this.message = message;
                this.store = store;
        }
        
        show(){
                // Only show the error if it's not already visible
                if(!this.store.state.notice.active){
                        this.store.dispatch('showNotice', {
                                message: this.message,
                                type: 'error',
                        });
                        
                        setTimeout(() =>{
                                // Only fire if it is still showing
                                if(this.store.state.notice.active)
                                this.store.dispatch('hideNotice');
                        }, 5000);
                }
        }
}