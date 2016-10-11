<template>
    <form class="shorten-form" action="/" method="POST" @submit.prevent="submit">
        <input type="text" class="shorten-form-url" :placeholder="placeholder" :value="formValue" @input="updateFormValue" />
        <input type="submit" class="shorten-form-submit" :disabled="linkSuccess.active" value="Shorten" />
    </form>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    .shorten-form{
        height: 50px;
        display: flex;
        justify-content: space-between;
        
        & > input{
            width: 100%;
            height: 100%;
            border: none;
            outline: 0px;
            transition: background-color 0.2s ease-in-out;
            
            font-family: $pt-sans;
            @include fontSize(18px);
        }
        
        & > .shorten-form-url{
            max-width: 500px;
            height: 100%;
            padding: 0px 10px;
            background-color: rgba(255,255,255,0.8);
            box-shadow: $elm-shadow;
            color: $dark-text;
            
            &:focus{
                background-color: #fff;
            }
        }
        
        & > .shorten-form-submit{
            max-width: 130px;
            background-color: #fff;
            box-shadow: $elm-shadow;
            cursor: pointer;
            color: $tny-blue;
            transition: background-color 0.3s ease-out, color 0.3s ease-out;
            
            /* Remove the dotted line FireFox */
            &::-moz-focus-inner{
                border: 0;
            }
            
            &:disabled{
                background-color: #f6f3f3;
                color: #d5d7d8;
                cursor: default;
                
                &:active{
                    background-color: #f6f3f3;
                }
            }
            
            &:active{
                margin-top: 1px; 
                background-color: #f0f0f0;
            }
        }
    }
</style>

<script>
    import { mapGetters } from 'vuex';
    import * as Constants from '../filters/constants';
    import {ValidateUrl, JsonPostRequest} from '../filters/index';
    import {NoticeUserError} from '../filters/ErrorClass.js';
    
    export default{
        name: 'shorten-form',
        
        data(){
            return{
                placeholder: 'Add your url here...'
            }
        },
        
        computed: mapGetters(['linkSuccess', 'formValue']),
        
        methods:{
            /*
             * Update the form value state
            */
            updateFormValue: function(event){
                this.$store.dispatch('updateFormValue', event.target.value);
            },
            /*
             * The submit method handles the submit event
             */
            submit: function(){
                try{
                    // Create the object which will be sent to the db
                    let linkData = {
                        source: ValidateUrl(this.$store.state.formValue),
                        userid: window.localStorage.getItem(Constants.USERID_LOCALSTORAGE),
                        userkey: window.localStorage.getItem(Constants.USERID_KEY_LOCALSTORAGE),
                        created: ~~(Date.now() / 1000),
                    };
                    
                    JsonPostRequest('/api/add', linkData).then(function(response){
                        // There was an error with the request
                        if(response !== false){
                            // Json was recieved check if it is valid
                            if(response.code !== 200){
                                const apiError = new NoticeUserError(response.message, true);
                                apiError.show();
                            }else{
                                this.$store.dispatch('showLinkSuccess', {
                                    title: 'Your short link: ',
                                    linkHref: `${location.origin}/${response.linkid}`,
                                });
                            }
                        }
                    }.bind(this));
                } catch(e){
                    const linkError = new NoticeUserError(e.message, true);
                    linkError.show();
                }
            }
        }
    }
</script>