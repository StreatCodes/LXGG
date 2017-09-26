export const getImages = () => {
	return fetch("/api/images", {
		credentials: "include",
	})
	.then((res) => {
		if(res.status == 200){
			return res.json();
		} else {
			throw new Error(res.body);
		}
	})
}

export const getContainers = () => {
	fetch("/api/containers", {
		credentials: "include",
	})
	.then((res) => {
		if(res.status == 200){
			return res.json();
		} else {
			throw new Error(res.body);
		}
	})
	.then((json) => {
		this.containers = json;
		this.loading = false;
		console.log(this.containers);
	})
	.catch((err) => {
		//TODO show fetch error
		this.loading = false;
	});
}


export const createContainer = (name, image) => {
	const payload = {
		Name: name,
		Image: image
	};

	return fetch("/api/containers/new",
	{
		method: "POST",
		headers: {'Content-Type': 'application/json'},
		credentials: 'include',
		body: JSON.stringify(payload)
	}).then((res) => {
		console.log(res);
		if(res.status == 200){
			return null
		}else {
			throw new Error(res);
		}
	});
}