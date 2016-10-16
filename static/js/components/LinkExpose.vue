<template>
    <div class="link-expose" :class="{'active': show}">
        <p>The shortened url points to: <br />
        <a class="source" :href="data.source" target="_blank">{{data.source}}</a></p><br />
        <p>The url was created on: <span class="date">{{FormatDate(data.created)}}</span></p><br />
        <p>The url has been reported <span class="flags">{{data.flags}}</span> {{pluralize(data.flags, 'time')}}</p>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    .link-expose{
        width: 100%;
        padding: 25px;
        display: none;
        background-color: rgba(46,145,246, 0.2);
        
        &.active{
            display:block;
        }
        
        & > p{
            font-family: $pt-sans;
            @include fontSize(18px);
            color: $dark-text;
            
            a,
            span{
                color: $tny-blue;
            }
        }
    }
</style>

<script>
    import * as Constants from '../filters/constants';
    import {FormatDate} from '../filters/';
    
    export default{
        name: 'link-expose',
        
        props: ['data', 'show'],
        
        methods:{
            FormatDate: FormatDate,
            
            pluralize: function(num, word){
                if(num === 0 || num > 1){
                    return word + 's';
                }
                else{
                    return word;
                }
            }
        }
    }
</script>