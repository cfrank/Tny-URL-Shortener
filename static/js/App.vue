<!-- The Application view -->
<template>
    <div class="viewport">
        <!-- Error handler -->
        <notice-header></notice-header>
        <!-- View component -->
        <router-view></router-view>
        <!-- Menu -->
    </div>
</template>

<style lang="sass">
    @import '../css/header.scss';
    *{
        margin: 0;
        padding: 0;
        box-sizing: border-box;
    }
    
    html,body,
    .viewport{
        height: 100%;
    }
    
    .viewport{
        background-color: $bg-color;
        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>

<script>
    import NoticeHeader from './components/NoticeHeader.vue';
    
    export default{
        name: 'AppView',
        
        components:{
            NoticeHeader
        },
    }
</script>