<template>
    <main>
        <div v-if="loading" class="loading">
            loading TODO spinner
        </div>
        <div v-else>
            <div v-if="containers.length < 1" class="empty-containers">
                <p>No containers found</p>
                <router-link tag="button" to="/new">Create New Container</router-link>
            </div>
            <div v-else class="split-view">
                <div class="containers panel">
                    <div class="search-bar">
                        <input type="text" placeholder="Filter Containers" v-model="filter">
                    </div>
                    <div
                        v-for="container in filteredContainers()"
                        :key="container.id"
                        :class="['container-item', {selected: selected == container}]"
                        @click="selectContainer(container)">
                            {{container.name}}
                    </div>
                </div>
                <container-details :data="selected" />
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
            filter: ""
        }
    },
    methods: {
        fetchContainers: function() {
            fetch("/1.0/containers")
            .then((res) => {
                if(res.status == 200){
                    return res.json();
                } else {
                    this.loading = false;
                }
            })
            .then((json) => {
                const reqs = [];
                for(const url of json.metadata) {
                    this.fetchContainerData(url); //TODO we can make this async
                }

                this.loading = false;
            })
        },
        fetchContainerData: function(url) {
            fetch(url)
            .then((res) => {
                if(res.status == 200){
                    return res.json();
                } else {
                    this.loading = false;
                }
            }).then(container => {
                this.containers.push(container.metadata);
            });
        },
        filteredContainers: function() {
            const containers = [];
            //Loop through containers and only include containers with filter in name
            for(let container of this.containers) {
                const filter = this.filter.toLowerCase();

                if (container.name.toLowerCase().indexOf(filter) !== -1 ||
                    container.config['image.description'].toLowerCase().indexOf(filter) !== -1 ||
                    container.status.toLowerCase().indexOf(filter) !== -1) {
                    containers.push(container);
                }
            }

            return containers;
        },
        selectContainer: function(container) {
            this.selected = this.selected === container ? null : container;
        }
    },
    components: {
        'container-details': ContainerDetails
    }
}
</script>

<style scoped>
</style>
