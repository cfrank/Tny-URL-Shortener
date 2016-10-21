<template>
    <div class="footer-menu">
        <ul>
            <li><router-link to="/" exact>Shorten</router-link></li>
            <li><router-link to="/pages/history">History</router-link></li>
            <li><router-link to="/pages/unshorten">Unshorten</router-link></li>
            <li><router-link to="/pages/report">Report</router-link></li>
            <li><router-link to="/pages/about">About</router-link></li>
        </ul>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    .footer-menu{
        width: 100%;
        height: auto;
        background-color: $tny-blue;
        padding-bottom: 10px;
        position: absolute;
        bottom: 0;
        left: 0;
        text-align: center;
        
        & > ul{
            display: flex;
            justify-content: center;
            list-style: none;
            
            & > li{
                & > a{
                    font-family: $pt-sans;
                    @include fontSize(18px);
                    text-decoration: none;
                    color: #fff;
                    outline: none;
                    
                    &.router-link-active{
                        color: $active-blue;
                    }
                }
                &:not(:last-child){
                    padding-right: 10px;
                }
            }
        }
    }
</style>

<script>
    name: 'footer-menu'
</script>