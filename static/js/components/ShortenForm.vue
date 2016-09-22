<template>
    <form class="shorten-form" action="/" method="POST" @submit.prevent="submit">
        <input type="text" class="shorten-form-url" :placeholder="placeholder" v-model="value" />
        <input type="submit" class="shorten-form-submit" value="Submit" />
    </form>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    .shorten-form{
        height: 50px;
        display: flex;
        justify-content: space-between;
        
        & > .shorten-form-url{
            width: 100%;
            max-width: 500px;
            height: 100%;
            padding: 0px 10px;
            border: none;
            background-color: rgba(255,255,255,0.8);
            box-shadow: $elm-shadow;
            transition: background-color 0.2s ease-in-out;
            
            font-family: $pt-sans;
            @include fontSize(18px);
            color: $grey-text;
            
            &:focus{
                background-color: #fff;
            }
        }
        
        & > .shorten-form-submit{
            width: 100%;
            max-width: 130px;
            height: 100%;
            border: none;
            background-color: #fff;
            box-shadow: $elm-shadow;
            cursor: pointer;
            transition: background-color 0.2s ease-in-out;
            
            font-family: $pt-sans;
            @include fontSize(18px);
            color: $teal;
            
            /* Remove the dotted line FireFox */
            &::-moz-focus-inner{
                border: 0;
            }
            
            &:active{
                background-color: #f0f0f0;
            }
        }
    }
</style>

<script>
    import {ValidateUrl} from '../filters/index.js';
    import store from '../vuex/store';
    
    export default{
        name: 'shorten-form',
        
        data(){
            return{
                placeholder: 'Add your url here...',
                value: ''
            }
        },
        
        methods:{
            /*
             * The submit method handles the submit event
             */
            submit: function(){
                try{
                    // Create the object which will be sent to the db
                    let linkData = {
                        encodedUrl: ValidateUrl(this.value),
                        userid: this.$store.state.userId,
                        date: ~~(Date.now() / 1000),
                    };
                    console.log(linkData);
                } catch(e){
                    // Only show the error if it's not already visible
                    if(!this.$store.state.notice.show){
                        store.dispatch('showNotice', {
                            message: e.message,
                            type: 'error',
                        });
                        
                        setTimeout(() =>{
                            // Only fire if it is still showing
                            if(this.$store.state.notice.show)
                                store.dispatch('hideNotice');
                        }, 5000);
                    }
                }
            }
        }
    }
</script>