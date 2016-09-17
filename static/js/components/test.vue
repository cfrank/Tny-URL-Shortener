<template>
    <div class="testing">
        <p>HEllo world!</p>
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