<template>
    <div class="unshorten-view">
        <section class="intro">
            <p>
                Use this form to fetch the location of any url shortened with
                this service. Also see when the link was created, and how many
                times it has been reported as harmful.
                <br /><br />
                If you find that a link is pointing to a location which breaks our
                <router-link to="/pages/rules">simple rules</router-link>, please
                <router-link to="/pages/report">report</router-link> it so that it can be looked into and
                removed. Thanks!
            </p>
        </section>
        <unshorten-form></unshorten-form>
        <span class="unshorten-instructions">Enter a linkid and press enter</span>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    .unshorten-view{
        width: 100%;
        max-width: 650px;
        padding: 25px;
        background-color: #fff;
        box-shadow: $elm-shadow;
        
        & > .intro{
            margin-bottom: 25px;
            
            & > p{
                font-family: $pt-sans;
                color: $dark-text;
                
                & > a{
                    font-weight: bold;
                    text-decoration: none;
                    color: $tny-blue;
                }
            }
        }
        
        & > .unshorten-instructions{
            width: 100%;
            margin-top: 20px;
            display: block;
            font-family: $pt-sans;
            @include fontSize(13px);
            color: $grey-text;
            text-align: center;
        }
    }
</style>

<script>
    import UnshortenForm from '../components/UnshortenForm.vue';
    
    export default{
        name: 'unshorten-view',
    
        components:{
            UnshortenForm
        }
    }
</script>