<template>
    <div class="history-view">
        <div v-if="historyLoading === false && this.historyUnavailable === false" class="history-container">
            <history-entry v-for="item in historyItems" :item="item"></history-entry>
        </div>
        <div v-else class="history-status">
            <!-- Loading -->
            <div class="history-loading" :class="{'active': this.historyLoading}"></div>
            <!-- No history items -->
            <div class="history-empty" :class="{'active': this.historyUnavailable}">
                <p>You have no history items!</p>
            </div>
        </div>
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
        max-height: 95vh;
        overflow-y: auto;
        
        &::-webkit-scrollbar{
            display: none;
        }
    }
    
    .history-loading.active{
        width: 24px;
        height: 24px;
        margin: 0 auto;
        border: .25rem solid rgba(255,255,255,0.3);
        border-top-color: #fff;
        border-radius: 50%;
        animation: spin 755ms infinite linear;
    }
    
    .history-empty{
        width: 100%;
        text-align: center;
        display: none;
        
        &.active{
            display: block;
        }
    }
</style>

<script>
    import { mapGetters } from 'vuex';
    import * as Constants from '../filters/constants';
    import {JsonPostRequest} from '../filters/';
    import {NoticeUserError} from '../filters/ErrorClass';
    
    import HistoryEntry from '../components/HistoryEntry.vue';
    
    export default{
        name: "history-view",
        
        data(){
            return{
                historyLoading: true,
                historyUnavailable: false,
                historyItems: []
            }
        },
        
        components:{
            HistoryEntry
        },
        
        computed: mapGetters(['historyCache']),
        
        /* Run when the component is created */
        created: function(){
            let userId = window.localStorage.getItem(Constants.USERID_LOCALSTORAGE);
            let userKey = window.localStorage.getItem(Constants.USERID_KEY_LOCALSTORAGE);
            
            if(this.historyCache.stale === true){
                JsonPostRequest('/api/history', {userId, userKey}).then(function(response){
                    if(response !== false && response.length > 0 && response.code === 200){
                        this.historyLoading = false;
                        this.historyItems = response.links;
                        this.$store.dispatch('cacheHistory', {
                            links: response.links,
                            length: response.length,
                            stale: false
                        });
                    }
                    else if(response.code === 204){
                        // No history items
                        this.historyLoading = false;
                        this.historyUnavailable = true;
                    }
                    else{
                        const historyApiError = new NoticeUserError("Error retrieving history from server", true);
                        historyApiError.show();
                    }
                }.bind(this));
            }
            else if(!this.historyCache.length > 0){
                this.historyLoading = false;
                this.historyUnavailable = true;
            }
            else{
                this.historyLoading = false;
                this.historyItems = this.historyCache.links;
            }
        }
    }
</script>