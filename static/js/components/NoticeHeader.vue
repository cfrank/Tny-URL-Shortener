<template>
    <div class="notice-header" :class="{'showing': notice.show}">
        <p :class="notice.type">{{notice.message}}</p>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    
    .notice-header{
        width: 100%;
        height: 50px;
        background-color: white;
        text-align: center;
        line-height: 50px;
        position: absolute;
        top: 0px;
        left: 0px;
        box-shadow: $elm-shadow;
        transform: translateY(-100%);
        transition: transform 0.3s ease-out;
        
        &.showing{
            transform: translateY(0%);
        }
        
        & > p{
            font-family: $pt-sans;
            @include fontSize(20px);
            
            &.error{
                color: $error;
            }
            
            &.success{
                color: $success;
            }
        }
    }
</style>

<script>
    import { mapGetters } from 'vuex';
    
    export default{
        name: 'notice-header',
        // Get state
        computed: mapGetters([ 'notice' ]),
    }
</script>