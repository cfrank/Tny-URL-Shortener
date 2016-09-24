<template>
    <div class="shorten">
        <shorten-form></shorten-form>
        <link-success 
            :active="linkSuccess.active" 
            :title="linkSuccess.title" 
            :linkhref="linkSuccess.linkHref">
        </link-success>
    </div>
</template>

<style lang="sass">
    .shorten{
        width: 100%;
        max-width: 650px;
        margin: 0 auto;
    }
</style>

<script>
    import {mapGetters} from 'vuex';
    
    import ShortenForm from '../components/ShortenForm.vue';
    import LinkSuccess from '../components/LinkSuccess.vue';
    
    export default{
        name: 'ShortenView',
        
        components:{
            ShortenForm,
            LinkSuccess
        },
        
        computed: mapGetters([ 'linkSuccess' ])
    }
</script>