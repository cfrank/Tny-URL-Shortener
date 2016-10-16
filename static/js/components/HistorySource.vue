<template>
    <div class="source-link">
        <a :href="source" target="_blank">
            {{source}}
        </a>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    
    .source-link{
        padding: 25px;
        text-align: center;
        
        & > a{
            font-family: $pt-sans;
            @include fontSize(24px);
            color: $tny-blue;
            text-decoration: none;
        }
    }
</style>

<script>
    export default{
        name: "history-source",
        
        props: ['source']
    }
</script>