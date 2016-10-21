<template>
    <div class="about-cont interior-page">
        <section class="about-content">
            <p>
                This is a simple project by <a href="https://github.com/cfrank" target="_blank">Chris Frank.</a>
                It is a URL shortener which does exactly that, it takes a long URL and gives you
                a more manageable one which can be shared amongst your friends or co-workers.
                <br /><br />
                The site is built with <a href="https://vuejs.org/">Vue JS</a> and <a href="https://golang.org/">Go.</a>
                All code is open source and available on <a href="https://github.com/cfrank/Tny-URL-Shortener">Github</a>
                <br /><br />
                This project has taught me a lot, and I hope you find it useful. Please make sure to check out the simple
                <router-link to="/pages/rules">rules</router-link> so you don't give me extra work. Thanks!
            </p>
        </section>
    </div>
</template>

<style lang="sass">
</style>

<script>
    export default{
        name: 'about-view'
    }
</script>