<template>
    <div class="unshorten-view interior-page">
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
        <linkid-form :submit="submit"></linkid-form>
        <link-expose :data="linkData" :show="dataAvailable"></link-expose>
        <span class="unshorten-instructions">Enter a linkid and press enter</span>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    .unshorten-view{
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
    import LinkidForm from '../components/LinkidForm.vue';
    import LinkExpose from '../components/LinkExpose.vue';
    import {ValidateLinkid, JsonPostRequest} from '../filters/index';
    import {NoticeUserError} from '../filters/ErrorClass';
    
    export default{
        name: 'unshorten-view',
    
        components:{
            LinkidForm,
            LinkExpose
        },
        
        data(){
            return{
                dataAvailable: false,
                linkData: {}
            }
        },
        
        methods:{
            submit: function(event){
                let linkid = event.target.elements[0].value;
                // Validate the linkid
                try{
                    let formValue = ValidateLinkid(linkid);
                    JsonPostRequest('/api/unshorten', formValue).then(function(response){
                        if(response !== true){
                            if(response.code !== 200){
                                const apiError = new NoticeUserError(response.message, true);
                                apiError.show();
                            }else{
                                this.linkData = response;
                                this.dataAvailable = true;
                            }
                        }
                    }.bind(this));
                }catch(e){
                    let linkidError = new NoticeUserError(e.message, true);
                    linkidError.show();
                }
            }
        }
    }
</script>