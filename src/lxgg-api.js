//TODO add LXGG auth here
export function getImages() {
	//Get image list
	return fetch('/1.0/images')
	.then(res => {
		return res.json();
	}).then(data => {
		const images = data.metadata;
		const reqs = [];

		//Fetch info for each image
		for(const image of images) {
			let request = fetch(image)
			.then(res => {
				return res.json();
			}).then(data => {
				return data.metadata;
			});

			reqs.push(request);
		}

		//When all requests complete
		const result = Promise.all(reqs).then(images => {
			//Sort images by OS name
			return images.sort((a,b) => {
				if(a.properties.os > b.properties.os){
					return true
				} else {
					return !(b.properties.os > a.properties.os)
				}
			});
		});

		return result;
	});
}