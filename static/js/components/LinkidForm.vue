<template>
    <div class="linkid-cont">
        <span class="site-name">{{site_name}}</span>
        <form class="linkid-form" action="/" method="POST" @submit.prevent="submit">
            <input type="text" class="linkid" placeholder="DhvQnR" name="linkid" maxlength="6" />
        </form>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    
    .linkid-cont{
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: flex-end;
        
        & > .site-name{
            font-family: $pt-sans;
            @include fontSize(25px);
            color: $dark-text;
        }
        
        & > .linkid-form{
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
    export default{
        name: 'linkid-form',
        
        // Function from parent which handles data and api call
        props: ['submit'],
        
        data(){
            return{
                site_name: `${location.origin}/`
            }
        }
    }
</script>