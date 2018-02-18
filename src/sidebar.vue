<template>
    <div :class="['sidebar',{hidebar: $route.name == 'login'}]">
        <router-link :class="['new', {active: $route.name == 'new'}]" to="/new"><span>New Container</span></router-link>
        <router-link :class="['containers', {active: $route.name == 'containers'}]" to="/containers"><span>Containers</span></router-link>
        <router-link :class="['images', {active: $route.name == 'images'}]" to="/images"><span>Images</span></router-link>
        <router-link :class="['networking', {active: $route.name == 'networking'}]" to="/networking"><span>Networking</span></router-link>
        <router-link :class="['host', {active: $route.name == 'host'}]" to="/host"><span>Host</span></router-link>
        <router-link :class="['storage', {active: $route.name == 'storage'}]" to="/storage"><span>Storage</span></router-link>
        <router-link :class="['user', {active: $route.name == 'user'}]" to="/user"><span>User</span></router-link>
    </div>
</template>

<script>
export default {
    props: {
    }
}
</script>