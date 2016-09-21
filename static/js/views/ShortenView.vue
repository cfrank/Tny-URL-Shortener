<template>
    <div class="shorten-container">
        <shorten-form></shorten-form>
        <p :data-test="testing">{{testing}}</p>
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
    import store from '../vuex/store';
    import { mapGetters } from 'vuex';
    
    import ShortenForm from '../components/ShortenForm.vue';
    
    export default{
        name: 'ShortenView',
        
        components:{
            ShortenForm
        },
        
        created(){
            console.log(this);
        },
        
        computed: mapGetters([ 'testing' ]),
        
        data(){
            return{
                message: 'Hello from ShortenView'
            }
        }
    }
</script>