<template>
    <div class="panel login">
        <form @submit="login">
            <h1 style="text-align: center;">LXGG</h1>
            <label>Username</label>
            <input type="text" v-model="username">
            <label>Password</label>
            <input type="password" v-model="password">
            <button :class="{loading: loading}" style="width: 100%;" @click="login"><span>Login</span></button>
        </form>
    </div>
</template>

<script>
export default {
    props: {
    },
    data: () => {
        return {
            username: "",
            password: "",
            loading: false
        }
    },
    methods: {
        login: function(e) {
            e.preventDefault();
            this.loading = true;

            const payload = {
                Username: this.username,
                Password: this.password,
            };

            fetch("/lxgg/login",
            {
                method: "POST",
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify(payload)
            }).then((res) => {
                if(res.status == 200){
                    return res.text();
                }else {
                    //TODO show login error
                    this.loading = false;
                    console.log(res);
                }
            }).then(key => {
                console.log("set key");
                window.localStorage.setItem("key", key);
                console.log("Redirect");
                this.$router.push('/containers');
            });
        }
    }
}
</script>