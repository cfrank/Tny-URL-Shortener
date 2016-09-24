import {USERID_LOCALHOST} from '../filters/';

export default{
        SHOW_NOTICE(state, message){
                state.notice = message;
        },
        
        HIDE_NOTICE(state){
                state.notice = {
                        active: false,
                        message: '',
                        type: '',
                }
        },
        
        SHOW_LINK_SUCCESS(state, message){
                state.linkSuccess = message;
        },
        
        HIDE_LINK_SUCCESS(state){
                state.linkSuccess = {
                        active: false,
                        message: '',
                        linkHref: '',
                }
        },
        
        REMOVE_UID(state){
                window.localStorage.removeItem(USERID_LOCALHOST);
                state.userId = '';
        }
}