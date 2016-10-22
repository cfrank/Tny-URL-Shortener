<template>
    <div class="rules-cont interior-page">
        <section class="rules-content">
            <p>
                In order to keep this service running, and not get blacklisted by service
                providers, there needs to be a few rules regarding the content which can
                be pointed to by our links.
                <br /><br />
                <ul>
                    <li>No affiliate links, or links where you directly profit from each click.</li>
                    <li>Nothing NSFW.</li>
                    <li>Don't point anywhere that automatically executes downloading of files.</li>
                    <li>Don't give me any extra work because you link to something stupid.</li>
                    <li class="disclamer"><i>This list can change without notice</i></li>
                </ul>
                <br />
                If you find a link which is breaking these simple rules, please
                <router-link to="/pages/report">report</router-link> it and I will look
                into it. If you really need to contact me send me an <a href="mailto:chris@cfrank.org">email</a>
            </p>
        </section>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    
    .rules-content > p > ul{
        margin-left: 30px;
        list-style: square;
        
        & > li{
            &:not(:last-child){
                margin-bottom: 20px;
            }
            
            &.disclamer{
                color: $grey-text;
            }
        }
    }
</style>

<script>
    export default{
        name: 'rules-view'
    }
</script>