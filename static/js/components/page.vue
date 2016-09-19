<!-- The main shortener component -->
<template>
    <div class="page-container flex-box">
        <shorten></shorten>
    </div>
</template>

<style lang="sass">
    /* Main styling of the page with globals */
    @import '../../css/header.scss';
    
    *{
        margin: 0;
        padding: 0;
        box-sizing: border-box;
    }

    html,
    body{
        width: 100%;
        height: 100%;
    }

    body{
        background-color: $bg-color;
    }
    
    .container,
    .page-container{
        width: 100%;
        height: 100%;
    }
</style>

<script>
    import Shorten from './shorten.vue';
    
    export default {
        name: 'Page',
        
        components: {
            Shorten
        },
        
        data() {
            return {
                message: 'Hello from a component'
            }
        }
    }
</script>