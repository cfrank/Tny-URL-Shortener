<template>
    <div class="shorten-container">
        <shorten-form></shorten-form>
    </div>
</template>

<style lang="sass">
    .shorten-container{
        width: 100%;
        max-width: 650px;
        margin: 0 auto;
    }
</style>

<script>
    import ShortenForm from '../components/ShortenForm.vue';
    
    export default{
        components:{
            ShortenForm
        },
        
        data(){
            return{
                message: 'Hello from ShortenView'
            }
        }
    }
</script>