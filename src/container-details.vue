<template>
    <div class="container-details panel" v-if="data">
        <h2>{{data.name}}</h2>
        <p>OS: <span>{{data.config['image.description']}}</span></p>
        <p>State: <span>{{data.status}}</span></p>
        <p>Created at: <span>{{new Date(data.created_at).toLocaleString()}}</span></p>
        <p>Last used: <span>{{new Date(data.last_used_at).toLocaleString()}}</span></p>
        <p>Profiles: <span>{{data.profiles.join(', ')}}</span></p>
        <p>Description: <span>{{data.description}}</span></p>
        <template v-for="(device, key) in data.expanded_devices">
            <p :key="key">{{key}}:</p>
            <template v-for="(value, field) in device">
                <p :key="field" class="padded">{{field}}: <span>{{value}}</span></p>
            </template>
        </template>
    </div>
</template>

<script>
import API from './lxgg-api.js';

export default {
    created () {
    },
    props: {
        data: Object
    },
    methods: {
    }
}
</script>
