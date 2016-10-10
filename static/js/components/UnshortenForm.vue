<template>
    <div class="unshorten-form">
        <span class="site-name">{{site_name}}</span>
        <form class="unshorten" action="/" method="POST" @submit.prevent="submit">
            <input type="text" class="unshorten-linkid" placeholder="DhvQnR" name="linkid" maxlength="6" v-model="value" />
        </form>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    
    .unshorten-form{
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: flex-end;
        
        & > .site-name{
            font-family: $pt-sans;
            @include fontSize(25px);
            color: $dark-text;
        }
        
        & > .unshorten{
            & > input[type="text"]{
                width: 100%;
                max-width: 130px;
                background-color: rgba(46,145,246, 0.2);
                margin-left: 5px;
                border: none;
                outline: none;
                font-family: $pt-sans;
                @include fontSize(25px);
                color: $dark-text;
                text-align: center;
            }
        }
    }
</style>

<script>
    import {ValidateLinkid, JsonPostRequest} from '../filters/index';
    import {NoticeUserError} from '../filters/ErrorClass';
    
    export default{
        name: 'unshorten-form',
        
        methods:{
            submit: function(){
                // Validate the linkid
                try{
                    let formValue = ValidateLinkid(this.value);
                    JsonPostRequest('/api/unshorten', formValue).then(function(response){
                        if(response !== true){
                            if(response.code !== 200){
                                const apiError = new NoticeUserError(response.message, true);
                                apiError.show();
                            }else{
                                console.log(response);
                            }
                        }
                    }.bind(this));
                }catch(e){
                    let linkidError = new NoticeUserError(e.message, true);
                    linkidError.show();
                }
            }
        },
        
        data(){
            return{
                site_name: `${location.origin}/`,
                value: ''
            }
        }
    }
</script>