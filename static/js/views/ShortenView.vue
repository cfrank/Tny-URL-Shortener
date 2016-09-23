<template>
    <div class="shorten">
        <shorten-form></shorten-form>
        <link-success title="Your short link:" linkhref="https://www.google.com"></link-success>
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
    import ShortenForm from '../components/ShortenForm.vue';
    import LinkSuccess from '../components/LinkSuccess.vue';
    
    export default{
        name: 'ShortenView',
        
        components:{
            ShortenForm,
            LinkSuccess
        },
        
        data(){
            return{
                message: 'Hello from ShortenView'
            }
        }
    }
</script>