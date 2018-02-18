<template>
    <main>
        <div class="panel">
			<h2>Images</h2>
            <table>
                <thead>
                    <tr>
                        <td>OS</td>
                        <td>Release</td>
                        <td>Architecture</td>
                        <td>Auto Update</td>
                        <td>Public</td>
                        <td>Size</td>
                        <td>Last Used</td>
                        <td>Update Source</td>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="image in getSortedImages()" :key="image.fingerprint">
                        <td>{{image.properties.os}}</td>
                        <td>{{image.properties.release}}</td>
                        <td>{{image.properties.architecture}}</td>
                        <td>{{image.auto_update}}</td>
                        <td>{{image.public}}</td>
                        <td>{{image.size}}</td>
                        <td>{{new Date(image.last_used_at).toLocaleString()}}</td>
                        <td><a rel="noopener noreferrer" target="_blank" :href="image.update_source.server">{{image.update_source.server}}</a></td>
                    </tr>
                </tbody>
            </table>
        </div>
    </main>
</template>

<script>
export default {
    mounted: function() {
        this.getImageList().then(images => {
            const reqs = [];

            for(const image of images) {
                let request = fetch(image)
                .then(res => {
                    return res.json();
                }).then(data => {
                    return data.metadata;
                });

                reqs.push(request);
            }

            Promise.all(reqs).then(values => {
                this.images = values;
            });
        });
    },
    data: function(){
        return {
            images: []
        }
    },
    methods: {
        getImageList: function() {
            return fetch('/1.0/images')
            .then(res => {
                return res.json();
            }).then(data => {
                return data.metadata;
            })
        },
        getSortedImages: function() {
            const images = this.images.slice();
            return images.sort((a,b) => {
                if(a.properties.os > b.properties.os){
                    return true
                } else {
                    return !(b.properties.os > a.properties.os)
                }
            });
        }
    }
}
</script>