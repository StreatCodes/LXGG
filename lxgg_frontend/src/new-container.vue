<template>
    <main>
        <div class="panel">
            <h1>New Container</h1>
            <form @submit="createContainer">
                <div class="grid">
                    <div>
                        <label>Container Name</label>
                        <input type="text" v-model="name">
                    </div>
                    <div>
                        <label>Folder</label>
                        <select v-model="folder">
                            <option :value="option.Value" :key="index" v-for="(option, index) in folders">{{option.Name}}</option>
                        </select>
                    </div>
                    <div class="new-row">
                        <label>IPV4 Address</label>
                        <input type="text" v-model="ip">
                    </div>
                    <div class="new-row">
                        <button :class="{loading: loading}" @click="createContainer"><span>Deploy Container!</span></button>
                    </div>
                </div>
            </form>
        </div>
    </main>
</template>

<script>
module.exports = {
    props: {
    },
    data: function(){
        return {
            name: "",
            folder: "",
            ip: "",

            folders: [
                {Name: "Folder One", Value: "1"},
                {Name: "Folder 2", Value: "2"},
                {Name: "Folder Three", Value: "3"},
                {Name: "Folder 4", Value: "4"},
            ],
            loading: false
        }
    },
    methods: {
        createContainer: function(e){
            e.preventDefault();
            this.loading = true;

            const payload = {
                name: this.name,
                folder: this.folder,
                ip: this.ip
            };

            fetch("/api/newcontainer",
            {
                method: "POST",
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify(payload)
            }).then((res) => {
                    console.log(this);
                if(res.status == 200){
                    this.$router.push('/containers');
                }else {
                    console.log(res.body);
                }
            });
        }
    }
}
</script>

<style scoped>
</style>