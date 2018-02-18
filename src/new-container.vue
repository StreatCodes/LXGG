<template>
	<main>
		<div class="panel">
			<h1>New Container</h1>
			<form @submit.prevent="createContainer">
				<div class="grid">
					<div>
						<label> <!-- TODO 64 chars max, ASCII, no slash, no colon and no comma -->
							<p>Container Name</p>
							<input type="text" name="name" v-model="name" placeholder="Bernard Web">
						</label>
					</div>
					<div>
						<label>
							<p>Architecture</p>
							<select name="arch" v-model="arch">
								<option>x86_64</option>
							</select>
						</label>
					</div>
					<div class="new-row">
						<label>
							<p>Description</p>
							<textarea v-model="description" name="description" placeholder="Staging server for Bernard Web"></textarea>
						</label>
					</div>
					<div>
						<label><p>Profiles</p></label>
						<label class="cust-check">
							<input type="checkbox" v-model="profiles" name="profiles" value="default">
							<p>default</p>
						</label>
					</div>
					<div class="new-row">
						<label>
							<p>Image</p>
							<select v-model="image">
								<option v-for="image in imageList" :key="image.fingerprint" :value="image.fingerprint">{{image.properties.description}}</option>
							</select>
						</label>
					</div>
					<div class="new-row">
						<label class="cust-check">
							<input type="checkbox" v-model="ephemeral" name="ephemeral" value="default">
							<p>Destroy container on shutdown</p>
						</label>
					</div>
					<div class="new-row">
						<button :class="{loading: loading}" type="submit"><span>Deploy Container!</span></button>
					</div>
				</div>
			</form>
		</div>
	</main>
</template>

<script>
//TODO
// "config": {"limits.cpu": "2"},                                      # Config override.
// "devices": {                                                        # optional list of devices the container should have
//     "kvm": {
//         "path": "/dev/kvm",
//         "type": "unix-char"
//     },
// },
// "instance_type": "c2.micro",                                        # An optional instance type to use as basis for limits

import * as API from './lxgg-api';

export default {
	created: function() {
		API.getImages()
		.then((images) => {
			this.imageList = images;
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
			arch: "x86_64",
			description: "",
			profiles: ['default'],
			ip: "",
			image: "",
			ephemeral: false,

			imageList: [],
			loading: false
		}
	},
	methods: {
		createContainer: function(e) {
			this.loading = true;

			const payload = {
				name: this.name,
				architecture: this.arch,
				profiles: this.profiles,
				description: this.description,
				ephemeral: this.ephemeral,
				source: {
					type: "image",
					fingerprint: this.image
				}
			};

			return fetch('/1.0/containers', {
				method: "POST",
				body: JSON.stringify(payload)
			}).then(res => {
				return res.json();
			}).then(data => {
				console.log(data);
				// this.$router.push('/containers');
			});
		}
	}
}
</script>