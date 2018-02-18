<template>
    <main>
        <div class="panel">
			<h2>Networking</h2>
            <table>
                <thead>
                    <tr>
                        <td>Name</td>
                        <td>Type</td>
                        <td>Description</td>
                        <td>IPV4 Addr</td>
                        <td>IPV4 NAT</td>
                        <td>IPV6 Addr</td>
                        <td>IPV6 NAT</td>
                        <td>Used By</td>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="network in networks" :key="network.name">
                        <td>{{network.name}}</td>
                        <td>{{network.type}}</td>
                        <td>{{network.description}}</td>
                        <td>{{network.config['ipv4.address']}}</td>
                        <td>{{network.config['ipv4.nat']}}</td>
                        <td>{{network.config['ipv6.address']}}</td>
                        <td>{{network.config['ipv6.nat']}}</td>
                        <td>
                            <template v-for="container in network.used_by">
                                <p :key="container">{{container.replace('/1.0/containers/', '')}}</p>
                            </template>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </main>
</template>

<script>
export default {
    mounted: function(){
        this.getNetworkList().then(networks => {
            const reqs = [];

            for(const network of networks) {
                let request = fetch(network)
                .then(res => {
                    return res.json();
                }).then(data => {
                    return data.metadata;
                });

                reqs.push(request);
            }

            Promise.all(reqs).then(values => {
                this.networks = values;
            });
        });
    },
    data: function(){
        return {
            networks: []
        }
    },
    methods: {
        getNetworkList: function() {
            return fetch('/1.0/networks')
            .then(res => {
                return res.json();
            }).then(data => {
                return data.metadata;
            })
        }
    }
}
</script>