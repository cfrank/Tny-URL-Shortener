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
        <span class="linkid-instructions">Enter a linkid and press enter</span>
    </div>
</template>

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
                let formValue = event.target.elements[0].value;
                // Validate the linkid
                try{
                    let linkid = ValidateLinkid(formValue);
                    JsonPostRequest('/api/unshorten', linkid).then(function(response){
                        if(response !== false){
                            if(response.code !== 200){
                                const apiError = new NoticeUserError(response.message, true);
                                apiError.show();
                            }else{
                                this.linkData = response.link;
                                this.dataAvailable = true;
                            }
                        }
                    }.bind(this));
                }catch(e){
                    const linkidError = new NoticeUserError(e.message, true);
                    linkidError.show();
                }
            }
        }
    }
</script>