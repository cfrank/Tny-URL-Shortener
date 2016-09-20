<!-- The main shortener component -->
<template>
    <div class="hello-world">
        <p>Hello from App.vue</p>
    </div>
</template>

<style lang="sass">
    .hello-world{
        & > p{
            color: green;
        }
    }
</style>

<script>
    export default{
        name: 'app'
    }
</script>