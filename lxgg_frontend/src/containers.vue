<template>
    <main>
        <div class="empty-containers" v-if="loading || containers == null">
            <p>No containers found</p>
            <router-link tag="button" to="/new">Create New Container</router-link>
        </div>
        <div v-else>
            <div class="search-bar">
                <input type="text" placeholder="Filter Containers">
            </div>
            <div class="loading-panel" v-if="loading">
                Loading.............
            </div>
            <div class="containers" v-else>
                <div
                    v-for="(container, index) in filteredContainers()"
                    :key="container.id"
                    @click="selectContainer(container)">
                    <div :class="['container', 'panel', {selected: selected == container}]">
                        {{container.name}}
                    </div>
                    <container-details
                        v-if="container == selected" :data="container"></container-details>
                    </div>
                </div>
            </div>
        </div>
    </main>
</template>

<script>
import API from './lxgg-api.js';
import ContainerDetails from './container-details.vue';

export default {
    created () {
        this.fetchContainers()
    },
    props: {
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
                this.containers = json;
                this.loading = false;
				console.log(this.containers);
            })
            .catch((err) => {
                //TODO show fetch error
                this.loading = false;
            });
        },
        filteredContainers: function() {
            return this.containers;
        },
        selectContainer: function(container) {
            console.log("BOo")
            this.selected = container;
        }
    },
    components: {
        'container-details': ContainerDetails
    }
}
</script>

<style scoped>
</style>
