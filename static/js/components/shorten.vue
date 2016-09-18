<template>
    <!-- The main shortener component -->
    <p>{{ message }}</p>
</template>

<script>
    export default {
        name: 'Shorten',
        
        data(){
            return {
                message: 'Hello from shorten'
            }
        }
    }
</script>