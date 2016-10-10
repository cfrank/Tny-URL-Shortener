import store from '../vuex/store';

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
        constructor(message, hide){
                this.message = message;
                this.hide = hide;
        }
        
        show(){
                // If there is an active notice, close it first
                if(store.state.notice.active){
                        store.dispatch('hideNotice');
                }
                store.dispatch('showNotice', {
                        message: this.message,
                        type: 'error',
                });
                
                if(this.hide){
                        setTimeout(() =>{
                                // Only fire if it is still showing
                                if(store.state.notice.active)
                                        store.dispatch('hideNotice');
                        }, 5000);
                }
        }
}