<template>
    <div class="report-cont interior-page">
        <section class="intro">
            <p>
                We have a set of very simple <router-link to="/pages/rules">rules.</router-link>
                If while browsing the internet you happen to find a link using
                our service which breaks these rules please use the following form to
                report it. We will look into it and remove it if necessary. Thanks!
            </p>
        </section>
        <linkid-form :submit="submit"></linkid-form>
        <report-success :show="reportSuccessful"></report-success>
        <span class="linkid-instructions">Enter a linkid and press enter</span>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
</style>

<script>
    import LinkidForm from '../components/LinkidForm.vue';
    import ReportSuccess from '../components/ReportSuccess.vue';
    import {ValidateLinkid, JsonPostRequest} from '../filters/index';
    import {NoticeUserError} from '../filters/ErrorClass';
    
    export default{
        name: 'report-view',
        
        components:{
            LinkidForm,
            ReportSuccess
        },
        
        data(){
            return{
                reportSuccessful: false
            }
        },
        
        methods:{
            submit: function(event){
                let formValue = event.target.elements[0].value;
                try{
                    let linkid = ValidateLinkid(formValue);
                    JsonPostRequest('/api/report', linkid).then(function(response){
                        if(response !== false){
                            if(response.code !== 200){
                                const apiError = new NoticeUserError(response.message, true);
                                apiError.show();
                                this.reportSuccessful = false;
                            }
                            else{
                                this.reportSuccessful = true;
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