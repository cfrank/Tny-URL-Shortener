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
        }
}