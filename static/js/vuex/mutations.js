import {USERID_LOCALHOST} from '../filters/';

export default{
        SHOW_NOTICE(state, message){
                state.notice = message;
        },
        
        HIDE_NOTICE(state){
                state.notice = {
                        show: false,
                        message: '',
                        type: '',
                }
        },
        
        REMOVE_UID(state){
                window.localStorage.removeItem(USERID_LOCALHOST);
                state.userId = '';
        }
}