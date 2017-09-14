<template>
    <div :class="['sidebar',{hidebar: $route.name == 'login'}]">
        <router-link :class="['new', {active: $route.name == 'new'}]" to="/new"></router-link>
        <router-link :class="['containers', {active: $route.name == 'containers'}]" to="/containers"></router-link>
        <router-link :class="['images', {active: $route.name == 'images'}]" to="/images"></router-link>
        <router-link :class="['networking', {active: $route.name == 'networking'}]" to="/networking"></router-link>
        <router-link :class="['host', {active: $route.name == 'host'}]" to="/host"></router-link>
        <router-link :class="['settings', {active: $route.name == 'settings'}]" to="/settings"></router-link>
        <router-link :class="['user', {active: $route.name == 'user'}]" to="/user"></router-link>
    </div>
</template>

<script>
export default {
    props: {
    }
}
</script>

<style scoped>
</style>