<template>
    <!-- The main shortener component -->
    <div class="page-container">
        <shorten></shorten>
    </div>
</template>

<style lang="sass">
    @import '../../css/variables.scss';
    /* Main styling of the page with globals */
    *{
        margin: 0;
        padding: 0;
        box-sizing: border-box;
    }
    
    body{
        width: 100%;
        height: 100%;
        background-color: $test;
    }
    
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