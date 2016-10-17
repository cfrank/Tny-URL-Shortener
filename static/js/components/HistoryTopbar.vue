<template>
    <div class="top-bar flex-box">
        <div class="history-short-link left-side">
            <a :href="getShortUrl(linkid)" target="_blank">{{getShortUrl(linkid)}}</a>
        </div>
        <div class="history-favicon right-side">
            <img :src="`https://s2.googleusercontent.com/s2/favicons?alt=s32&sz=32&domain_url=${source}`" height="32" width="32" />
        </div>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    
    .top-bar{
        width: 100%;
        height: 62px;
        padding: 0px 25px;
        background-color: #fff;
        border-bottom: 1px solid #e3e2e2;
        align-items: center;
        
        & > .history-short-link > a{
            font-family: $pt-sans;
            @include fontSize(21px);
            color: $tny-blue;
            text-decoration: none;
        }
    }
</style>

<script>
    import * as Constants from '../filters/constants';
    
    export default{
        name: 'history-topbar',
        
        props: ['linkid', 'source'],
        
        methods:{
            getShortUrl: function(linkid){
                return `${Constants.HOST_NAME}${linkid}`;
            }
        }
    }
</script>