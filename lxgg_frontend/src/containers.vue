<template>
    <main>
        <div class="search-bar">
            <input type="text" placeholder="Filter Containers">
        </div>
        <div class="loading-panel" v-if="loading">
            Loading............
        </div>
        <div class="panel" v-else>
            {{filteredContainers()}}
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
            containers: []
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
                }else {
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
        }
    }
}
</script>

<style scoped>
</style>