<template>
    <main>
        <div class="panel">
            <h1>New Container</h1>
            <form @submit="createContainer">
                <div class="grid">
                    <div>
                        <label>Container Name</label>
                        <input type="text" v-model="name" placeholder="Bernard Web">
                    </div>
                    <div>
                        <label>Tags (Seperated by a comma)</label>
                        <input type="text" v-model="tags" placeholder="monitoring, amazing, state of the art">
                    </div>
                    <div class="new-row">
                        <label>IPV4 Address</label>
                        <input type="text" v-model="ip" placeholder="192.168.10.69">
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
export default {
    props: {
    },
    data: function(){
        return {
            name: "",
            tags: "",
            ip: "",

            loading: false
        }
    },
    methods: {
        createContainer: function(e){
            e.preventDefault();
            this.loading = true;

            const payload = {
                name: this.name,
                tags: this.tags,
                ip: this.ip
            };

            fetch("/api/containers/new",
            {
                method: "POST",
                headers: {'Content-Type': 'application/json'},
				credentials: 'include',
                body: JSON.stringify(payload)
            }).then((res) => {
                    console.log(this);
                if(res.status == 200){
                    this.$router.push('/containers');
                }else {
                    console.log(res);
                }
            });
        }
    }
}
</script>

<style scoped>
</style>
