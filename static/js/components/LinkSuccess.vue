<template>
    <div class="success-link" :class="{'active': active}">
        <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" view-box="0 0 48 48">
            <path d="M24 4C12.95 4 4 12.95 4 24s8.95 20 20 20 20-8.95 20-20S35.05 4 24 4zm-2 35.86C14.11 38.88 8 32.16 8 24c0-1.23.15-2.43.42-3.58L18 30v2c0 2.21 1.79 4 4 4v3.86zm13.79-5.07C35.28 33.17 33.78 32 32 32h-2v-6c0-1.1-.9-2-2-2H16v-4h4c1.1 0 2-.9 2-2v-4h4c2.21 0 4-1.79 4-4v-.83c5.86 2.37 10 8.11 10 14.83 0 4.16-1.6 7.94-4.21 10.79z"></path>
        </svg>
        <p class="title">{{title}} <a :href="linkhref" class="link" target="_blank">{{linkhref}}</a></p>
        <span @click="hideLinkSuccess" class="success-link-close">x</span>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    
    .success-link{
        width: 100%;
        height: 0px;
        position: relative;
        text-align: center;
        background-color: #fff;
        box-shadow: $elm-shadow;
        overflow: hidden;
        transition:
            height 0.3s ease-out,
            margin-top 0.3s ease-out,
            padding 0.3s ease-out;
        
        & > svg {
            opacity: 0;
            transition: opacity 0.5s ease-out;
            
            & > path{
                fill: $success;
            }
        }
        
        & > .title{
            font-family: $pt-sans;
            @include fontSize(25px);
            color: $dark-text;
            
            & > a{
                color: $teal-text;
                text-decoration: none;
            }
        }
        
        & > .success-link-close{
            display: block;
            padding: 3px 10px;
            color: $grey-text;
            font-family: $pt-sans;
            @include fontSize(18px);
            position: absolute;
            top: 0px;
            right: 0px;
            cursor: pointer;
            
        }
    }
    
    .success-link.active{
        height: 125.5px;
        margin-top: 20px;
        padding: 20px;
        
        & > svg{
            opacity: 1;
        }
    }
</style>

<script>
    export default{
        name: 'link-success',
        
        props:[
            'active',
            'title',
            'linkhref',
        ],
        
        methods:{
            hideLinkSuccess: function(){
                this.$store.dispatch('hideLinkSuccess');
                this.$store.dispatch('updateFormValue', '');
            }
        }
    }
</script>