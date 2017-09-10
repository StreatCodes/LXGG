<template>
    <div class="app">
        <sidebar></sidebar>
        <transition name="fade">
            <router-view></router-view>
        </transition>
        <footer>footer</footer>
    </div>
</template>

<script>
import SideBar from './sidebar.vue'

module.exports = {
    props: {
    },
    components: {
        'sidebar': SideBar
    }
}
</script>

<style scoped>
</style>