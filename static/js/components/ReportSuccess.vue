<template>
    <div class="report-success" :class="{'active': show}">
        <p>Thank you for your report! We will look into it shortly :)</p>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    
    .report-success{
        width: 100%;
        padding-top: 25px;
        padding-bottom: 25px;
        display: none;
        background-color: rgba(46,145,246, 0.2);
        
        &.active{
            display: block;
        }
        
        & > p{
            font-family: $pt-sans;
            color: $dark-text;
            text-align:center;
        }
    }
</style>

<script>
    export default{
        name: 'report-success',
        
        props: ['show']
    }
</script>