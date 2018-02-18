<template>
    <main>
        <div v-if="host" class="panel host-details">
            <h2>Host info</h2>
            <p>Host Kernel: <span>{{host.environment.kernel}}</span></p>
            <p>Kernel Version: <span>{{host.environment.kernel_version}} ({{host.environment.kernel_architecture}})</span></p>
            <p>Driver: <span>{{host.environment.driver}} ({{host.environment.driver_version}})</span></p>
            <p>Server: <span>{{host.environment.server}} ({{host.environment.server_version}})</span></p>
            <p>Storage: <span>{{host.environment.storage}} ({{host.environment.storage_version}})</span></p>
            <p>Supported Architectures: <span>{{host.environment.architectures.join(', ')}}</span></p>
            <p>API Extensions: <span>{{host.api_extensions.join(', ')}}</span></p>
        </div>
        <div v-else-if="loading" class="loading"></div>
    </main>
</template>

<script>
export default {
    mounted: function(){
        this.getHostInfo().then(info => {
            this.host = info;
        });
    },
    data: function(){
        return {
            host: null,
            loading: false,
        }
    },
    methods: {
        getHostInfo: function() {
            return fetch("/1.0")
            .then((res) => {
                if(res.status == 200){
                    return res.json();
                } else {
                    this.loading = false;
                }
            })
            .then((json) => {
                this.loading = false;
                return json.metadata;
            })
        }
    }
}
</script>