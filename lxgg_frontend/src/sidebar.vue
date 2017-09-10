<template>
    <div class="sidebar">
        <router-link class="new" to="/new"></router-link>
        <router-link class="containers active" to="/containers"></router-link>
        <router-link class="images" to="/images"></router-link>
        <router-link class="networking" to="/networking"></router-link>
        <router-link class="host" to="/host"></router-link>
        <router-link class="settings" to="/settings"></router-link>
        <router-link class="user" to="/user"></router-link>
    </div>
</template>

<script>
module.exports = {
    props: {
    }
}
</script>

<style scoped>
</style>