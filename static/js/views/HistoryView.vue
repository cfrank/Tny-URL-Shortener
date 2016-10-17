<template>
    <div class="history-view">
        <div v-if="historyAvailable" class="history-container">
            <history-entry v-for="item in historyItems" :item="item"></history-entry>
        </div>
        <div v-else class="history-loading"></div>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    
    @keyframes spin {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }
    
    .history-view{
        width: 100%;
        max-width: 650px;
        max-height: 80vh;
        margin: 0 auto;
        overflow-y: auto;
        
        &::-webkit-scrollbar{
            display: none;
        }
    }
    
    .history-loading{
        width: 24px;
        height: 24px;
        margin: 0 auto;
        border: .25rem solid rgba(255,255,255,0.3);
        border-top-color: #fff;
        border-radius: 50%;
        animation: spin 755ms infinite linear;
    }
</style>

<script>
    import * as Constants from '../filters/constants';
    import {JsonPostRequest} from '../filters/';
    import {NoticeUserError} from '../filters/ErrorClass';
    
    import HistoryEntry from '../components/HistoryEntry.vue';
    
    export default{
        name: "history-view",
        
        data(){
            return{
                historyAvailable: false,
                historyUnavailable: false,
                historyItems: []
            }
        },
        
        components:{
            HistoryEntry
        },
        
        /* Run when the component is created */
        created: function(){
            let userId = window.localStorage.getItem(Constants.USERID_LOCALSTORAGE);
            let userKey = window.localStorage.getItem(Constants.USERID_KEY_LOCALSTORAGE);
            
            if(this.historyAvailable !== true){
                JsonPostRequest('/api/history', {userId, userKey}).then(function(response){
                    if(response !== false && response.length > 0 && response.code === 200){
                        this.historyAvailable = true;
                        this.historyItems = response.links;
                    }
                    else if(response.length === 0){
                        this.historyUnavailable = true;
                    }
                    else{
                        const historyApiError = new NoticeUserError(response.message, false);
                        historyApiError.show();
                    }
                }.bind(this));
            }
        }
    }
</script>