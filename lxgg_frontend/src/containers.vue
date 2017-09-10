<template>
    <main>
        <div class="search-bar">
            <input type="text" placeholder="Filter Containers">
        </div>
        <div class="panel">
            {{$route.name}}
            Data
        </div>
    </main>
</template>

<script>
module.exports = {
    props: {
    }
}
</script>

<style scoped>
</style>