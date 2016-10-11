<!-- The Application view -->
<template>
    <div class="viewport">
        <!-- Error handler -->
        <notice-header></notice-header>
        <!-- View component -->
        <router-view></router-view>
        <!-- Menu -->
        <footer-menu></footer-menu>
    </div>
</template>

<style lang="sass">
    @import '../css/header.scss';
    @import '../css/global.scss';
    
    .viewport{
        height: 100%;
        background-color: $tny-blue;
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>

<script>
    import {CallFetchJson} from './filters/';
    import * as Constants from './filters/constants';
    import NoticeHeader from './components/NoticeHeader.vue';
    import FooterMenu from './components/FooterMenu.vue';
    
    export default{
        name: 'AppView',
        
        components:{
            NoticeHeader,
            FooterMenu
        },
        
        created: function(){
            if(window.localStorage.getItem(Constants.USERID_LOCALSTORAGE) === null
            || window.localStorage.getItem(Constants.USERID_KEY_LOCALSTORAGE) === null){
                // Uid not set
                try{
                    CallFetchJson('/api/uid').then(function(response){
                        if(response !== false){
                            window.localStorage.setItem(Constants.USERID_LOCALSTORAGE, response.uid);
                            window.localStorage.setItem(Constants.USERID_KEY_LOCALSTORAGE, response.key);
                        }
                    });
                }catch(e){
                    // There was a problem recieving the uid data
                    console.log(e);
                }
            }
        }
    }
</script>