<template>
    <div class="container-details">
        <h3>Description</h3>
        <p>{{data.description}}</p>
        <h3>Config</h3>
        <p><span class="bold">OS: </span>{{data.config['image.os']}}</p>
        <p><span class="bold">Release: </span>{{data.config['image.release']}}</p>
        <p><span class="bold">Architecture: </span>{{data.config['image.architecture']}}</p>
        <h3>Devices</h3>
        <p>TODO maybe..</p>

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
