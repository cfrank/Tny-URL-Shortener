<template>
    <div class="history-entry">
        <history-topbar :linkid="item.linkid" :source="this.hostname"></history-topbar>
        <div class="history-body">
            <history-info :created="item.created" :clicks="item.clicks" :flags="item.abuseflags"></history-info>
            <history-source :source="item.source"></history-source>
        </div>
    </div>
</template>

<style lang="sass">
    @import '../../css/header.scss';
    
    .history-entry{
        width: 100%;
        max-height: 200px;
        background-color: #f4f4f4;
        box-shadow: $elm-shadow;
        
        & > .history-body{
            padding: 5px 25px;
        }
        
        &:not(:last-child){
            margin-bottom: 25px;
        }
    }
</style>

<script>
    import HistoryTopbar from '../components/HistoryTopbar.vue';
    import HistoryInfo from '../components/HistoryInfo.vue';
    import HistorySource from '../components/HistorySource.vue';
    
    export default{
        name: 'history-entry',
        
        props: ['item'],
        
        components: {
            HistoryTopbar,
            HistoryInfo,
            HistorySource
        },
        
        data(){
            return{
                hostname: ''
            }
        },
        
        created: function(){
            let tempParser = document.createElement('a');
            tempParser.href = this.item.source;
            
            this.hostname = tempParser.host;
            tempParser = null;
        }
    }
</script>