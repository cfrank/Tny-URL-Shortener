<template>
    <div class="testing">
        <p>{{message}}</p>
    </div>
</template>

<style lang="sass">
    *{
        margin: 0;
        padding: 0;
        box-sizing: border-box;
    }
    
    .testing{
        display: flex;
        & > p{
            color: blue;
        }
    }
</style>

<script>
    export default {
        data() {
            return {
                message: 'Hello from a component'
            }
        }
    }
</script>