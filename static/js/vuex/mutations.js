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
        
        UPDATE_FORM_VALUE(state, message){
                state.formValue = message;
        },
        
        REMOVE_UID(state){
                window.localStorage.removeItem(USERID_LOCALHOST);
                state.userId = '';
        },
        
        CACHE_HISTORY(state, message){
                state.historyCache = message;
        },
        
        EXPIRE_HISTORY_CACHE(state){
                state.historyCache = {
                        links: [],
                        length: 0,
                        stale: true,
                }
        }
}