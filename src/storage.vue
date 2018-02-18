<template>
    <main>
        <div class="panel">
			<h2>Storage</h2>

        </div>
    </main>
</template>

<script>
export default {
    mounted: function(){
        this.getStorageList().then(pools => {
            const reqs = [];

            for(const pool of pools) {
                let request = fetch(pool)
                .then(res => {
                    return res.json();
                }).then(data => {
                    return data.metadata;
                });

                reqs.push(request);
            }

            Promise.all(reqs).then(values => {
                this.pools = values;
            });
        });
    },
    data: function(){
        return {
            pools: []
        }
    },
    methods: {
        getStorageList: function() {
            return fetch('/1.0/storage-pools')
            .then(res => {
                return res.json();
            }).then(data => {
                return data.metadata;
            })
        }
    }
}
</script>