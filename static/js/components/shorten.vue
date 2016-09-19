<!-- The main shortener component -->
<template>
    <div class="shorten-container flex-box">
        <form class="shorten-form" method="POST">
            <input type="text" class="shorten-input" :placeholder="placeholder" autofocus />
            <input type="submit" class="shorten-submit" value="Submit" />
        </form>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    .shorten-container{
        height: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
        
        & > .shorten-form{
            width: 100%;
            max-width: 650px;
            height: 50px;
            display: flex;
            justify-content: space-between;
            
            & > .shorten-input{
                width: 100%;
                max-width: 500px;
                height: 100%;
                padding: 0px 10px;
                border: none;
                box-shadow: $elm-shadow;
                
                font-family: 'PT Sans', sans-serif;
                font-size: 18px;
                color: $grey-text;
            }
            
            & > .shorten-submit{
                width: 100%;
                max-width: 130px;
                height: 100%;
                background-color: #fff;
                border: none;
                box-shadow: $elm-shadow;
                outline: none;
                cursor: pointer;
                transition: background-color 0.5s ease-out;
                
                font-family: 'PT Sans', sans-serif;
                font-size: 18px;
                color: $teal;
                
                &::-moz-focus-inner{
                    border: 0;
                }
                
                &:focus{
                    background-color: #f0f0f0;
                }
            }
        }
    }
</style>

<script>
    export default {
        name: 'Shorten',
        
        data(){
            return {
                message: 'Hello from shorten',
                placeholder: 'Add your url here...'
            }
        }
    }
</script>