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
					<!-- <div class="new-row">
						<label>IPV4 Address</label>
						<input type="text" v-model="ip" placeholder="192.168.10.69">
					</div> -->
					<div class="new-row">
						<label>Image</label>
						<select v-model="image">
							<option v-for="image in imageList" :key="image.fingerprint" :value="image.fingerprint">{{image.properties.description}}</option>
						</select>
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
import * as API from './lxgg-api.js';

export default {
	created: function() {
		API.getImages()
		.then((json) => {
			this.imageList = json;
			console.log(this.imageList);
		})
		.catch((err) => {
			//TODO show fetch error
			console.log(err);
		});
	},
	props: {
	},
	data: function(){
		return {
			name: "",
			tags: "",
			ip: "",
			image: "",

			imageList: [],
			loading: false
		}
	},
	methods: {
		createContainer: function(e) {
			e.preventDefault();
			this.loading = true;

			API.createContainer(this.name, this.image)
			.then((json) => {
				this.$router.push('/containers');
			})
			.catch((err) => {
				//TODO show fetch error
				this.loading = false;
				console.log("Error start===========");
				console.log(err);
				console.log("Error end=============");
			});
		}
	}
}
</script>

<style scoped>
</style>
