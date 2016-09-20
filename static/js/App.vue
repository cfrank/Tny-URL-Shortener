<!-- The Application view -->
<template>
    <div class="viewport">
        <!-- Error handler -->
        <!-- Shortener -->
        <shorten-view></shorten-view>
        <!-- Anouncements -->
        <!-- Menu -->
    </div>
</template>

<style lang="sass">
    @import '../css/header.scss';
    *{
        margin: 0;
        padding: 0;
        box-sizing: border-box;
    }
    
    /* XXX: Not sure anyother way to do this... */
    html,
    body,
    .root,
    .viewport{
        height: 100%;
    }
    
    .viewport{
        background-color: $bg-color;
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>

<script>
    import ShortenView from './views/ShortenView.vue';
    
    export default{
        name: 'app',
        
        components:{
            ShortenView
        }
    }
</script>