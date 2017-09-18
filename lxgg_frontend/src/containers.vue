<template>
    <main>
        <div class="search-bar">
            <input type="text" placeholder="Filter Containers">
        </div>
        <div class="loading-panel" v-if="loading">
            Loading............
        </div>
        <div class="containers" v-else>
            <div class="list">
                <div
                    v-for="(container, index) in filteredContainers()"
                    :class="['container', 'panel', {selected: selected == container}]"
                    :key="container.id"
                    @click="selectContainer(container)">
                    {{container.Name}}
                </div>
            </div>
            <div class="container-board":class="['panel', {active: selected !== null}]">

            </div>
        </div>
    </main>
</template>

<script>
export default {
    props: {
    },
    created () {
        this.fetchContainers()
    },
    data: function(){
        return {
            loading: true,
            containers: [],
            selected: null,
        }
    },
    methods: {
        fetchContainers: function() {
            fetch("/api/containers", {
                credentials: "include",
            })
            .then((res) => {
                if(res.status == 200){
                    return res.json();
                } else {
                    throw new Error(res.body);
                }
            })
            .then((json) => {
                console.log(json);
                this.containers = json;
                this.loading = false;
            })
            .catch((err) => {
                //TODO show fetch error
                this.loading = false;
                console.log(err);
            });
        },
        filteredContainers: function() {
            return this.containers;
        },
        selectContainer: function(container) {
            console.log("BOo")
            this.selected = container;
        }
    }
}
</script>

<style scoped>
</style>